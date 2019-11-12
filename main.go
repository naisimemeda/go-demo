package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"go-demo/route"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := route.Router()

	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
