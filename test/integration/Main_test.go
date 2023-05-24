package integration

import (
	"cursoGo/test/test_utils/containers"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"testing"
)

type IntegrationContext struct {
	mongoClient *mongo.Client
}

var context IntegrationContext

func TestMain(m *testing.M) {
	log.Println("before all package tests")
	context = initGlobalContext()
	exitCode := m.Run()
	log.Println("after all package tests")

	os.Exit(exitCode)
}

func initGlobalContext() IntegrationContext {
	return IntegrationContext{
		mongoClient: initDb(),
	}
}

func initDb() *mongo.Client {
	user := "myuser"
	password := "mysupersecretpassword"
	mongoContainer := containers.MongoContainer{Version: "latest", Env: map[string]string{
		"MONGO_INITDB_ROOT_USERNAME": user,
		"MONGO_INITDB_ROOT_PASSWORD": password,
	}}
	mongoContainer.Run()
	return mongoContainer.Client()
}
