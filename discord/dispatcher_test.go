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
	"io"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/cgear-go/bot/discord/clientmock"
	"github.com/cgear-go/bot/discord/command"
	"github.com/cgear-go/bot/discord/commandmock"
	"github.com/cgear-go/bot/discord/reaction"
	"github.com/cgear-go/bot/discord/reactionmock"
	"github.com/franela/goblin"
	"github.com/golang/mock/gomock"
)

func TestDispatcher__AddCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	g := goblin.Goblin(t)
	g.Describe("dispatcher.AddCommand", func() {
		g.It("Should add command to commands", func() {
			dispatcher := &dispatcher{
				session:   nil,
				client:    nil,
				commands:  make(map[string]command.Command),
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			command := commandmock.NewMockCommand(ctrl)
			command.
				EXPECT().
				Name().
				Return("test")

			dispatcher.AddCommand(command)
			g.Assert(dispatcher.commands["test"]).Eql(command)
		})
	})
}

func TestDispatcher__AddReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	g := goblin.Goblin(t)
	g.Describe("dispatcher.AddReaction", func() {
		g.It("Should add reaction to reactions", func() {
			dispatcher := &dispatcher{
				session:   nil,
				client:    nil,
				commands:  make(map[string]command.Command),
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			reaction := reactionmock.NewMockReaction(ctrl)
			reaction.
				EXPECT().
				Emoji().
				Return("🙏")

			dispatcher.AddReaction(reaction)
			g.Assert(dispatcher.reactions["🙏"]).Eql(reaction)
		})
	})
}

func TestDispatcher__commandEndIndex(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("dispatcher.commandEndIndex", func() {
		g.It("Return end index if using whitespace as split", func() {
			dispatcher := &dispatcher{
				session:   nil,
				client:    nil,
				commands:  make(map[string]command.Command),
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			g.Assert(dispatcher.commandEndIndex("test a b c")).Eql(4)
		})

		g.It("Return end index using carriage return as split", func() {
			dispatcher := &dispatcher{
				session:   nil,
				client:    nil,
				commands:  make(map[string]command.Command),
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			g.Assert(dispatcher.commandEndIndex(`test
a b c`)).Eql(4)
		})

		g.It("Return length if no separator is found", func() {
			dispatcher := &dispatcher{
				session:   nil,
				client:    nil,
				commands:  make(map[string]command.Command),
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			g.Assert(dispatcher.commandEndIndex("test")).Eql(4)
			g.Assert(dispatcher.commandEndIndex("")).Eql(0)
		})
	})
}

func TestDispatcher__executeCommand(t *testing.T) {
	ctrl := gomock.NewController(t)

	g := goblin.Goblin(t)
	g.Describe("dispatcher.executeCommand", func() {
		g.It("Execute command", func() {
			testCommand := commandmock.NewMockCommand(ctrl)
			client := clientmock.NewMockClient(ctrl)
			dispatcher := &dispatcher{
				session: nil,
				client:  client,
				commands: map[string]command.Command{
					"test": testCommand,
				},
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			client.
				EXPECT().
				UserChannelPermissions("456", "78").
				Return(int64(8), nil)

			client.
				EXPECT().
				ChannelMessageDelete("78", "90").
				Return(nil)

			testCommand.
				EXPECT().
				Execute(gomock.Eq(client), gomock.Eq(command.Event{
					GuildID:         "123",
					UserID:          "456",
					UserPermissions: 8,
					ChannelID:       "78",
					MessageID:       "90",
					Params:          "a b c",
				})).
				Return(nil)

			g.Assert(dispatcher.executeCommand(&discordgo.MessageCreate{
				Message: &discordgo.Message{
					ID:      "90",
					GuildID: "123",
					Author: &discordgo.User{
						ID: "456",
					},
					ChannelID: "78",
				},
			}, "test a b c ")).IsNil()
		})

		g.It("Fail silently if command is not valid", func() {
			testCommand := commandmock.NewMockCommand(ctrl)
			client := clientmock.NewMockClient(ctrl)
			dispatcher := &dispatcher{
				session: nil,
				client:  client,
				commands: map[string]command.Command{
					"test": testCommand,
				},
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			client.
				EXPECT().
				ChannelMessageDelete("78", "90").
				Return(nil)

			g.Assert(dispatcher.executeCommand(&discordgo.MessageCreate{
				Message: &discordgo.Message{
					ID:      "90",
					GuildID: "123",
					Author: &discordgo.User{
						ID: "456",
					},
					ChannelID: "78",
				},
			}, "")).IsNil()
		})

		g.It("Fail silently if command is not found", func() {
			testCommand := commandmock.NewMockCommand(ctrl)
			client := clientmock.NewMockClient(ctrl)
			dispatcher := &dispatcher{
				session: nil,
				client:  client,
				commands: map[string]command.Command{
					"test": testCommand,
				},
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			client.
				EXPECT().
				ChannelMessageDelete("78", "90").
				Return(nil)

			g.Assert(dispatcher.executeCommand(&discordgo.MessageCreate{
				Message: &discordgo.Message{
					ID:      "90",
					GuildID: "123",
					Author: &discordgo.User{
						ID: "456",
					},
					ChannelID: "78",
				},
			}, "test1 a b c ")).IsNil()
		})

		g.It("Return permission retrieve error", func() {
			testCommand := commandmock.NewMockCommand(ctrl)
			client := clientmock.NewMockClient(ctrl)
			dispatcher := &dispatcher{
				session: nil,
				client:  client,
				commands: map[string]command.Command{
					"test": testCommand,
				},
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			client.
				EXPECT().
				UserChannelPermissions("456", "78").
				Return(int64(0), io.ErrClosedPipe)

			client.
				EXPECT().
				ChannelMessageDelete("78", "90").
				Return(nil)

			g.Assert(dispatcher.executeCommand(&discordgo.MessageCreate{
				Message: &discordgo.Message{
					ID:      "90",
					GuildID: "123",
					Author: &discordgo.User{
						ID: "456",
					},
					ChannelID: "78",
				},
			}, "test a b c ")).Eql(io.ErrClosedPipe)
		})

		g.It("Return command error", func() {
			testCommand := commandmock.NewMockCommand(ctrl)
			client := clientmock.NewMockClient(ctrl)
			dispatcher := &dispatcher{
				session: nil,
				client:  client,
				commands: map[string]command.Command{
					"test": testCommand,
				},
				reactions: make(map[string]reaction.Reaction),
				closers:   make([]func(), 0, 3),
			}

			client.
				EXPECT().
				UserChannelPermissions("456", "78").
				Return(int64(8), nil)

			client.
				EXPECT().
				ChannelMessageDelete("78", "90").
				Return(nil)

			testCommand.
				EXPECT().
				Execute(gomock.Eq(client), gomock.Eq(command.Event{
					GuildID:         "123",
					UserID:          "456",
					UserPermissions: 8,
					ChannelID:       "78",
					MessageID:       "90",
					Params:          "a b c",
				})).
				Return(io.ErrNoProgress)

			g.Assert(dispatcher.executeCommand(&discordgo.MessageCreate{
				Message: &discordgo.Message{
					ID:      "90",
					GuildID: "123",
					Author: &discordgo.User{
						ID: "456",
					},
					ChannelID: "78",
				},
			}, "test a b c ")).Eql(io.ErrNoProgress)
		})

	})
}

func TestDispatcher__reactionAdded(t *testing.T) {

}

func TestDispatcher__reactionRemoved(t *testing.T) {

}
func TestNewDispatcher(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("TestNewDispatcher", func() {
		g.It("create dispatcher")
	})
}
