/*


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
	apps "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GroupDeploymentSpec defines the desired state of GroupDeployment
type GroupDeploymentSpec struct {

	// Label selector for managing deployments.
	// Not needed, name can be enough
	// Selector *metav1.LabelSelector `json:"selector"`

	// List of active deployments name.
	// +optional
	ActiveDeployments []string `json:"activeDeployments,omitempty"`

	// Common deployment strategy to use.
	// If empty then fall back to
	// each active deployment.
	// +optional
	Strategy string `json:"strategy,omitempty"`

	// Max concurrent deployment in parallel.
	// Defaults to 1
	// +optional
	// +kubebuilder:default=1
	Concurrency int `json:"concurrency,omitempty"`

	// define order in which deployments
	// should be applied
	// If empty then all deployments are
	// until concurrency is hit.
	// +optional
	Order []string `json:"order,omitempty"`

	// Template lists deployments that will be created.
	Items []GroupDeploymentTemplate `json:"items"`
}

// GroupDeploymentTemplate defines the deployment and name.
type GroupDeploymentTemplate struct {
	// Name of this deployment.
	Name string `json:"name"`
	// Deployment.
	Template apps.Deployment `json:"template"`
}

// GroupDeploymentStatus defines the observed state of GroupDeployment
type GroupDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// GroupDeployment is the Schema for the groupdeployments API
type GroupDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GroupDeploymentSpec   `json:"spec,omitempty"`
	Status GroupDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupDeploymentList contains a list of GroupDeployment
type GroupDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GroupDeployment{}, &GroupDeploymentList{})
}
