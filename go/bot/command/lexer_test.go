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
)

func TestNewLexer(t *testing.T) {
	lexer := NewLexer(`raid`).(*lexer)
	if lexer == nil {
		t.Fatalf(`Expected "Lexer", got: "%v"`, lexer)
	}

	if lexer.cursor != 0 {
		t.Fatalf(`Expected "0", got: "%v"`, lexer.cursor)
	}

	if lexer.length != 4 {
		t.Fatalf(`Expected "4", got: "%v"`, lexer.length)
	}
}

func TestLexer_HasNext(t *testing.T) {
	{
		lexer := &lexer{[]rune("raid"), 0, 4}
		if hasNext := lexer.HasNext(); !hasNext {
			t.Fatalf(`Expected "true", got: "%v"`, hasNext)
		}
	}

	{
		lexer := &lexer{[]rune("raid"), 4, 4}
		if hasNext := lexer.HasNext(); hasNext {
			t.Fatalf(`Expected "false", got: "%v"`, hasNext)
		}
	}
}

func TestLexer_Next(t *testing.T) {
	{
		lexer := &lexer{[]rune(`raid  5   11h20 "Parc Sainte Marie" "Test with 🙂"`), 0, 49}

		{
			expected := `raid`
			token, err := lexer.Next()
			if err != nil || token != expected {
				t.Fatalf(`Expected ("%v", "<nil>"), got: ("%v", "%v")`, expected, token, err)
			}
		}

		{
			expected := `5`
			token, err := lexer.Next()
			if err != nil || token != expected {
				t.Fatalf(`Expected ("%v", "<nil>"), got: ("%v", "%v")`, expected, token, err)
			}
		}

		{
			expected := `11h20`
			token, err := lexer.Next()
			if err != nil || token != expected {
				t.Fatalf(`Expected ("%v", "<nil>"), got: ("%v", "%v")`, expected, token, err)
			}
		}

		{
			expected := `Parc Sainte Marie`
			token, err := lexer.Next()
			if err != nil || token != expected {
				t.Fatalf(`Expected ("%v", "<nil>"), got: ("%v", "%v")`, expected, token, err)
			}
		}

		{
			expected := `Test with 🙂`
			token, err := lexer.Next()
			if err != nil || token != expected {
				t.Fatalf(`Expected ("%v", "<nil>"), got: ("%v", "%v")`, expected, token, err)
			}
		}

		{
			token, err := lexer.Next()
			if err != io.EOF || token != "" {
				t.Fatalf(`Expected ("", "%v"), got: ("%v", "%v")`, io.EOF, token, err)
			}
		}
	}

	{
		lexer := &lexer{[]rune(`"Parc Sainte Marie`), 0, 18}

		{
			token, err := lexer.Next()
			if err != io.ErrUnexpectedEOF || token != "" {
				t.Fatalf(`Expected ("", "%v"), got: ("%v", "%v")`, io.ErrUnexpectedEOF, token, err)
			}
		}
	}

	{
		lexer := &lexer{[]rune(`"Parc Sainte Marie"a`), 0, 20}

		{
			token, err := lexer.Next()
			if err != io.ErrUnexpectedEOF || token != "" {
				t.Fatalf(`Expected ("", "%v"), got: ("%v", "%v")`, io.ErrUnexpectedEOF, token, err)
			}
		}
	}
}

func TestLexer_ScanNext(t *testing.T) {
	{
		lexer := &lexer{[]rune(`5 11h20`), 0, 7}

		{
			var invitations uint16
			err := lexer.ScanNext("%d", &invitations)
			if invitations != 5 || err != nil {
				t.Fatalf(`Expected ("5", "%v"), got: ("%v", "%v")`, io.EOF, invitations, err)
			}
		}

		{
			var hour, minute uint8
			err := lexer.ScanNext("%dh%d", &hour, &minute)
			if hour != 11 || minute != 20 || err != nil {
				t.Fatalf(`Expected ("11", "20", "%v"), got: ("%v", "%v", "%v")`, io.EOF, hour, minute, err)
			}
		}

		{
			var invitations int16
			err := lexer.ScanNext("%d", &invitations)
			if err != io.EOF {
				t.Fatalf(`Expected ("0", "%v"), got: ("%v", "%v")`, io.EOF, invitations, err)
			}
		}
	}
}

func TestLexer_Rest(t *testing.T) {
	{
		lexer := &lexer{[]rune(`raid Parc Sainte Marie`), 0, 22}

		{
			expected := `raid`
			token, err := lexer.Next()
			if err != nil || token != expected {
				t.Fatalf(`Expected ("%v", "<nil>"), got: ("%v", "%v")`, expected, token, err)
			}
		}

		{
			expected := `Parc Sainte Marie`
			token, err := lexer.Rest()
			if err != nil || token != expected {
				t.Fatalf(`Expected ("%v", "<nil>"), got: ("%v", "%v")`, expected, token, err)
			}
		}

		{
			token, err := lexer.Rest()
			if err != io.EOF || token != "" {
				t.Fatalf(`Expected ("", "%v"), got: ("%v", "%v")`, io.EOF, token, err)
			}
		}
	}
}
