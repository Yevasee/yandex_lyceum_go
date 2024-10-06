package main

import (
	"testing"
)

var cases = []struct {
	name  string
	value []byte
	want  int
	err   error
}{
	{"english", []byte("Hello"), 5, nil},
	{"russian", []byte("Привет"), 6, nil},
	{"japanese", []byte("こんにちは"), 5, nil},
	{"invalid", []byte{0xFF, 0xFE, 0xFD}, 0, ErrInvalidUTF8},
}

func TestGetUTFLength(t *testing.T) {
	for _, tc := range cases {
		tc := tc
		// запуск отдельного теста
		t.Run(tc.name, func(t *testing.T) {
			// тестируем функцию Sum
			got, err := GetUTFLength(tc.value)
			// проверим полученное значение
			if got != tc.want || err != tc.err {
				t.Errorf("GetUTFLength(%v) = (%v, %v); want (%v, %v)", tc.value, got, err, tc.want, tc.err)
			}
		})
	}
}
