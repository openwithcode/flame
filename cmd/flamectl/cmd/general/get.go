// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package general

import (
	"github.com/spf13/cobra"
)

func (c *container) GetCommands() *cobra.Command {
	var command = &cobra.Command{
		Use:   "get <resource> <name>",
		Short: "Command to get resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	command.AddCommand(
		c.codeCommands.GetCodeForDesign(),
		c.datasetCommands.GetDataset(),
		c.designCommands.GetDesign(),
		c.designCommands.GetDesigns(),
		c.jobCommands.GetJob(),
		c.jobCommands.GetJobs(),
		c.schemaCommands.GetDesignSchema(),
		c.schemaCommands.GetDesignSchemas(),
	)

	return command
}
