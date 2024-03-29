package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url ="https://catalogapi.nso.go.th/api/index?table=SES_OS_19&format=json"

func main()  {
/*
	content := "สวัสดี จากภาษา Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
	defer readFile("./fromString.txt")
*/
	// อ่านข้อความไฟล์จากเว็บ
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	// fmt.Print(content)

	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}
}
/*
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
*/

func toursFromJson(content string) []Tour  {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}