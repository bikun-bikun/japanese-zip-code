package main

import (
	"flag"
	"fmt"
	"os"

	japanese_zip_code "github.com/bikun-bikun/japanese-zip-code"
)

func main() {
	z := flag.String("z", "", "zip code")
	flag.Parse()

	c := japanese_zip_code.NewClient()
	zipCode, err := c.GetZipCode(*z)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(zipCode)
	os.Exit(0)
}
