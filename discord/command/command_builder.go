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
	"github.com/cgear-go/bot/discord/client"
)

// FilterFn is a function used to filter commands
type FilterFn func(event Event) (skip bool, err error)

// CommandFn is the resolver function for a command
type CommandFn func(client client.Client, event Event, arguments Arguments) error

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

	// Resolver sets the command resolver
	Resolver(CommandFn) CommandBuilder

	// Build the command
	Build() Command
}

// commandBuilder is an implementation of `CommandBuilder`
type commandBuilder struct {

	// name holds the command name
	name string

	// parameters holds the command parameters
	parameters []parameter

	// filters holds the command filters
	filters []FilterFn

	// resolver corresponds to the function that will be executed for the command
	resolver CommandFn
}

func (c *commandBuilder) AddInt(name string) CommandBuilder {
	c.parameters = append(c.parameters, parameter{name: name, parameterType: parameterTypeInt})
	return c
}

func (c *commandBuilder) AddString(name string) CommandBuilder {
	c.parameters = append(c.parameters, parameter{name: name, parameterType: parameterTypeString})
	return c
}

func (c *commandBuilder) AddRest(name string) CommandBuilder {
	c.parameters = append(c.parameters, parameter{name: name, parameterType: parameterTypeRest})
	return c
}

func (c *commandBuilder) AddFilter(filter FilterFn) CommandBuilder {
	c.filters = append(c.filters, filter)
	return c
}

func (c *commandBuilder) Resolver(resolver CommandFn) CommandBuilder {
	c.resolver = resolver
	return c
}

func (c *commandBuilder) Build() Command {
	return nil
}

// NewCommandBuilder creates a command builder
func NewCommandBuilder(name string) CommandBuilder {
	return &commandBuilder{
		name:       name,
		parameters: make([]parameter, 0, 1),
		filters:    make([]FilterFn, 0),
		resolver:   nil,
	}
}
