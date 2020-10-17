package TestZGo

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//zad 1
func StringLength(s string) int {
	len := 0
	for range s {
		len += 1
	}
	return len
}

//zad 2
func StringLengthFlag() int {
	stringFlag := flag.String("napis", "", "flaga z napisem")
	flag.Parse()
	return len(*stringFlag)
}

//zad 3
func MakeMapFromFile(path string) map[string]int {
	mapa := map[string]int{}
	for _, i := range ReadLinesFromFile(path) {
		i = strings.ToLower(i)
		mapa[i]++
	}
	return mapa
}

func ReadLinesFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)

	results := []string{}
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return results
}

//zad 4
func MakeMapFromFileAndSaveIt(path string) (map[string]int, error) {
	mapa := map[string]int{}
	for _, i := range ReadLinesFromFile(path) {
		i = strings.ToLower(i)
		mapa[i]++
	}

	f, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer f.Close()

	for key, val := range mapa {
		fmt.Fprintln(f, key, " - ", val)
	}

	if err != nil {
		fmt.Println(err)
		f.Close()
		return nil, err
	}

	return mapa, nil
}

//zad 5

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!", r.URL.Path[1:])
}

//zad 6
func GetPage(path string) []byte {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
