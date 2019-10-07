package main

import "testing"

func Test_translate(t *testing.T) {
	type args struct {
		input string
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
			if got := translate(tt.args.input); got != tt.want {
				t.Errorf("translate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_argswitch(t *testing.T) {
	t.Errorf("Error thrown")
	type args struct {
		input []string
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
			if got := argswitch(tt.args.input); got != tt.want {
				t.Errorf("argswitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_help(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			help()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
