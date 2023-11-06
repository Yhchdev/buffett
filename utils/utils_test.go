package utils

import "testing"

func Test_convertToBillions(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.

		{
			"",
			args{4761082555.91},
			47.61,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBillions(tt.args.n); got != tt.want {
				t.Errorf("convertToBillions() = %v, want %v", got, tt.want)
			}
		})
	}
}
