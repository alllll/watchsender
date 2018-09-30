package watcher

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/alllll/watchsender/models"
)

//Watch return string from file in directory
func Watch(path string, extension string, c chan models.Email) {
	ticker := time.NewTicker(time.Millisecond * 3000)
	go func() {
		for range ticker.C {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				if !file.IsDir() && filepath.Ext(file.Name()) == extension {
					//c <- strings.Join([]string{path, file.Name()}, "")
					//c <- file.Name()
					Reader(strings.Join([]string{path, file.Name()}, ""), c)
				}
			}
		}
	}()

}
