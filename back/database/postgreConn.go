package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type authConfig struct {
	DBname   string
	Username string
	Password string
	Hostname string
	Port     string
}

func getAuthConfig() (ret *authConfig, err error) {
	ret = new(authConfig)
	data, err := os.Open("postgredb.auth.json")
	if err != nil {
		return nil, errors.Wrap(err, "could not read postgredb.auth.json")
	}

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, errors.Wrap(err, "could not read byte data")
	}
	json.Unmarshal(byteValue, ret)
	return ret, nil
}

var PostgreDB *gorm.DB = nil

func init() {
	authConfig, err := getAuthConfig()
	if err != nil {
		log.Fatalln("get authconfig fail")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable", authConfig.Hostname, authConfig.Username, authConfig.Password, authConfig.DBname, authConfig.Port)
	PostgreDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("db connection fail")
	}
}
