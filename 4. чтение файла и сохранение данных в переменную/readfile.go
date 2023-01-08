package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("!!!READING FILE, WAIT A BIT 0_0!!!")
	f, err := os.Open("temp/чтение файла.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//Чтение файла с ридером
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	fmt.Println("@@@IM READING A FILE@@@")
	//для более мелких файлов
	fContent, err := os.ReadFile("temp/чтение файла.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))
}
