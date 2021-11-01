package main

import (
	"fmt"
	"log"
)

var banner = `

$ cd ../../../ZipCreater 
$ # Directory traversal/path traversal exploit creater.
$ # Blog: https://payloads.online

`

func CheckErrorOnExit(err error) {
	if err != nil {
		panic(err)
	}
	return
}
func CheckErrorOnPrint(err error) {
	if err != nil {
		log.Println(err)
	}
	return
}

func ShowBanner() {
	fmt.Println(banner)
}
