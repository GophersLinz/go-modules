package main

import (
	"io/ioutil"

	bf "github.com/russross/blackfriday/v2"
	log "github.com/sirupsen/logrus"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	data, err := ioutil.ReadFile("markdown.md")
	check(err)

	log.Info(string(data))

	output := bf.Run(data)

	log.Info(string(output))
}
