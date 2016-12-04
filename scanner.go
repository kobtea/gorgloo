package gorgloo

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// Scanner represents a lexical scanner
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Scan return the next token and literal value
func (s *Scanner) Scan() (tok Token, lit string) {
	if tok, lit := s.scanWhitespace(); tok == WS {
		return tok, lit
	}
	return s.scanLetter()
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	if ch := s.read(); !isWhitespace(ch) {
		s.unread()
		return ILLEGAL, ""
	}
	s.unread()
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

func (s *Scanner) scanLetter() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitespace(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	str := buf.String()
	// asterisk
	if strings.Count(str, "*") == len(str) {
		return ASTERISK, str
	}

	// state
	if isState(str) {
		return STATE, str
	}

	// priority
	if isPriority(str) {
		return PRIORITY, str
	}

	// tags
	if isTag(str) {
		return TAG, str
	}

	return HEADLINE, str
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func isAsterisk(ch rune) bool {
	return ch == '*'
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\n'
}

func isState(str string) bool {
	states := []string{"TODO", "DONE"}
	for _, state := range states {
		if str == state {
			return true
		}
	}
	return false
}

func isPriority(str string) bool {
	priorities := []string{"[#A]", "[#B]", "[#C]"}
	for _, priority := range priorities {
		if str == priority {
			return true
		}
	}
	return false
}

func isTag(str string) bool {
	if strings.HasPrefix(str, ":") && strings.HasSuffix(str, ":") {
		return true
	}
	return false
}

var eof = rune(0)
