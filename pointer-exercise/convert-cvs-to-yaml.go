package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Ward struct {
	Id      int    `yaml:"id"`
	Name    string `yaml:"name"`
	Deleted bool   `yaml:"deleted,omitempty"`
}

type District struct {
	Id      int     `yaml:"id"`
	Name    string  `yaml:"name"`
	Deleted bool    `yaml:"deleted,omitempty"`
	Wards   []*Ward `yaml:"wards"`
}

type Region struct {
	Id        int         `yaml:"id"`
	Name      string      `yaml:"name"`
	Deleted   bool        `yaml:"deleted,omitempty"`
	Districts []*District `yaml:"districts"`
}

var mapRegion map[int]*Region = map[int]*Region{}
var mapDistrict map[int]*District = map[int]*District{}

func convertCSVToYAML(fn string) (data []*Region, err error) {
	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return nil, err
	}

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
		if len(parts) < 6 {
			panic("Line is not correct | data: " + line)
		}
		if parts[0] == "" || parts[0] == "Tỉnh Thành Phố" {
			continue
		}
		regionId, _ := strconv.Atoi(parts[1])
		regionName := parts[0]
		districtId, _ := strconv.Atoi(parts[3])
		districtName := parts[2]
		wardId, _ := strconv.Atoi(parts[5])
		wardName := parts[4]
		region, ok := mapRegion[regionId]
		if !ok {
			region = &Region{
				Id:   regionId,
				Name: regionName,
			}
			mapRegion[regionId] = region
			data = append(data, region)
		}
		district, ok := mapDistrict[districtId]
		if !ok {
			district = &District{
				Id:   districtId,
				Name: districtName,
			}
			mapDistrict[districtId] = district
			region.Districts = append(region.Districts, district)
		}

		ward := &Ward{
			Id:   wardId,
			Name: wardName,
		}
		district.Wards = append(district.Wards, ward)
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return data, nil
}

func main() {
	file := "./data.csv"
	regions, _ := convertCSVToYAML(file)
	data, err := yaml.Marshal(&regions)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
