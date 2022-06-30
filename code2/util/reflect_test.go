package util

import (
	"reflect"
	"testing"
)

type ZeroT struct {
	A int
}

func (z *ZeroT) IsZero() bool {
	return z.A == 0
}

// TestIsEmptyValue 测试不同类型的检测为空
func TestIsEmptyValue(t *testing.T) {

	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Array test not empty",
			args: args{
				v: []interface{}{"213", "234"},
			},
			want: false,
		},
		{
			name: "Array test empty",
			args: args{
				v: []string{},
			},
			want: true,
		},
		{
			name: "Bool test not empty",
			args: args{
				v: true,
			},
			want: false,
		},
		{
			name: "Bool test empty",
			args: args{
				v: false,
			},
			want: true,
		},
		{
			name: "Int test not empty",
			args: args{
				v: 1,
			},
			want: false,
		},
		{
			name: "Int test empty",
			args: args{
				v: 0,
			},
			want: true,
		},
		{
			name: "UInt test not empty",
			args: args{
				v: uint(20),
			},
			want: false,
		},
		{
			name: "UInt test empty",
			args: args{
				v: uint(0),
			},
			want: true,
		},
		{
			name: "Ptr test not empty",
			args: args{
				v: new(args),
			},
			want: false,
		},
		{
			name: "Ptr test empty",
			args: args{
				v: (*int)(nil),
			},
			want: true,
		},
		{
			name: "TypeZero test not empty",
			args: args{
				v: &ZeroT{
					A: 20,
				},
			},
			want: false,
		},
		{
			name: "TypeZero test empty",
			args: args{
				v: &ZeroT{
					A: 0,
				},
			},
			want: true,
		},
		{
			name: "Float test not empty",
			args: args{
				v: 100.1,
			},
			want: false,
		},
		{
			name: "Float test empty",
			args: args{
				v: 0.0,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isEmptyValue(reflect.ValueOf(tt.args.v))
			if res != tt.want {
				t.Errorf("isEmptyValue got = %v, want %v", res, tt.want)
			}
		})
	}

}
