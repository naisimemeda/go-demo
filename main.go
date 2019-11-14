package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-demo/route"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Getenv("mysql_database"))

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
