package main

import (
	"fmt"
	"log"

	"github.com/ProstoyVadila/fuckgit/config"
)

func main() {
	log.Println("Loading config...")
	config, err := config.New("./")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config.GinMode)

	log.Println("Starting server...")
	app, err := NewApp(config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(app.config.Addr)
}
