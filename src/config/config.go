package config

import (
	"errors"
	"flag"

	"fmt"
	"github.com/spf13/pflag"
)

var (
	address        = pflag.String("address", "0.0.0.0", "bind http address")
	port           = pflag.String("port", "9096", "http listen port")
	harborprotocol = pflag.String("harborprotocol", "http", "The protocol to access harbor server ")
	harborhost     = pflag.String("harborhost", "", "The address of the harbor server ")
	harborport     = pflag.String("harborport", "80", "The port of the harbor server.")
	username       = pflag.String("username", "", "harbor user name.")
	password       = pflag.String("password", "", "password of the user")
	email          = pflag.String("email", "test@example.com", "email of the user.")
	dockertcpport  = pflag.String("dockertcpport", "2375", "a tcp port to connect docker server")
	dockerprotocol = pflag.String("dockerprotocol", "http", "The protocol to access docker server")
)

var conf *ImageConfig

// ImageConfig store config which telnet kube-apiserver and etcd
type ImageConfig struct {
	Address        string `json:"address"`
	Port           string `json:"port"`
	HarborPortocol string `json:"harborprotocol"`
	HarborHost     string `json:"harborhost"`
	HarborPort     string `json:"harborport"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	DockerTCPPort  string `json:"dockertcpport"`
	DockerPortocol string `json:"dockerprotocol"`
}

func init() {
	conf = loadConfig()
}

func loadConfig() *ImageConfig {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	return &ImageConfig{
		Address:        *address,
		Port:           *port,
		HarborPortocol: *harborprotocol,
		HarborHost:     *harborhost,
		HarborPort:     *harborport,
		Username:       *username,
		Password:       *password,
		Email:          *email,
		DockerTCPPort:  *dockertcpport,
		DockerPortocol: *dockerprotocol,
	}
}

// GetConfig return config
func GetConfig() *ImageConfig {
	return conf
}

// GetConfig return config
func CheckConfig() error {
	if len(conf.HarborHost) == 0 {
		errmsg := fmt.Sprintf("--harborhost must be specified.")
		return errors.New(errmsg)
	}

	if len(conf.Username) == 0 {
		errmsg := fmt.Sprintf("--username must be specified.")
		return errors.New(errmsg)
	}

	if len(conf.Password) == 0 {
		errmsg := fmt.Sprintf("--password must be specified.")
		return errors.New(errmsg)
	}

	return nil
}
