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
	"time"

	"github.com/bwmarrin/discordgo"
)

type Raid struct {
	// Channel corresponds to the raid channel
	Channel *discordgo.Channel

	// Gym holds the gym name
	Gym string

	// Operator corresponds to the raid operator
	Operator string

	// Invited corresponds to the users invited to the raid
	Invited []string

	// Start holds the raid launch time
	Start time.Time
}
