/*
Пакет cong считывает кофигурацию для подключения к базе данных и к серверу
*/
package conf

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// DBConf считывает конфигурацию для базы данных и возвращает строку которую можно использовать для подключения к бд
func DbConf() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Error loading .env file")
	}

	dbType := os.Getenv("DBTYPE")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbAddr := os.Getenv("DBADDR")
	dbName := os.Getenv("DBNAME")

	return fmt.Sprintf("%v://%v:%v@%v/%v", dbType, dbUser, dbPass, dbAddr, dbName), nil

}

// ServerConf считывает конфигурацию для сервера и возвращает строку которую можно использовать для подключения к серверу
func ServerConf() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Error loading .env file")
	}

	sHost := os.Getenv("SERVHOST")
	sPort := os.Getenv("SERVPORT")

	return fmt.Sprintf("%v:%v", sHost, sPort), nil
}
