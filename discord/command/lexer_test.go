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
	"io"
	"testing"

	"github.com/franela/goblin"
)

func TestLexer__min(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("lexer.min", func() {
		g.It("Should return the minimum value of the given integers", func() {
			lexer := &lexer{command: []rune{}, cursor: 0, length: 0}
			g.Assert(lexer.min(1, 2)).Eql(1)
			g.Assert(lexer.min(3, 1)).Eql(1)
			g.Assert(lexer.min(1, 1)).Eql(1)
		})
	})
}

func TestLexer__advanceCursor(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("lexer.advanceCursor", func() {
		g.It("Should advance cursor if the end of input has not been reached", func() {
			lexer := &lexer{command: []rune{}, cursor: 0, length: 1}
			lexer.advanceCursor()
			g.Assert(lexer.cursor).Eql(1)
		})

		g.It("Should set cursor to length if the end of input has been reached", func() {
			lexer := &lexer{command: []rune{}, cursor: 1, length: 1}
			lexer.advanceCursor()
			g.Assert(lexer.cursor).Eql(1)
		})
	})
}

func TestLexer__HasNext(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("lexer.HasNext", func() {
		g.It("Should return true if end of input has not been reached", func() {
			lexer := &lexer{command: []rune{}, cursor: 0, length: 1}
			g.Assert(lexer.HasNext()).IsTrue()
		})

		g.It("Should return false if the end of input has been reached", func() {
			lexer := &lexer{command: []rune{}, cursor: 1, length: 1}
			g.Assert(lexer.HasNext()).IsFalse()
		})
	})
}

func TestLexer__Next(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("lexer.Next", func() {
		g.It("Should return the next token", func() {
			lexer := &lexer{
				command: []rune{'&', 'é', '§', '^', 'c', '+', ' ', 'a', 'b', 'c', ' ', ' ', '\n', '😃'},
				cursor:  0,
				length:  14,
			}

			{
				token, err := lexer.Next()
				g.Assert(err).IsNil()
				g.Assert(token).Eql("&é§^c+")
			}

			{
				token, err := lexer.Next()
				g.Assert(err).IsNil()
				g.Assert(token).Eql("abc")
			}

			{
				token, err := lexer.Next()
				g.Assert(err).IsNil()
				g.Assert(token).Eql("😃")
			}
		})

		g.It("Should advance cursor", func() {
			lexer := &lexer{
				command: []rune{'&', 'é', '§', '^', 'c', '+', ' ', 'a', 'b', 'c', '\n', '😃'},
				cursor:  0,
				length:  12,
			}

			lexer.Next()
			g.Assert(lexer.cursor).Eql(7)
		})

		g.It("Should return EOF if the end of input has been reached", func() {
			lexer := &lexer{command: []rune{'😃'}, cursor: 1, length: 1}
			token, err := lexer.Next()
			g.Assert(err).Eql(io.EOF)
			g.Assert(token).IsZero()
		})
	})
}
