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

package code

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cisco-open/flame/cmd/flamectl/resources"
	"github.com/cisco-open/flame/pkg/openapi/constants"
	"github.com/cisco-open/flame/pkg/restapi"
	"github.com/cisco-open/flame/pkg/util"
)

type Params struct {
	resources.CommonParams

	DesignId string
	CodePath string
	CodeVer  string
}

func Create(params Params) error {
	// construct URL
	uriMap := map[string]string{
		constants.ParamUser:     params.User,
		constants.ParamDesignID: params.DesignId,
	}
	url := restapi.CreateURL(params.Endpoint, restapi.CreateDesignCodeEndPoint, uriMap)

	// "fileName", "fileVer" and "fileData" are names of variables used in openapi specification
	kv := map[string]io.Reader{
		"fileName": strings.NewReader(filepath.Base(params.CodePath)),
		"fileData": mustOpen(params.CodePath),
	}

	// create multipart/form-data
	buf, writer, err := restapi.CreateMultipartFormData(kv)
	if err != nil {
		fmt.Printf("Failed to create multipart/form-data: %v\n", err)
		return nil
	}

	// send post request
	resp, err := http.Post(url, writer.FormDataContentType(), buf)
	if err != nil {
		var msg string
		body, _ := io.ReadAll(resp.Body)
		_ = json.Unmarshal(body, &msg)
		fmt.Printf("Failed to create a code: %s\n", msg)
		return nil
	}
	defer resp.Body.Close()

	if err != nil || restapi.CheckStatusCode(resp.StatusCode) != nil {
		var msg string
		body, _ := io.ReadAll(resp.Body)
		_ = json.Unmarshal(body, &msg)
		fmt.Printf("Failed to create a code - code: %d; %s\n", resp.StatusCode, msg)
		return nil
	}

	fmt.Printf("Code created successfully for design '%s'\n", params.DesignId)

	return nil
}

func Get(params Params) error {
	// construct URL
	uriMap := map[string]string{
		constants.ParamUser:     params.User,
		constants.ParamDesignID: params.DesignId,
		constants.ParamVersion:  params.CodeVer,
	}
	url := restapi.CreateURL(params.Endpoint, restapi.GetDesignCodeEndPoint, uriMap)

	code, body, err := restapi.HTTPGet(url)
	if err != nil || restapi.CheckStatusCode(code) != nil {
		var msg string
		_ = json.Unmarshal(body, &msg)
		fmt.Printf("Failed to retrieve a code of version '%s' - code: %d; %s\n", params.CodeVer, code, msg)
		return nil
	}

	fileName := params.DesignId + "_ver-" + params.CodeVer + ".zip"
	err = os.WriteFile(fileName, body, util.FilePerm0644)
	if err != nil {
		fmt.Printf("Failed to save %s: %v\n", fileName, err)
		return nil
	}

	fmt.Printf("Downloaded %s successfully\n", fileName)

	return nil
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", f, err)
	}

	return r
}

func Remove(params Params) error {
	// construct URL
	uriMap := map[string]string{
		constants.ParamUser:     params.User,
		constants.ParamDesignID: params.DesignId,
		constants.ParamVersion:  params.CodeVer,
	}
	url := restapi.CreateURL(params.Endpoint, restapi.DeleteDesignCodeEndPoint, uriMap)

	statusCode, body, err := restapi.HTTPDelete(url, nil, "")
	if err != nil || restapi.CheckStatusCode(statusCode) != nil {
		var msg string
		_ = json.Unmarshal(body, &msg)
		fmt.Printf("Failed to delete code version '%s' - code: %d; %s\n", params.CodeVer, statusCode, msg)
		return nil
	}

	fmt.Printf("Deleted code version '%s' successfully\n", params.CodeVer)

	return nil
}
