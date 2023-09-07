package routines

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"sync"

	"github.com/kevinronu/email-indexer/server/models"
	"github.com/kevinronu/email-indexer/server/utils"
	"github.com/kevinronu/email-indexer/server/zinc"
)

func parseEmailFiles(pathsChan <-chan string, emailsChan chan<- *models.Email) {
	for path := range pathsChan {
		email, err := utils.FileToEmail(path)
		if err != nil {
			log.Printf("WARNING: failed to parse %v: %v\n", path, err)
		} else {
			emailsChan <- email
		}
	}
}

func uploadEmails(emailsChan <-chan *models.Email, bulkUploadQuantity int, zincAuth *models.ZincAuth) {
	bulk := &models.BulkEmails{
		Index:   "emails",
		Records: make([]models.Email, bulkUploadQuantity),
	}
	parsed := 0
	total := 0

	for email := range emailsChan {
		bulk.Records[parsed] = *email
		parsed++
		if parsed == bulkUploadQuantity {
			log.Printf("TRACE: uploading %d emails\n", parsed)
			err := zinc.UploadEmails(bulk, zincAuth)
			if err != nil {
				log.Fatal("FATAL: failed to upload emails: ", err)
			}
			total += parsed
			parsed = 0
		}
	}
	if parsed > 0 {
		bulk.Records = bulk.Records[:parsed]
		err := zinc.UploadEmails(bulk, zincAuth)
		if err != nil {
			log.Fatal("FATAL: failed to upload emails: ", err)
		}
		total += parsed
	}
	log.Printf("INFO: goroutine uploaded %d emails\n", total)
}

func ParseAndUploadEmails(emailsDir string, numUploaderWorkers int, numParserWorkers int, bulkUploadQuantity int, zincAuth *models.ZincAuth) {
	pathsChan := make(chan string)
	emailsChan := make(chan *models.Email)
	// pathsChan := make(chan string, numParserWorkers)
	// emailsChan := make(chan *models.Email, bulkUploadQuantity)

	log.Printf("TRACE: spawning %d uploader goroutines", numUploaderWorkers)
	var wgUploaders sync.WaitGroup
	for i := 0; i < numUploaderWorkers; i++ {
		wgUploaders.Add(1)
		go func() {
			defer wgUploaders.Done()
			uploadEmails(emailsChan, bulkUploadQuantity, zincAuth)
		}()
	}

	log.Printf("TRACE: spawning %d parser goroutines", numParserWorkers)
	var wgParsers sync.WaitGroup
	for i := 0; i < numParserWorkers; i++ {
		wgParsers.Add(1)
		go func() {
			defer wgParsers.Done()
			parseEmailFiles(pathsChan, emailsChan)
		}()
	}

	err := filepath.WalkDir(emailsDir, func(path string, entry fs.DirEntry, errWalk error) error {
		if errWalk != nil {
			log.Fatal(errWalk)
		}
		if !entry.IsDir() {
			pathsChan <- path
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	close(pathsChan)
	wgParsers.Wait()
	close(emailsChan)
	wgUploaders.Wait()
}
