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

import "github.com/jonathanarnault/cgear-go/go/discord"

// Reaction represents a reaction on a message
type Reaction struct {

	// Guild is the guild where the action was performed
	Guild string

	// User is the user that performed the action
	User string

	// Channel is the channel where the action was performed
	Channel string

	// Message holds the message where the action was performed
	Message string

	// Emoji holds the reaction emoji ID
	Emoji string
}

// ReactionFn is the callback used to listen for reactions events
type ReactionFn func(session discord.Session, reaction Reaction) (err error)