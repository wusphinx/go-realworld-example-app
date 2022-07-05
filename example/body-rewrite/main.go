package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func OverWriteBody(newBody []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		body, err := ioutil.ReadAll(io.TeeReader(c.Request.Body, &buf))
		if err != nil {
			c.AbortWithStatusJSON(500, err)
			return
		}
		log.Printf("old body is %s", body)

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(newBody))
	}
}

func main() {
	r := gin.New()

	r.POST("/overwrite", OverWriteBody([]byte(`go`)), func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(500, err)
			return
		}

		c.JSON(200, gin.H{"body": string(body)})
	})

	//nolint
	r.Run()
}
