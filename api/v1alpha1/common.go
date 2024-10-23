/*
Copyright 2024 The KCP Authors.

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
	corev1 "k8s.io/api/core/v1"
)

// ImageSpec defines settings for using a specific image.
type ImageSpec struct {
	// The repository to use for all KCP containers. Defaults to `ghcr.io/kcp-dev/kcp`.
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
	// Optional: A list of secret references that should be used as image pull secrets (e.g. when a private registry is used).
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
}

type RootShardConfig struct {
	Reference *corev1.ObjectReference `json:"ref,omitempty"`
}

type EtcdConfig struct {
	Endpoints []string        `json:"endpoints"`
	Cert      EtcdCertificate `json:"cert"`
}

type EtcdCertificate struct {
	SecretRef corev1.LocalObjectReference `json:"secretRef"`
}
