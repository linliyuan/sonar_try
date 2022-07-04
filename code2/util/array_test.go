package util

import (
	"testing"
)

func TestInArrayFloat32(t *testing.T) {
	type args struct {
		needle   float32
		haystack []float32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试在数组内",
			args: args{
				needle:   1.2,
				haystack: []float32{1.2, 1.3},
			},
			want: true,
		},
		{
			name: "测试在数组外",
			args: args{
				needle:   1.1,
				haystack: []float32{1.2, 1.3},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := InArrayFloat32(tt.args.needle, tt.args.haystack)
			if res != tt.want {
				t.Errorf("InArrayFloat32 got = %v, want %v", res, tt.want)
			}
		})
	}
}
