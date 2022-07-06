package main

import (
	_ "embed" //nolint
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App string `yaml:"app" json:"app,omitempty"`
}

//go:embed config.yml
var rawFile string

func main() {
	r := gin.New()

	var conf Config

	if err := yaml.Unmarshal([]byte(rawFile), &conf); err != nil {
		log.Fatalf("unmarshal failed err: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		//output: {"app":"example"}
		c.JSON(200, conf)
	})

	//nolint
	r.Run()
}
