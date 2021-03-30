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

//go:generate mockgen -destination command_builder_mock_test.go -package command . CommandBuilder
package command

import (
	"context"

	"github.com/jonathanarnault/cgear-go/go/discord"
)

// FilterFn is a function used to filter commands
type FilterFn func(event *CommandEvent) (skip bool, err error)

// CommandFn is the resolver function for a command
type CommandFn func(context.Context, discord.Session, Arguments) error

// CommandBuilder allows to build commands
type CommandBuilder interface {

	// AddInt adds an integer argument to the command
	AddInt(name string) CommandBuilder

	// AddString adds a string argument to the command
	AddString(name string) CommandBuilder

	// AddRest add the rest of the arguments concatenated with a whitespace to the command
	AddRest(name string) CommandBuilder

	// AddFilter adds a channel filter for the command
	AddFilter(filter FilterFn) CommandBuilder

	// AddResolver sets the command resolver
	AddResolver(CommandFn)

	// execute command with the provided context and parser
	execute(context.Context, discord.Session, Parser) error
}

// commandBuilder is an implementation of `Command`
type commandBuilder struct {

	// parameters holds the command parameters
	parameters []parameter

	// filters holds the command filters
	filters []FilterFn

	// resolver corresponds to the function that will be executed for the command
	resolver CommandFn
}

func (c *commandBuilder) AddInt(name string) CommandBuilder {
	c.parameters = append(c.parameters, parameter{name: name, tpe: parameterTypeInt})
	return c
}

func (c *commandBuilder) AddString(name string) CommandBuilder {
	c.parameters = append(c.parameters, parameter{name: name, tpe: parameterTypeString})
	return c
}

func (c *commandBuilder) AddRest(name string) CommandBuilder {
	c.parameters = append(c.parameters, parameter{name: name, tpe: parameterTypeRest})
	return c
}

func (c *commandBuilder) AddFilter(filter FilterFn) CommandBuilder {
	c.filters = append(c.filters, filter)
	return c
}

func (c *commandBuilder) AddResolver(resolver CommandFn) {
	c.resolver = resolver
}

func (c commandBuilder) execute(ctx context.Context, session discord.Session, parser Parser) error {
	arguments := &arguments{values: make(map[string]interface{})}

	for _, parameter := range c.parameters {
		var (
			value interface{}
			err   error
		)

		switch parameter.tpe {
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

	return c.resolver(ctx, session, arguments)
}
