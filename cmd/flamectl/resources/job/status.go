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
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/cisco-open/flame/cmd/flamectl/models"
	"github.com/cisco-open/flame/pkg/openapi"
	"github.com/cisco-open/flame/pkg/restapi"
)

func (c *container) GetStatus(params *models.JobParams) error {
	// construct URL
	uriMap := map[string]string{
		"user":  params.User,
		"jobId": params.Id,
	}
	url := restapi.CreateURL(params.Endpoint, restapi.GetJobStatusEndPoint, uriMap)

	code, body, err := restapi.HTTPGet(url)
	if err != nil || restapi.CheckStatusCode(code) != nil {
		var msg string
		_ = json.Unmarshal(body, &msg)
		fmt.Printf("Failed to Get job status - code: %d; %s\n", code, msg)
		return nil
	}

	fmt.Printf("Request to get job status successful\n")

	// convert the response into list of struct
	jobStatus := openapi.JobStatus{}
	err = json.Unmarshal(body, &jobStatus)
	if err != nil {
		fmt.Printf("Failed to unmarshal get job status: %v\n", err)
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Job ID", "State", "created At", "started At", "ended At"})
	table.Append([]string{jobStatus.Id, string(jobStatus.State), jobStatus.CreatedAt.String(),
		jobStatus.StartedAt.String(), jobStatus.EndedAt.String()})
	table.Render() // Send output

	return nil
}