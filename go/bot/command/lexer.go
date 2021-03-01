//   Copyright 2020 Pokémon GO Nancy
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package command

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

// Lexer tokenizes commands
type Lexer interface {

	// HasNext returns true while the lexer has more token to read.
	HasNext() bool

	// Next returns the next token as string
	Next() (string, error)

	// ScanNext reads the next token and scans then token user `fmt.Sscanf`
	ScanNext(format string, args ...interface{}) error

	// Rest reads all tokens that have not been read and returns them as a space-separated string.
	Rest() (string, error)
}

// lexer in an implementation of `Lexer`
type lexer struct {
	// The command to tokenize
	command []rune

	// The current cursor position
	cursor int

	// The command length
	length int
}

func (l *lexer) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (l *lexer) advanceCursor() {
	l.cursor = l.min(l.cursor+1, l.length)
}

func (l *lexer) isCurrentQuote() bool {
	return l.command[l.cursor] == '"'
}

func (l *lexer) isCurrentSpace() bool {
	return unicode.IsSpace(l.command[l.cursor])
}

func (l *lexer) HasNext() bool {
	return l.cursor < l.length
}

func (l *lexer) Next() (string, error) {
	var (
		sb       strings.Builder
		inString bool = false
	)

	if !l.HasNext() {
		return "", io.EOF
	}

	for {
		readString := false
		if !l.HasNext() {
			l.advanceCursor()
			if inString {
				return "", io.ErrUnexpectedEOF
			}

			str := sb.String()
			if str == "" {
				return "", io.EOF
			}
			return str, nil
		}

		if l.isCurrentQuote() {
			l.advanceCursor()
			readString = inString
			inString = !inString

			if !l.HasNext() {
				return sb.String(), nil
			}
		}

		if l.isCurrentSpace() {
			l.advanceCursor()

			if inString {
				sb.WriteRune(' ')
			} else {
				str := sb.String()
				if str == "" {
					continue
				}
				return str, nil
			}
		} else if readString {
			return "", io.ErrUnexpectedEOF
		}

		sb.WriteRune(l.command[l.cursor])
		l.advanceCursor()
	}

}

func (l *lexer) ScanNext(format string, args ...interface{}) error {
	token, err := l.Next()
	if err != nil {
		return err
	}

	_, err = fmt.Sscanf(token, format, args...)
	return err
}

func (l *lexer) Rest() (string, error) {
	var sb strings.Builder

	for {
		token, err := l.Next()

		if err == io.EOF {
			str := sb.String()
			if str == "" {
				return "", io.EOF
			}

			return str[:len(str)-1], nil
		} else if err != nil {
			return "", err
		}

		sb.WriteString(token)
		sb.WriteRune(' ')
	}
}

// NewLexer returns a lexer for the command.
func NewLexer(command string) Lexer {
	runes := []rune(command)
	return &lexer{
		command: runes,
		cursor:  0,
		length:  len(runes),
	}
}
