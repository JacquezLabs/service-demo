package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getPort() string {

	value := os.Getenv("PORT")
	if value == "" {
		value = ":8080"
	}
	return value
}

func getTimeSeconds() float64 {

	value := os.Getenv("TIME_SECONDS")
	seconds, err := strconv.ParseFloat(value, 64)
	if err != nil {

		seconds = 10
	}
	return seconds
}

func home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to k8s Demo service!!")
}

func consumeCPU(c *gin.Context) {
	x := 0.0001
	for i := 1; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	msg := fmt.Sprintf("OK: %v", x)
	c.String(http.StatusOK, msg)
}

func getHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s", hostname)
}
func main() {
	started := time.Now()

	TIME_SECONDS := getTimeSeconds()
	PORT := getPort()
	getHostname()
	router := gin.Default()

	router.GET("/", home)
	router.GET("/consumeCPU", consumeCPU)
	router.GET("/started", func(c *gin.Context) {
		data := (time.Since(started)).String()
		c.String(http.StatusOK, data)
	})
	router.GET("/healthz", func(c *gin.Context) {
		duration := time.Since(started)

		if duration.Seconds() > TIME_SECONDS {
			message := fmt.Sprintf("error: %v", duration.Seconds())
			c.String(http.StatusInternalServerError, message)
			return
		}
		c.String(http.StatusOK, "ok")
	})

	router.Run(PORT)

}
