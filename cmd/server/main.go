package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	srv := &http.Server{
		Addr:    ":80",
		Handler: InitGin(),
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println("Server forced to shutdown")
	}
}

func InitGin() *gin.Engine {
	ginEngine := gin.New()

	ginEngine.GET("/print", func(context *gin.Context) {
		data, err := httputil.DumpRequest(context.Request, true)
		if err != nil {
			_, _ = context.Writer.WriteString(err.Error())
			return
		}

		_, _ = context.Writer.Write(data)
	})

	return ginEngine
}
