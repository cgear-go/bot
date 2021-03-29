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

//go:generate mockgen -destination mock/dispatcher_mock.go -package mock . Dispatcher
package reaction

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Dispatcher represents a command dispatcher
type Dispatcher interface {

	// AddCommand registers a command to the dispatcher
	// Warning: this method is not supposed to be thread safe
	AddReaction(reactions ...string) ReactionBuilder

	// ListenReactions creates a reaction listener on the bot
	Listen(channels ...string)

	// Close closes dispatcher listeners
	Close()
}

// dispatcher is an implmentation of `Dispatcher`
type dispatcher struct {

	// session holds the discord session
	session *discordgo.Session

	// reactions
	reactions []*reactionBuilder

	// cancellers holds the listener cancel functions
	cancellers []func()
}

func (d *dispatcher) AddReaction(reactions ...string) ReactionBuilder {
	builder := &reactionBuilder{
		reactions:       reactions,
		filters:         make([]FilterFn, 0),
		reactionAdded:   nil,
		reactionRemoved: nil,
	}
	d.reactions = append(d.reactions, builder)
	return builder
}

func (d *dispatcher) shouldSkip(reaction *reactionBuilder, event *Reaction) (skip bool, err error) {
	if len(reaction.reactions) > 0 {
		found := false
		for _, name := range reaction.reactions {
			found = found || name == event.EmojiID
		}

		if !found {
			return true, nil
		}
	}

	for _, filter := range reaction.filters {
		skip, err := filter(event)
		if err != nil {
			return false, err
		}

		if skip {
			return true, nil
		}
	}

	return false, nil
}

func (d *dispatcher) Listen(channels ...string) {
	d.cancellers = append(d.cancellers,
		d.session.AddHandler(func(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
			if reaction.UserID == session.State.User.ID {
				return
			}

			permissions, err := d.session.UserChannelPermissions(reaction.UserID, reaction.ChannelID)
			if err != nil {
				log.Printf("Failed to execute fetch user permissions for channel (%s): %v", reaction.ChannelID, err)
				return
			}

			event := &Reaction{
				UserID:          reaction.UserID,
				UserPermissions: permissions,
				GuildID:         reaction.GuildID,
				ChannelID:       reaction.ChannelID,
				MessageID:       reaction.MessageID,
				EmojiID:         reaction.Emoji.Name,
			}

			for _, reaction := range d.reactions {
				if reaction.reactionAdded == nil {
					continue
				}

				skip, err := d.shouldSkip(reaction, event)
				if err != nil {
					log.Printf("Failed to execute reaction filter: %v", err)
					continue
				}

				if skip {
					continue
				}

				if err := reaction.reactionAdded(d.session, event); err != nil {
					log.Printf("Failed to execute reaction: %v", err)
				}
			}
		}))

	d.cancellers = append(d.cancellers,
		d.session.AddHandler(func(session *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
			if reaction.UserID == session.State.User.ID {
				return
			}

			permissions, err := d.session.UserChannelPermissions(reaction.UserID, reaction.ChannelID)
			if err != nil {
				log.Printf("Failed to execute fetch user permissions for channel (%s): %v", reaction.ChannelID, err)
				return
			}

			event := &Reaction{
				UserID:          reaction.UserID,
				UserPermissions: permissions,
				GuildID:         reaction.GuildID,
				ChannelID:       reaction.ChannelID,
				MessageID:       reaction.MessageID,
				EmojiID:         reaction.Emoji.Name,
			}

			for _, reaction := range d.reactions {
				if reaction.reactionRemoved == nil {
					continue
				}

				skip, err := d.shouldSkip(reaction, event)
				if err != nil {
					log.Printf("Failed to execute reaction filter: %v", err)
					continue
				}

				if skip {
					continue
				}

				if err := reaction.reactionRemoved(d.session, event); err != nil {
					log.Printf("Failed to execute reaction: %v", err)
				}
			}
		}))
}

// Close closes dispatcher listeners
func (d *dispatcher) Close() {
	for _, canceller := range d.cancellers {
		canceller()
	}
}

// NewDispatcher creates a `Dispatcher`
func NewDispatcher(session *discordgo.Session) Dispatcher {
	return &dispatcher{
		session:    session,
		reactions:  make([]*reactionBuilder, 0),
		cancellers: make([]func(), 0, 2),
	}
}
