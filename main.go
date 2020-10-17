package main

import (
	"fmt"
	"github/filip/test1/TestZGo"
	"net/http"
)

func main() {
	//fmt.Println(TestZGo.NewUser("Adam", "Kowalski", 18).ToString())
	//fmt.Println(TestZGo.StringLengthFlag())
	//fmt.Println(TestZGo.StringLength("łąć"))
	//fmt.Println(TestZGo.MakeMapFromFileAndSaveIt("txt.txt"))
	fmt.Println(TestZGo.GetPage("https://google.com"))

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}
