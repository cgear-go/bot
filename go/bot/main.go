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

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jonathanarnault/cgear-go/go/bot/raid"
	"github.com/jonathanarnault/cgear-go/go/discord"
	"github.com/jonathanarnault/cgear-go/go/discord/command"
)

func main() {
	session, err := discord.NewBot(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Failed to start Discord bot: %v", err)
	}
	defer session.Close()

	dispatcher := command.NewDispatcher(session)
	engine := raid.NewEngine(session)

	raidChannelId := os.Getenv("RAID_CHANNEL_ID")
	dispatcher.AddCommand("raid").
		AddString("level").
		AddString("time").
		AddInt("invites").
		AddRest("gym").
		AddResolver(func(ctx context.Context, bot *discordgo.Session, args command.Arguments) error {
			command := ctx.Value(discord.ContextMessageKey).(*discordgo.MessageCreate)

			if command.ChannelID != raidChannelId {
				log.Printf("Skipping command (%s != %s)", command.ChannelID, raidChannelId)
				return nil
			}

			startTime := args.GetString("time")

			var hours, minutes int
			_, err := fmt.Sscanf(startTime, "%dh%d", &hours, &minutes)
			if err != nil {
				return err
			}

			now := time.Now()

			raid := raid.Raid{
				Level:     args.GetString("level"),
				Gym:       args.GetString("gym"),
				Operator:  command.Author,
				Invites:   args.GetInt("invites"),
				Attendees: make([]*discordgo.User, 0),
				Start:     time.Date(now.Year(), now.Month(), now.Day(), hours, minutes, 0, 0, time.Local),
			}

			id, err := engine.SubmitRaid(ctx, raid)
			if err != nil {
				return err
			}

			go func() {
				time.Sleep(5 * time.Second)
				engine.EndRaid(ctx, id)
			}()

			return nil
		})

	cancelCommandListener := dispatcher.ListenMessages(raidChannelId)
	fmt.Println("CGear Bot connected to Discord, press Ctrl + C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	cancelCommandListener()
}
