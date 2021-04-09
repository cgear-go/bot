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
	"fmt"
	"strings"
	"time"

	"github.com/cgear-go/bot/discord"
	"github.com/cgear-go/bot/discord/client"
	"github.com/cgear-go/bot/discord/command"
)

func registerRaidCommand(dispatcher discord.Dispatcher, engine *engine, config map[string]Config) {
	dispatcher.AddCommand(
		command.NewCommandBuilder("raid").
			AddString("level").
			AddString("start").
			AddRest("gym").
			AddFilter(func(event command.Event) (bool, error) {
				guild, ok := config[event.GuildID]
				if !ok {
					return true, nil
				}
				return event.ChannelID != guild.RaidChannelID, nil
			}).
			Resolver(func(client client.Client, event command.Event, arguments command.Arguments) error {
				start := arguments.GetString("start")
				if strings.HasSuffix(start, "h") {
					start = start + "00"
				}

				var hours, minutes int
				_, err := fmt.Sscanf(start, "%dh%d", &hours, &minutes)
				if err != nil {
					return err
				}

				now := time.Now()
				return engine.createRaid(client, raidInfo{
					guildID:     event.GuildID,
					organizerID: event.UserID,
					level:       arguments.GetString("level"),
					gym:         arguments.GetString("gym"),
					start:       time.Date(now.Year(), now.Month(), now.Day(), hours, minutes, 0, 0, time.Local),
				})
			}).
			Build())
}
