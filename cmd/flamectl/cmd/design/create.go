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

package design

import (
	"github.com/spf13/cobra"

	"github.com/cisco-open/flame/cmd/flamectl/models"
)

func (c *container) CreateDesign() *cobra.Command {
	var command = &cobra.Command{
		Use:   "design <designId>",
		Short: "Create a new design template",
		Long:  "This command creates a new design template",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			designId := args[0]

			flags := cmd.Flags()
			description, err := flags.GetString("desc")
			if err != nil {
				return err
			}

			return c.service.Create(&models.DesignParams{
				CommonParams: models.CommonParams{
					Endpoint: c.config.ApiServer.Endpoint,
					User:     c.config.User,
				},
				Id:   designId,
				Desc: description,
			})
		},
	}

	command.Flags().StringP("desc", "d", "", "Design description")

	return command
}
