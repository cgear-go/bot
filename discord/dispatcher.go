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

package discord

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/cgear-go/bot/discord/command"
	"github.com/cgear-go/bot/discord/reaction"
)

type Dispatcher interface {

	// AddCommand registers a command for the dispatcher
	AddCommand(command command.Command) Dispatcher

	// AddReaction registers a reaction for the dispatcher
	AddReaction(reaction reaction.Reaction) Dispatcher

	// Listen for reactions and commands
	Listen()

	// Close dispatcher
	Close()
}

// dispatcher is an implmentation of `Dispatcher`
type dispatcher struct {
	// commands
	commands map[string]command.Command

	// reactions
	reactions map[string]reaction.Reaction

	// closers holds close functions
	closers []func()
}

func (d *dispatcher) AddCommand(command command.Command) Dispatcher {
	d.commands[command.Name()] = command
	return d
}

func (d *dispatcher) AddReaction(reaction reaction.Reaction) Dispatcher {
	d.reactions[reaction.Emoji()] = reaction
	return d
}

func (d *dispatcher) Listen() {

}

func (d *dispatcher) Close() {
	for _, closer := range d.closers {
		closer()
	}
}

func NewDispatcher(token string) (Dispatcher, error) {
	if token == "" {
		return nil, errors.New("invalid discord token")
	}

	connection, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, fmt.Errorf("failed to create Discord connection: %v", err)
	}

	if err := connection.Open(); err != nil {
		return nil, fmt.Errorf("failed to open Discord connection: %v", err)
	}

	return &dispatcher{
		commands:  make(map[string]command.Command),
		reactions: make(map[string]reaction.Reaction),
		closers:   make([]func(), 0, 3),
	}, nil
}
