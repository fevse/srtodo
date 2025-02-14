/*
пакет storage содержит модель данных Task для хранения в бд и функцию для подключения к бд
*/

package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var Storage *pgx.Conn

// NewStorage создает новое подключение к базе данных
func NewStorage(conf string) error {
	conn, err := pgx.Connect(context.Background(), conf)
	if err != nil {
		return err
	}
	Storage = conn
	return nil
}
