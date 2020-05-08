package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	japanese_zip_code "github.com/bikun-bikun/japanese-zip-code"
)

func main() {
	z := flag.String("z", "", "zip code")
	flag.Parse()

	if *z == "" {
		fmt.Println("-z は必須です。")
		os.Exit(1)
	}

	c := japanese_zip_code.NewClient()
	zipCode, err := c.GetZipCode(*z)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	json, err := json.Marshal(zipCode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(json))
	os.Exit(0)
}
