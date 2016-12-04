package gorgloo

import (
	"reflect"
	"strings"
	"testing"
)

func TestParserParse(t *testing.T) {
	tests := []struct {
		s    string
		node *Node
		err  string
	}{
		{
			s:    "* hoge",
			node: &Node{Headline: "hoge"},
		},
		{
			s:    "* TODO something",
			node: &Node{State: "TODO", Headline: "something"},
		},
		{
			s:    "* TODO [#A] something",
			node: &Node{State: "TODO", Priority: A, Headline: "something"},
		},
		{
			s:    "* TODO [#A] something :tag:",
			node: &Node{State: "TODO", Priority: A, Headline: "something", Tag: ":tag:"},
		},
		// {
		// 	s:    "** DONE something todo",
		// 	node: &Node{State: "DONE", Headline: "something todo"},
		// },
	}

	for i, tt := range tests {
		node, err := NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(errstring(err), tt.err) {
			t.Errorf("%d. %q error mismatch:\n exp=%s\n got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(node, tt.node) {
			t.Errorf("%d. %q node mistatch:\n exp=%#v\n got=%#v\n\n", i, tt.s, tt.node, node)
		}

	}
}

func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
