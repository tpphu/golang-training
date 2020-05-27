package main

import (
	"fmt"
	"./convert"
	"gopkg.in/yaml.v2"
)



func main() {
	file := "./data.csv"
	cities, err := convert.ConvertCSV2Yaml(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := yaml.Marshal(&cities)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}