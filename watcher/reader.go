package watcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/alllll/watchsender/models"
)

type attachJson struct {
	File string
}

type emailJson struct {
	To       string
	Subject  string
	Body     string
	Attaches []attachJson
}

func Reader(path string, emailchan chan models.Email) {

	emjson := decode(path)
	email := models.Email{}
	email.To = emjson.To
	email.Subject = emjson.Subject
	email.Body = emjson.Body
	for _, attach := range emjson.Attaches {
		attachpath := strings.Join([]string{filepath.Dir(path), "\\", attach.File}, "")
		dat, _ := ioutil.ReadFile(attachpath)
		at := models.Attach{}
		at.File = dat
		at.Name = filepath.Base(attach.File)
		email.Attaches = append(email.Attaches, at)
		//remove attach
		os.Remove(attachpath)
	}
	//remove files
	os.Remove(path)

	emailchan <- email

}

func decode(path string) emailJson {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening config file:", err)
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error:", err)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) //look BOE
	email := emailJson{}
	err = json.Unmarshal(body, &email)
	if err != nil {
		fmt.Println("error:", err)
	}
	return email
}
