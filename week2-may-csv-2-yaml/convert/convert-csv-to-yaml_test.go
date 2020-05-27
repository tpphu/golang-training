package convert

import (
	"reflect"
	"testing"
)

func TestConvertCSV2Yaml(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name       string
		args       args
		wantCities []*City
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCities, err := ConvertCSV2Yaml(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertCSV2Yaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCities, tt.wantCities) {
				t.Errorf("ConvertCSV2Yaml() = %v, want %v", gotCities, tt.wantCities)
			}
		})
	}
}
