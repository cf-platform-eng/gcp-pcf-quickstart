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

package commands

import (
	"fmt"
	"log"

	"omg-cli/config"
	"omg-cli/omg/setup"

	"path/filepath"

	"os"

	"github.com/alecthomas/kingpin"
)

type RemoteCommand struct {
	logger  *log.Logger
	command string
	envDir  string
	quiet   bool
}

const (
	RemoteName = "remote"
	Username   = "omg"
)

func (bc *RemoteCommand) register(app *kingpin.Application) {
	c := app.Command(RemoteName, "Run an OMG command from outside of the network").Action(bc.run)
	registerEnvConfigFlag(c, &bc.envDir)
	registerQuietFlag(c, &bc.quiet)
	c.Arg("command", "command and arguments to execute").Required().StringVar(&bc.command)
}

func (bc *RemoteCommand) run(c *kingpin.ParseContext) error {
	sshKeyPath := filepath.Join(bc.envDir, "keys", "jumpbox_ssh")

	cfg, err := config.TerraformFromEnvDirectory(bc.envDir)
	if err != nil {
		return fmt.Errorf("load terraform config: %v", err)
	}

	jb, err := setup.NewJumpbox(bc.logger, os.Stdout, cfg.JumpboxIp, Username, sshKeyPath, bc.envDir, bc.quiet)
	if err != nil {
		return fmt.Errorf("connect to jumpbox: %v", err)
	}

	return run([]step{
		{jb.PoolTillStarted, "PoolTillStarted"},
		{jb.UploadDependencies, "UploadDependencies"},
		{func() error {
			return jb.RunOmg(bc.command)
		}, "RunOmg"},
	})
}
