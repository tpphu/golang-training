package main

import (
	"fmt"

	"./convert"
)

func main() {
	file := "./data.csv"
	converter := convert.NewCSVToYamlConverter()
	cities, err := converter.ConvertCSV2Yaml(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Cach 1 => Work
	var haNoiCity *convert.City
	haNoiCity = cities[0]
	haNoiCity.ID = 10
	fmt.Println("City ID:", haNoiCity.ID)
	fmt.Println("City Name:", haNoiCity.Name)
	//
	fmt.Println("City ID:", cities[0].ID)
	fmt.Println("City Name:", cities[0].Name)

	// Cach 2 => Work
	// var haNoiCity convert.City
	// haNoiCity = *cities[0]
	// haNoiCity.ID = 10
	// fmt.Println("City ID:", haNoiCity.ID)
	// fmt.Println("City Name:", haNoiCity.Name)
	// //
	// fmt.Println("City ID:", cities[0].ID)
	// fmt.Println("City Name:", cities[0].Name)
	//
	// for i := 0; i < len(cities); i++ {
	// 	city := cities[i]
	// 	fmt.Println("==========================")
	// 	fmt.Println("City ID:", city.ID)
	// 	fmt.Println("City Name:", city.Name)
	// 	for _, district := range city.Districts {
	// 		fmt.Println("------------")
	// 		fmt.Println("District ID:", district.ID)
	// 		fmt.Println("District Name:", district.Name)
	// 	}
	// }
	// data, err := yaml.Marshal(&cities)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))
}
