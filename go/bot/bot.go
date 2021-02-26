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
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
)

func messageHandler(session *discord.Session, message *discord.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "+ping" {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}

func main() {

	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal("Missing discord token, please set DISCORD_TOKEN environment")
	}

	connection, err := discord.New(fmt.Sprintf("Bot %s", discordToken))
	if err != nil {
		log.Fatalf("Failed to create Discord connection: %v", err)
	}

	if err := connection.Open(); err != nil {
		log.Fatalf("Failed to open Discord connection: %v", err)
	}
	defer connection.Close()

	connection.AddHandler(messageHandler)
	connection.Identify.Intents = discord.IntentsGuildMessages

	fmt.Println("CGear Bot connected to Discord, press Ctrl + C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
