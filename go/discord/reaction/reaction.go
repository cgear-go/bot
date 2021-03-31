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
	"github.com/jonathanarnault/cgear-go/go/discord/client"
)

// Reaction represents a reaction on a message
type Reaction interface {

	// Emoji returns the reaction emoji
	Emoji() (emoji string)

	// Added executes the added callback
	Added(discord client.Client) (err error)

	// Removed executes the removed callback
	Removed(discord client.Client) (err error)
}

// reaction is an implmentation of `Reaction`
type reaction struct {

	// emoji holds the supported emoji
	emoji string

	// filters holds the filters to apply to the reaction
	filters []FilterFn

	// onAdded holds the callback function when a reaction is added
	onAdded ReactionFn

	// reactionAdded holds the callback function when a reaction is removed
	onRemoved ReactionFn
}

func (r reaction) Emoji() string {
	return ""
}

// Added executes the added callback
func (r reaction) Added(discord client.Client) error {
	return nil
}

// Removed executes the removed callback
func (r reaction) Removed(discord client.Client) error {
	return nil
}
