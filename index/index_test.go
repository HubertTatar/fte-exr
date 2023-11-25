package index

import (
	"fts-exr/document"
	"reflect"
	"testing"
)

func Test_index_none(t *testing.T) {
	idx := make(Index)
	doc := []document.Document{{"", "", "", 0}}
	idx.Add(doc)

	if len(idx.Search("a")) > 0 {
		t.Errorf("shoudln't find anything")
	}
	if len(idx.Search("b")) > 0 {
		t.Errorf("shoudln't find anything")
	}
	if len(idx.Search("c")) > 0 {
		t.Errorf("shoudln't find anything")
	}
}

func Test_index_signle(t *testing.T) {
	idx := make(Index)
	doc := []document.Document{{"", "", "quick brown fox jumps over the fence", 0}}
	idx.Add(doc)

	if len(idx.Search("quick")) != 1 {
		t.Errorf("shoudln find")
	}
	if len(idx.Search("jump")) != 1 {
		t.Errorf("shoudln't find anything")
	}
	if len(idx.Search("jumps")) != 1 {
		t.Errorf("shoudln't find anything")
	}
	if len(idx.Search("over")) != 1 {
		t.Errorf("shoudln't find anything")
	}
}

func Test_index_multiple(t *testing.T) {
	idx := make(Index)
	doc := []document.Document{
		{"", "", "quick brown fox jumps over the fence", 0},
		{"", "", "history about old wise fox", 1},
		{"", "", "cats cant swim", 2},
	}
	idx.Add(doc)

	exp := []int{1}
	if got := idx.Search("jump"); !reflect.DeepEqual(got, []int{0}) {
		t.Errorf("Search() = %v, want %v", got, exp)
	}
	exp = []int{1, 2}
	if got := idx.Search("fox"); !reflect.DeepEqual(got, []int{0, 1}) {
		t.Errorf("Search() = %v, want %v", got, exp)
	}
	exp = []int{1}
	if got := idx.Search("cat"); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("Search() = %v, want %v", got, exp)
	}
}
