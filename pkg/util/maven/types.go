/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package maven

import (
	"encoding/xml"
)

// Integration --
type Integration struct {
	Project     Project
	JavaSources map[string]string
	Resources   map[string]string
	Env         map[string]string // TODO: should we deprecate it ? env are set on deployment
}

// Project represent a maven project
type Project struct {
	XMLName              xml.Name
	XMLNs                string               `xml:"xmlns,attr"`
	XMLNsXsi             string               `xml:"xmlns:xsi,attr"`
	XsiSchemaLocation    string               `xml:"xsi:schemaLocation,attr"`
	ModelVersion         string               `xml:"modelVersion"`
	GroupID              string               `xml:"groupId"`
	ArtifactID           string               `xml:"artifactId"`
	Version              string               `xml:"version"`
	DependencyManagement DependencyManagement `xml:"dependencyManagement"`
	Dependencies         Dependencies         `xml:"dependencies"`
}

// DependencyManagement represent maven's dependency management block
type DependencyManagement struct {
	Dependencies Dependencies `xml:"dependencies"`
}

// Dependencies --
type Dependencies struct {
	Dependencies []Dependency `xml:"dependency"`
}

// Add a dependency to maven's dependencies
func (deps *Dependencies) Add(dep Dependency) {
	deps.Dependencies = append(deps.Dependencies, dep)
}

// AddGAV a dependency to maven's dependencies
func (deps *Dependencies) AddGAV(groupID string, artifactID string, version string) {
	deps.Add(NewDependency(groupID, artifactID, version))
}

// AddEncodedGAV a dependency to maven's dependencies
func (deps *Dependencies) AddEncodedGAV(gav string) {
	if d, err := ParseGAV(gav); err == nil {
		// TODO: error handling
		deps.Add(d)
	}
}

// Dependency represent a maven's dependency
type Dependency struct {
	GroupID    string `xml:"groupId"`
	ArtifactID string `xml:"artifactId"`
	Version    string `xml:"version,omitempty"`
	Type       string `xml:"type,omitempty"`
	Classifier string `xml:"classifier,omitempty"`
	Scope      string `xml:"scope,omitempty"`
}

// NewDependency create an new dependency from the given gav info
func NewDependency(groupID string, artifactID string, version string) Dependency {
	return Dependency{
		GroupID:    groupID,
		ArtifactID: artifactID,
		Version:    version,
		Type:       "jar",
		Classifier: "",
	}
}
