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

	"github.com/cgear-go/bot/discord/client"
	"github.com/franela/goblin"
)

func TestReactionBuilder__AddFilter(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reactionBuilder.AddFilter", func() {
		g.It("Should add filter to reactionBuilder", func() {
			builder := &reactionBuilder{
				emoji:     "",
				filters:   make([]FilterFn, 0, 1),
				onAdded:   nil,
				onRemoved: nil,
			}

			builder.AddFilter(func(Event) (bool, error) {
				return false, io.EOF
			})

			builder.AddFilter(func(Event) (bool, error) {
				return true, nil
			})

			g.Assert(len(builder.filters)).Eql(2)

			{
				skip, err := builder.filters[0](Event{})
				g.Assert(skip).Eql(false)
				g.Assert(err).Eql(io.EOF)
			}

			{
				skip, err := builder.filters[1](Event{})
				g.Assert(skip).Eql(true)
				g.Assert(err).Eql(nil)
			}
		})
	})
}

func TestReactionBuilder__OnAdded(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reactionBuilder.OnAdded", func() {
		g.It("Should register reactionAdded callback", func() {
			builder := &reactionBuilder{
				emoji:     "",
				filters:   make([]FilterFn, 0),
				onAdded:   nil,
				onRemoved: nil,
			}

			builder.OnAdded(func(client.Client, Event) error {
				return io.EOF
			})

			g.Assert(builder.onAdded).IsNotNil()
			g.Assert(builder.onAdded(nil, Event{})).Eql(io.EOF)
		})

		g.It("Should register unregister previous callback when called muliple times", func() {
			builder := &reactionBuilder{
				emoji:     "",
				filters:   make([]FilterFn, 0),
				onAdded:   nil,
				onRemoved: nil,
			}

			builder.OnAdded(func(client.Client, Event) error {
				return nil
			})

			builder.OnAdded(func(client.Client, Event) error {
				return io.EOF
			})

			g.Assert(builder.onAdded).IsNotNil()
			g.Assert(builder.onAdded(nil, Event{})).Eql(io.EOF)
		})
	})
}

func TestReactionBuilder__OnRemoved(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reactionBuilder.OnRemoved", func() {
		g.It("Should register reactionAdded callback", func() {
			builder := &reactionBuilder{
				emoji:     "",
				filters:   make([]FilterFn, 0),
				onAdded:   nil,
				onRemoved: nil,
			}

			builder.OnRemoved(func(client.Client, Event) error {
				return io.EOF
			})

			g.Assert(builder.onRemoved).IsNotNil()
			g.Assert(builder.onRemoved(nil, Event{})).Eql(io.EOF)
		})

		g.It("Should register unregister previous callback when called muliple times", func() {
			builder := &reactionBuilder{
				emoji:     "",
				filters:   make([]FilterFn, 0),
				onAdded:   nil,
				onRemoved: nil,
			}

			builder.OnRemoved(func(client.Client, Event) error {
				return nil
			})

			builder.OnRemoved(func(client.Client, Event) error {
				return io.EOF
			})

			g.Assert(builder.onRemoved).IsNotNil()
			g.Assert(builder.onRemoved(nil, Event{})).Eql(io.EOF)
		})
	})
}

func TestReactionBuilder__Build(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reactionBuilder.Build", func() {
		g.It("Should return reaction", func() {
			builder := &reactionBuilder{
				emoji:     "🙏",
				filters:   make([]FilterFn, 0),
				onAdded:   nil,
				onRemoved: nil,
			}

			builder.OnRemoved(func(client.Client, Event) error {
				return io.EOF
			})

			_, ok := builder.Build().(reaction)
			g.Assert(ok).IsTrue()
		})

		g.It("Should initialize reaction fields", func() {
			builder := &reactionBuilder{
				emoji: "🙏",
				filters: []FilterFn{
					func(event Event) (skip bool, err error) {
						return true, nil
					},
				},
				onAdded:   func(discord client.Client, event Event) (err error) { return io.EOF },
				onRemoved: func(discord client.Client, event Event) (err error) { return io.ErrUnexpectedEOF },
			}

			r, ok := builder.Build().(reaction)
			g.Assert(ok).IsTrue()
			g.Assert(r.emoji).Eql("🙏")
			g.Assert(len(r.filters)).Eql(1)
			g.Assert(r.onAdded(nil, Event{})).Eql(io.EOF)
			g.Assert(r.onRemoved(nil, Event{})).Eql(io.ErrUnexpectedEOF)
		})
	})
}

func TestNewReactionBuilder(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("NewReactionBuilder", func() {
		g.It("Should return *reactionBuilder", func() {
			_, ok := NewReactionBuilder("🙏").(*reactionBuilder)
			g.Assert(ok).IsTrue()
		})

		g.It("Should set init *reactionBuilder fields", func() {
			builder, ok := NewReactionBuilder("🙏").(*reactionBuilder)
			g.Assert(ok).IsTrue()
			g.Assert(builder.emoji).Eql("🙏")
			g.Assert(len(builder.filters)).Eql(0)
			g.Assert(builder.onAdded == nil).IsTrue()
			g.Assert(builder.onRemoved == nil).IsTrue()
		})
	})
}
