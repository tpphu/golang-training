package main

import (
	"bufio"
	"bytes"
	// "encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Ward struct {
	Id   uint32
	Name string
}

type District struct {
	Id    uint32
	Name  string
	Wards []Ward
}

type City struct {
	Id         uint32
	Name       string
	Disctricts []District
}

var aMapOfCities map[int]City = make(map[int]City)
var aMapOfDistrics map[int]District = make(map[int]District)

func readFileAndGetResult(fn string) []City {
	var output []City
	file, err := os.Open(fn)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		var buffer bytes.Buffer

		var l []byte
		var isPrefix bool
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)
			if !isPrefix {
				break
			}
			if err != nil {
				break
			}
		}
		if err == io.EOF {
			break
		}
		line := buffer.String()
		parts := strings.Split(line, ",")
		if parts[0] == "Tỉnh Thành Phố" {
			continue
		}
		// 3. Lay city
		id, _ := strconv.Atoi(parts[1])
		if _, ok := aMapOfCities[id]; !ok {
			city := City{
				Id:   uint32(id),
				Name: parts[0],
			}
			aMapOfCities[id] = city
			output = append(output, city)
		}
		var city *City
		for i := 0; i < len(output); i++ {
			temp := output[i] // @TOTO: Can hieu cho nay 1
			if temp.Id == uint32(id) {
				city = &temp // @TOTO: Can hieu cho nay 1
				break
			}
		}
		// 2. Lay huyen
		id, _ = strconv.Atoi(parts[3])
		if _, ok := aMapOfDistrics[id]; !ok {
			district := District{
				Id:    uint32(id),
				Name:  parts[2],
				Wards: []Ward{},
			}
			aMapOfDistrics[id] = district
			city.Disctricts = append(city.Disctricts, district)
		}
		var district *District
		for i := 0; i < len(city.Disctricts); i++ {
			temp := city.Disctricts[i] // @TOTO: Can hieu cho nay 2
			if temp.Id == uint32(id) {
				district = &temp // @TOTO: Can hieu cho nay 2
				break
			}
		}

		// 1. Lay xa
		id, _ = strconv.Atoi(parts[5])
		ward := Ward{
			Id:   uint32(id),
			Name: parts[4],
		}
		// 2.1 Append cai xa vo huyen
		fmt.Println("city:", city)
		fmt.Println("district:", district)
		district.Wards = append(district.Wards, ward)
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}
	return output
}

func main() {
	file := "./small.csv"
	data := readFileAndGetResult(file)
	// jsonByte, _ := json.Marshal(data)
	// fmt.Println(string(jsonByte))

	dataByte, _ := yaml.Marshal(data)
	fmt.Println(string(dataByte))
}
