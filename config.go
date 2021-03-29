package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
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

func (c *config) getConfig() *config {

	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
