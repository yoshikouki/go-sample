package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"case1", "Hello, world", "dlrow ,olleH"},
		{"case2", "Hello, 世界", "界世 ,olleH"},
		{"case3", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseRunes(tt.args); got != tt.want {
				t.Errorf("ReverseRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
