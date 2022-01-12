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

// JobsApiService is a service that implents the logic for the JobsApiServicer
// This service should implement the business logic for every endpoint for the JobsApi API.
// Include any external packages or services that will be required by this service.
type JobsApiService struct {
}

// NewJobsApiService creates a default api service
func NewJobsApiService() openapi.JobsApiServicer {
	return &JobsApiService{}
}

// CreateJob - Create a new job specification
func (s *JobsApiService) CreateJob(ctx context.Context, user string, jobSpec openapi.JobSpec) (openapi.ImplResponse, error) {
	// TODO: validate the input
	zap.S().Debugf("New job request received for user: %s | jobSpec: %v", user, jobSpec)

	// create controller request
	uriMap := map[string]string{
		"user": user,
	}
	url := restapi.CreateURL(HostEndpoint, restapi.CreateJobEndpoint, uriMap)

	// send post request
	code, responseBody, err := restapi.HTTPPost(url, jobSpec, "application/json")

	// response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("create new job request failed")
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), err
	}

	// everything went well and response is a job id
	resp := openapi.JobStatus{}
	err = util.ByteToStruct(responseBody, &resp)
	if err != nil {
		errMsg := fmt.Sprintf("failed to construct response message: %v", err)
		zap.S().Errorf(errMsg)
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf(errMsg)
	}

	return openapi.Response(http.StatusCreated, resp), nil
}

// DeleteJob - Delete job specification
func (s *JobsApiService) DeleteJob(ctx context.Context, user string, jobId string) (openapi.ImplResponse, error) {
	// TODO - update DeleteJob with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("DeleteJob method not implemented")
}

// GetJob - Get a job specification
func (s *JobsApiService) GetJob(ctx context.Context, user string, jobId string) (openapi.ImplResponse, error) {
	// TODO - update GetJob with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, JobSpec{}) or use other options such as http.Ok ...
	//return Response(200, JobSpec{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetJob method not implemented")
}

// GetJobStatus - Get job status of a given jobId
func (s *JobsApiService) GetJobStatus(ctx context.Context, user string, jobId string) (openapi.ImplResponse, error) {
	// TODO - update GetJobStatus with the required logic for this service method.
	// Add api_jobs_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, JobStatus{}) or use other options such as http.Ok ...
	//return Response(200, JobStatus{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetJobStatus method not implemented")
}

// GetJobs - Get status info on all the jobs owned by user
func (s *JobsApiService) GetJobs(ctx context.Context, user string, limit int32) (openapi.ImplResponse, error) {
	zap.S().Debugf("Get status of all jobs for user: %s", user)

	//create controller request
	uriMap := map[string]string{
		"user":  user,
		"limit": strconv.Itoa(int(limit)),
	}
	url := restapi.CreateURL(HostEndpoint, restapi.GetJobsEndPoint, uriMap)

	//send get request
	code, responseBody, err := restapi.HTTPGet(url)

	// response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, err), fmt.Errorf("get jobs status failed: %v", err)
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		var errMsg error
		_ = util.ByteToStruct(responseBody, &errMsg)
		return openapi.Response(code, errMsg), err
	}

	var resp []openapi.JobStatus
	err = util.ByteToStruct(responseBody, &resp)

	return openapi.Response(http.StatusOK, resp), err
}

// GetTask - Get a job task for a given job and agent
func (s *JobsApiService) GetTask(ctx context.Context, jobId string, agentId string) (openapi.ImplResponse, error) {
	uriMap := map[string]string{
		"jobId":   jobId,
		"agentId": agentId,
	}
	url := restapi.CreateURL(HostEndpoint, restapi.GetTaskEndpoint, uriMap)

	code, taskMap, err := restapi.HTTPGetMultipart(url)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to fetch task")
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), err
	}

	return openapi.Response(http.StatusOK, taskMap), nil
}

// GetTasksInfo - Get the info of tasks in a job
func (s *JobsApiService) GetTasksInfo(ctx context.Context, user string, jobId string, limit int32) (openapi.ImplResponse, error) {
	uriMap := map[string]string{
		"user":  user,
		"jobId": jobId,
		"limit": strconv.Itoa(int(limit)),
	}
	url := restapi.CreateURL(HostEndpoint, restapi.GetTasksInfoEndpoint, uriMap)

	code, responseBody, err := restapi.HTTPGet(url)

	// response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, err), fmt.Errorf("get tasks info failed: %v", err)
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		var errMsg error
		_ = util.ByteToStruct(responseBody, &errMsg)
		return openapi.Response(code, errMsg), err
	}

	var resp []openapi.TaskInfo
	err = util.ByteToStruct(responseBody, &resp)

	return openapi.Response(http.StatusOK, resp), err
}

// UpdateJob - Update a job specification
func (s *JobsApiService) UpdateJob(ctx context.Context, user string, jobId string,
	jobSpec openapi.JobSpec) (openapi.ImplResponse, error) {
	zap.S().Debugf("Job update request received for user: %s | jobSpec: %v", user, jobSpec)

	// create controller request
	uriMap := map[string]string{
		"user":  user,
		"jobId": jobId,
	}
	url := restapi.CreateURL(HostEndpoint, restapi.UpdateJobEndPoint, uriMap)

	// send put request
	code, _, err := restapi.HTTPPut(url, jobSpec, "application/json")

	// response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("update job request failed")
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), err
	}

	return openapi.Response(http.StatusOK, nil), err
}

// UpdateJobStatus - Update the status of a job
func (s *JobsApiService) UpdateJobStatus(ctx context.Context, user string, jobId string,
	jobStatus openapi.JobStatus) (openapi.ImplResponse, error) {
	zap.S().Debugf("Job status update request received for user: %s | jobStatus: %v", user, jobStatus)

	// create controller request
	uriMap := map[string]string{
		"user":  user,
		"jobId": jobId,
	}
	url := restapi.CreateURL(HostEndpoint, restapi.UpdateJobStatusEndPoint, uriMap)

	// send put request
	code, _, err := restapi.HTTPPut(url, jobStatus, "application/json")
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("job status update request failed")
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

// UpdateTaskStatus - Update the status of a task
func (s *JobsApiService) UpdateTaskStatus(ctx context.Context, jobId string, agentId string,
	taskStatus openapi.TaskStatus) (openapi.ImplResponse, error) {
	zap.S().Debugf("Task status update request received for job %s | agent: %s", jobId, agentId)

	// create controller request
	uriMap := map[string]string{
		"jobId":   jobId,
		"agentId": agentId,
	}
	url := restapi.CreateURL(HostEndpoint, restapi.UpdateTaskStatusEndPoint, uriMap)

	// send put request
	code, _, err := restapi.HTTPPut(url, taskStatus, "application/json")
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("agent status update request failed")
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}
