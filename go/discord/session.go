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

package discord

import "github.com/bwmarrin/discordgo"

// Session
type Session interface {
	AddHandler(handler interface{}) func()

	// GuildChannelCreateComplex
	GuildChannelCreateComplex(guildID string, data discordgo.GuildChannelCreateData) (*discordgo.Channel, error)

	// ChannelDelete
	ChannelDelete(channelID string) (*discordgo.Channel, error)

	// ChannelPermissionSet
	ChannelPermissionSet(channelID, targetID string, targetType discordgo.PermissionOverwriteType, allow, deny int64) error

	// ChannelPermissionDelete
	ChannelPermissionDelete(channelID, targetID string) error

	// ChannelMessageSend
	ChannelMessageSend(channelID string, content string) (*discordgo.Message, error)

	// ChannelMessageDelete
	ChannelMessageDelete(channelID, messageID string) error

	// UserChannelPermissions
	UserChannelPermissions(userID, channelID string) (int64, error)

	// MessageReactionAdd
	MessageReactionAdd(channelID, messageID, emojiID string) error

	// MessageReactionRemove
	MessageReactionRemove(channelID, messageID, emojiID, userID string) error
}
