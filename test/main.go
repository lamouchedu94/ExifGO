package main

import (
	"fmt"
	"path/filepath"

	decode "github.com/lamouchedu94/ExifGO"
)

func main() {
	//fmt.Println("hello world")
	files := "/home/paul/Images/Test/3H2A0105.CR3"
	run(files)
	/*
		for _, image := range files {
			fmt.Println(image)
			run(image)
			fmt.Println()
		}
	*/
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
	_, _ = name, date
	//fmt.Println(name)
	fmt.Println(date)
}
