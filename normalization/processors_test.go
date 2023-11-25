package normalization

import (
	"reflect"
	"testing"
)

func TestToLowercase(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   []string
	}{
		{
			name:   "Lowercase",
			tokens: []string{"Quick", "Brown", "foX", "jUmping", "OVER", "the", "FENce"},
			want:   []string{"quick", "brown", "fox", "jumping", "over", "the", "fence"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowercase(tt.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterLowercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStopWords(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   []string
	}{
		{
			name:   "Stop words 1",
			tokens: []string{"i", "am", "the", "night"},
			want:   []string{"night"},
		},
		{
			name:   "Stop words 2",
			tokens: []string{"quick", "brown", "fox", "jumps", "over", "fence"},
			want:   []string{"quick", "brown", "fox", "jumps", "over", "fence"},
		},
		{
			name: "Stop words - all removed",
			tokens: []string{"a", "and", "be", "have", "i", "am",
				"in", "of", "that", "the", "to", "is",
				"it", "for", "not", "on", "with", "are",
				"he", "as", "you", "do", "at", "she",
				"this", "but", "his", "by", "from", "they"},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterStopWords(tt.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterStopWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stemmer(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   []string
	}{
		{
			name:   "first",
			tokens: []string{"quick", "brown", "fox", "jumps", "over", "fence"},
			want:   []string{"quick", "brown", "fox", "jump", "over", "fenc"},
		},
		{
			name:   "second",
			tokens: []string{"jumping", "swimming", "boxing", "dancing", "cars", "boxes"},
			want:   []string{"jump", "swim", "box", "danc", "car", "box"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stemmer(tt.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stemmerFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
