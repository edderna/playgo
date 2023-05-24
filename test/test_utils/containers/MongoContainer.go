package containers

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type MongoContainer struct {
	Version   string
	Env       map[string]string
	container testcontainers.Container
}

func (m *MongoContainer) Run() {
	ctx := context.Background()
	port := "27017"
	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{port + "/tcp"},
		WaitingFor:   wait.ForListeningPort(nat.Port(port)).WithStartupTimeout(30 * time.Second),
		Env:          m.Env,
	}
	mongoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Printf("Can't start mongo container %s", err.Error())
		os.Exit(-1)
	}
	m.container = mongoContainer
}

func (m *MongoContainer) exposedPort() string {
	port, err := m.container.MappedPort(context.Background(), "27017")
	if err != nil {
		fmt.Printf("Can't start mongo container %s", err.Error())
		os.Exit(-1)
	}
	return port.Port()
}

func (m *MongoContainer) Client() *mongo.Client {
	ctx := context.Background()
	mongoUri := fmt.Sprintf("mongodb://%s:%s@localhost:%s", m.Env["MONGO_INITDB_ROOT_USERNAME"],
		m.Env["MONGO_INITDB_ROOT_PASSWORD"], m.exposedPort())
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		fmt.Printf("Can't connect to mongo: %s", err.Error())
		os.Exit(-1)
	}
	return client
}
