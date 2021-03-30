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

package discord

import (
	"github.com/jonathanarnault/cgear-go/go/discord/command"
	"github.com/jonathanarnault/cgear-go/go/discord/reaction"
)

type Dispatcher interface {

	// AddCommand registers a command for the dispatcher
	AddCommand(cmd command.Command)

	// AddReaction registers a reaction for the dispatcher
	AddReaction(cmd reaction.Reaction)

	// ListenCommands listen for commands
	ListenCommands()

	// ListenCommands listen for reactions
	ListenReactions()

	// Close dispatcher
	Close()
}
