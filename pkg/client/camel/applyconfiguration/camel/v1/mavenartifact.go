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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// MavenArtifactApplyConfiguration represents an declarative configuration of the MavenArtifact type for use
// with apply.
type MavenArtifactApplyConfiguration struct {
	GroupID    *string `json:"groupId,omitempty"`
	ArtifactID *string `json:"artifactId,omitempty"`
	Type       *string `json:"type,omitempty"`
	Version    *string `json:"version,omitempty"`
	Classifier *string `json:"classifier,omitempty"`
}

// MavenArtifactApplyConfiguration constructs an declarative configuration of the MavenArtifact type for use with
// apply.
func MavenArtifact() *MavenArtifactApplyConfiguration {
	return &MavenArtifactApplyConfiguration{}
}

// WithGroupID sets the GroupID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GroupID field is set to the value of the last call.
func (b *MavenArtifactApplyConfiguration) WithGroupID(value string) *MavenArtifactApplyConfiguration {
	b.GroupID = &value
	return b
}

// WithArtifactID sets the ArtifactID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ArtifactID field is set to the value of the last call.
func (b *MavenArtifactApplyConfiguration) WithArtifactID(value string) *MavenArtifactApplyConfiguration {
	b.ArtifactID = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *MavenArtifactApplyConfiguration) WithType(value string) *MavenArtifactApplyConfiguration {
	b.Type = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *MavenArtifactApplyConfiguration) WithVersion(value string) *MavenArtifactApplyConfiguration {
	b.Version = &value
	return b
}

// WithClassifier sets the Classifier field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Classifier field is set to the value of the last call.
func (b *MavenArtifactApplyConfiguration) WithClassifier(value string) *MavenArtifactApplyConfiguration {
	b.Classifier = &value
	return b
}
