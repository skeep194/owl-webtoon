package database

import (
	"context"
	"io/ioutil"
	"os"

	"encoding/json"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	data, err := os.Open("mongodb.auth.json")
	if err != nil {
		return nil, errors.Wrap(err, "could not read mongodb.auth.json")
	}

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, errors.Wrap(err, "could not read byte data")
	}
	json.Unmarshal(byteValue, ret)
	return ret, nil
}

func mongoConn() (client *mongo.Client, err error) {
	authInfo, err := getAuthConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not get auth config")
	}

	credential := options.Credential{
		Username: authInfo.Username,
		Password: authInfo.Password,
	}
	clientOptions := options.Client().ApplyURI("mongodb://" + authInfo.Hostname + authInfo.Port).SetAuth(credential)

	// ctx, _ := context.WithTimeout(context.Background(), time.Second*30)

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect mongodb instance")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not ping mongo client")
	}
	return client, nil
}

func GetCollection(colName string) (ret *mongo.Collection, err error) {
	authInfo, err := getAuthConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not get auth config")
	}

	conn, err := mongoConn()
	if err != nil {
		return nil, errors.Wrap(err, "mongodb connection error")
	}
	return conn.Database(authInfo.DBname).Collection(colName), nil
}
