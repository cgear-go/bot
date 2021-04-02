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

type Event struct {

	// GuildID is the guild where the action was performed
	GuildID string

	// UserID is the user that performed the action
	UserID string

	// UserPermission holds the user permission
	UserPermissions int64

	// ChannelID is the channel where the action was performed
	ChannelID string

	// MessageID holds the message where the action was performed
	MessageID string
}
