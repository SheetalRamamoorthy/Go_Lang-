package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"os"
)

type personaldetails struct {
	Experience int    `json:"Experience"`
	Name       string `json:"Name"`
	Skills     []struct {
		Name        string        `json:"Name"`
		Experience  string        `json:"experience"`
		Projectlist []interface{} `json:"projectlist"`
		Projects    bool          `json:"projects"`
	} `json:"Skills"`
	Team string      `json:"Team"`
	Phd  interface{} `json:"phd"`
}

func main() {
	jsonFile, err := os.Open("sample.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)
}
