/*
пакет storage содержит модель данных Task для хранения в бд и функцию для подключения к бд
*/

package storage

import (
	_ "github.com/jackc/pgx/stdlib" // driver
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

var Storage *sqlx.DB

// NewStorage создает новое подключение к базе данных
func NewStorage(conf string) error {
	conn, err := sqlx.Connect("pgx", conf)
	if err != nil {
		return err
	}
	Storage = conn
	return nil
}

func Migrate(dir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(Storage.DB, dir); err != nil {
		return err
	}
	return nil
}
