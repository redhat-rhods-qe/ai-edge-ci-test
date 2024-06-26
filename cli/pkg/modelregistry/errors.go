/*
Copyright 2024. Open Data Hub Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package modelregistry

import (
	"errors"
	"strings"

	"github.com/kubeflow/model-registry/pkg/openapi"
)

var (
	// ErrModelExists is returned when a model already exists
	ErrModelExists = errors.New("model already exists")
	// ErrModelNotFound is returned when a model is not found
	ErrModelNotFound = errors.New("no registered model found")

	// ErrModelNotFoundForName is returned when a model is not found for a given name using FindRegisteredModel
	ErrModelNotFoundForName = errors.New("no registered models found")

	// ErrVersionExists is returned when a version already exists for a given model
	ErrVersionExists = errors.New("version already exists")
	// ErrVersionNotFound is returned when a version is not found
	ErrVersionNotFound = errors.New("no model version found")
	// ErrVersionNotFoundForName is returned when a version is not found for a given name & model ID using FindModelVersion
	ErrVersionNotFoundForName = errors.New("no model versions found")

	// ErrArtifactExists is returned when a model version artifact already exists for a given model version
	ErrArtifactExists = errors.New("artifact already exists")

	// ErrAlreadyExists is a generic error to check the model registry returned errors when an entity (Model, Version,
	// Artifact) already exists
	ErrAlreadyExists = errors.New("already exists")
)

// isOpenAPIErrorOfKind checks if the error is of the given kind (targetErr). It checks if sourceErr is an
// openapi.GenericOpenAPIError and if the error message contains the targetErr message.
//
// This is a workaround to handle the error until the model registry supports returning standard HTTP status codes for
// errors with known status codes.
func isOpenAPIErrorOfKind(sourceErr, targetErr error) bool {
	var e *openapi.GenericOpenAPIError
	if errors.As(sourceErr, &e) {
		if me, ok := e.Model().(openapi.Error); ok {
			if msg, ok := me.GetMessageOk(); ok {
				return strings.Contains(*msg, targetErr.Error())
			}
		}
	}
	return false
}
