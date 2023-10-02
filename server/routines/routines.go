package routines

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"sync"

	"github.com/alitto/pond"
	"github.com/kevinronu/email-indexer/server/models"
	"github.com/kevinronu/email-indexer/server/utils"
	"github.com/kevinronu/email-indexer/server/zinc"
)

func parseEmailFiles(pathsChan <-chan string, emailsChan chan<- models.Email) {
	for path := range pathsChan {
		email, err := utils.FileToEmail(path)
		if err != nil {
			log.Printf("WARNING: failed to parse %v: %v\n", path, err)
		} else {
			emailsChan <- email
		}
	}

	// for {
	// 	path, ok := <-pathsChan
	// 	if !ok {
	// 		break
	// 	}
	// 	email, err := utils.FileToEmail(path)
	// 	if err != nil {
	// 		log.Printf("WARNING: failed to parse %v: %v\n", path, err)
	// 	} else {
	// 		emailsChan <- email
	// 	}
	// }
}

func uploadEmails(emailsChan <-chan models.Email, bulkUploadQuantity int, zincCredentials zinc.ZincService) {
	bulk := models.BulkEmails{
		Index:   "emails",
		Records: make([]models.Email, bulkUploadQuantity),
	}
	count := 0
	totalUploaded := 0

	for email := range emailsChan {
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

// func walkDirectories(dir string, pathsChan chan<- string, wgWalkers *sync.WaitGroup) {
// 	defer wgWalkers.Done()

// 	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, errWalk error) error {
// 		if errWalk != nil {
// 			fmt.Printf("WARNING: Error walking '%s': %v\n", dir, errWalk)
// 			return errWalk
// 		}
// 		if entry.IsDir() && path != dir {
// 			wgWalkers.Add(1)
// 			go walkDirectories(path, pathsChan, wgWalkers)
// 			return filepath.SkipDir
// 		}
// 		if !entry.IsDir() {
// 			pathsChan <- path
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Printf("WARNING: Error walking directory '%s': %v\n", dir, err)
// 	}
// }

func ParseAndUploadEmails(emailsDir string, numUploaderWorkers int, bulkUploadQuantity int, zincAuth zinc.ZincService) {
	fileCount := 0
	err := filepath.WalkDir(emailsDir, func(path string, entry fs.DirEntry, errWalk error) error {
		if errWalk != nil {
			fmt.Printf("WARNING: Error walking '%s': %v\n", emailsDir, errWalk)
		}
		if !entry.IsDir() {
			fileCount++
		}
		return nil
	})
	if err != nil {
		log.Fatalf("FATAL: Error walking directory '%s': %v", emailsDir, err)
	}
	log.Printf("INFO: the number of files in the directory '%s' is %d\n", emailsDir, fileCount)

	pathsChan := make(chan string, 4)
	emailsChan := make(chan models.Email)

	parserPool := pond.New(100, fileCount)
	for i := 0; i < fileCount; i++ {
		parserPool.Submit(func() {
			parseEmailFiles(pathsChan, emailsChan)
		})
	}

	log.Printf("TRACE: spawning %d uploader goroutines\n", numUploaderWorkers)
	var wgUploaders sync.WaitGroup
	for i := 0; i < numUploaderWorkers; i++ {
		wgUploaders.Add(1)
		go func() {
			defer wgUploaders.Done()
			uploadEmails(emailsChan, bulkUploadQuantity, zincAuth)
		}()
	}

	err = filepath.WalkDir(emailsDir, func(path string, entry fs.DirEntry, errWalk error) error {
		if errWalk != nil {
			fmt.Printf("WARNING: Error walking '%s': %v\n", emailsDir, errWalk)
		}
		if !entry.IsDir() {
			pathsChan <- path
		}
		return nil
	})
	if err != nil {
		log.Fatalf("FATAL: Error walking directory '%s': %v", emailsDir, err)
	}

	// var wgWalkers sync.WaitGroup
	// wgWalkers.Add(1)
	// go walkDirectories(emailsDir, pathsChan, &wgWalkers)
	// wgWalkers.Wait()

	close(pathsChan)
	parserPool.StopAndWait()
	close(emailsChan)
	wgUploaders.Wait()
}
