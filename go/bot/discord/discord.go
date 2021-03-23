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
)

// Bot represents a Discord bot
type Bot interface {

	// MessageCreate sends a message to the channel
	MessageCreate(channel int64, message string) error

	// Close discord connecton
	Close()
}

// bot is an implementation of `Bot`
type bot struct {
	session *discordgo.Session
}

func (b bot) MessageCreate(channel int64, message string) error {
	return errors.New("not implmented")
}

func (b *bot) Close() {
	b.session.Close()
}

// Connect to a discord server
func Connect(token string) (Bot, error) {
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
