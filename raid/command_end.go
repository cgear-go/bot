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
	"github.com/cgear-go/bot/discord"
	"github.com/cgear-go/bot/discord/client"
	"github.com/cgear-go/bot/discord/command"
)

func registerEndCommand(dispatcher discord.Dispatcher, engine *engine, config map[string]Config) {
	dispatcher.AddCommand(
		command.NewCommandBuilder("fin").
			AddFilter(func(event command.Event) (bool, error) {
				lobby := engine.getLobbyByChannel(event.ChannelID)
				return lobby == nil, nil
			}).
			Resolver(func(client client.Client, event command.Event, arguments command.Arguments) error {
				return engine.endRaid(client, event.GuildID, event.ChannelID)
			}).
			Build())
}
