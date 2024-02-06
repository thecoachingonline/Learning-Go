package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	content := "สวัสดี จากภาษา Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
	defer readFile("./fromString.txt")
}

func readFile(fileName string)  {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("ข้อความที่อ่านจากไฟล์:", string(data))
}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}