package client

import (
	"fmt"
	"log"

	"github.com/docker/docker/client"

	"hubimage/src/config"
)

func getDockerClient() (*client.Client, error) {
	conf := config.GetConfig()
	dockerUrl := fmt.Sprintf("%s://%s:%s", conf.DockerPortocol, conf.HarborHost, conf.DockerTCPPort)
	log.Printf("Connect docker: %s", dockerUrl)

	cli, err := client.NewClient(dockerUrl, "", nil, nil)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func NewDockerClient() (*client.Client, error) {
	dockerClient, err := getDockerClient()
	if err != nil {
		return nil, err
	}

	log.Printf("Create docker client successfully.")
	return dockerClient, nil
}
