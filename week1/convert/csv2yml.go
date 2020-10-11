package convert

import (
	"encoding/csv"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type City struct {
	ID        string
	Name      string
	Districts []District
}

type District struct {
	ID    string
	Name  string
	Wards []Ward
}

type Ward struct {
	ID   string
	Name string
}

// ConvertCSV2YML returns nil if success and not nil if failed.
func ConvertCSV2YML(csvFile string) error {
	f, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	arrCities := []City{}
	mapCities := map[string]bool{}
	mapDistricts := map[string]bool{}
	for i, record := range records {
		// Skip header
		if i < 1 {
			continue
		}
		cityName := record[0]
		cityID := record[1]
		districtName := record[2]
		districtID := record[3]
		if _, ok := mapDistricts[districtID]; !ok {
			district := districtID{
				ID:   cityID,
				Name: districtName,
			}
		}
		if _, ok := mapCities[cityID]; !ok {
			city := City{
				ID:   cityID,
				Name: cityName,
			}
			arrCities = append(arrCities, city)
			mapCities[cityID] = true
		}
	}
	data, err := yaml.Marshal(&arrCities)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("out.yml", data, 0o644)
	if err != nil {
		return err
	}
	return nil
}
