package infrastructure

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TxKey string

const txKey TxKey = "auth_tx"

type DB struct {
	*gorm.DB
}

var d *DB

// アプリケーション起動時に呼ばれる関数
func RDBConnect(dsn string) error {
	var db *gorm.DB = nil // TODO:アクセス作成する
	// var db *gorm.DB = createDB() // TODO:アクセス作成する

	d = &DB{db}
	return nil
}

func GetDB() *DB {
	return d
}

func (d *DB) Transaction(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	tx := d.Begin()
	defer func(tx *gorm.DB) {
		if r := recover(); r != nil {
			tx.Rollback()
			err = errors.New(fmt.Sprintf("recovered from panic, err: %s", r))
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
