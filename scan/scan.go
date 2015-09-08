package scan

import (
	"github.com/rthornton128/calc/token"
	"unicode"
)

type Scanner struct {
	ch      rune
	offset  int
	roffset int
	src     string
	file    *token.File
}

func (s *Scanner) Init(file *token.File, src string) {
	s.file = file
	s.offset, s.roffset = 0, 0
	s.src = src
	s.file.AddLine(s.offset)
	s.next()
}

func (s *Scanner) Scan() (lit string, tok *token.Token, pos token.Pos) {
	s.skipWhitespace()
	if unicode.IsDigit(s.ch) {
		return s.scanNumber()
	}
	lit, pos = string(s.ch), s.file.Pos(s.offset)
	switch s.ch {
	case '(':
		tok = token.LPAREN
	case ')':
		tok = token.RPAREN
	case '+':
		tok = token.ADD
	case '-':
		tok = token.SUB
	case '*':
		tok = token.MUL
	case '/':
		tok = token.QUO
	case '%':
		tok = token.REM
	case ';':
		s.skipComment()
		s.next()
		return s.Scan()
	default:
		if s.offset >= len(s.src)-1 {
			tok = token.EOF
		} else {
			tok = token.ILLEGAL
		}
	}
	s.next()
	return
}

func (s *Scanner) next() {
	s.ch = rune(0)
	if s.roffset < len(s.src) {
		s.offset = s.roffset
		s.ch = rune(s.src[s.offset])
		if s.ch == '\n' {
			s.file.AddLine(s.offset)
		}
		s.roffset++
	}
}

func (s *Scanner) scanNumber() (string, token.Token, token.Pos) {
	start := s.offset
	for unicode.IsDigit(s.ch) {
		s.next()
	}
	offset := s.offset
	if s.ch == rune(0) {
		offset++
	}
	return s.src[start:offset], token.INTEGER, s.file.Pos(start)
}

func (s *Scanner) skipComment() {
	for s.ch != '\n' && s.offset < len(s.src)-1 {
		s.next()
	}
}

func (s *Scanner) skipWhitespace() {
	for unicode.IsSpace(s.ch) {
		s.next()
	}
}
