package convert

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type City struct {
	ID        int
	Name      string
	Districts []*District
}

type District struct {
	ID    int
	Name  string
	Wards []Ward
}

type Ward struct {
	ID   int
	Name string
}

var InvalidCSVData = errors.New("Invalid CSV Data")
var InvalidParseIDCity = errors.New("Invalid Parse ID City")
var InvalidParseIDDistrict = errors.New("Invalid Parse ID District")
var InvalidParseIDWard = errors.New("Invalid Parse ID Ward")

var citiMap = map[int]*City{}
var districtMap = map[int]*District{}

func processCity(cities *[]*City, id string, name string) (*City, error) {
	ID, err := strconv.Atoi(id) //int
	if err != nil {
		return nil, InvalidParseIDCity
	}
	city, ok := citiMap[ID]
	if !ok {
		city = &City{
			Name: name,
			ID:   ID,
		}
		citiMap[ID] = city
		// Dong nay het suc quan trong
		// Can phai dc hieu ro rang
		*cities = append(*cities, city)
	}
	return city, nil
}

func processDistrict(city *City, id string, name string) (*District, error) {
	ID, err := strconv.Atoi(id) //int
	if err != nil {
		return nil, InvalidParseIDDistrict
	}
	district, ok := districtMap[ID]
	if !ok {
		district = &District{
			Name: name,
			ID:   ID,
		}
		districtMap[ID] = district
		city.Districts = append(city.Districts, district)
	}
	return district, nil
}

func processWard(district *District, id string, name string) error {
	ID, err := strconv.Atoi(id) //int
	if err != nil {
		return InvalidParseIDWard
	}
	ward := Ward{
		Name: name,
		ID:   ID,
	}
	district.Wards = append(district.Wards, ward)
	return nil
}

// Private
func ConvertCSV2Yaml(file string) (cities []*City, err error) {
	inFile, err := os.Open(file)
	if err != nil {
		return cities, err
	}
	// Hoc bua sau
	defer inFile.Close()
	//
	scanner := bufio.NewScanner(inFile)
	isFirstLine := false
	for scanner.Scan() {
		line := scanner.Text()
		if isFirstLine == false {
			isFirstLine = true
			continue
		}
		cells := strings.Split(line, ",")
		if len(cells) < 6 {
			return cities, InvalidCSVData
		}
		// Xu ly cai city
		city, err := processCity(&cities, cells[1], cells[0])
		if err != nil {
			return cities, err
		}
		// Xu ly district
		district, err := processDistrict(city, cells[3], cells[2])
		if err != nil {
			return cities, InvalidParseIDDistrict
		}
		// Xu ly ward
		err = processWard(district, cells[5], cells[4])
		if err != nil {
			return cities, InvalidParseIDWard
		}
	}
	return cities, nil
}
