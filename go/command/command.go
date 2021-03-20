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
	"context"
)

// CommandFn is the resolver function for a command
type CommandFn func(context.Context, Arguments) error

// CommandBuilder allows to build commands
type CommandBuilder interface {

	// AddInt adds an integer argument to the command
	AddInt(name string) CommandBuilder

	// AddString adds a string argument to the command
	AddString(name string) CommandBuilder

	// AddRest add the rest of the arguments concatenated with a whitespace to the command
	AddRest(name string) CommandBuilder

	// AddResolver sets the command resolver
	AddResolver(CommandFn)
}

// commandBuilder is an implementation of `CommandBuilder`
type commandBuilder struct {
}

func (d *commandBuilder) AddInt(name string) CommandBuilder {
	return nil
}

func (d *commandBuilder) AddString(name string) CommandBuilder {
	return nil
}

func (d *commandBuilder) AddRest(name string) CommandBuilder {
	return nil
}

func (d *commandBuilder) AddResolver(CommandFn) {

}
