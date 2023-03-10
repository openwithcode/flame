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
	"fmt"
	"os"

	"github.com/cisco-open/flame/cmd/controller/app/database"
	"github.com/cisco-open/flame/cmd/controller/app/objects"
	"github.com/cisco-open/flame/cmd/controller/config"
	"github.com/cisco-open/flame/pkg/openapi"
	"github.com/cisco-open/flame/pkg/util"
)

const (
	defaultGroup   = "default"
	groupByTypeTag = "tag"
	taskKeyLen     = 32
)

////////////////////////////////////////////////////////////////////////////////
// Job Builder related code
////////////////////////////////////////////////////////////////////////////////

// JobBuilder is a struct used to build a job.
type JobBuilder struct {
	// dbService is the database service used to interact with the database.
	dbService database.DBService
	// jobSpec contains the details of the job to be built.
	jobSpec *openapi.JobSpec
	// jobParams contains the parameters needed to construct the jobSpec.
	jobParams config.JobParams

	// schema contains information for OpenAPI design schema.
	schema openapi.DesignSchema
	// roleCode maps each role to their respective code.
	roleCode map[string][]byte
	// datasets contains a list of available datasets.
	// structure is: trainer role => group name => list of datasets
	datasets map[string]map[string][]openapi.DatasetInfo

	// groupAssociations stores which groups have access to which resources.
	groupAssociations map[string][]map[string]string

	channels map[string]openapi.Channel
}

func NewJobBuilder(dbService database.DBService, jobParams config.JobParams) *JobBuilder {
	return &JobBuilder{
		dbService:         dbService,
		jobParams:         jobParams,
		roleCode:          make(map[string][]byte),
		groupAssociations: make(map[string][]map[string]string),
		datasets:          make(map[string]map[string][]openapi.DatasetInfo),
	}
}

func (b *JobBuilder) GetTasks(jobSpec *openapi.JobSpec) (
	tasks []objects.Task, roles []string, err error,
) {
	b.jobSpec = jobSpec
	if b.jobSpec == nil {
		return nil, nil, fmt.Errorf("job spec is nil")
	}

	err = b.setup()
	if err != nil {
		return nil, nil, err
	}

	tasks, roles, err = b.build()
	if err != nil {
		return nil, nil, err
	}

	return tasks, roles, nil
}

// A function named setup which belongs to the struct JobBuilder is defined, which takes no arguments and returns an error.

func (b *JobBuilder) setup() error {
	// Retrieving data from jobSpec field of the JobBuilder structure.
	spec := b.jobSpec
	// Extracting user ID, design ID, schema version and code version from the JobSpec structure.
	userId, designId, schemaVersion, codeVersion := spec.UserId, spec.DesignId, spec.SchemaVersion, spec.CodeVersion

	// Accessing database service to retrieve the design schema using user ID, design ID and schema version.
	schema, err := b.dbService.GetDesignSchema(userId, designId, schemaVersion)
	if err != nil {
		return err
	}
	// Assigning acquired schema into the schema field of JobBuilder.
	b.schema = schema

	// Getting zipped design code using the user ID, design ID and code version with the help of a database service
	zippedCode, err := b.dbService.GetDesignCode(userId, designId, codeVersion)
	if err != nil {
		return err
	}

	// Create a temporary file with the name util.ProjectName using OS package.
	f, err := os.CreateTemp("", util.ProjectName)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer f.Close()

	// Writing zip code into previously created temporary file.
	if _, err = f.Write(zippedCode); err != nil {
		return fmt.Errorf("failed to save zipped code: %v", err)
	}

	// Unzipping extracted zip code information using unzipFile function call and storing files into slice with fdList.
	fdList, err := util.UnzipFile(f)
	if err != nil {
		return fmt.Errorf("failed to unzip file: %v", err)
	}

	// Creating zip code by top level directory from files present in fdList.
	zippedRoleCode, err := util.ZipFileByTopLevelDir(fdList)
	if err != nil {
		return fmt.Errorf("failed to do zip file by top level directory: %v", err)
	}
	// Saving generated zip code by top-level directory in roleCode field of JobBuilder.
	b.roleCode = zippedRoleCode

	// Iterating for each dataset id to fetch dataset info and update the datasets array.
	for roleName, trainerGrups := range b.jobSpec.DataSpec.FromSystem {
		if len(trainerGrups) == 0 {
			return fmt.Errorf("no dataset group specified for trainer role %s", roleName)
		}

		b.datasets[roleName] = make(map[string][]openapi.DatasetInfo)

		for groupName, datasetIds := range trainerGrups {
			if len(datasetIds) == 0 {
				return fmt.Errorf("no dataset specified for trainer role %s, group %s", roleName, groupName)
			}

			for _, datasetId := range datasetIds {
				datasetInfo, err := b.dbService.GetDatasetById(datasetId)
				if err != nil {
					return err
				}

				b.datasets[roleName][groupName] = append(b.datasets[roleName][groupName], datasetInfo)
			}
		}
	}

	for _, role := range b.schema.Roles {
		b.groupAssociations[role.Name] = role.GroupAssociation
	}

	for i, channel := range b.schema.Channels {
		b.channels[channel.Name] = b.schema.Channels[i]
	}

	// Return nil if there are no errors encountered during the execution of declared functions.
	return nil
}

func (b *JobBuilder) build() ([]objects.Task, []string, error) {
	dataRoles, templates := b.getTaskTemplates()
	if err := b.preCheck(dataRoles, templates); err != nil {
		return nil, nil, err
	}

	var tasks []objects.Task

	for roleName := range templates {
		tmpl := templates[roleName]

		if !tmpl.isDataConsumer {
			var count int
			for i, associations := range b.groupAssociations[roleName] {
				task := tmpl.Task

				task.ComputeId = util.DefaultRealm
				task.Type = openapi.SYSTEM
				task.Key = util.RandString(taskKeyLen)
				task.JobConfig.GroupAssociation = associations

				index := i + count
				count++

				task.GenerateTaskId(index)

				tasks = append(tasks, task)
			}
			continue
		}

		// TODO: this is absolete and should be removed
		for group, count := range b.jobSpec.DataSpec.FromUser {
			for i := 0; i < int(count); i++ {
				task := tmpl.Task

				task.Type = openapi.USER
				task.JobConfig.Realm = group
				task.JobConfig.GroupAssociation = b.getGroupAssociationByGroup(roleName, group)

				task.GenerateTaskId(i)

				tasks = append(tasks, task)
			}
		}

		var count int
		for groupName, datasets := range b.datasets[roleName] {
			for i, dataset := range datasets {
				task := tmpl.Task

				task.ComputeId = dataset.ComputeId
				task.Type = openapi.SYSTEM
				task.Key = util.RandString(taskKeyLen)
				task.JobConfig.DatasetUrl = dataset.Url
				task.JobConfig.GroupAssociation = b.getGroupAssociationByGroup(roleName, groupName)

				index := count + i
				count++

				task.GenerateTaskId(index)

				tasks = append(tasks, task)
			}
		}
	}

	if err := b.postCheck(dataRoles, templates); err != nil {
		return nil, nil, err
	}

	var roles []string

	for _, template := range templates {
		roles = append(roles, template.Role)
	}

	return tasks, roles, nil
}

func (b *JobBuilder) getGroupAssociationByGroup(roleName, groupName string) map[string]string {
	for _, associations := range b.groupAssociations[roleName] {
		for _, association := range associations {
			if association == groupName {
				return associations
			}
		}
	}
	return nil
}

func (b *JobBuilder) getTaskTemplates() ([]string, map[string]*taskTemplate) {
	var dataRoles []string
	templates := make(map[string]*taskTemplate)

	for _, role := range b.schema.Roles {
		template := &taskTemplate{}
		jobConfig := &template.JobConfig

		jobConfig.Configure(b.jobSpec, b.jobParams.Brokers, b.jobParams.Registry, role)

		// check channels and set default group if channels don't have groupBy attributes set
		for i := range b.schema.Channels {
			channel := b.schema.Channels[i]

			if len(channel.GroupBy.Value) == 0 {
				// since there is no groupBy attribute, set default
				jobConfig.Channels[i].GroupBy.Type = groupByTypeTag
				jobConfig.Channels[i].GroupBy.Value = append(jobConfig.Channels[i].GroupBy.Value, defaultGroup)
			}

			jobConfig.Channels = append(jobConfig.Channels, channel)
		}

		template.isDataConsumer = role.IsDataConsumer
		if role.IsDataConsumer {
			dataRoles = append(dataRoles, role.Name)
		}

		template.ZippedCode = b.roleCode[role.Name]
		template.Role = role.Name
		template.JobId = jobConfig.Job.Id

		templates[role.Name] = template
	}

	return dataRoles, templates
}

// preCheck checks sanity of templates
func (b *JobBuilder) preCheck(dataRoles []string, templates map[string]*taskTemplate) error {
	// This function will evolve as more invariants are defined
	// Before processing templates, the following invariants should be met:
	// 1. At least one data consumer role should be defined.
	// 2. a role should be associated with a code.
	// 3. template should be connected.
	// 4. when graph traversal starts at a data role template, the depth of groupBy tag
	//    should strictly decrease from one channel to another.
	// 5. two different data roles cannot be connected directly.

	if len(dataRoles) == 0 {
		return fmt.Errorf("no data consumer role found")
	}

	for _, role := range b.schema.Roles {
		if _, ok := b.roleCode[role.Name]; !ok {
			// rule 1 violated
			return fmt.Errorf("no code found for role %s", role.Name)
		}
	}

	if !b.isTemplatesConnected(templates) {
		// rule 2 violated
		return fmt.Errorf("templates not connected")
	}

	if !b.isConverging(dataRoles, templates) {
		// rule 3 violated
		return fmt.Errorf("groupBy length violated")
	}

	// TODO: implement invariant 4

	return nil
}

func (b *JobBuilder) isTemplatesConnected(templates map[string]*taskTemplate) bool {
	return true
}

func (b *JobBuilder) isConverging(dataRoles []string, templates map[string]*taskTemplate) bool {
	return true
}

func (b *JobBuilder) postCheck(dataRoles []string, templates map[string]*taskTemplate) error {
	// This function will evolve as more invariants are defined
	// At the end of processing templates, the following invariants should be met:
	//

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Task Template related code
////////////////////////////////////////////////////////////////////////////////

type taskTemplate struct {
	objects.Task

	isDataConsumer bool
}
