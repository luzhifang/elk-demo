package main

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func init() {
	logger := &lumberjack.Logger{
		Filename:   "logs/elk-demo.log",
		MaxSize:    1,
		MaxBackups: 7,
		MaxAge:     30,
	}
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	log.SetOutput(logger)

	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = logger
}

func main() {
	engine := gin.Default()

	engine.GET("log", func(context *gin.Context) {
		log.Info("单条日志测试")
		context.JSON(200, gin.H{
			"msg": "log create success",
		})
	})

	engine.GET("logs", func(context *gin.Context) {
		count, _ := strconv.Atoi(context.Query("count"))
		if count == 0 {
			count = 100
		}

		for i := 0; i < count; i++ {
			log.WithField("i", i).Info("批量日志测试")
		}
		context.JSON(200, gin.H{
			"msg": "logs create success",
		})
	})

	err := engine.Run(":5000")
	if err != nil {
		panic(err)
	}
}
