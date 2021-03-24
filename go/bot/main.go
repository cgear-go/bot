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
	"github.com/jonathanarnault/cgear-go/go/bot/command"
	"github.com/jonathanarnault/cgear-go/go/bot/discord"
	"github.com/jonathanarnault/cgear-go/go/raid"
)

func main() {
	bot, err := discord.NewBot(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Failed to start Discord bot: %v", err)
	}
	defer bot.Close()

	dispatcher := command.NewDispatcher(bot)
	dispatcher.AddCommand("hello").
		AddString("name").
		AddResolver(func(ctx context.Context, bot discord.Bot, args command.Arguments) error {
			_, err := bot.MessageCreate(
				ctx.Value(discord.ContextChannelID).(string),
				fmt.Sprintf("Hello, %s!", args.GetString("name")))
			return err
		})
	dispatcher.AddCommand("raid").
		AddInt("invites").
		AddString("time").
		AddRest("gym").
		AddResolver(func(ctx context.Context, bot discord.Bot, args command.Arguments) error {
			var hours, minutes int
			_, err := fmt.Sscanf(args.GetString("time"), "%dh%d", &hours, &minutes)
			if err != nil {
				return err
			}

			now := time.Now()
			start := time.Date(now.Year(), now.Month(), now.Day(), hours, minutes, 0, 0, time.Local)

			channel, err := bot.ChannelCreateWithPermissions(
				ctx.Value(discord.ContextGuildID).(string),
				os.Getenv("CATEGORY_ID"),
				fmt.Sprintf("raid-%s", start.Format("02-01-15h04")),
				[]*discordgo.PermissionOverwrite{})
			if err != nil {
				return err
			}

			log.Printf("Creating raid: %v", raid.Raid{
				ID:    channel,
				Gym:   args.GetString("gym"),
				Users: []string{},
				Start: start,
			})
			return nil
		})

	cancel := dispatcher.ListenMessages()
	fmt.Println("CGear Bot connected to Discord, press Ctrl + C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	cancel()
}
