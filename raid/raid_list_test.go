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
	"testing"
	"time"

	"github.com/franela/goblin"
)

func TestRaidList__Create(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("raidList.Create", func() {
		g.It("Should append createRaid event to events", func() {
			events := make([]event, 0, 1)
			raidList := &raidList{
				events: &events,
			}

			info := RaidInfo{
				ID:          "456",
				LobbyID:     "789",
				OrganizerID: "123",
				Level:       "5",
				Gym:         "Place Stanislas",
				Start:       time.Now(),
			}

			raidList.Create(info)
			g.Assert(len(events)).Eql(1)

			g.Assert(events[0].eventType).Eql(eventTypeRaidCreate)
			payload, ok := events[0].payload.(RaidInfo)
			g.Assert(ok).IsTrue()
			g.Assert(payload).Eql(info)
		})
	})
}
