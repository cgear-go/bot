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
	"log"

	"github.com/cgear-go/bot/discord"
	"github.com/cgear-go/bot/discord/client"
	"github.com/cgear-go/bot/discord/reaction"
)

func registerJoinReaction(dispatcher discord.Dispatcher, engine *engine, config map[string]Config) {
	dispatcher.AddReaction(
		reaction.NewReactionBuilder("🙏").
			AddFilter(func(event reaction.Event) (bool, error) {
				engine.lock.Lock()
				defer engine.lock.Unlock()

				lobby := engine.getLobbyByMessage(event.MessageID)
				if lobby == nil {
					log.Println("Lobby is nil")
					return true, nil
				}
				return lobby.info.organizerID == event.UserID, nil
			}).
			OnAdded(func(client client.Client, event reaction.Event) error {
				return engine.joinRemotely(client, event.ChannelID, event.MessageID, event.UserID)
			}).
			OnRemoved(func(client client.Client, event reaction.Event) error {
				return engine.leaveRemotely(client, event.ChannelID, event.MessageID, event.UserID)
			}).
			Build())
}
