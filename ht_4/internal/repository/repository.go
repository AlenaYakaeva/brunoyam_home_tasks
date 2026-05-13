package repository

import (
	"ToDoList/internal/repository/db"
	"ToDoList/internal/repository/memstorage"
	"errors"

	"github.com/golang-migrate/migrate"
)

type Repositury struct {
	DB         *db.Storage
	MemStorage *memstorage.Storage
	IsRemote   bool
}

func New(dbDSN string) *Repositury {
	MemStorage := memstorage.New()
	isRemote := true
	repo, err := db.New(dbDSN)
	if err != nil {
		isRemote = false
	} else {
		if err = db.RunMigrations(dbDSN); err != nil {
			if !errors.Is(err, migrate.ErrNoChange) {
				isRemote = false
			}
		}
	}

	return &Repositury{
		DB:         repo,
		MemStorage: MemStorage,
		IsRemote:   isRemote,
	}
}
