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
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cgear-go/bot/discord"
	"github.com/cgear-go/bot/discord/client"
	"github.com/cgear-go/bot/discord/command"
)

func main() {
	token, ok := os.LookupEnv("DISCORD_TOKEN")
	if !ok {
		log.Fatalf("Failed to retrieve discord token")
	}

	dispatcher, err := discord.NewDispatcher(token)
	if err != nil {
		log.Fatalf("Failed to connect to discord: %v", err)
	}

	dispatcher.AddCommand(
		command.NewCommandBuilder("raid").
			AddString("level").
			AddString("start").
			AddString("gym").
			Resolver(func(client client.Client, event command.Event, arguments command.Arguments) error {
				log.Printf("%v", event)
				return nil
			}).
			Build())

	dispatcher.Listen()
	defer dispatcher.Close()

	log.Println("CGear Bot connected to Discord, press Ctrl + C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}