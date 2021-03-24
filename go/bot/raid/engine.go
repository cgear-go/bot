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

package raid

import "context"

// Engine represents the raid engine
type Engine interface {

	// SubmitRaid submits a raid with the given users
	SubmitRaid(context.Context, Raid)
}

// engine is and implmentation of `Engine`
type engine struct {

	// raids holds the raids that are ongoing
	raids map[string]Raid
}

func (e *engine) SubmitRaid(_ context.Context, raid Raid) {
	e.raids[raid.ID] = raid
}

// NewEngine creates a new raid engine
func NewEngine() Engine {
	return &engine{
		raids: make(map[string]Raid),
	}
}
