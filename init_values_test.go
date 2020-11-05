package main

import (
	"reflect"
	"testing"
)

func Test_readFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "existFile", args: args{"initDataTest.json"}, want: []byte{
			91, 13, 10, 32, 32, 49, 52, 44, 13, 10, 32, 32, 49, 48, 44, 13, 10, 32, 32, 50, 49, 13, 10, 93,
		},
			wantErr: false},
		{name: "errFile", args: args{"initData1.json"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInitValues(t *testing.T) {
	type args struct {
		values []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "correct_parse", args: args{values: []byte{
				91, 13, 10, 32, 32, 49, 52, 44, 13, 10, 32, 32, 49, 48, 44, 13, 10, 32, 32, 50, 49, 13, 10, 93,
			}},
			want:    []int{14, 10, 21},
			wantErr: false,
		},
		{
			name: "bad_json", args: args{values: []byte{
				91, 13, 10, 32, 32, 49, 52, 11, 13, 10, 32, 32, 49, 48, 44, 13, 10, 32, 32, 50, 49, 13, 10, 93,
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInitValues(tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInitValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInitValues() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInitValues(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "good_test", args: args{filePath: "initDataTest.json"}, want: []int{14, 10, 21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInitValues(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInitValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
