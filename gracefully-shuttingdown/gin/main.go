package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	prodMod := true
	if !prodMod {
		_, err := os.Create("status/live")
		fmt.Println("err", err)
		defer os.Remove("status/live")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/super-slow-process", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "finally i'm done",
		})
	})

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	// // run srv in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// // create channel of os signal for waiting signal
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// // close `quit` channel when app're closed
	// defer func() {
	// 	close(quit)
	// }()
	s := <-quit
	log.Println("signal is: ", s)
	log.Println("Shutting down app...")

	// // The context is used to inform the server it has 5 seconds to finish
	// // the request it is currently handling
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// // create context -> waiting every process in server is done
	ctx := context.Background()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("App forced to shutdown:", err)
	}

	log.Println("App exiting")
	os.Remove("/status/live")

}
