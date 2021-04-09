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

//go:generate mockgen -destination ../commandmock/command.go -package commandmock . Command
package command

import "github.com/cgear-go/bot/discord/client"

// Command represents a command
type Command interface {

	// Name returns the command name
	Name() (name string)

	// Execute the command
	Execute(client client.Client, event Event) (err error)
}

type command struct {
	// name holds the command name
	name string

	// parameters holds the command parameters
	parameters []parameter

	// filters holds the command filters
	filters []FilterFn

	// resolver corresponds to the function that will be executed for the command
	resolver CommandFn
}

func (c command) Name() (name string) {
	return c.name
}

func (c command) Execute(client client.Client, event Event) (err error) {
	for _, filter := range c.filters {
		if filter(event) {
			return nil
		}
	}

	runes := []rune(event.Params)
	parser := &parser{
		lexer: &lexer{
			command: runes,
			cursor:  0,
			length:  len(runes),
		},
	}
	arguments := &arguments{values: make(map[string]interface{})}

	for _, parameter := range c.parameters {
		var (
			value interface{}
			err   error
		)

		switch parameter.parameterType {
		case parameterTypeInt:
			value, err = parser.ReadInt()
		case parameterTypeString:
			value, err = parser.ReadString()
		default:
			value, err = parser.ReadRest()
		}

		if err != nil {
			return err
		}
		arguments.values[parameter.name] = value
	}

	return c.resolver(client, event, arguments)
}
