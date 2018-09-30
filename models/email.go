package models

//Email struct for email :)
type Email struct {
	To       string
	Subject  string
	Body     string
	Attaches []Attach
}

//Attach struct for attaches
type Attach struct {
	Name string
	File []byte
}
