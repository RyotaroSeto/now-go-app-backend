package infrastructure

import (
	"context"
	"fmt"
	"log"
	"now-go-kon/pkg/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TxKey string

const txKey TxKey = "auth_tx"

type DB struct {
	*gorm.DB
}

var d *DB

func RDBConnect(config util.Config) error {
	db, err := createDB(config)
	if err != nil {
		return err
	}
	d = &DB{db}
	return nil
}

func createDB(config util.Config) (*gorm.DB, error) {
	dsn := "host=" + config.DBHost + " user=" + config.DBUser + " password=" + config.DBPassword + " dbname=" + config.DBName + " port=" + config.DBPort + " sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to db", err)
		return nil, err
	}

	return db, nil
}

func GetDB() *DB {
	return d
}

func (d *DB) Transaction(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	tx := d.Begin()
	defer func(tx *gorm.DB) {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("recovered from panic, err: %s", r)
		}
	}(tx)

	txCtx := context.WithValue(ctx, txKey, tx)
	if err = fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
