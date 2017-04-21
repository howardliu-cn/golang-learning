package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	ServerName string
	ServerIp   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	file, err := os.Open("file/json/unmarshal/servers.json")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var s Serverslice
	err = json.Unmarshal(data, &s)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(s)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
