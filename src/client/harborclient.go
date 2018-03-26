package client

import (
	"fmt"
	"log"

	"hubimage/src/config"
	"hubimage/swagger"
)

func newConfiguration() *swagger.Configuration {
	conf := config.GetConfig()
	harbor_url := fmt.Sprintf("%s://%s:%s/api", conf.HarborPortocol, conf.HarborHost, conf.HarborPort)
	log.Printf("Connect harbor: %s", harbor_url)

	cfg := &swagger.Configuration{
		BasePath:      harbor_url,
		DefaultHeader: make(map[string]string),
		APIKey:        make(map[string]string),
		APIKeyPrefix:  make(map[string]string),
		APIClient:     &swagger.APIClient{},
		Host:          harbor_url,
		Username:      conf.Username,
		Password:      conf.Password,
	}

	cfg.APIClient.SetConfig(cfg)
	return cfg
}

func NewHarborClient() *swagger.ProductsApi {
	configuration := newConfiguration()
	log.Printf("Create harbor client successfully.")
	return &swagger.ProductsApi{
		Configuration: configuration,
	}
}
