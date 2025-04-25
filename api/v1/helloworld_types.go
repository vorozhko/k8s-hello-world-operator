package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HelloWorldSpec defines the desired state of HelloWorld
type HelloWorldSpec struct {
	Message string `json:"message,omitempty"`
}

// HelloWorldStatus defines the observed state of HelloWorld
type HelloWorldStatus struct {
	SeenMessage string `json:"seenMessage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// HelloWorld is the Schema for the helloworlds API
type HelloWorld struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelloWorldSpec   `json:"spec,omitempty"`
	Status HelloWorldStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HelloWorldList contains a list of HelloWorld
type HelloWorldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelloWorld `json:"items"`
}

