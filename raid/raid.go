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

// Raid allows to perform operations on raid
type Raid interface {

	// ID returns the raid ID
	ID() (id string)

	// Lobby returns the raid lobby ID
	Lobby() (lobbyID string)

	// Info returns the raid info
	Info() (info RaidInfo)

	// PlayerCount returns the number of players in the raid
	PlayerCount() (count int)

	// InvitesRemaining returns the number of invites that remain
	InvitesRemaining() (count int)

	// GetLocalPlayers returns the player that will join the raid by their own means
	GetLocalPlayers() (users []string)

	// AddLocalPlayer add a player that will join the raid by its own means
	TryAddLocalPlayer(userID string)

	// GetRemotePlayers returns the player that will join the raid remotely
	GetRemotePlayers() (users []string)

	// AddRemotePlayer add a player that will join the raid remotely
	TryAddRemotePlayer(userID string)

	// End raid
	End() (err error)
}
