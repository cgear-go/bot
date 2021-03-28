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

import (
	"github.com/jonathanarnault/cgear-go/go/discord/session"
)

// ReactionBuilder allows to build Reactions handlers
type ReactionBuilder interface {

	// AddChannelFilter adds a channel filter for the ReactionBuilder
	AddChannelFilter(filter session.ChannelFilter) (builder ReactionBuilder)

	// OnReactionAdded is executed when a reaction is added to a message
	OnReactionAdded(callback ReactionFn) (builder ReactionBuilder)

	// OnReactionRemoved is executed when a reaction is removed from a message
	OnReactionRemoved(callback ReactionFn) (builder ReactionBuilder)
}

// reactionBuilder is an implmentation of `ReactionBuilder`
type reactionBuilder struct {

	// channelFilters holds the filters to apply to the reaction
	channelFilters []session.ChannelFilter

	// reactionAdded holds the callback function when a reaction is added
	reactionAdded ReactionFn

	// reactionAdded holds the callback function when a reaction is removed
	reactionRemoved ReactionFn
}

func (r *reactionBuilder) AddChannelFilter(filter session.ChannelFilter) ReactionBuilder {
	r.channelFilters = append(r.channelFilters, filter)
	return r
}

func (r *reactionBuilder) OnReactionAdded(callback ReactionFn) ReactionBuilder {
	r.reactionAdded = callback
	return r
}

func (r *reactionBuilder) OnReactionRemoved(callback ReactionFn) ReactionBuilder {
	r.reactionRemoved = callback
	return r
}
