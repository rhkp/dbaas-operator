/*
Copyright 2021.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PlatformsName name of platform
type PlatformsName string

// PlatformsInstlnStatus status of platform installation
type PlatformsInstlnStatus string

// PlatformsType platform type
type PlatformsType int

// Supported platforms
const (
	CrunchyBridgeInstallation      PlatformsName = "crunchy-bridge"
	MongoDBAtlasInstallation       PlatformsName = "mongodb-atlas"
	DBaaSDynamicPluginInstallation PlatformsName = "dbaas-dynamic-plugin"
	CockroachDBInstallation        PlatformsName = "cockroachdb-cloud"
	ObservabilityInstallation      PlatformsName = "observability"
	DBaaSQuickStartInstallation    PlatformsName = "dbaas-quick-starts"
	RDSProviderInstallation        PlatformsName = "rds-provider"
)

// Platform types
const (
	TypeQuickStart PlatformsType = iota
	TypeConsolePlugin
	TypeOperator
)

// Platform status values
const (
	ResultSuccess    PlatformsInstlnStatus = "success"
	ResultFailed     PlatformsInstlnStatus = "failed"
	ResultInProgress PlatformsInstlnStatus = "in progress"
)

// PlatformConfig defines parameters for a platform
type PlatformConfig struct {
	Name           string
	CSV            string
	DeploymentName string
	Image          string
	PackageName    string
	Channel        string
	DisplayName    string
	Envs           []corev1.EnvVar
	Type           PlatformsType
}

// ObservabilityConfig defines parameters for a observatorium
type ObservabilityConfig struct {
	AuthType        string
	RemoteWritesURL string
	RHSSOTokenURL   string
	AddonName       string
	RHOBSSecretName string
}

// DBaaSPlatformSpec defines the desired state of DBaaSPlatform
type DBaaSPlatformSpec struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1440
	// The SyncPeriod set The minimum interval at which the provider operator controllers reconcile, the default value is 180 minutes.
	SyncPeriod *int `json:"syncPeriod,omitempty"`
}

// DBaaSPlatformStatus defines the observed state of DBaaSPlatform
type DBaaSPlatformStatus struct {
	PlatformName   PlatformsName         `json:"platformName"`
	PlatformStatus PlatformsInstlnStatus `json:"platformStatus"`
	LastMessage    string                `json:"lastMessage,omitempty"`
	Conditions     []metav1.Condition    `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DBaaSPlatform is the Schema for the dbaasplatforms API
//+operator-sdk:csv:customresourcedefinitions:displayName="DBaaSPlatform"
type DBaaSPlatform struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DBaaSPlatformSpec   `json:"spec,omitempty"`
	Status DBaaSPlatformStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DBaaSPlatformList contains a list of DBaaSPlatform
type DBaaSPlatformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBaaSPlatform `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBaaSPlatform{}, &DBaaSPlatformList{})
}
