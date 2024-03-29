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
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cgear-go/bot/discord"
	"github.com/cgear-go/bot/raid"
)

func main() {
	addr, ok := os.LookupEnv("SERVER_ADDR")
	if !ok {
		log.Fatalf("Failed to retrieve server addr")
	}

	token, ok := os.LookupEnv("DISCORD_TOKEN")
	if !ok {
		log.Fatalf("Failed to retrieve discord token")
	}

	raidChannelID, ok := os.LookupEnv("RAID_CHANNEL_ID")
	if !ok {
		log.Fatalf("Failed to retrieve raid channel ID")
	}

	dispatcher, err := discord.NewDispatcher(token)
	if err != nil {
		log.Fatalf("Failed to connect to discord: %v", err)
	}

	raid.CreateEngine(dispatcher, map[string]raid.Config{
		"827454700743426069": {
			RaidChannelID:  "827457292605325323",
			RaidCategoryId: "827457054323114004",
		},
		"339731167630852097": {
			RaidChannelID:  "825316713515450368",
			RaidCategoryId: "825316949033877544",
		},
	})

	dispatcher.OnMessage(func(message *discord.Message) {
		if (message.UserPermissions & discord.PermissionManageChannels) > 0 {
			return
		}

		if message.ChannelID == raidChannelID {
			dispatcher.Client().ChannelMessageDelete(message.ChannelID, message.MessageID)
		}
	})

	dispatcher.Listen()
	defer dispatcher.Close()

	server := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprintln(w, "OK")
		}),
	}

	go server.ListenAndServe()

	log.Println("CGear Bot connected to Discord, press Ctrl + C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	server.Close()
}
