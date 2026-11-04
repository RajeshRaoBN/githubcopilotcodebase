package main

import (
	"reflect"
	"testing"
)

func TestFindDates(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "example",
			input: "I have a dentist appointment on 11/14/2023 and book club on 12-1-23",
			want:  []string{"11/14/2023", "12-1-23"},
		},
		{
			name: "supported formats",
			input: "05/02/24 05/02/2024 5/2/24 5/2/2024 " +
				"05-02-24 05-02-2024 5-2-24 5-2-2024",
			want: []string{
				"05/02/24", "05/02/2024", "5/2/24", "5/2/2024",
				"05-02-24", "05-02-2024", "5-2-24", "5-2-2024",
			},
		},
		{
			name:  "ignores invalid dates and numbers",
			input: "02/30/2024, 13-1-24, and 123",
			want:  []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := findDates(test.input); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findDates(%q) = %v, want %v", test.input, got, test.want)
			}
		})
	}
}
