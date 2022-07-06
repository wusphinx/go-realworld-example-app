package main

import (
	_ "embed"
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
	yaml.Unmarshal([]byte(rawFile), &conf)

	r.GET("/", func(c *gin.Context) {
		//output: {"app":"example"}
		c.JSON(200, conf)
	})

	//nolint
	r.Run()
}
