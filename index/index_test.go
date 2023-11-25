package index

import (
	"fts-exr/document"
	"reflect"
	"testing"
)

func Test_index_none(t *testing.T) {
	idx := make(index)
	doc := []document.Document{{"", "", "", 0}}
	idx.add(doc)

	if len(idx.search("a")) > 0 {
		t.Errorf("shoudln't find anything")
	}
	if len(idx.search("b")) > 0 {
		t.Errorf("shoudln't find anything")
	}
	if len(idx.search("c")) > 0 {
		t.Errorf("shoudln't find anything")
	}
}

func Test_index_signle(t *testing.T) {
	idx := make(index)
	doc := []document.Document{{"", "", "quick brown fox jumps over the fence", 0}}
	idx.add(doc)

	if len(idx.search("quick")) != 1 {
		t.Errorf("shoudln find")
	}
	if len(idx.search("jump")) != 1 {
		t.Errorf("shoudln't find anything")
	}
	if len(idx.search("jumps")) != 1 {
		t.Errorf("shoudln't find anything")
	}
	if len(idx.search("over")) != 1 {
		t.Errorf("shoudln't find anything")
	}
}

func Test_index_multiple(t *testing.T) {
	idx := make(index)
	doc := []document.Document{
		{"", "", "quick brown fox jumps over the fence", 0},
		{"", "", "history about old wise fox", 1},
		{"", "", "cats cant swim", 2},
	}
	idx.add(doc)

	exp := []int{1}
	if got := idx.search("jump"); !reflect.DeepEqual(got, []int{0}) {
		t.Errorf("search() = %v, want %v", got, exp)
	}
	exp = []int{1, 2}
	if got := idx.search("fox"); !reflect.DeepEqual(got, []int{0, 1}) {
		t.Errorf("search() = %v, want %v", got, exp)
	}
	exp = []int{1}
	if got := idx.search("cat"); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("search() = %v, want %v", got, exp)
	}
}
