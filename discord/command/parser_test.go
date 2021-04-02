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
	"github.com/golang/mock/gomock"

	"github.com/cgear-go/bot/discord/commandmock"
)

func TestParser__ReadInt(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	g := goblin.Goblin(t)

	g.Describe("Parser.ReadInt", func() {
		g.It("Should return an integer value when lexer returns a valid int token", func() {
			l := commandmock.NewMockLexer(mockCtrl)
			l.
				EXPECT().
				Next().
				Return("5", nil)

			n, err := (&parser{lexer: l}).ReadInt()
			g.Assert(err).IsNil()
			g.Assert(n).Equal(5)
		})

		g.It("Should return an error when lexer returns a invalid int token", func() {
			l := commandmock.NewMockLexer(mockCtrl)
			l.
				EXPECT().
				Next().
				Return("a string", nil)

			n, err := (&parser{lexer: l}).ReadInt()
			g.Assert(err).IsNotNil()
			g.Assert(n).IsZero()
		})

		g.It("Should return an error when lexer is done", func() {
			l := commandmock.NewMockLexer(mockCtrl)
			l.
				EXPECT().
				Next().
				Return("", io.EOF)

			n, err := (&parser{lexer: l}).ReadInt()
			g.Assert(err).Equal(io.EOF)
			g.Assert(n).IsZero()
		})
	})

	g.Describe("Parser.ReadString", func() {
		g.It("Should return a string value when lexer returns a valid string token", func() {
			l := commandmock.NewMockLexer(mockCtrl)
			l.
				EXPECT().
				Next().
				Return("str", nil)

			n, err := (&parser{lexer: l}).ReadString()
			g.Assert(err).IsNil()
			g.Assert(n).Equal("str")
		})

		g.It("Should return an error when lexer is done", func() {
			l := commandmock.NewMockLexer(mockCtrl)
			l.
				EXPECT().
				Next().
				Return("", io.EOF)

			n, err := (&parser{lexer: l}).ReadString()
			g.Assert(err).Equal(io.EOF)
			g.Assert(n).IsZero()
		})
	})

	g.Describe("Parser.ReadRest", func() {
		g.It("Should return a string value when lexer returns a valid string token", func() {
			l := commandmock.NewMockLexer(mockCtrl)
			gomock.InOrder(
				l.EXPECT().HasNext().Return(true),
				l.EXPECT().HasNext().Return(true),
				l.EXPECT().HasNext().Return(false))

			gomock.InOrder(
				l.EXPECT().Next().Return("some", nil),
				l.EXPECT().Next().Return("str", nil))

			n, err := (&parser{lexer: l}).ReadRest()
			g.Assert(err).IsNil()
			g.Assert(n).Equal("some str")
		})

		g.It("Should return an error when next token is an error", func() {
			l := commandmock.NewMockLexer(mockCtrl)
			gomock.InOrder(
				l.EXPECT().HasNext().Return(true))

			gomock.InOrder(
				l.EXPECT().Next().Return("", io.ErrUnexpectedEOF))

			n, err := (&parser{lexer: l}).ReadRest()
			g.Assert(err).Equal(io.ErrUnexpectedEOF)
			g.Assert(n).Equal("")
		})

		g.It("Should return an error when lexer is done", func() {
			l := commandmock.NewMockLexer(mockCtrl)
			l.
				EXPECT().
				HasNext().
				Return(false)

			n, err := (&parser{lexer: l}).ReadRest()
			g.Assert(err).Equal(io.EOF)
			g.Assert(n).IsZero()
		})
	})
}
