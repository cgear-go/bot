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

//go:generate mockgen -destination ../commandmock/parser.go -package commandmock . Parser
package command

import (
	"io"
	"strconv"
	"strings"
)

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

func (p *parser) ReadInt() (int, error) {
	token, err := p.lexer.Next()
	if err != nil {
		return 0, err
	}

	n, err := strconv.Atoi(token)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (p *parser) ReadString() (string, error) {
	token, err := p.lexer.Next()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (p *parser) ReadRest() (string, error) {
	sb := strings.Builder{}

	i := 0
	for p.lexer.HasNext() {
		str, err := p.lexer.Next()
		if err != nil {
			return "", err
		}

		if i > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(str)
		i++
	}

	if i == 0 {
		return "", io.EOF
	}
	return sb.String(), nil
}
