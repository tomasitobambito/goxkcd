package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://imgs.xkcd.com/comics/exoplanet_system.png"

	response, err := http.Get(url)

	fmt.Println(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	file, err := os.Create("./temp.png")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// either get full body extract image or get image url
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success1")
}
