package main

import (
	"reflect"
	"testing"
)

func TestRemoveUser(t *testing.T) {
	type args struct {
		u User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateUser(t *testing.T) {
	tests := []struct {
		name string
		want User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
