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
	count := 0
	totalUploaded := 0

	for email := range emailsChan {
		bulk.Records[count] = *email
		count++
		if count == bulkUploadQuantity {
			log.Printf("TRACE: uploading %d emails\n", count)
			err := zinc.UploadEmails(bulk, zincAuth)
			if err != nil {
				log.Fatal("FATAL: failed to upload emails: ", err)
			}
			totalUploaded += count
			count = 0
		}
	}
	// Upload any remaining emails (less than bulkUploadQuantity) in the last batch
	if count > 0 {
		log.Printf("TRACE: uploading %d emails\n", count)
		bulk.Records = bulk.Records[:count]
		err := zinc.UploadEmails(bulk, zincAuth)
		if err != nil {
			log.Fatal("FATAL: failed to upload emails: ", err)
		}
		totalUploaded += count
	}
	log.Printf("INFO: goroutine uploaded %d emails\n", totalUploaded)
}

func ParseAndUploadEmails(emailsDir string, numUploaderWorkers int, numParserWorkers int, bulkUploadQuantity int, zincAuth *models.ZincAuth) {
	pathsChan := make(chan string)
	emailsChan := make(chan *models.Email)

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
