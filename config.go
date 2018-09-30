package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Config configuration app
type Config struct {
	Port         int
	Path         string
	Extension    string
	Smtphost     string
	Smtptls      bool
	Smtpport     string
	Smtplogin    string
	Smtppassword string
	From         string
}

//Load loading configuration app from file
func (config *Config) Load(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening config file:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
}
