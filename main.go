package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v32/github"
)

func main() {
	pwd, _ := os.Getwd()
	host, _ := os.Hostname()
	fmt.Printf("🌋: Hello World!\n💻: %s\n📂: %s\n⏰: %s\n", host, pwd, time.Now().Format("2006-01-02T15:04:05-0700"))

	client := github.NewClient(nil)
	fmt.Println(client)
}
