package test

import "testing"

func TestHelloGo(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelloGo(); got != tt.want {
				t.Errorf("HelloGo() = %v, want %v", got, tt.want)
			}
		})
	}
}
