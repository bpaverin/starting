package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {

// 	f, err := os.Open("temp/file.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	c, err := io.ReadAll(f)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("### File content ###\n%s\n", string(c))
// 	f.Close()

// 	f, err = os.OpenFile("temp/text.txt", os.O_CREATE|os.O_RDWR, 0777)
// 	if err != nil {
// 		panic(err)
// 	}
// 	io.WriteString(f, "Hello, my name is Boris, im from Togliatti, im 18 years old and im starting learn programming")
// 	f.Close()

// 	f, err = os.Open("temp/text.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	c, err = io.ReadAll(f)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("!!!FILE CONTENT BRO ^_^!!!\n%s\n", string(c))
// 	f.Close()
// }
