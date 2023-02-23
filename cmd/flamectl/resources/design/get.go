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
	"github.com/cisco-open/flame/cmd/flamectl/models"
	"github.com/cisco-open/flame/pkg/restapi"
	"github.com/cisco-open/flame/pkg/util"
	"go.uber.org/zap"
)

func (c *container) Get(params *models.DesignParams) error {
	logger := c.logger.Sugar().WithOptions(zap.Fields(
		zap.String("user", params.User),
		zap.String("designId", params.Id),
	))

	// construct URL
	uriMap := map[string]string{
		"user":     params.User,
		"designId": params.Id,
	}
	url := restapi.CreateURL(params.Endpoint, restapi.GetDesignEndPoint, uriMap)

	// send get request
	code, responseBody, err := restapi.HTTPGet(url)
	if err != nil || restapi.CheckStatusCode(code) != nil {
		logger.Errorf("Failed to retrieve design: code %d\n\n %s", code, string(responseBody))

		return nil
	}

	// format the output into prettyJson format
	prettyJSON, err := util.FormatJSON(responseBody)
	if err != nil {
		logger.Errorf("Failed while formating json: %v\n\n %s", err, string(responseBody))
		return err
	}

	logger.Info(string(prettyJSON))

	return nil
}