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

// Session is an interface
type Session interface {

	// GuildChannelTextCreate
	GuildChannelTextCreate(guildID, categoryID string, name string) (channelID string, err error)

	// ChannelDelete
	ChannelDelete(channelID string) (err error)

	// ChannelPermissionDelete
	ChannelPermissionDelete(channelID, targetID string) (err error)

	// ChannelMessageSend
	ChannelMessageSend(channelID string, content string) (messageID string, err error)

	// ChannelMessageDelete
	ChannelMessageDelete(channelID, messageID string) (err error)

	// UserChannelPermissionSet
	UserChannelPermissionSet(channelID, targetID string, allow, deny int64) (err error)

	// UserChannelPermissions
	UserChannelPermissions(userID, channelID string) (permissions int64, err error)

	// MessageReactionAdd
	MessageReactionAdd(channelID, messageID, emojiID string) (err error)

	// MessageReactionRemove
	MessageReactionRemove(channelID, messageID, emojiID, userID string) (err error)
}
