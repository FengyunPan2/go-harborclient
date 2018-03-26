package client

import (
	"log"

	dockerclient "github.com/docker/docker/client"

	"hubimage/swagger"
)

// Clients for docker and harbor
type Clients struct {
	dockerClient *dockerclient.Client
	harborClient *swagger.ProductsApi
}

var clients = &Clients{}

func init() {
	initDockerClient()
	initHarborClient()
}

// GetDockerClient return docker client
func GetDockerClient() *dockerclient.Client {
	if clients.dockerClient == nil {
		log.Printf("init docker client again.")
		initDockerClient()
	}

	return clients.dockerClient
}

// GetHarborClient return harbor client
func GetHarborClient() *swagger.ProductsApi {
	if clients.harborClient == nil {
		log.Printf("init harbor client again.")
		initHarborClient()
	}

	return clients.harborClient
}

func initDockerClient() {
	dockerClient, err := NewDockerClient()
	if err != nil {
		log.Fatalf("Failed to create docker client: %v", err)
	}

	clients.dockerClient = dockerClient
}

func initHarborClient() {
	harborClient := NewHarborClient()

	clients.harborClient = harborClient
}
