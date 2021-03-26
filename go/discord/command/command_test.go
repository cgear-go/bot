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
	"github.com/jonathanarnault/cgear-go/go/discord"
)

func TestCommand__AddInt(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command.AddInt", func() {
		g.It("Should append int parameter with the given name to the command parameters", func() {
			command := &command{
				parameters: make([]parameter, 0, 1),
				resolver:   nil,
			}

			command.AddInt("count")
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(command.parameters[0].name).Eql("count")
			g.Assert(command.parameters[0].tpe).Eql(parameterTypeInt)
		})

		g.It("Should return self", func() {
			command := &command{
				parameters: make([]parameter, 0, 1),
				resolver:   nil,
			}

			g.Assert(command.AddInt("count")).Eql(command)
		})
	})
}

func TestCommand__AddString(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command.AddString", func() {
		g.It("Should append string parameter with the given name to the command parameters", func() {
			command := &command{
				parameters: make([]parameter, 0, 1),
				resolver:   nil,
			}

			command.AddString("name")
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(command.parameters[0].name).Eql("name")
			g.Assert(command.parameters[0].tpe).Eql(parameterTypeString)
		})

		g.It("Should return self", func() {
			command := &command{
				parameters: make([]parameter, 0, 1),
				resolver:   nil,
			}

			g.Assert(command.AddString("count")).Eql(command)
		})
	})
}

func TestCommand__AddRest(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command.AddRest", func() {
		g.It("Should append string parameter with the given name to the command parameters", func() {
			command := &command{
				parameters: make([]parameter, 0, 1),
				resolver:   nil,
			}

			command.AddRest("gym")
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(command.parameters[0].name).Eql("gym")
			g.Assert(command.parameters[0].tpe).Eql(parameterTypeRest)
		})

		g.It("Should return self", func() {
			command := &command{
				parameters: make([]parameter, 0, 1),
				resolver:   nil,
			}

			g.Assert(command.AddRest("gym")).Eql(command)
		})
	})
}

func TestCommand__AddResolver(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command.AddResolver", func() {
		g.It("Should set command resolver", func() {
			command := &command{
				parameters: make([]parameter, 0, 1),
				resolver:   nil,
			}

			resolver := func(context.Context, discord.Session, Arguments) error { return io.ErrUnexpectedEOF }
			command.AddResolver(resolver)
			g.Assert(command.resolver(nil, nil, nil)).Eql(io.ErrUnexpectedEOF)
		})
	})
}

func TestCommand(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command", func() {
		g.It("Should support chaining", func() {
			command := &command{
				parameters: make([]parameter, 0, 1),
				resolver:   nil,
			}

			resolver := func(context.Context, discord.Session, Arguments) error { return io.ErrUnexpectedEOF }
			command.
				AddInt("count").
				AddString("name").
				AddRest("gym").
				AddResolver(resolver)
			g.Assert(len(command.parameters)).Eql(3)
			g.Assert(len(command.parameters)).Eql(3)
			g.Assert(command.parameters[0].name).Eql("count")
			g.Assert(command.parameters[0].tpe).Eql(parameterTypeInt)
			g.Assert(command.parameters[1].name).Eql("name")
			g.Assert(command.parameters[1].tpe).Eql(parameterTypeString)
			g.Assert(command.parameters[2].name).Eql("gym")
			g.Assert(command.parameters[2].tpe).Eql(parameterTypeRest)
			g.Assert(command.resolver(nil, nil, nil)).Eql(io.ErrUnexpectedEOF)
		})
	})
}
func TestCommand__execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	g := goblin.Goblin(t)
	g.Describe("command.AddRest", func() {
		g.It("Should execute command", func() {
			parser := NewMockParser(ctrl)
			parser.EXPECT().
				ReadInt().
				Return(1, nil)
			parser.EXPECT().
				ReadString().
				Return("ImBagheera", nil)
			parser.EXPECT().
				ReadRest().
				Return("Fontaine Pépinière", nil)

			command := &command{
				parameters: []parameter{
					{
						name: "count",
						tpe:  parameterTypeInt,
					},
					{
						name: "name",
						tpe:  parameterTypeString,
					},
					{
						name: "gym",
						tpe:  parameterTypeRest,
					},
				},
				resolver: func(context.Context, discord.Session, Arguments) error {
					return errors.New("1 - ImBagheera - Fontaine Pépinière")
				},
			}

			err := command.execute(context.Background(), nil, parser)
			g.Assert(err.Error()).Eql("1 - ImBagheera - Fontaine Pépinière")
		})

		g.It("Should return an error if argument parsing fails", func() {
			command := &command{
				parameters: []parameter{
					{
						name: "count",
						tpe:  parameterTypeInt,
					},
				},
				resolver: nil,
			}
			parser := NewMockParser(ctrl)
			parser.EXPECT().
				ReadInt().
				Return(0, io.EOF)

			g.Assert(command.execute(context.Background(), nil, parser)).Eql(io.EOF)
		})

		g.It("Should return an error if command fails", func() {
			command := &command{
				parameters: []parameter{
					{
						name: "count",
						tpe:  parameterTypeInt,
					},
				},
				resolver: func(context.Context, discord.Session, Arguments) error {
					return io.ErrUnexpectedEOF
				},
			}
			parser := NewMockParser(ctrl)
			parser.EXPECT().
				ReadInt().
				Return(1, nil)

			g.Assert(command.execute(context.Background(), nil, parser)).Eql(io.ErrUnexpectedEOF)
		})
	})
}
