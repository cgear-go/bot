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

package raid

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jonathanarnault/cgear-go/go/discord"
)

// Engine represents the raid engine
type Engine interface {

	// SubmitRaid submits a raid with the given users
	SubmitRaid(context.Context, Raid) (string, error)

	// EndRaid clears raid resources
	EndRaid(ctx context.Context) error

	// ListenReactionAdd creates a listener for added reactions for raid gates
	ListenReactions() func()
}

// engine is and implmentation of `Engine`
type engine struct {

	// lock used for concurrency
	lock *sync.Mutex

	// session holds the Discord Bot connection
	session discord.Session

	// raids holds the raids that are ongoing
	raids map[string]Raid

	// channelID holds the channel ID where raid commands are listened
	channelID string

	// categoryID holds the category where lobbies will be created
	categoryID string
}

func (e *engine) SubmitRaid(ctx context.Context, raid Raid) (string, error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	command := ctx.Value(discord.ContextMessageKey).(*discordgo.MessageCreate)

	channel, err := e.session.GuildChannelCreateComplex(
		command.GuildID,
		discordgo.GuildChannelCreateData{
			Name:                 fmt.Sprintf("raid-%s", raid.Start.Format("02-01-15h04")),
			ParentID:             e.categoryID,
			PermissionOverwrites: []*discordgo.PermissionOverwrite{},
		})
	if err != nil {
		return "", err
	}

	gate, err := e.session.ChannelMessageSend(command.ChannelID, raid.FormatGate())
	if err != nil {
		return "", err
	}

	if err := e.session.MessageReactionAdd(gate.ChannelID, gate.ID, "🙏"); err != nil {
		return "", err
	}

	if err := e.session.MessageReactionAdd(gate.ChannelID, gate.ID, "👍"); err != nil {
		return "", err
	}

	if err := e.session.ChannelPermissionSet(
		channel.ID, command.Author.ID, discordgo.PermissionOverwriteTypeMember,
		discordgo.PermissionViewChannel, 0); err != nil {

		return "", err
	}

	announce, err := e.session.ChannelMessageSend(channel.ID, raid.FormatAnnounce())
	if err != nil {
		return "", err
	}

	raid.Channel = channel
	raid.Gate = gate
	raid.Announce = announce

	log.Printf("Creating raid: %v", raid)
	e.raids[raid.Gate.ID] = raid
	return raid.Gate.ID, nil
}

func (e *engine) isOperator(user *discordgo.User, raid Raid) (bool, error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	permissions, err := e.session.UserChannelPermissions(user.ID, raid.Channel.ID)
	if err != nil {
		return false, err
	}

	if (permissions & discordgo.PermissionManageChannels) > 0 {
		return true, nil
	}

	return user.ID == raid.Operator.ID, nil
}

func (e *engine) endRaid(raid Raid) error {
	if _, err := e.session.ChannelDelete(raid.Channel.ID); err != nil {
		return err
	}

	if err := e.session.ChannelMessageDelete(raid.Gate.ChannelID, raid.Gate.ID); err != nil {
		return err
	}

	delete(e.raids, raid.Gate.ID)
	log.Printf("Ending raid: %v", raid)
	return nil
}

func (e *engine) EndRaid(ctx context.Context) error {
	command := ctx.Value(discord.ContextMessageKey).(*discordgo.MessageCreate)

	for _, r := range e.raids {
		if r.Channel.ID != command.ChannelID {
			continue
		}

		operator, err := e.isOperator(command.Author, r)
		if !operator || err != nil {
			return err
		}

		return e.endRaid(r)
	}

	return nil
}

func (e engine) ListenReactions() func() {
	addListenerCancel := e.session.AddHandler(func(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
		if reaction.UserID == session.State.User.ID {
			return
		}

		if reaction.ChannelID != e.channelID {
			return
		}

		if reaction.Emoji.Name != "🙏" && reaction.Emoji.Name != "👍" {
			session.MessageReactionRemove(reaction.ChannelID, reaction.MessageID, reaction.Emoji.Name, reaction.UserID)
			return
		}

		_, ok := e.raids[reaction.MessageID]
		if !ok {
			return
		}

		switch reaction.Emoji.ID {
		case "🙏":
			return
		case "👍":
			return
		}
	})

	return func() {
		addListenerCancel()
	}
}

// NewEngine creates a new raid engine
func NewEngine(session discord.Session, channelID, categoryID string) Engine {
	return &engine{
		lock:       &sync.Mutex{},
		session:    session,
		raids:      make(map[string]Raid),
		channelID:  channelID,
		categoryID: categoryID,
	}
}
