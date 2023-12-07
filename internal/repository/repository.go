package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Repository struct {
	Mail MailRepository
}

func New() *Repository {
	return &Repository{
		Mail: &Mail{},
	}
}

func Connect() error {
	var err error
	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=postgres dbname=hermes-mailer sslmode=disable timezone=UTC connect_timeout=5"), &gorm.Config{})
	if err != nil {
		return err
	}
	db.DB()

	return nil
}
