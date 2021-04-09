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

import "time"

// Raid allows to perform operations on raid
type Raid interface {

	// ID returns the raid ID
	ID() (id string)

	// Lobby returns the raid lobby ID
	Lobby() (lobbyID string)

	// OrganizerID returns the ID of the raid organizer
	OrganizerID() (user string)

	// Level returns the raid level
	Level() (level string)

	// Gym returns the name of the gym where the raid will take place
	Gym() (name string)

	// StartTime returns the raid launch time
	StartTime() (start time.Time)

	// End raid
	End()

	// PlayerCount returns the number of players in the raid
	PlayerCount() (count int)

	// InvitesRemaining returns the number of invites that remain
	InvitesRemaining() (count int)

	// GetLocalPlayers returns the player that will join the raid by their own means
	GetLocalPlayers() (users []string)

	// AddLocalPlayer add a player that will join the raid by its own means
	TryAddLocalPlayer(userID string) (added bool)

	// GetRemotePlayers returns the player that will join the raid remotely
	GetRemotePlayers() (users []string)

	// AddRemotePlayer add a player that will join the raid remotely
	TryAddRemotePlayer(userID string) (added bool)

	// TryRemovePlayer removes a player from the raid
	TryRemovePlayer(userID string) (removed bool)
}
