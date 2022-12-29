package main

import (
	"reflect"
	"testing"
)

func Test_computeEuler(t *testing.T) {
	tests := []struct {
		name    string
		want    EulerSolution
		wantErr bool
	}{
		{
			name:    "empty",
			want:    EulerSolution{133, 110, 84, 27, 144},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := computeEuler()
			if (err != nil) != tt.wantErr {
				t.Errorf("computeEuler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeEuler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeBeastEuler(t *testing.T) {
	tests := []struct {
		name    string
		want    EulerSolution
		wantErr bool
	}{
		{
			name:    "empty",
			want:    EulerSolution{133, 110, 84, 27, 144},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := computeBeastEuler()
			if (err != nil) != tt.wantErr {
				t.Errorf("computeBeastEuler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeBeastEuler() = %v, want %v", got, tt.want)
			}
		})
	}
}
