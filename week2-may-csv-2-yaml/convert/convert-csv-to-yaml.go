package convert

import (
	"bufio"
	"fmt"
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
	Wards []*Ward
}

type Ward struct {
	ID   int
	Name string
}

type InvalidParserError interface {
	error
	GetTypeData() string
	GetRawData() []string
}

type invalidParserErrorImp struct {
	typeData string
	rawData  []string
}

func (e invalidParserErrorImp) Error() string {
	return fmt.Sprintf("Invalid parser for %s, raw data are: %s",
		e.typeData,
		strings.Join(e.rawData, ","))
}

func (e invalidParserErrorImp) GetTypeData() string {
	return e.typeData
}

func (e invalidParserErrorImp) GetRawData() []string {
	return e.rawData
}

func NewInvalidParserError(typeData string, data ...string) InvalidParserError {
	return invalidParserErrorImp{
		typeData: typeData,
		rawData:  data,
	}
}

var InvalidCSVData = func(data ...string) InvalidParserError {
	return invalidParserErrorImp{
		typeData: "line",
		rawData:  data,
	}
}
var InvalidParseIDCity = func(data ...string) InvalidParserError {
	return invalidParserErrorImp{
		typeData: "city",
		rawData:  data,
	}
}
var InvalidParseIDDistrict = func(data ...string) InvalidParserError {
	return invalidParserErrorImp{
		typeData: "district",
		rawData:  data,
	}
}
var InvalidParseIDWard = func(data ...string) InvalidParserError {
	return invalidParserErrorImp{
		typeData: "ward",
		rawData:  data,
	}
}

type CSVToYaml struct {
	citiMap     map[int]*City
	districtMap map[int]*District
	result      []*City
}

func NewCSVToYamlConverter() CSVToYaml {
	converter := CSVToYaml{
		citiMap:     map[int]*City{},
		districtMap: map[int]*District{},
	}
	return converter
}
func (c *CSVToYaml) ConvertCSV2Yaml(file string) ([]*City, error) {
	inFile, err := os.Open(file)
	if err != nil {
		return c.result, err
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
		err := c.process(line)
		if err != nil {
			return c.result, err
		}
	}
	return c.result, nil
}

func (c *CSVToYaml) process(line string) InvalidParserError {
	cells := strings.Split(line, ",")
	if len(cells) < 6 {
		return InvalidCSVData(cells...)
	}
	// Xu ly cai city
	city, err := c.processCity(cells[1], cells[0])
	if err != nil {
		return err
	}
	// Xu ly district
	district, err := c.processDistrict(city, cells[3], cells[2])
	if err != nil {
		return err
	}
	// Xu ly ward
	_, err = c.processWard(district, cells[5], cells[4])
	return err
}

func (c *CSVToYaml) processCity(id string, name string) (*City, InvalidParserError) {
	ID, err := strconv.Atoi(id) //int
	if err != nil {
		return nil, InvalidParseIDCity(id, name)
	}
	city, ok := c.citiMap[ID]
	if !ok {
		city = &City{
			Name: name,
			ID:   ID,
		}
		c.citiMap[ID] = city
		c.result = append(c.result, city)
	}
	return city, nil
}

func (c *CSVToYaml) processDistrict(city *City, id string, name string) (*District, InvalidParserError) {
	ID, err := strconv.Atoi(id) //int
	if err != nil {
		return nil, InvalidParseIDDistrict(id, name)
	}
	district, ok := c.districtMap[ID]
	if !ok {
		district = &District{
			Name: name,
			ID:   ID,
		}
		c.districtMap[ID] = district
		city.Districts = append(city.Districts, district)
	}
	return district, nil
}

func (c *CSVToYaml) processWard(district *District, id string, name string) (*Ward, InvalidParserError) {
	ID, err := strconv.Atoi(id) //int
	if err != nil {
		return nil, InvalidParseIDWard(id, name)
	}
	ward := &Ward{
		Name: name,
		ID:   ID,
	}
	district.Wards = append(district.Wards, ward)
	return ward, nil
}
