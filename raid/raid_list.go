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

// RaidList holds the list of raids
type RaidList interface {

	// FindRaidByID returns a raid using its ID
	FindRaidByID(ID string) (raid *Raid, ok bool)

	// FindRaidByLobby returns a raid using its lobby
	FindRaidByLobby(lobbyID string) (raid *Raid, ok bool)

	// CreateRaid creates a raid
	CreateRaid(ID, lobbyID, info RaidInfo) (err error)
}
