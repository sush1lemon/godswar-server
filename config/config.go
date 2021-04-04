package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type AppConfig struct {
	Server     server   `yaml:"server"`
	GameServer server   `yaml:"gameserver"`
	Database   database `yaml:"database"`
}

type server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Type string `yaml:"type"`
}

type database struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Port     string `yaml:port`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (c *AppConfig) GetConfig() *AppConfig {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
