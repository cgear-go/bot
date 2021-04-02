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

//go:generate mockgen -destination mock/reaction_builder_mock.go -package mock . ReactionBuilder
package reaction

import "github.com/cgear-go/bot/discord/client"

type FilterFn func(event Event) (skip bool, err error)

type ReactionFn func(discord client.Client, event Event) (err error)

// ReactionBuilder allows to build Reactions handlers
type ReactionBuilder interface {

	// AddFilter adds a filter for the ReactionBuilder
	AddFilter(filter FilterFn) (builder ReactionBuilder)

	// OnAdded is executed when a reaction is added to a message
	OnAdded(callback ReactionFn) (builder ReactionBuilder)

	// OnRemoved is executed when a reaction is removed from a message
	OnRemoved(callback ReactionFn) (builder ReactionBuilder)

	Build() (reaction Reaction)
}

// reactionBuilder is an implmentation of `ReactionBuilder`
type reactionBuilder struct {

	// emoji holds the supported emoji
	emoji string

	// filters holds the filters to apply to the reaction
	filters []FilterFn

	// onAdded holds the callback function when a reaction is added
	onAdded ReactionFn

	// reactionAdded holds the callback function when a reaction is removed
	onRemoved ReactionFn
}

func (r *reactionBuilder) AddFilter(filter FilterFn) ReactionBuilder {
	r.filters = append(r.filters, filter)
	return r
}

func (r *reactionBuilder) OnAdded(callback ReactionFn) ReactionBuilder {
	r.onAdded = callback
	return r
}

func (r *reactionBuilder) OnRemoved(callback ReactionFn) ReactionBuilder {
	r.onRemoved = callback
	return r
}

func (r *reactionBuilder) Build() Reaction {
	reaction := reaction{
		emoji:     r.emoji,
		filters:   make([]FilterFn, len(r.filters)),
		onAdded:   r.onAdded,
		onRemoved: r.onRemoved,
	}
	copy(reaction.filters, r.filters[:])
	return reaction
}

func NewReactionBuilder(emoji string) ReactionBuilder {
	return &reactionBuilder{
		emoji:     emoji,
		filters:   make([]FilterFn, 0),
		onAdded:   nil,
		onRemoved: nil,
	}
}
