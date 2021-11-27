// Copyright (c) 2021 Cisco Systems, Inc. and its affiliates
// All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/cisco/fledge/pkg/openapi"
	"github.com/cisco/fledge/pkg/restapi"
	"github.com/cisco/fledge/pkg/util"
)

// DatasetsApiService is a service that implents the logic for the DatasetsApiServicer
// This service should implement the business logic for every endpoint for the DatasetsApi API.
// Include any external packages or services that will be required by this service.
type DatasetsApiService struct {
}

// NewDatasetsApiService creates a default api service
func NewDatasetsApiService() openapi.DatasetsApiServicer {
	return &DatasetsApiService{}
}

// CreateDataset - Create meta info for a new dataset.
func (s *DatasetsApiService) CreateDataset(ctx context.Context, user string,
	datasetInfo openapi.DatasetInfo) (openapi.ImplResponse, error) {
	//TODO input validation
	zap.S().Debugf("New dataset request received for user: %s | datasetInfo: %v", user, datasetInfo)

	// create controller request
	uriMap := map[string]string{
		"user": user,
	}
	url := restapi.CreateURL(HostEndpoint, restapi.CreateDatasetEndPoint, uriMap)

	// send post request
	code, resp, err := restapi.HTTPPost(url, datasetInfo, "application/json")

	// response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("create new dataset request failed")
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), err
	}

	return openapi.Response(http.StatusCreated, string(resp)), nil
}

// GetAllDatasets - Get the meta info on all the datasets
func (s *DatasetsApiService) GetAllDatasets(ctx context.Context, limit int32) (openapi.ImplResponse, error) {
	// TODO - update GetAllDatasets with the required logic for this service method.
	// Add api_datasets_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []DatasetInfo{}) or use other options such as http.Ok ...
	//return Response(200, []DatasetInfo{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetAllDatasets method not implemented")
}

// GetDataset - Get dataset meta information
func (s *DatasetsApiService) GetDataset(ctx context.Context, user string, datasetId string) (openapi.ImplResponse, error) {
	// TODO - update GetDataset with the required logic for this service method.
	// Add api_datasets_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DatasetInfo{}) or use other options such as http.Ok ...
	//return Response(200, DatasetInfo{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetDataset method not implemented")
}

// GetDatasets - Get the meta info on all the datasets owned by user
func (s *DatasetsApiService) GetDatasets(ctx context.Context, user string, limit int32) (openapi.ImplResponse, error) {
	zap.S().Debugf("get list of datasets for user: %s | limit: %d", user, limit)

	//create controller request
	//construct URL
	uriMap := map[string]string{
		"user":  user,
		"limit": strconv.Itoa(int(limit)),
	}
	url := restapi.CreateURL(HostEndpoint, restapi.GetDatasetsEndPoint, uriMap)

	//send get request
	code, responseBody, err := restapi.HTTPGet(url)

	//response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("get datasets information request failed")
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), err
	}

	var resp []openapi.DatasetInfo
	err = util.ByteToStruct(responseBody, &resp)
	return openapi.Response(http.StatusOK, resp), err
}

// UpdateDataset - Update meta info for a given dataset
func (s *DatasetsApiService) UpdateDataset(ctx context.Context, user string, datasetId string,
	datasetInfo openapi.DatasetInfo) (openapi.ImplResponse, error) {
	// TODO - update UpdateDataset with the required logic for this service method.
	// Add api_datasets_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("UpdateDataset method not implemented")
}
