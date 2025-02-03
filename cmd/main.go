package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	r "github.com/rcsolis/endoflife_client/internal/api/routes"
)

const (
	READ_TIMEOUT  = 5 * time.Second
	WRITE_TIMEOUT = 5 * time.Second
)

func main() {

	gin.ForceConsoleColor()

	router := gin.Default()
	r.GetRoutes(router)

	s := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    READ_TIMEOUT,
		WriteTimeout:   WRITE_TIMEOUT,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
