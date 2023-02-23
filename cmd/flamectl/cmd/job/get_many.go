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

func (c *container) GetJobs() *cobra.Command {
	var command = &cobra.Command{
		Use:   "jobs",
		Short: "Get status of all jobs",
		Long:  "This command retrieves the status of all jobs",
		Args:  cobra.RangeArgs(0, 0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.service.GetMany(&models.JobParams{
				CommonParams: models.CommonParams{
					Endpoint: c.config.ApiServer.Endpoint,
					User:     c.config.User,
				},
			})
		},
	}

	return command
}