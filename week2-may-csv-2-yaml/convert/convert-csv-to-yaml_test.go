package convert

import (
	"reflect"
	"testing"
)

func TestCSVToYaml_processWard(t *testing.T) {
	type fields struct {
		citiMap     map[int]*City
		districtMap map[int]*District
		result      []*City
	}
	type args struct {
		district *District
		id       string
		name     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Ward
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSVToYaml{
				citiMap:     tt.fields.citiMap,
				districtMap: tt.fields.districtMap,
				result:      tt.fields.result,
			}
			got, err := c.processWard(tt.args.district, tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CSVToYaml.processWard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CSVToYaml.processWard() = %v, want %v", got, tt.want)
			}
		})
	}
}
