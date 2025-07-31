// copy paste from MR - only thing to change is GenerateAddLdapGroupLinkOptions
//
// ---
//
// try to find the user-runner interface in the go sdk from gitlab
// -> runners.go
// -> users.go
//
// cmd+f for runner
//

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

package groups

import (
	"strings"
	"time"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/crossplane-contrib/provider-gitlab/apis/groups/v1alpha1"
	"github.com/crossplane-contrib/provider-gitlab/pkg/clients"
)

const (
	errGroupNotFound = "404 Group Not Found"
)

// Client defined User Runer operations
type Client interface {
	CreateUserRunner(opts *CreateUserRunnerOptions, options ...RequestOptionFunc) (*UserRunner, *Response, error)
}

// NewUserRunnerClient returns a new UserRunner service
func NewUserRunnerClient(c gitlab.Client) Client {
  git := clients.NewClient(cfg)
  return git.UserRunners
}

func IsErrorUserRunnerNotFound(err error) bool {
  if err == nil {
    return false
  }
  return strings.Contains(err.Error(), errGroupNotFound)
}

//TODO: get this one right, compare to group
func GenerateUserRunnerObservation(ur *UserRunner) v1alpha1.UserRunnerObservation {
  if ur == nil {
    return v1alpha1.UserRunnerObservation{}
  }

  return v1alpha1.UserRunnerObservation{
    ID:                ur.ID,
    Description:       ur.Description,
    AccessLevel:       ur.AccessLevel,
    GroupID:           ur.GroupID,
    CreatedAt:         metav1.NewTime(ur.CreatedAt),
    LastUsedAt:        metav1.NewTime(ur.LastUsedAt),
    MaintenanceNote:   ur.MaintenanceNote,
    MaximumTimeout:    ur.MaximumTimeout,
    PendingDelete:     ur.PendingDelete,
    Paused:            ur.Paused,
    Locked:            ur.Locked,
  }
}
