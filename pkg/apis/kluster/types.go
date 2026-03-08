package kluster

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Kluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   KlusterSpec
	Status KlsuterStatus
}

type KlsuterStatus struct {
	KlusterID  string
	Progress   string
	KubeConfig string
}

type KlusterSpec struct {
	Name        string
	Region      string
	Version     string
	TokenSecret string

	NodePools []NodePool `json:"nodePools,omitempty"`
}

type NodePool struct {
	Size  string
	Name  string
	Count int
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KlusterList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Kluster
}
