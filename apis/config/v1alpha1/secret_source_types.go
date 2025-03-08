/*
Copyright AppsCode Inc. and Contributors

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindSecretSource = "SecretSource"
	ResourceSecretSource     = "secretsource"
	ResourceSecretSources    = "secretsources"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=secretsources,singular=secretsource,scope=Cluster,shortName=scsource,categories={meta,virtual-secrets,appscode}
// +kubebuilder:printcolumn:name="SecretManagerType",type="string",JSONPath=".spec.secretManager"
type SecretSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SecretSourceSpec `json:"spec,omitempty"`
}

// SecretSourceSpec defines the desired state of SecretSource
type SecretSourceSpec struct {
	Vault *Vault `json:"vault,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecretSourceList contains a list of SecretSource
type SecretSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretSource `json:"items"`
}

type Vault struct {
	// Connection url to the secret manager
	URL string `json:"url"`

	// Name of the vault role to use for the operator
	// +optional
	RoleName string `json:"roleName,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SecretSource{})
	SchemeBuilder.Register(&SecretSourceList{})
}
