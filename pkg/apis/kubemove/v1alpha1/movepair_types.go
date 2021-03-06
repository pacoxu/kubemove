package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//======================================== Developer Note ===================================================
// After updating any fields of this file you must run the code generator.
// To prepare code generator/developer environment, run: make dev-image REGISTRY=<your docker registry>
// Then, each time you modify any field, run: make gen REGISTRY=<your docker registry>
// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
//============================================================================================================

// MovePairSpec defines the desired state of MovePair
// +k8s:openapi-gen=true
type MovePairSpec struct {
	// Config specifies the cluster config
	Config clientcmdapi.Config `json:"config"`
}

type PairState string

const (
	PairStateErrored    PairState = "Errored"
	PairStateConnection PairState = "Connected"
)

// MovePairStatus defines the observed state of MovePair
// +k8s:openapi-gen=true
type MovePairStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	State PairState `json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MovePair is the Schema for the movepairs API
// +k8s:openapi-gen=true
type MovePair struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MovePairSpec   `json:"spec,omitempty"`
	Status MovePairStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MovePairList contains a list of MovePair
type MovePairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MovePair `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MovePair{}, &MovePairList{})
}
