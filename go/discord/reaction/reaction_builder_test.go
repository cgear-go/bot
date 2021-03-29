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
	"io"
	"testing"

	"github.com/franela/goblin"
	"github.com/jonathanarnault/cgear-go/go/discord/session"
)

func TestReactionBuilder__AddFilter(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reactionBuilder.AddFilter", func() {
		g.It("Should add filter to reactionBuilder", func() {
			builder := &reactionBuilder{
				filters:         make([]FilterFn, 0, 1),
				reactionAdded:   nil,
				reactionRemoved: nil,
			}

			builder.AddFilter(func(*Reaction) (bool, error) {
				return false, io.EOF
			})

			builder.AddFilter(func(*Reaction) (bool, error) {
				return true, nil
			})

			g.Assert(len(builder.filters)).Eql(2)

			{
				skip, err := builder.filters[0](nil)
				g.Assert(skip).Eql(false)
				g.Assert(err).Eql(io.EOF)
			}

			{
				skip, err := builder.filters[1](nil)
				g.Assert(skip).Eql(true)
				g.Assert(err).Eql(nil)
			}
		})
	})
}

func TestReactionBuilder__OnReactionAdded(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reactionBuilder.OnReactionAdded", func() {
		g.It("Should register reactionAdded callback", func() {
			builder := &reactionBuilder{
				filters:         make([]FilterFn, 0),
				reactionAdded:   nil,
				reactionRemoved: nil,
			}

			builder.OnReactionAdded(func(session.Session, *Reaction) error {
				return io.EOF
			})

			g.Assert(builder.reactionAdded).IsNotNil()
			g.Assert(builder.reactionAdded(nil, nil)).Eql(io.EOF)
		})

		g.It("Should register unregister previous callback when called muliple times", func() {
			builder := &reactionBuilder{
				filters:         make([]FilterFn, 0),
				reactionAdded:   nil,
				reactionRemoved: nil,
			}

			builder.OnReactionAdded(func(session.Session, *Reaction) error {
				return nil
			})

			builder.OnReactionAdded(func(session.Session, *Reaction) error {
				return io.EOF
			})

			g.Assert(builder.reactionAdded).IsNotNil()
			g.Assert(builder.reactionAdded(nil, nil)).Eql(io.EOF)
		})
	})
}

func TestReactionBuilder__OnReactionRemoved(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reactionBuilder.OnReactionRemoved", func() {
		g.It("Should register reactionAdded callback", func() {
			builder := &reactionBuilder{
				filters:         make([]FilterFn, 0),
				reactionAdded:   nil,
				reactionRemoved: nil,
			}

			builder.OnReactionRemoved(func(session.Session, *Reaction) error {
				return io.EOF
			})

			g.Assert(builder.reactionRemoved).IsNotNil()
			g.Assert(builder.reactionRemoved(nil, nil)).Eql(io.EOF)
		})

		g.It("Should register unregister previous callback when called muliple times", func() {
			builder := &reactionBuilder{
				filters:         make([]FilterFn, 0),
				reactionAdded:   nil,
				reactionRemoved: nil,
			}

			builder.OnReactionRemoved(func(session.Session, *Reaction) error {
				return nil
			})

			builder.OnReactionRemoved(func(session.Session, *Reaction) error {
				return io.EOF
			})

			g.Assert(builder.reactionRemoved).IsNotNil()
			g.Assert(builder.reactionRemoved(nil, nil)).Eql(io.EOF)
		})
	})
}
