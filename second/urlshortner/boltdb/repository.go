package boltdb

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

const (
	urlShortnerBuckerName = "urlshortner"
)

// Repository implements the urlshortner Repository interface
// and encapsulates the logic for a boltdb interaction
type Repository struct {
	db *bolt.DB
}

// NewRepository attemps to connect to a boltdb with a given
// filepath and, if successfull, returns a Repository instance
func NewRepository(filePath string) (Repository, error) {
	repo := Repository{}
	if err := repo.openConnection(filePath); err != nil {
		return repo, err
	}

	repo.insertMockValues()
	return repo, nil
}

// openConnection opens a connection to a boltdb given a filepath
// and sets a ref to that connection in the Repository instance
func (repo *Repository) openConnection(filePath string) error {
	db, err := bolt.Open(filePath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	repo.db = db
	return nil
}

func (repo Repository) insertMockValues() {
	errs := make(chan error, len(mockData))
	var wg sync.WaitGroup

	for _, data := range mockData {
		wg.Add(1)
		go func(data map[string]string) {
			errs <- repo.db.Batch(func(tx *bolt.Tx) error {
				if bucket, err := tx.CreateBucketIfNotExists([]byte(urlShortnerBuckerName)); err != nil {
					return err
				} else if bucket != nil {
					return bucket.Put(
						[]byte(data["path"]),
						[]byte(data["url"]),
					)
				}

				return errors.New("no bucket exist with this name")
			})
			wg.Done()
		}(data)
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (repo Repository) matchShortURL(shortURL string) string {
	var url string

	repo.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(urlShortnerBuckerName))
		if bucket == nil {
			return nil
		}

		if bytes := bucket.Get([]byte(shortURL)); bytes != nil {
			url = string(bytes)
		}

		return nil
	})

	return url
}
