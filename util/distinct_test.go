package util

import (
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	type args struct {
		tempItem []string
	}
	tests := []struct {
		name       string
		args       args
		wantNewArr []string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewArr := Distinct(tt.args.tempItem); !reflect.DeepEqual(gotNewArr, tt.wantNewArr) {
				t.Errorf("Distinct() = %v, want %v", gotNewArr, tt.wantNewArr)
			}
		})
	}
}
