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

func TestCommand__Name(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command.Name", func() {
		g.It("Should return command name", func() {
			command := command{
				name: "test",
				parameters: []parameter{
					{name: "a", parameterType: parameterTypeInt},
					{name: "b", parameterType: parameterTypeString},
					{name: "c", parameterType: parameterTypeRest},
				},
				filters: []FilterFn{
					func(event Event) (skip bool) {
						return true
					},
				},
				resolver: func(client.Client, Event, Arguments) (err error) { return io.EOF },
			}

			g.Assert(command.Name()).Eql("test")
		})
	})
}

func TestCommand__Execute(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("command.Execute", func() {
		g.It("Should execute added callback", func() {
			command := command{
				name: "test",
				parameters: []parameter{
					{name: "a", parameterType: parameterTypeInt},
					{name: "b", parameterType: parameterTypeString},
					{name: "c", parameterType: parameterTypeRest},
				},
				filters:  []FilterFn{},
				resolver: func(client.Client, Event, Arguments) (err error) { return io.EOF },
			}

			g.Assert(command.Execute(nil, Event{Params: "1 test some text"})).Eql(io.EOF)
		})

		g.It("Should pass arguments", func() {
			command := command{
				name: "test",
				parameters: []parameter{
					{name: "a", parameterType: parameterTypeInt},
					{name: "b", parameterType: parameterTypeString},
					{name: "c", parameterType: parameterTypeRest},
				},
				filters: []FilterFn{},
				resolver: func(_ client.Client, _ Event, args Arguments) (err error) {
					g.Assert(args.GetInt("a")).Eql(1)
					g.Assert(args.GetString("b")).Eql("test")
					g.Assert(args.GetString("c")).Eql("some text")
					return nil
				},
			}

			g.Assert(command.Execute(nil, Event{Params: "1 test some text"})).Eql(nil)
		})

		g.It("Should apply filters", func() {
			command := command{
				name: "test",
				parameters: []parameter{
					{name: "a", parameterType: parameterTypeInt},
					{name: "b", parameterType: parameterTypeString},
					{name: "c", parameterType: parameterTypeRest},
				},
				filters: []FilterFn{
					func(event Event) (skip bool) {
						return true
					},
				},
				resolver: func(client.Client, Event, Arguments) (err error) { return io.EOF },
			}
			g.Assert(command.Execute(nil, Event{Params: "1 test some text"})).Eql(nil)
		})

		g.It("Should return parsing errors", func() {
			command := command{
				name: "test",
				parameters: []parameter{
					{name: "a", parameterType: parameterTypeInt},
					{name: "b", parameterType: parameterTypeString},
					{name: "c", parameterType: parameterTypeRest},
				},
				filters:  []FilterFn{},
				resolver: func(client.Client, Event, Arguments) (err error) { return io.ErrUnexpectedEOF },
			}

			g.Assert(command.Execute(nil, Event{Params: "1 test"})).Eql(io.EOF)
		})
	})
}
