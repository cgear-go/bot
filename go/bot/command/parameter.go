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

// parameterType represents a parameter type
type parameterType int

const (
	// parameterTypeInt represents an int parameter
	parameterTypeInt parameterType = iota

	// parameterTypeString represents a string parameter
	parameterTypeString

	// parameterTypeRest represents a string parameter that holds the rest of the command arguments
	parameterTypeRest
)

// parameter represents a command parameter
type parameter struct {

	// name holds the parameter name
	name string

	// t holds the parameter type
	tpe parameterType
}
