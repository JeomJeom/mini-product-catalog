package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"thalesapi/configs"
	"thalesapi/middleware"
	"thalesapi/router"
	"time"
)

func main() {

	config := configs.LoadProjectConfig()

	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{config.UIUrl},
		AllowMethods: []string{
			http.MethodPut,
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodDelete},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	engine.Use(middleware.MemoryUsageLogger())
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	r := router.NewRouter(engine, config)
	r.LoadRouter()

	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	// goroutine to launch a server for http request
	go func() {
		defer wg.Done()
		log.Infof("running on port: %s\n", config.Port)

		if err := http.ListenAndServe(":"+config.Port, engine); err != nil {
			log.Fatal(err)
		}
	}()
	// wait until WaitGroup is done
	wg.Wait()
}
