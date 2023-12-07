package repository

import "gorm.io/gorm"

type MailRepository interface {
	GetAll() ([]Mail, error)
}

type Mail struct {
	gorm.Model
	To   string
	From string
	Body string
	Type string
}

func (m *Mail) GetAll() ([]Mail, error) {
	// var mails []Mail
	// err := DB.Find(&mails).Error
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}
