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

// Parser represents a command parser
type Parser interface {

	// ReadInt reads an int from the input command
	ReadInt() (int, error)

	// ReadString reads a string from the input
	ReadString() (string, error)

	// ReadRest reads the rest of the input command and returns a string that concatenates remaining tokens using space
	ReadRest() (string, error)
}

// parser is an implementation of `Parser`
type parser struct {
	lexer Lexer
}

func (parser *parser) ReadInt() (int, error) {
	return 0, errors.New("Not Implemented")
}

func (parser *parser) ReadString() (string, error) {
	return "", errors.New("Not Implemented")
}

func (parser *parser) ReadRest() (string, error) {
	return "", errors.New("Not Implemented")
}
