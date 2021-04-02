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
	"io"
	"testing"

	"github.com/cgear-go/bot/discord/session"
	"github.com/franela/goblin"
)

func TestCommand__AddInt(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command.AddInt", func() {
		g.It("Should append int parameter with the given name to the command parameters", func() {
			command := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			command.AddInt("count")
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(command.parameters[0].name).Eql("count")
			g.Assert(command.parameters[0].parameterType).Eql(parameterTypeInt)
		})

		g.It("Should return self", func() {
			command := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
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
			command := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			command.AddString("name")
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(command.parameters[0].name).Eql("name")
			g.Assert(command.parameters[0].parameterType).Eql(parameterTypeString)
		})

		g.It("Should return self", func() {
			command := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
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
			command := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			command.AddRest("gym")
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(len(command.parameters)).Eql(1)
			g.Assert(command.parameters[0].name).Eql("gym")
			g.Assert(command.parameters[0].parameterType).Eql(parameterTypeRest)
		})

		g.It("Should return self", func() {
			command := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			g.Assert(command.AddRest("gym")).Eql(command)
		})
	})
}

func TestCommandBuilder__AddFilter(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("commandBuilder.AddFilter", func() {
		g.It("Should add filter to commandBuilder", func() {
			builder := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			builder.AddFilter(func(Event) (bool, error) {
				return false, io.EOF
			})

			builder.AddFilter(func(Event) (bool, error) {
				return true, nil
			})

			g.Assert(len(builder.filters)).Eql(2)

			{
				skip, err := builder.filters[0](Event{})
				g.Assert(skip).Eql(false)
				g.Assert(err).Eql(io.EOF)
			}

			{
				skip, err := builder.filters[1](Event{})
				g.Assert(skip).Eql(true)
				g.Assert(err).Eql(nil)
			}
		})
	})
}

func TestCommand__Resolver(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command.Resolver", func() {
		g.It("Should set command resolver", func() {
			command := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			resolver := func(context.Context, session.Session, Arguments) error { return io.ErrUnexpectedEOF }
			command.Resolver(resolver)
			g.Assert(command.resolver(nil, nil, nil)).Eql(io.ErrUnexpectedEOF)
		})
	})
}

func TestCommand(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command", func() {
		g.It("Should support chaining", func() {
			command := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			resolver := func(context.Context, session.Session, Arguments) error { return io.ErrUnexpectedEOF }
			command.
				AddInt("count").
				AddString("name").
				AddRest("gym").
				Resolver(resolver)
			g.Assert(len(command.parameters)).Eql(3)
			g.Assert(len(command.parameters)).Eql(3)
			g.Assert(command.parameters[0].name).Eql("count")
			g.Assert(command.parameters[0].parameterType).Eql(parameterTypeInt)
			g.Assert(command.parameters[1].name).Eql("name")
			g.Assert(command.parameters[1].parameterType).Eql(parameterTypeString)
			g.Assert(command.parameters[2].name).Eql("gym")
			g.Assert(command.parameters[2].parameterType).Eql(parameterTypeRest)
			g.Assert(command.resolver(nil, nil, nil)).Eql(io.ErrUnexpectedEOF)
		})
	})
}
