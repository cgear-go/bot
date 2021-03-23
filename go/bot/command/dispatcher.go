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
	"errors"

	"github.com/jonathanarnault/cgear-go/go/bot/discord"
)

// Dispatcher represents a command dispatcher
type Dispatcher interface {

	// AddCommand registers a command to the dispatcher
	// Warning: this method is not supposed to be thread safe
	AddCommand(name string) Command

	// Execute a command
	// This method is supposed to be thread safe
	Execute(ctx context.Context, command string) error
}

// dispatcher is an implmentation of `Dispatcher`
type dispatcher struct {

	// bot holds the discord bot
	bot discord.Bot

	// commands holds the available commands
	commands map[string]Command
}

func (d *dispatcher) AddCommand(name string) Command {
	if _, ok := d.commands[name]; !ok {
		d.commands[name] = &command{
			parameters: make([]parameter, 0, 8),
			resolver:   nil,
		}
	}
	return d.commands[name]
}

func (d *dispatcher) Execute(ctx context.Context, command string) error {
	runes := []rune(command)
	parser := &parser{
		lexer: &lexer{
			command: runes,
			cursor:  0,
			length:  len(runes),
		},
	}

	name, err := parser.ReadString()
	if err != nil {
		return err
	}

	cmd, ok := d.commands[name]
	if !ok {
		return errors.New("unexiting command")
	}

	return cmd.execute(ctx, d.bot, parser)
}

// NewDispatcher creates a `Dispatcher`
func NewDispatcher(bot discord.Bot) Dispatcher {
	return &dispatcher{
		bot:      bot,
		commands: make(map[string]Command),
	}
}
