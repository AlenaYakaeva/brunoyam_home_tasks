package repository

import (
	"ToDoList/internal/repository/db"
	"ToDoList/internal/repository/memstorage"
	"errors"
	"time"

	"github.com/golang-migrate/migrate"
)

type Repository struct {
	DB         *db.Storage
	MemStorage *memstorage.Storage
}

func New(dbDSN string, maxAttempts int) (*Repository, error) {
	MemStorage := memstorage.New()
	if dbDSN == "" {
		return &Repository{
			DB:         nil,
			MemStorage: MemStorage,
		}, errors.New("DBDSN is empty")
	}
	repo, err := ConnectWithRetry(dbDSN, maxAttempts)
	if err != nil {
		return &Repository{
			DB:         nil,
			MemStorage: MemStorage,
		}, errors.New("Connection faild")
	} else {
		if err = db.RunMigrations(dbDSN); err != nil {
			if !errors.Is(err, migrate.ErrNoChange) {

				return &Repository{
					DB:         nil,
					MemStorage: MemStorage,
				}, errors.New("Migration faild")
			}
		}
	}

	return &Repository{
		DB:         repo,
		MemStorage: MemStorage,
	}, nil
}

func ConnectWithRetry(dbDSN string, maxAttempts int) (*db.Storage, error) {
	delay := 1 * time.Second

	for attempt := 1; attempt <= maxAttempts; attempt++ {

		repo, err := db.New(dbDSN)
		if err != nil {

			if attempt == maxAttempts {
				return nil, err
			}

			time.Sleep(delay)
			delay *= 2

		} else {
			return repo, nil
		}
	}
	return nil, errors.New("Connection faild")
}
