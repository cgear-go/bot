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

	// FindByID returns a raid using its ID
	FindByID(id string) (raid *Raid, ok bool)

	// FindByLobby returns a raid using its lobby
	FindByLobby(lobbyID string) (raid *Raid, ok bool)

	// Create creates a raid
	Create(id, lobbyID string, info RaidInfo)
}

// raidList is an implementation of `RaidList`
type raidList struct {

	// events holds the events to execute
	events *[]event
}

func (l *raidList) FindByID(id string) (raid *Raid, ok bool) {
	return nil, false
}

func (l *raidList) FindByLobby(lobbyID string) (raid *Raid, ok bool) {
	return nil, false
}

func (l *raidList) Create(id, lobbyID string, info RaidInfo) {
	*(l.events) = append(*(l.events), event{
		eventType: eventTypeRaidCreate,
		payload: &createRaidPayload{
			id:      id,
			lobbyID: lobbyID,
			info:    info,
		},
	})
}
