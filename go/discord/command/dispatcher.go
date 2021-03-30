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
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jonathanarnault/cgear-go/go/discord"
)

// Dispatcher represents a command dispatcher
type Dispatcher interface {

	// AddCommand registers a command to the dispatcher
	// Warning: this method is not supposed to be thread safe
	AddCommand(name string) CommandBuilder

	// Execute a command
	// This method is supposed to be thread safe
	Execute(ctx context.Context, command string) error

	// ListenMessages creates a message listener on the bot, automatically cleaning messages that are not commands in the given
	// channels
	ListenMessages(channels ...string) func()
}

// dispatcher is an implmentation of `Dispatcher`
type dispatcher struct {

	// session holds the discord session
	session *discordgo.Session

	// commands holds the available commands
	commands map[string]CommandBuilder
}

func (d *dispatcher) AddCommand(name string) CommandBuilder {
	if _, ok := d.commands[name]; !ok {
		d.commands[name] = &commandBuilder{
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

	return cmd.execute(ctx, d.session, parser)
}

func (d *dispatcher) ListenMessages(channels ...string) func() {
	return d.session.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		if message.Author.ID == session.State.User.ID {
			return
		}

		content := strings.TrimSpace(message.Content)
		if len(content) == 0 {
			return
		}

		if content[0] != '+' {
			for _, channel := range channels {
				if channel == message.ChannelID {
					session.ChannelMessageDelete(message.ChannelID, message.ID)
				}
			}
			return
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, discord.ContextMessageKey, message)
		err := d.Execute(ctx, content[1:])
		if err != nil {
			log.Printf("Failed to execute command (%s): %v", content, err)
		}

		session.ChannelMessageDelete(message.ChannelID, message.ID)
	})
}

// NewDispatcher creates a `Dispatcher`
func NewDispatcher(session *discordgo.Session) Dispatcher {
	return &dispatcher{
		session:  session,
		commands: make(map[string]CommandBuilder),
	}
}
