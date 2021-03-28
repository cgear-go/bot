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

//go:generate mockgen -destination command_builder_mock_test.go -package command . Command
package command

import (
	"context"

	"github.com/jonathanarnault/cgear-go/go/discord"
	"github.com/jonathanarnault/cgear-go/go/discord/session"
)

// CommandFn is the resolver function for a command
type CommandFn func(context.Context, discord.Session, Arguments) error

// Command allows to build commands
type Command interface {

	// AddInt adds an integer argument to the command
	AddInt(name string) Command

	// AddString adds a string argument to the command
	AddString(name string) Command

	// AddRest add the rest of the arguments concatenated with a whitespace to the command
	AddRest(name string) Command

	// AddChannelFilter adds a channel filter for the command
	AddChannelFilter(filter session.ChannelFilter) Command

	// AddResolver sets the command resolver
	AddResolver(CommandFn)

	// execute command with the provided context and parser
	execute(context.Context, discord.Session, Parser) error
}

// command is an implementation of `Command`
type command struct {

	// parameters holds the command parameters
	parameters []parameter

	// resolver corresponds to the function that will be executed for the command
	resolver CommandFn
}

func (c *command) AddInt(name string) Command {
	c.parameters = append(c.parameters, parameter{name: name, tpe: parameterTypeInt})
	return c
}

func (c *command) AddString(name string) Command {
	c.parameters = append(c.parameters, parameter{name: name, tpe: parameterTypeString})
	return c
}

func (c *command) AddRest(name string) Command {
	c.parameters = append(c.parameters, parameter{name: name, tpe: parameterTypeRest})
	return c
}

func (c *command) AddChannelFilter(filter session.ChannelFilter) Command {
	return c
}

func (c *command) AddResolver(resolver CommandFn) {
	c.resolver = resolver
}

func (c command) execute(ctx context.Context, session discord.Session, parser Parser) error {
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