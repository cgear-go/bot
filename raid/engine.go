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
	"log"
	"sync"

	"github.com/cgear-go/bot/discord"
	"github.com/cgear-go/bot/discord/client"
)

// engine is an implementation of `Engine`
type engine struct {

	// lock used for concurrency
	lock *sync.Mutex

	// config holds engine config
	config map[string]Config

	// lobbies holds the raids that are ongoing
	lobbies map[string]*lobby
}

func (e *engine) addUser(client client.Client, channelID, userID string) error {
	return client.UserChannelPermissionSet(
		channelID, userID, discord.PermissionViewChannel, 0)
}

func (e *engine) removeUser(client client.Client, channelID, userID string) error {
	return client.UserChannelPermissionSet(
		channelID, userID, 0, 0)
}

func (e *engine) getLobbyByMessage(messageID string) *lobby {
	lobby, ok := e.lobbies[messageID]
	if !ok {
		return nil
	}
	return lobby
}

func (e *engine) getLobbyByChannel(channelId string) *lobby {
	for _, l := range e.lobbies {
		if l.channelID == channelId {
			return l
		}
	}
	return nil
}

func (e *engine) createRaid(client client.Client, info raidInfo) (err error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	lobby := &lobby{
		info:             info,
		invitesRemaining: 5,
		remoteAttendees:  make([]string, 0),
	}

	raidChannel := e.config[info.guildID].RaidChannelID

	channel, err := client.GuildChannelTextCreate(
		info.guildID,
		e.config[info.guildID].RaidCategoryId,
		fmt.Sprintf("raid-%s", info.start.Format("02-01-15h04")))
	if err != nil {
		return err
	}

	message, err := client.ChannelMessageSend(raidChannel, lobby.formatMessage())
	if err != nil {
		return err
	}

	if err := client.MessageReactionAdd(raidChannel, message, "🙏"); err != nil {
		return err
	}

	if err := e.addUser(client, channel, info.organizerID); err != nil {
		return err
	}

	announce, err := client.ChannelMessageSend(channel, lobby.formatAnnounce())
	if err != nil {
		return err
	}

	lobby.messageID = message
	lobby.channelID = channel
	lobby.announceID = announce

	e.lobbies[lobby.messageID] = lobby

	log.Printf("Creating raid: %v", lobby.messageID)
	return nil
}

func (e *engine) endRaid(client client.Client, guildID, channelId string) (err error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	var lobby *lobby = e.getLobbyByChannel(channelId)
	if lobby == nil {
		return nil
	}

	if err := client.ChannelDelete(lobby.channelID); err != nil {
		return err
	}

	if err := client.ChannelMessageDelete(e.config[guildID].RaidChannelID, lobby.messageID); err != nil {
		return err
	}
	delete(e.lobbies, lobby.messageID)

	log.Printf("Ending raid: %v", lobby.messageID)
	return nil
}

func (e engine) isUserRemote(l *lobby, userID string) bool {
	for _, attendee := range l.remoteAttendees {
		if attendee == userID {
			return true
		}
	}
	return false
}

func (e engine) hasUserJoined(l *lobby, userID string) bool {
	return e.isUserRemote(l, userID)
}

func (e engine) joinRemotely(client client.Client, channelID, messageID, userID string) (err error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	lobby := e.getLobbyByMessage(messageID)
	if lobby == nil {
		return nil
	}

	if lobby.invitesRemaining == 0 {
		return nil
	}

	if e.hasUserJoined(lobby, userID) {
		return nil
	}

	if err := e.addUser(client, lobby.channelID, userID); err != nil {
		return err
	}

	lobby.remoteAttendees = append(lobby.remoteAttendees, userID)
	lobby.invitesRemaining -= 1

	client.ChannelMessageEdit(channelID, messageID, lobby.formatMessage())
	client.ChannelMessageEdit(lobby.channelID, lobby.announceID, lobby.formatAnnounce())
	log.Printf("Player joining raid remotely (%v): %v", lobby.messageID, userID)
	return nil
}

func (e engine) leaveRemotely(client client.Client, channelID, messageID, userID string) (err error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	lobby := e.getLobbyByMessage(messageID)
	if lobby == nil {
		return nil
	}

	if !e.isUserRemote(lobby, userID) {
		return nil
	}

	if err := e.removeUser(client, lobby.channelID, userID); err != nil {
		return err
	}

	index := -1
	for i, attendee := range lobby.remoteAttendees {
		if attendee == userID {
			index = i
			break

		}
	}

	if index > -1 {
		lobby.remoteAttendees[index] = lobby.remoteAttendees[len(lobby.remoteAttendees)-1]
		lobby.remoteAttendees[len(lobby.remoteAttendees)-1] = ""
		lobby.remoteAttendees = lobby.remoteAttendees[:len(lobby.remoteAttendees)-1]
	}
	lobby.invitesRemaining += 1

	client.ChannelMessageEdit(channelID, messageID, lobby.formatMessage())
	client.ChannelMessageEdit(lobby.channelID, lobby.announceID, lobby.formatAnnounce())

	log.Printf("Player leaving raid remotely (%v): %v", lobby.messageID, userID)
	return nil
}

//
func CreateEngine(dispatcher discord.Dispatcher, config map[string]Config) {
	engine := &engine{
		lock:    &sync.Mutex{},
		config:  config,
		lobbies: make(map[string]*lobby),
	}

	registerRaidCommand(dispatcher, engine, config)
	registerEndCommand(dispatcher, engine, config)
	registerJoinReaction(dispatcher, engine, config)
}
