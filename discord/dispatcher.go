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
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cgear-go/bot/discord/client"
	"github.com/cgear-go/bot/discord/command"
	"github.com/cgear-go/bot/discord/reaction"
)

type Dispatcher interface {

	// Client returns the internal Discord client
	Client() client.Client

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
	// session holds discord connection
	session *discordgo.Session

	// client holds the client wrapper around discord session
	client client.Client

	// commands
	commands map[string]command.Command

	// reactions
	reactions map[string][]reaction.Reaction

	// closers holds close functions
	closers []func()
}

func (d *dispatcher) Client() client.Client {
	return d.client
}

func (d *dispatcher) AddCommand(command command.Command) Dispatcher {
	d.commands[command.Name()] = command
	return d
}

func (d *dispatcher) AddReaction(r reaction.Reaction) Dispatcher {
	emoji := r.Emoji()
	if _, ok := d.reactions[emoji]; !ok {
		d.reactions[emoji] = make([]reaction.Reaction, 0, 1)
	}
	d.reactions[emoji] = append(d.reactions[emoji], r)
	return d
}

func (d *dispatcher) commandEndIndex(content string) int {
	wsIndex := strings.Index(content, " ")
	crIndex := strings.Index(content, "\n")
	if wsIndex == -1 {
		wsIndex = len(content)
	}

	if crIndex == -1 {
		crIndex = len(content)
	}

	if wsIndex < crIndex {
		return wsIndex
	}
	return crIndex
}

func (d *dispatcher) executeCommand(message *discordgo.MessageCreate, content string) error {
	defer d.client.ChannelMessageDelete(message.ChannelID, message.ID)

	index := d.commandEndIndex(content)
	if index == 0 {
		return nil
	}

	cmd, ok := d.commands[strings.TrimSpace(content[:index])]
	if !ok {
		return nil
	}

	permissions, err := d.client.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if err != nil {
		return err
	}

	return cmd.Execute(d.client, command.Event{
		GuildID:         message.GuildID,
		UserID:          message.Author.ID,
		UserPermissions: permissions,
		ChannelID:       message.ChannelID,
		MessageID:       message.ID,
		Params:          strings.TrimSpace(content[index:]),
	})
}

func (d *dispatcher) reactionAdded(added *discordgo.MessageReactionAdd) error {
	reactions, ok := d.reactions[added.Emoji.Name]
	if !ok {
		return nil
	}

	permissions, err := d.client.UserChannelPermissions(added.UserID, added.ChannelID)
	if err != nil {
		return err
	}

	for _, r := range reactions {
		if err := r.Added(d.client, reaction.Event{
			GuildID:         added.GuildID,
			UserID:          added.UserID,
			UserPermissions: permissions,
			ChannelID:       added.ChannelID,
			MessageID:       added.MessageID,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (d *dispatcher) reactionRemoved(added *discordgo.MessageReactionRemove) error {
	reactions, ok := d.reactions[added.Emoji.Name]
	if !ok {
		return nil
	}

	permissions, err := d.client.UserChannelPermissions(added.UserID, added.ChannelID)
	if err != nil {
		return err
	}

	for _, r := range reactions {
		if err := r.Removed(d.client, reaction.Event{
			GuildID:         added.GuildID,
			UserID:          added.UserID,
			UserPermissions: permissions,
			ChannelID:       added.ChannelID,
			MessageID:       added.MessageID,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (d *dispatcher) Listen() {
	d.closers = append(d.closers, d.session.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		if message.Author.ID == session.State.User.ID {
			return
		}

		content := strings.TrimSpace(message.Content)
		if len(content) == 0 {
			return
		}

		if content[0] != '+' {
			return
		}

		if err := d.executeCommand(message, content[1:]); err != nil {
			log.Printf("Failed to execute command (%s): %v", content, err)
		}
	}))

	d.closers = append(d.closers, d.session.AddHandler(func(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
		if reaction.UserID == session.State.User.ID {
			return
		}

		if err := d.reactionAdded(reaction); err != nil {
			log.Printf("Failed to execute reaction added (%s): %v", reaction.Emoji.Name, err)
		}
	}))

	d.closers = append(d.closers, d.session.AddHandler(func(session *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
		if reaction.UserID == session.State.User.ID {
			return
		}

		if err := d.reactionRemoved(reaction); err != nil {
			log.Printf("Failed to execute reaction removed (%s): %v", reaction.Emoji.Name, err)
		}
	}))
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
		session:   connection,
		client:    &clientImpl{session: connection},
		commands:  make(map[string]command.Command),
		reactions: make(map[string][]reaction.Reaction),
		closers:   make([]func(), 0, 3),
	}, nil
}
