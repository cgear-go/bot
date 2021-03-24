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

//go:generate mockgen -destination lexer_mock_test.go -package command . Lexer
package command

import (
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
}

// lexer is an implementation of `Lexer`
type lexer struct {
	// command holds the command to tokenize
	command []rune

	// cursor holds the cursor position
	cursor int

	// length holds the command length
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

func (l *lexer) HasNext() bool {
	return l.cursor < l.length
}

func (l *lexer) Next() (string, error) {
	sb := strings.Builder{}
	for l.HasNext() {
		current := l.command[l.cursor]
		l.advanceCursor()

		if unicode.IsSpace(current) {
			break
		}
		sb.WriteRune(current)
	}

	if sb.Len() == 0 {
		if l.HasNext() {
			return l.Next()
		}
		return "", io.EOF
	}
	return sb.String(), nil
}
