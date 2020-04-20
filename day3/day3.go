package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func add_color(colors []string) []string {
	return append(colors, "sky blue")
}

func main() {

	// slice 1
	colors := []string{"red", "blue", "green", "yellow", "purple"}
	update_slice(colors)
	fmt.Println("length is ", len(colors))
	fmt.Println(colors)
	new_color := add_color(colors)
	fmt.Println(new_color)
	fmt.Println("updated len", len(new_color))

	//2 map
	m := map[int]string{1: "rahul", 2: "akshay"}
	fmt.Println("Original map=======")
	printmap(m)
	modifymap(m, 1, "bir")
	fmt.Println("modified map=========")
	printmap(m)
	deletemap(m, 1)
	fmt.Println("deleted map ========")
	printmap(m)

	//http.get implementation
	getresp("https://google.com")

}

func update_slice(colors []string) {
	colors[4] = "orange"
}

func printmap(m map[int]string) {
	for key, value := range m {
		fmt.Println(key, " : ", value)
	}
}
func modifymap(m map[int]string, rollno int, name string) {
	m[rollno] = name
}

func deletemap(m map[int]string, rollno int) {
	delete(m, rollno)
}

func getresp(url string) {
	resp, error := http.Get(url)
	if error != nil {
		fmt.Println(error)
	} else {
		defer resp.Body.Close()
		contents, error := ioutil.ReadAll(resp.Body)
		if error != nil {
			fmt.Println(error)
		}
		fmt.Println("contents --> ", string(contents))
		fmt.Println("response code --> ", resp.StatusCode)
	}
}
