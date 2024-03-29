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
)

// lobby holds a raid lobby
type lobby struct {

	// info holds the raids infos
	info raidInfo

	// messageID
	messageID string

	// channelID
	channelID string

	// channelID
	announceID string

	//
	invitesRemaining int

	// Players that join the raid remotely
	remoteAttendees []string

	// Players that join the raid without invites
	localAttendees []string
}

func (l lobby) String() string {
	return fmt.Sprintf(
		"Raid %s - Lancement %s par <@%s> - %s",
		l.info.level,
		l.info.start.Local().Format("15h04"),
		l.info.organizerID,
		l.info.gym)
}

func (l lobby) formatAttendees(attendees []string) string {
	sb := strings.Builder{}

	for _, attendee := range attendees {
		sb.WriteString("\n- <@")
		sb.WriteString(attendee)
		sb.WriteString(">")
	}

	return sb.String()
}

// formatMessage
func (l lobby) formatMessage() string {
	return fmt.Sprintf(
		`%s

Pour participer au raid :
🙏 pour demander une invitation à distance ;
👍 pour participer sans invitation.

%d invitations à distance disponibles.`,
		l.String(),
		l.invitesRemaining)
}

// formatAnnounce
func (l lobby) formatAnnounce() string {
	return fmt.Sprintf(
		`%s

Participants :%s
- <@%s>

Participants invités à distance :%s
`,
		l.String(),
		l.formatAttendees(l.localAttendees),
		l.info.organizerID,
		l.formatAttendees(l.remoteAttendees))
}
