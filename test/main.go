package main

import (
	"decode"
	"fmt"
	"path/filepath"
)

func main() {
	//fmt.Println("hello world")
	files := []string{"3H2A7200.MP4", "3H2A8080.JPG", "3H2A8080.CR3"}
	for _, image := range files {
		fmt.Println(image)
		run(image)
		fmt.Println()
	}

}

func run(path string) {
	img, err := decode.Read_img(path)

	if err != nil {
		fmt.Println(err)
	}
	ext := filepath.Ext(path)

	name, err := decode.Camera_name(img, ext)
	if err != nil {
		fmt.Print(path, ": ")
		fmt.Println(err)
	}

	date, err := decode.Image_date(img, ext)
	if err != nil {
		fmt.Print(path, ": ")
		fmt.Println(err)
	}
	fmt.Println(name)
	fmt.Println(date)
}
