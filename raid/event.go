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

// eventType reprents a type of event
type eventType int

const (

	// raidCreate is the event used to create raid
	eventTypeRaidCreate eventType = iota

	// raidEnd is the event used to end raid
	eventTypeRaidEnd
)

// event represents an event to execute to raid list
type event struct {

	// eventType is the event type
	eventType eventType

	// payload is the event payload
	payload interface{}
}

// createRaidPayload is the payload when creating raids
type createRaidPayload struct {
	// id is the raid ID
	id string

	// lobbyID is the raid lobby ID
	lobbyID string

	// info is the raid info
	info RaidInfo
}
