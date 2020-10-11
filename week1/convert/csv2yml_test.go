package convert

import "testing"

func TestConvertCSV2YML_HappyCase(t *testing.T) {
	csvFile := "../data.csv"
	err := ConvertCSV2YML(csvFile)
	if err != nil {
		t.Error("This should not error", err)
	}
}

func TestConvertCSV2YML_UnHappyCase_FileNotExist(t *testing.T) {
	csvFile := "../data_not_found.csv"
	err := ConvertCSV2YML(csvFile)
	if err == nil {
		t.Error("This should error")
	}
}

func TestConvertCSV2YML_UnHappyCase_FileExistButNotValidFormat(t *testing.T) {
	csvFile := "../image.png"
	err := ConvertCSV2YML(csvFile)
	if err == nil {
		t.Error("This should error")
	}
}
