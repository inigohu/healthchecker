package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	s, err := strconv.Atoi(os.Getenv("INTERVAL"))
	if err != nil {
		fmt.Println(err)
	}
	for range time.NewTicker(time.Duration(s) * time.Second).C {
		check()
	}
}

func check() {
	resp, err := http.Get(fmt.Sprintf("http://%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}
