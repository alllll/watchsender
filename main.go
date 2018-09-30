package main

import (
	"fmt"

	"github.com/alllll/watchsender/models"

	"github.com/alllll/watchsender/sender"
	"github.com/alllll/watchsender/watcher"
)

var config Config

func main() {
	config.Load("config.json")
	fileschan := make(chan models.Email)
	go watcher.Watch(config.Path, config.Extension, fileschan)
	go sender.Sender(fileschan, config.Smtphost, config.Smtpport, config.Smtplogin, config.Smtppassword, config.From)
	/*for s := range fileschan {
		fmt.Println(s)
	}*/
	var a string
	fmt.Scanln(&a)
	fmt.Println(a)
}
