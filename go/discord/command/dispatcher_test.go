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

package command

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/franela/goblin"
	"github.com/golang/mock/gomock"

	discordmock "github.com/jonathanarnault/cgear-go/go/bot/discord/mock"
)

func TestDispatcher__AddCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	g := goblin.Goblin(t)
	g.Describe("dispatcher.AddCommand", func() {
		g.It("Should create a command with the given name", func() {
			dispatcher := &dispatcher{commands: make(map[string]Command)}

			hello := dispatcher.AddCommand("hello")
			cmd, ok := hello.(*command)
			g.Assert(ok).IsTrue()
			g.Assert(cmd.parameters).IsNotNil()
			g.Assert(cap(cmd.parameters)).Eql(8)
			g.Assert(cmd.resolver == nil).IsTrue()
		})

		g.It("Should return existing command if name is already taken", func() {
			command := NewMockCommand(ctrl)
			dispatcher := &dispatcher{commands: map[string]Command{
				"hello": command,
			}}

			g.Assert(dispatcher.AddCommand("hello")).Eql(command)
		})
	})
}

func TestDispatcher__Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	g := goblin.Goblin(t)
	g.Describe("dispatcher.AddCommand", func() {
		g.It("Should execute the command with the given name", func() {
			bot := discordmock.NewMockBot(ctrl)
			command := NewMockCommand(ctrl)
			dispatcher := &dispatcher{
				commands: map[string]Command{
					"hello": command,
				},
				bot: bot,
			}

			ctx := context.Background()
			command.
				EXPECT().
				execute(gomock.Eq(ctx), gomock.Eq(bot), gomock.Any()).
				Return(nil)
			g.Assert(dispatcher.Execute(ctx, "hello")).IsNil()
		})

		g.It("Return an error if a parameter is not valid", func() {
			bot := discordmock.NewMockBot(ctrl)
			command := NewMockCommand(ctrl)
			dispatcher := &dispatcher{
				commands: map[string]Command{
					"hello": command,
				},
				bot: bot,
			}

			ctx := context.Background()
			command.
				EXPECT().
				execute(gomock.Eq(ctx), gomock.Eq(bot), gomock.Any()).
				Return(io.EOF)
			g.Assert(dispatcher.Execute(ctx, "hello")).Eql(io.EOF)
		})

		g.It("Return an error if execution fails", func() {
			bot := discordmock.NewMockBot(ctrl)
			command := NewMockCommand(ctrl)
			dispatcher := &dispatcher{
				commands: map[string]Command{
					"hello": command,
				},
				bot: bot,
			}

			ctx := context.Background()
			command.
				EXPECT().
				execute(gomock.Eq(ctx), gomock.Eq(bot), gomock.Any()).
				Return(io.ErrClosedPipe)

			g.Assert(dispatcher.Execute(ctx, "hello")).Eql(io.ErrClosedPipe)
		})

		g.It("Return an error if no command is provided", func() {
			command := NewMockCommand(ctrl)
			dispatcher := &dispatcher{commands: map[string]Command{
				"hello": command,
			}}

			g.Assert(dispatcher.Execute(context.Background(), "")).Eql(io.EOF)
		})

		g.It("Return an error if command does not exist", func() {
			command := NewMockCommand(ctrl)
			dispatcher := &dispatcher{commands: map[string]Command{
				"hello": command,
			}}

			g.Assert(dispatcher.Execute(context.Background(), "name")).Eql(errors.New("unexiting command"))
		})
	})
}
