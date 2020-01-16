package gojson

import (
	"fmt"
	"testing"
)

func Test_Marshal(t *testing.T) {
	data := map[string]interface{}{
		"name":    "gojson",
		"age":     18,
		"success": true,
		"error":   nil,
		"methods": map[string]interface{}{
			"one":   "Marshal",
			"two":   "MustMarshal",
			"other": 123456,
		},
		"color": []string{"blue", "green", "yellow"},
	}
	res, err := Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}

func TestMarshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "json marshal",
			args:    args{v: map[string]string{"name": "zhcppy"}},
			want:    "{\n  \"name\": \"zhcppy\"\n}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustMarshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "json must marshal",
			args: args{v: map[string]string{"name": "zhcppy"}},
			want: "{\n  \"name\": \"zhcppy\"\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustMarshal(tt.args.v); got != tt.want {
				t.Errorf("MustMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "json data format",
			args:    args{data: []byte(MustMarshal(map[string]int{"age": 18}))},
			want:    "{\n  \"age\": 18\n}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Format(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}
