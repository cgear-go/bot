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

//go:generate mockgen -destination mock/reaction_mock.go -package mock . Reaction
package reaction

import (
	"io"
	"testing"

	"github.com/franela/goblin"
	"github.com/jonathanarnault/cgear-go/go/discord/client"
)

func TestReaction__Emoji(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reaction.Emoji", func() {
		g.It("Should return reaction emoji", func() {
			reaction := reaction{
				emoji: "🙏",
				filters: []FilterFn{
					func(event Event) (skip bool, err error) {
						return true, nil
					},
				},
				onAdded:   func(discord client.Client, event Event) (err error) { return io.EOF },
				onRemoved: func(discord client.Client, event Event) (err error) { return io.ErrUnexpectedEOF },
			}

			g.Assert(reaction.Emoji()).Eql("🙏")
		})
	})
}

func TestReaction__Added(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reaction.Added", func() {

		g.It("Should execute added callback", func() {
			reaction := reaction{
				emoji: "🙏",
				filters: []FilterFn{
					func(event Event) (skip bool, err error) {
						return false, nil
					},
				},
				onAdded:   func(discord client.Client, event Event) (err error) { return io.EOF },
				onRemoved: nil,
			}

			g.Assert(reaction.Added(nil, Event{})).Eql(io.EOF)
		})

		g.It("Should apply filters", func() {
			reaction := reaction{
				emoji: "🙏",
				filters: []FilterFn{
					func(event Event) (skip bool, err error) {
						return true, nil
					},
				},
				onAdded:   func(discord client.Client, event Event) (err error) { return io.EOF },
				onRemoved: nil,
			}

			g.Assert(reaction.Added(nil, Event{})).Eql(nil)
		})

		g.It("Should return filters error", func() {
			reaction := reaction{
				emoji: "🙏",
				filters: []FilterFn{
					func(event Event) (skip bool, err error) {
						return false, io.ErrClosedPipe
					},
				},
				onAdded:   func(discord client.Client, event Event) (err error) { return io.EOF },
				onRemoved: nil,
			}

			g.Assert(reaction.Added(nil, Event{})).Eql(io.ErrClosedPipe)
		})
	})
}

func TestReaction__Removed(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("reaction.Removed", func() {

		g.It("Should execute removed callback", func() {
			reaction := reaction{
				emoji: "🙏",
				filters: []FilterFn{
					func(event Event) (skip bool, err error) {
						return false, nil
					},
				},
				onAdded:   nil,
				onRemoved: func(discord client.Client, event Event) (err error) { return io.EOF },
			}

			g.Assert(reaction.Removed(nil, Event{})).Eql(io.EOF)
		})

		g.It("Should apply filters", func() {
			reaction := reaction{
				emoji: "🙏",
				filters: []FilterFn{
					func(event Event) (skip bool, err error) {
						return true, nil
					},
				},
				onAdded:   nil,
				onRemoved: func(discord client.Client, event Event) (err error) { return io.EOF },
			}

			g.Assert(reaction.Removed(nil, Event{})).Eql(nil)
		})

		g.It("Should return filters error", func() {
			reaction := reaction{
				emoji: "🙏",
				filters: []FilterFn{
					func(event Event) (skip bool, err error) {
						return false, io.ErrClosedPipe
					},
				},
				onAdded:   nil,
				onRemoved: func(discord client.Client, event Event) (err error) { return io.EOF },
			}

			g.Assert(reaction.Removed(nil, Event{})).Eql(io.ErrClosedPipe)
		})
	})
}
