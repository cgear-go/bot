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

package reaction

import (
	"testing"

	"github.com/franela/goblin"
)

func TestDispatcher__AddReaction(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("dispatcher.AddReaction", func() {
		g.It("Should create reaction builder", func() {
			dispatcher := &dispatcher{
				session:    nil,
				reactions:  make([]*reactionBuilder, 0),
				cancellers: make([]func(), 0, 2),
			}

			reaction := dispatcher.AddReaction()

			g.Assert(reaction).IsNotNil()
			g.Assert(len(dispatcher.reactions)).Eql(1)
			g.Assert(reaction).Eql(dispatcher.reactions[0])

			g.Assert(len(dispatcher.reactions[0].reactions)).Eql(0)
			g.Assert(len(dispatcher.reactions[0].filters)).Eql(0)
			g.Assert(dispatcher.reactions[0].reactionAdded == nil).IsTrue()
			g.Assert(dispatcher.reactions[0].reactionRemoved == nil).IsTrue()
		})

		g.It("Should create add reactions to existing list", func() {
			dispatcher := &dispatcher{
				session:    nil,
				reactions:  make([]*reactionBuilder, 1, 2),
				cancellers: make([]func(), 0, 2),
			}

			dispatcher.AddReaction()
			g.Assert(len(dispatcher.reactions)).Eql(2)
		})

		g.It("Should create reaction builder with given reactions", func() {
			dispatcher := &dispatcher{
				session:    nil,
				reactions:  make([]*reactionBuilder, 0),
				cancellers: make([]func(), 0, 2),
			}

			dispatcher.AddReaction("🙏", "👍")
			g.Assert(len(dispatcher.reactions[0].reactions)).Eql(2)
			g.Assert(dispatcher.reactions[0].reactions[0]).Eql("🙏")
			g.Assert(dispatcher.reactions[0].reactions[1]).Eql("👍")
		})
	})
}
