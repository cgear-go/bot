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
	"testing"

	"github.com/franela/goblin"
)

var testArguments = arguments{
	values: map[string]interface{}{
		"count": 1,
		"name":  "ImBagheera",
	},
}

func TestArguments__GetInt(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("arguments.GetInt", func() {
		g.It("Should return the argument value if found", func() {
			g.Assert(testArguments.GetInt("count")).Eql(1)
		})

		g.It("Should return zero value if argument is not an int", func() {
			g.Assert(testArguments.GetInt("name")).IsZero()
		})

		g.It("Should return zero value if argument is not set", func() {
			g.Assert(testArguments.GetInt("unset")).IsZero()
		})
	})
}

func TestArguments__GetString(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("arguments.GetString", func() {
		g.It("Should return the argument value if found", func() {
			g.Assert(testArguments.GetString("name")).Eql("ImBagheera")
		})

		g.It("Should return zero value if argument is not a string", func() {
			g.Assert(testArguments.GetString("count")).IsZero()
		})

		g.It("Should return zero value if argument is not set", func() {
			g.Assert(testArguments.GetString("unset")).IsZero()
		})
	})
}
