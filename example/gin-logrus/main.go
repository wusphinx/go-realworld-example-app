package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	// 日志切割 refer:https://github.com/natefinch/lumberjack
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		logger.Infof("hello world")
		c.String(200, "")
	})

	//nolint
	r.Run()
}
