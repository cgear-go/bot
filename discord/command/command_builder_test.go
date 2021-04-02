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
	"io"
	"testing"

	"github.com/cgear-go/bot/discord/client"
	"github.com/franela/goblin"
)

func TestCommandBuilder__AddInt(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("builder.AddInt", func() {
		g.It("Should append int parameter with the given name to the command parameters", func() {
			builder := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			builder.AddInt("count")
			g.Assert(len(builder.parameters)).Eql(1)
			g.Assert(len(builder.parameters)).Eql(1)
			g.Assert(builder.parameters[0].name).Eql("count")
			g.Assert(builder.parameters[0].parameterType).Eql(parameterTypeInt)
		})

		g.It("Should return self", func() {
			builder := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			g.Assert(builder.AddInt("count")).Eql(builder)
		})
	})
}

func TestCommandBuilder__AddString(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("builder.AddString", func() {
		g.It("Should append string parameter with the given name to the command parameters", func() {
			builder := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			builder.AddString("name")
			g.Assert(len(builder.parameters)).Eql(1)
			g.Assert(len(builder.parameters)).Eql(1)
			g.Assert(builder.parameters[0].name).Eql("name")
			g.Assert(builder.parameters[0].parameterType).Eql(parameterTypeString)
		})

		g.It("Should return self", func() {
			builder := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			g.Assert(builder.AddString("count")).Eql(builder)
		})
	})
}

func TestCommandBuilder__AddRest(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("builder.AddRest", func() {
		g.It("Should append string parameter with the given name to the command parameters", func() {
			builder := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			builder.AddRest("gym")
			g.Assert(len(builder.parameters)).Eql(1)
			g.Assert(len(builder.parameters)).Eql(1)
			g.Assert(builder.parameters[0].name).Eql("gym")
			g.Assert(builder.parameters[0].parameterType).Eql(parameterTypeRest)
		})

		g.It("Should return self", func() {
			builder := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			g.Assert(builder.AddRest("gym")).Eql(builder)
		})
	})
}

func TestCommandBuilderBuilder__AddFilter(t *testing.T) {
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

func TestCommandBuilder__Resolver(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("builder.Resolver", func() {
		g.It("Should set command resolver", func() {
			builder := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			resolver := func(client.Client, Event, Arguments) error { return io.ErrUnexpectedEOF }
			builder.Resolver(resolver)
			g.Assert(builder.resolver(nil, Event{}, nil)).Eql(io.ErrUnexpectedEOF)
		})
	})
}

func TestCommandBuilder__Build(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("commandBuilder.build", func() {
		g.It("Should build command", func() {
			builder := &commandBuilder{
				name:       "test",
				parameters: make([]parameter, 0, 1),
				filters:    make([]FilterFn, 0, 1),
				resolver:   nil,
			}

			resolver := func(client.Client, Event, Arguments) error { return io.ErrUnexpectedEOF }
			builder.
				AddInt("count").
				AddString("name").
				AddRest("gym").
				Resolver(resolver)
			g.Assert(len(builder.parameters)).Eql(3)
			g.Assert(len(builder.parameters)).Eql(3)
			g.Assert(builder.parameters[0].name).Eql("count")
			g.Assert(builder.parameters[0].parameterType).Eql(parameterTypeInt)
			g.Assert(builder.parameters[1].name).Eql("name")
			g.Assert(builder.parameters[1].parameterType).Eql(parameterTypeString)
			g.Assert(builder.parameters[2].name).Eql("gym")
			g.Assert(builder.parameters[2].parameterType).Eql(parameterTypeRest)
			g.Assert(builder.resolver(nil, Event{}, nil)).Eql(io.ErrUnexpectedEOF)
		})
	})
}

func TestNewCommandBuilder(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("NewCommandBuilder", func() {
		g.It("Should create command builder", func() {
			builder, ok := NewCommandBuilder("test").(*commandBuilder)

			g.Assert(ok).IsTrue()
			g.Assert(builder.name).Eql("test")
			g.Assert(len(builder.parameters)).Eql(0)
			g.Assert(len(builder.filters)).Eql(0)
			g.Assert(builder.resolver == nil).IsTrue()
		})
	})
}
