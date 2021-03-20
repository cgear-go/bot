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

import "errors"

// Lexer tokenizes commands
type Lexer interface {

	// HasNext returns true while the lexer has more token to read.
	HasNext() bool

	// Next returns the next token as string
	Next() (string, error)
}

// lexer is an implementation of `Lexer`
type lexer struct {
}

func (l *lexer) HasNext() bool {
	return false
}

func (l *lexer) Next() (string, error) {
	return "", errors.New("Not Implemented")
}
