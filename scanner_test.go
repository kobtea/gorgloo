package gorgloo

import (
	"strings"
	"testing"
)

func TestScannerScan(t *testing.T) {
	tests := []struct {
		s   string
		tok Token
		lit string
	}{
		{s: " ", tok: WS, lit: " "},
		{s: "  ", tok: WS, lit: "  "},
		{s: "*", tok: ASTERISK, lit: "*"},
		{s: "**", tok: ASTERISK, lit: "**"},
		{s: "TODO", tok: STATE, lit: "TODO"},
		{s: "[#A]", tok: PRIORITY, lit: "[#A]"},
		{s: "hoge", tok: HEADLINE, lit: "hoge"},
		{s: ":hoge:", tok: TAG, lit: ":hoge:"},
	}

	for i, tt := range tests {
		s := NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tok != tt.tok {
			t.Errorf("%d. %q token mistatch: exp=%q got=%q", i, tt.s, tt.tok, tok)
		} else if lit != tt.lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}

}
