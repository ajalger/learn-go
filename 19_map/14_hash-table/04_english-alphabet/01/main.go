package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/fatih/color"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	blue := color.New(color.FgBlue)
	blue.Println(str)

}
