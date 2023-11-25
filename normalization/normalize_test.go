package normalization

import (
	"reflect"
	"testing"
)

func Test_normalize(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   []string
	}{
		{
			name:   "Normlize",
			tokens: []string{"I", "AM", "the", "niGHt"},
			want:   []string{"night"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
