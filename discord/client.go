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

type clientImpl struct {
	session *discordgo.Session
}

func (c clientImpl) GuildChannelTextCreate(guildID, categoryID string, name string) (channelID string, err error) {
	channel, err := c.session.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
		Name:     name,
		ParentID: categoryID,
	})
	if err != nil {
		return "", err
	}
	return channel.ID, nil
}

func (c clientImpl) ChannelDelete(channelID string) error {
	_, err := c.session.ChannelDelete(channelID)
	return err
}

func (c clientImpl) ChannelPermissionDelete(channelID, targetID string) error {
	return c.session.ChannelPermissionDelete(channelID, targetID)
}

func (c clientImpl) ChannelMessageSend(channelID string, content string) (string, error) {
	message, err := c.session.ChannelMessageSend(channelID, content)
	if err != nil {
		return "", err
	}
	return message.ID, nil
}

func (c clientImpl) ChannelMessageEdit(channelID string, messageID, content string) error {
	_, err := c.session.ChannelMessageEdit(channelID, messageID, content)
	return err
}

func (c clientImpl) ChannelMessageDelete(channelID, messageID string) error {
	return c.session.ChannelMessageDelete(channelID, messageID)
}

func (c clientImpl) UserChannelPermissionSet(channelID, targetID string, allow, deny int64) error {
	return c.session.ChannelPermissionSet(channelID, targetID, discordgo.PermissionOverwriteTypeMember, allow, deny)
}

func (c clientImpl) UserChannelPermissions(userID, channelID string) (int64, error) {
	return c.session.UserChannelPermissions(userID, channelID)
}

func (c clientImpl) MessageReactionAdd(channelID, messageID, emojiID string) error {
	return c.session.MessageReactionAdd(channelID, messageID, emojiID)
}

func (c clientImpl) MessageReactionRemove(channelID, messageID, emojiID, userID string) error {
	return c.session.MessageReactionRemove(channelID, messageID, emojiID, userID)
}
