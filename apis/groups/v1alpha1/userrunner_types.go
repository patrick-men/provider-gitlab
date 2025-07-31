/*
Copyright 2021 The Crossplane Authors.
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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UserRunnerParameters define the desired state of a User Runner.
type UserRunnerParameters struct {
  // AccessLevel is the access level of the user runner.
	// +kubebuilder:validation:Enum=not_protected;ref_protected
	// +optional
	AccessLevel string `json:"accessLevel,omitempty"`

	// Description is a description of the user runner.
	// +optional
	Description string `json:"description,omitempty"`

	// GroupID is the ID of the group that the user runner belongs to.
	// +optional
	GroupID *int `json:"groupId,omitempty"`

	// GroupIDRef is a reference to a Group used to populate groupId.
	// +optional
	GroupIDRef *xpv1.Reference `json:"groupIdRef,omitempty"`

	// GroupIDSelector selects a reference to a Group used to populate groupId.
	// +optional
	GroupIDSelector *xpv1.Selector `json:"groupIdSelector,omitempty"`

	// Locked indicates whether the user runner is locked.
	// +required
  Locked bool `json:"locked"`

	// MaintenanceNote is a note that will be displayed on the user runner page when the runner is in maintenance mode. Free-form, 1024 character limit.
	// +optional
	MaintenanceNote string `json:"maintenanceNote,omitempty"`

	// MaximumTimeout is the maximum amount of time a job can run for (minimum is 600 seconds).
	// +optional
	// +kubebuilder:validation:Minimum=600
	MaximumTimeout *int `json:"maximumTimeout,omitempty"`

  // Paused indicates whether the user runner should ignore new jobs.
  // +optional
	Paused bool `json:"paused,omitempty"`

  // ProjectID is the ID of the project that the user runner belongs to.
	// +optional
	ProjectID *int `json:"projectId,omitempty"`

	// ProjectIDRef is a reference to a Project used to populate projectId.
	// +optional
	ProjectIDRef *xpv1.Reference `json:"projectIdRef,omitempty"`

	// ProjectIDSelector selects a reference to a Project used to populate projectId.
	// +optional
	ProjectIDSelector *xpv1.Selector `json:"projectIdSelector,omitempty"`

	// RunnerType indicated the type of the user runner - either group or project.
	// +kubebuilder:validation:Enum=group;project
	// +required
	RunnerType string `json:"runnerType,omitempty"`

	// RunUntagged indicates whether the user runner can run untagged jobs.
	// +optional
	RunUntagged bool `json:"runUntagged,omitempty"`

	// TagList are a list of tagList to be applied to the user runner.
	// +optional
	TagList []string `json:tagList,omitempty`"`
}

// Shared fields for User Runner
type UserRunner struct {
  
}

// UserRunnerObservation is the observed state of a User Runner.
type UserRunnerObservation struct {
  ID            *int 			`json:"id,omitempty"`
  Status			  string 		`json:"status,omitempty"`
  Version       string 		`json:"version,omitempty"`
}

// GroupsRunnerSpec defines the desired state of a User Runner.
type UserRunnerSpec struct {
	xpv1.ResourceSpec 					`json:",inline"`
	ForProvider UserRunnerParameters `json:"forProvider"`
}

// A UserRunnerStatus represents the observed state of a User Runner.
type UserRunnerStatus struct {
	xpv1.ResourceStatus 							`json:",inline"`
	AtProvider UserRunnerObservation 	`json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserRunnerList contains a list of group items
type UserRunnerList struct {
	metav1.TypeMeta 		`json:",inline"`
	metav1.ListMeta 		`json:"metadata,omitempty"`
	Items []UserRunner  `json:"items"`
}
