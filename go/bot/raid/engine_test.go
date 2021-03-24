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
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/franela/goblin"
)

func TestRaid__SubmitRaid(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("raid.SubmitRaid", func() {
		g.It("Should append raid to the raids map", func() {
			raid := Raid{
				Message: &discordgo.Message{
					ID: "1234",
				},
			}
			engine := &engine{
				raids: make(map[string]Raid),
			}

			engine.SubmitRaid(context.Background(), raid)
			g.Assert(engine.raids["1234"]).Eql(raid)
		})
	})
}
