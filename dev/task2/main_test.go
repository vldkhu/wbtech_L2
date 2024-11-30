package main

import "testing"

func TestStringUnpack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Valid case with repetitions", args{"a4bc2d5e"}, "aaaabccddddde", false},
		{"Valid case without repetitions", args{"abcd"}, "abcd", false},
		{"Invalid case starting with digit", args{"45"}, "", true},
		{"Empty string", args{""}, "", false},
		{"Valid case with escape sequences", args{"qwe\\4\\5"}, "qwe45", false},
		{"Valid case with multiple repetitions", args{"qwe\\45"}, "qwe44444", false},
		{"Valid case with double escape", args{"qwe\\\\5"}, "qwe\\\\\\", false}, // Проверка на экранирование
		{"Valid case with escaped characters", args{"a\\b\\c"}, "abc", false},
		{"Invalid case with incomplete escape", args{"a\\"}, "", true},
		{"Valid case with single character", args{"a1"}, "a", false},
		{"Valid case with multiple single characters", args{"a1b1c1"}, "abc", false},
		{"Valid case with mixed escape and repetitions", args{"a\\2b3"}, "aabbb", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringUnpack(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringUnpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringUnpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
