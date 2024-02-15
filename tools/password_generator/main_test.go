package main

import (
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		characters string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Generate 8-character length password", args: args{characters: "8"}, want: 8},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := parseInput(tc.args.characters); got != tc.want {
				t.Errorf("parseInput() = %v, want %v", got, tc.want)
			}
		})
	}
}

func Test_generatePassword(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePassword(tt.args.length); got != tt.want {
				t.Errorf("generatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
