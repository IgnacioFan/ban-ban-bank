package conn

import (
	"go-bank-express/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=bank_express port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panicf("unable to connect postgres db, err: %v", err)
	}
	db.AutoMigrate(&entity.User{}) // should remove
	return db, nil
}
