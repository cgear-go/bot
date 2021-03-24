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
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Raid represents a raid
type Raid struct {

	// Message corresponds to the message used to register to raid
	Message *discordgo.Message

	// Channel corresponds to the raid channel
	Channel *discordgo.Channel

	// Level holds the raid level
	Level string

	// Gym holds the gym name
	Gym string

	// Operator corresponds to the raid operator
	Operator *discordgo.User

	// Invites contains the number of invites available
	Invites int

	// Attendees corresponds to the users invited to the raid
	Attendees []*discordgo.User

	// Start holds the raid launch time
	Start time.Time
}

func (r Raid) String() string {
	return fmt.Sprintf(
		"Raid %s - Lancement %s par %s - %s",
		r.Level,
		r.Start.Local().Format("15h04"),
		r.Operator.Mention(),
		r.Gym)
}
