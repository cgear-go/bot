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
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/jonathanarnault/cgear-go/go/discord"
)

var (
	raidCategoryId = os.Getenv("RAID_CATEGORY_ID")
)

// Engine represents the raid engine
type Engine interface {

	// SubmitRaid submits a raid with the given users
	SubmitRaid(context.Context, Raid) (string, error)

	// EndRaid clears raid resources
	EndRaid(ctx context.Context, id string) error
}

// engine is and implmentation of `Engine`
type engine struct {

	// session holds the Discord Bot connection
	session *discordgo.Session

	// raids holds the raids that are ongoing
	raids map[string]Raid
}

func (e *engine) SubmitRaid(ctx context.Context, raid Raid) (string, error) {
	command := ctx.Value(discord.ContextMessageKey).(*discordgo.MessageCreate)

	channel, err := e.session.GuildChannelCreateComplex(
		command.GuildID,
		discordgo.GuildChannelCreateData{
			Name:                 fmt.Sprintf("raid-%s", raid.Start.Format("02-01-15h04")),
			ParentID:             raidCategoryId,
			PermissionOverwrites: []*discordgo.PermissionOverwrite{},
		})
	if err != nil {
		return "", err
	}

	message, err := e.session.ChannelMessageSend(
		command.ChannelID,
		fmt.Sprintf(
			`%s

Pour participer au raid :
🙏 pour demander une invitation à distance ;
👍 pour participer.`,
			raid.String()))
	if err != nil {
		return "", err
	}

	if err := e.session.MessageReactionAdd(message.ChannelID, message.ID, "🙏"); err != nil {
		return "", err
	}

	if err := e.session.MessageReactionAdd(message.ChannelID, message.ID, "👍"); err != nil {
		return "", err
	}

	raid.Channel = channel
	raid.Message = message

	log.Printf("Creating raid: %v", raid)
	e.raids[raid.Message.ID] = raid
	return raid.Message.ID, nil
}

func (e *engine) EndRaid(_ context.Context, id string) error {
	raid, ok := e.raids[id]
	if !ok {
		return nil
	}

	if _, err := e.session.ChannelDelete(raid.Channel.ID); err != nil {
		return err
	}

	if err := e.session.ChannelMessageDelete(raid.Message.ChannelID, raid.Message.ID); err != nil {
		return err
	}

	delete(e.raids, id)
	log.Printf("Ending raid: %v", raid)
	return nil
}

// NewEngine creates a new raid engine
func NewEngine(session *discordgo.Session) Engine {
	return &engine{
		session: session,
		raids:   make(map[string]Raid),
	}
}
