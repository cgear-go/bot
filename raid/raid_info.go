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

import "time"

type raidInfo struct {

	// Guild is the ID of the guild where the raid will take place
	guildID string

	// Organizer is the ID of the raid organizer
	organizerID string

	// level is the raid level
	level string

	// gym is the name of the gym where the raid will take place
	gym string

	// start holds the raid launch time
	start time.Time
}
