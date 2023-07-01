package vin

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		vin string
	}
	tests := []struct {
		name    string
		args    args
		want    *Vin
		wantErr bool
	}{
		{
			name: "valid vin",
			args: args{
				vin: "1M8GDM9AXKP042788",
			},
			want: &Vin{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.vin)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
