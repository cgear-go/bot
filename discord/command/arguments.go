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

// Arguments represents a command arguments
type Arguments interface {

	// GetInt returns the integer value of an argument
	GetInt(name string) int

	// GetString returns the string value of an argument
	GetString(name string) string
}

// arguments is an implementation of `Arguments`
type arguments struct {

	// values holds the arguments map
	values map[string]interface{}
}

func (a arguments) GetInt(name string) int {
	switch v := a.values[name].(type) {
	case int:
		return v
	default:
		return 0
	}
}

func (a arguments) GetString(name string) string {
	switch v := a.values[name].(type) {
	case string:
		return v
	default:
		return ""
	}
}
