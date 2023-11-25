package tokenizer

import (
	"reflect"
	"testing"
)

func Test_tokenize(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		tokens []string
	}{
		{
			name:   "empty",
			text:   "",
			tokens: []string{},
		},
		{
			name:   "letter",
			text:   "a",
			tokens: []string{"a"},
		},
		{
			name:   "letter",
			text:   "quick brown fox",
			tokens: []string{"quick", "brown", "fox"},
		},
		{
			name:   "numbers",
			text:   "1111 222 312",
			tokens: []string{"1111", "222", "312"},
		},
		{
			name:   "mixed",
			text:   "call me 123-312-312",
			tokens: []string{"call", "me", "123", "312", "312"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tokenize(tt.text); !reflect.DeepEqual(got, tt.tokens) {
				t.Errorf("tokenize() = %v, want %v", got, tt.tokens)
			}
		})
	}
}
