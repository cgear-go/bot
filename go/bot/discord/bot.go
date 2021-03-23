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

//go:generate mockgen -destination mock/bot.go -package discordmock . Bot
package discord

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandListener func(user, channel, message, command string) error

// Bot represents a Discord bot
type Bot interface {

	// AddCommandListener registers a command listener to the bot
	AddCommandListener(listener CommandListener) func()

	// MessageCreate sends a message to the channel
	MessageCreate(channel, message string) (string, error)

	// MessageDelete deletes a message
	MessageDelete(channel, message string) error

	// Close discord connecton
	Close()
}

// bot is an implementation of `Bot`
type bot struct {
	session *discordgo.Session
}

func (b *bot) AddCommandListener(listener CommandListener) func() {
	return b.session.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
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
		err := listener(message.Author.ID, message.ChannelID, message.ID, content[1:])
		if err != nil {
			log.Printf("Failed to execute command (%s): %v", content, err)
		}

		session.ChannelMessageDelete(message.ChannelID, message.ID)
	})
}

func (b *bot) MessageCreate(channel, message string) (string, error) {
	msg, err := b.session.ChannelMessageSend(channel, message)
	if err != nil {
		return "", err
	}
	return msg.ID, nil
}

func (b *bot) MessageDelete(channel, message string) error {
	return b.session.ChannelMessageDelete(channel, message)
}

func (b *bot) Close() {
	b.session.Close()
}

// NewBot connects to a discord server as a bot
func NewBot(token string) (Bot, error) {
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

	return &bot{
		session: connection,
	}, nil
}
