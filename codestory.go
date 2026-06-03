// +build !heroku

package main

import (
	_ "bitbucket.org/tbruyelle/codestory/elevator"
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
