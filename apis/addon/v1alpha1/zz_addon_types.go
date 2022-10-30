/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AddonObservation struct {

	// The ID of the add-on.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AddonParameters struct {

	// The name of the add-on.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The source URL to display in a frame in the PagerDuty UI. HTTPS is required.
	// +kubebuilder:validation:Required
	Src *string `json:"src" tf:"src,omitempty"`
}

// AddonSpec defines the desired state of Addon
type AddonSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddonParameters `json:"forProvider"`
}

// AddonStatus defines the observed state of Addon.
type AddonStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddonObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Addon is the Schema for the Addons API. Creates and manages an add-on in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Addon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddonSpec   `json:"spec"`
	Status            AddonStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddonList contains a list of Addons
type AddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Addon `json:"items"`
}

// Repository type metadata.
var (
	Addon_Kind             = "Addon"
	Addon_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Addon_Kind}.String()
	Addon_KindAPIVersion   = Addon_Kind + "." + CRDGroupVersion.String()
	Addon_GroupVersionKind = CRDGroupVersion.WithKind(Addon_Kind)
)

func init() {
	SchemeBuilder.Register(&Addon{}, &AddonList{})
}