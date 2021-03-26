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
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Raid represents a raid
type Raid struct {
	// Channel corresponds to the raid channel
	Channel *discordgo.Channel

	// Gate corresponds to the message used to participate to raid
	Gate *discordgo.Message

	// Announce corresponds to the message that lists players in the raid channel
	Announce *discordgo.Message

	// Level holds the raid level
	Level string

	// Gym holds the gym name
	Gym string

	// Operator corresponds to the raid operator
	Operator *discordgo.User

	// Invites contains the number of invites available
	Invites int

	// RemoteAttendees corresponds to the attendees that need to be invited
	RemoteAttendees []*discordgo.User

	// LocalAttendees corresponds to the users that do not need to be invited
	LocalAttendees []*discordgo.User

	// Start holds the raid launch time
	Start time.Time
}

func (f Raid) formatAttendees(attendees []*discordgo.User) string {
	sb := strings.Builder{}

	for _, attendee := range attendees {
		sb.WriteString("- ")
		sb.WriteString(attendee.Mention())
		sb.WriteString("\n")
	}

	return sb.String()
}

func (r Raid) String() string {
	return fmt.Sprintf(
		"Raid %s - Lancement %s par %s - %s",
		r.Level,
		r.Start.Local().Format("15h04"),
		r.Operator.Mention(),
		r.Gym)
}

func (r Raid) FormatGate() string {
	return fmt.Sprintf(
		`%s

Pour participer au raid :
🙏 pour demander une invitation à distance ;
👍 pour participer sans invitation.

%d invitations à distance disponibles.`,
		r.String(),
		r.Invites)
}

func (r Raid) FormatAnnounce() string {
	return fmt.Sprintf(
		`%s

Participants invités à distance : %s

Participants : %s
- %s
`,
		r.String(),
		r.formatAttendees(r.LocalAttendees),
		r.formatAttendees(r.RemoteAttendees),
		r.Operator.Mention())
}
