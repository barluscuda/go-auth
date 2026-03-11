package main

import (
	"log"
	"os"
)

func main() {
    log.Println("server is starting...")
    log.Println("server is running in PID:", os.Getpid())
}