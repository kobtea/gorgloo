package gorgloo

import (
	"errors"
	"io"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (*Node, error) {
	if tok, _ := p.scanIgnoreWhitespace(); tok == ASTERISK {
		node := &Node{}

		// need ws scan
		if tok, lit := p.scanIgnoreWhitespace(); tok == STATE {
			node.State = lit
		} else {
			p.unscan()
		}

		if tok, lit := p.scanIgnoreWhitespace(); tok == PRIORITY {
			s, err := String2Priority(lit)
			if err != nil {
				return nil, err
			}
			node.Priority = s
		} else {
			p.unscan()
		}

		// need ws scan
		if tok, lit := p.scanIgnoreWhitespace(); tok == HEADLINE {
			node.Headline = lit
		} else {
			p.unscan()
		}
		// for {
		// 	tok, lit := p.scan()
		// 	if tok == HEADLINE || (tok == WS && lit != "\n") {
		// 		node.Headline += lit
		// 		continue
		// 	}
		// 	p.unscan()
		// 	break
		// }

		if tok, lit := p.scanIgnoreWhitespace(); tok == TAG {
			node.Tag = lit
		} else {
			p.unscan()
		}

		return node, nil
	}

	return nil, errors.New("hoge")
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}
	tok, lit = p.s.Scan()
	p.buf.tok, p.buf.lit = tok, lit
	return
}

func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) unscan() {
	p.buf.n = 1
}
