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

func parseEmailFiles(pathsChan <-chan string, emailsChan chan<- models.Email) {
	// for path := range pathsChan {
	// 	email, err := utils.FileToEmail(path)
	// 	if err != nil {
	// 		log.Printf("WARNING: failed to parse %v: %v\n", path, err)
	// 	} else {
	// 		emailsChan <- email
	// 	}
	// }

	for {
		path, ok := <-pathsChan
		if !ok {
			break
		}
		email, err := utils.FileToEmail(path)
		if err != nil {
			log.Printf("WARNING: failed to parse %v: %v\n", path, err)
		} else {
			emailsChan <- email
		}
	}
}

func uploadEmails(emailsChan <-chan models.Email, bulkUploadQuantity int, zincCredentials zinc.ZincService) {
	bulk := models.BulkEmails{
		Index:   "emails",
		Records: make([]models.Email, bulkUploadQuantity),
	}
	count := 0
	totalUploaded := 0

	// for email := range emailsChan {
	// 	bulk.Records[count] = email
	// 	count++
	// 	if count == bulkUploadQuantity {
	// 		log.Printf("TRACE: uploading %d emails\n", count)
	// 		err := zinc.UploadEmails(bulk, zincCredentials)
	// 		if err != nil {
	// 			log.Fatal("FATAL: Failed to upload emails:", err)
	// 		}
	// 		totalUploaded += count
	// 		count = 0
	// 	}
	// }

	for {
		email, ok := <-emailsChan
		if !ok {
			break
		}
		bulk.Records[count] = email
		count++
		if count == bulkUploadQuantity {
			log.Printf("TRACE: uploading %d emails\n", count)
			err := zinc.UploadEmails(bulk, zincCredentials)
			if err != nil {
				log.Fatal("FATAL: Failed to upload emails:", err)
			}
			totalUploaded += count
			count = 0
		}
	}
	// Upload any remaining emails (less than bulkUploadQuantity) in the last batch
	if count > 0 {
		log.Printf("TRACE: uploading %d emails\n", count)
		bulk.Records = bulk.Records[:count]
		err := zinc.UploadEmails(bulk, zincCredentials)
		if err != nil {
			log.Fatal("FATAL: Failed to upload emails:", err)
		}
		totalUploaded += count
	}
	log.Printf("INFO: goroutine uploaded %d emails\n", totalUploaded)
}

func walkDirectories(dir string, pathsChan chan<- string, wgWalkers *sync.WaitGroup) {
	defer wgWalkers.Done()

	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, errWalk error) error {
		if errWalk != nil {
			fmt.Printf("WARNING: Error walking '%s': %v\n", dir, errWalk)
			return errWalk
		}
		if entry.IsDir() && path != dir {
			wgWalkers.Add(1)
			go walkDirectories(path, pathsChan, wgWalkers)
			return filepath.SkipDir
		}
		if !entry.IsDir() {
			pathsChan <- path
		}
		return nil
	})

	if err != nil {
		fmt.Printf("WARNING: Error walking directory '%s': %v\n", dir, err)
	}
}

func ParseAndUploadEmails(emailsDir string, numUploaderWorkers int, numParserWorkers int, bulkUploadQuantity int, zincAuth zinc.ZincService) {
	pathsChan := make(chan string, bulkUploadQuantity)
	emailsChan := make(chan models.Email, bulkUploadQuantity)

	log.Printf("TRACE: spawning %d uploader goroutines\n", numUploaderWorkers)
	var wgUploaders sync.WaitGroup
	for i := 0; i < numUploaderWorkers; i++ {
		wgUploaders.Add(1)
		go func() {
			defer wgUploaders.Done()
			uploadEmails(emailsChan, bulkUploadQuantity, zincAuth)
		}()
	}

	log.Printf("TRACE: spawning %d parser goroutines\n", numParserWorkers)
	var wgParsers sync.WaitGroup
	for i := 0; i < numParserWorkers; i++ {
		wgParsers.Add(1)
		go func() {
			defer wgParsers.Done()
			parseEmailFiles(pathsChan, emailsChan)
		}()
	}

	var wgWalkers sync.WaitGroup
	wgWalkers.Add(1)
	go walkDirectories(emailsDir, pathsChan, &wgWalkers)
	wgWalkers.Wait()

	close(pathsChan)
	wgParsers.Wait()
	close(emailsChan)
	wgUploaders.Wait()
}
