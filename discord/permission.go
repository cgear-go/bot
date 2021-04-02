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

// General permissions
const (
	PermissionAdministrator       = 8
	PermissionViewAuditLog        = 128
	PermissionViewServerInsights  = 524288
	PermissionManageServer        = 32
	PermissionManageRoles         = 268435456
	PermissionManageChannels      = 16
	PermissionKickMembers         = 2
	PermissionBanMembers          = 4
	PermissionCreateInstantInvite = 1
	PermissionChangeNickname      = 67108864
	PermissionManageNicknames     = 134217728
	PermissionManageEmojis        = 1073741824
	PermissionManageWebhooks      = 536870912
	PermissionViewChannel         = 1024
)

// Text channels
const (
	PermissionSendMessages       = 2048
	PermissionSendTTSMessages    = 4096
	PermissionManageMessages     = 8192
	PermissionEmbedLinks         = 16384
	PermissionAttachFiles        = 32768
	PermissionReadMessageHistory = 65536
	PermissionMentionEveryone    = 131072
	PermissionUseExternalEmojis  = 262144
	PermissionAddReactions       = 64
	PermissionUseSlashCommands   = 2147483648
)
