/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ert

import (
	"omg-cli/config"
)

var fullRuntime = config.Tile{
	config.PivnetMetadata{
		"elastic-runtime",
		220833,
		254457,
		"5540900a3626b092bffdb01b530791a116cf5f1022fd1b048edaeea4424318fd",
	},
	product,
	&stemcell,
}

var smallRuntime = config.Tile{
	config.PivnetMetadata{
		"elastic-runtime",
		220833,
		254473,
		"1ab242bff8f95598193b0c742b7d6a520628ebeb682fd949d18da5ef6c8e5c7a",
	},
	product,
	&stemcell,
}

var product = config.OpsManagerMetadata{
	"cf",
	"2.3.3",
}

var stemcell = config.StemcellMetadata{
	config.PivnetMetadata{
		"stemcells-ubuntu-xenial",
		225315,
		259369,
		"6b1c333f382770a5f351362f2c8142b7fd165fe9dd96a902bd4e40006f87a650"},
	"light-bosh-stemcell-97.31-google-kvm-ubuntu-xenial-go_agent",
}

type Tile struct{}

func (*Tile) Definition(envConfig *config.EnvConfig) config.Tile {
	if envConfig.SmallFootprint {
		return smallRuntime
	} else {
		return fullRuntime
	}
}

func (*Tile) BuiltIn() bool {
	return false
}
