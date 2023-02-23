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

package job

import (
	"github.com/spf13/cobra"

	"github.com/cisco-open/flame/cmd/flamectl/models"
)

const flagStatus = "status"

func (c *container) GetJob() *cobra.Command {
	var command = &cobra.Command{
		Use:   "job <jobId>",
		Short: "Get a job specification or status",
		Long:  "This command retrieves a job specification or status",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			jobId := args[0]

			params := &models.JobParams{
				CommonParams: models.CommonParams{
					Endpoint: c.config.ApiServer.Endpoint,
					User:     c.config.User,
				},
				Id: jobId,
			}

			flagStatusValue, _ := cmd.Flags().GetBool(flagStatus)
			if flagStatusValue {
				return c.service.GetStatus(params)
			}

			return c.service.Get(params)
		},
	}

	command.Flags().BoolP(flagStatus, "s", false, "Retrieve status of a job")

	return command
}