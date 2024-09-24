// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MonitorPrivateLinkScopeInitParameters struct {

	// The default ingestion access mode for the associated private endpoints in scope. Possible values are Open and PrivateOnly. Defaults to Open.
	IngestionAccessMode *string `json:"ingestionAccessMode,omitempty" tf:"ingestion_access_mode,omitempty"`

	// The default query access mode for hte associated private endpoints in scope. Possible values are Open and PrivateOnly. Defaults to Open.
	QueryAccessMode *string `json:"queryAccessMode,omitempty" tf:"query_access_mode,omitempty"`

	// A mapping of tags which should be assigned to the Azure Monitor Private Link Scope.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorPrivateLinkScopeObservation struct {

	// The ID of the Azure Monitor Private Link Scope.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The default ingestion access mode for the associated private endpoints in scope. Possible values are Open and PrivateOnly. Defaults to Open.
	IngestionAccessMode *string `json:"ingestionAccessMode,omitempty" tf:"ingestion_access_mode,omitempty"`

	// The default query access mode for hte associated private endpoints in scope. Possible values are Open and PrivateOnly. Defaults to Open.
	QueryAccessMode *string `json:"queryAccessMode,omitempty" tf:"query_access_mode,omitempty"`

	// The name of the Resource Group where the Azure Monitor Private Link Scope should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the Azure Monitor Private Link Scope.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorPrivateLinkScopeParameters struct {

	// The default ingestion access mode for the associated private endpoints in scope. Possible values are Open and PrivateOnly. Defaults to Open.
	// +kubebuilder:validation:Optional
	IngestionAccessMode *string `json:"ingestionAccessMode,omitempty" tf:"ingestion_access_mode,omitempty"`

	// The default query access mode for hte associated private endpoints in scope. Possible values are Open and PrivateOnly. Defaults to Open.
	// +kubebuilder:validation:Optional
	QueryAccessMode *string `json:"queryAccessMode,omitempty" tf:"query_access_mode,omitempty"`

	// The name of the Resource Group where the Azure Monitor Private Link Scope should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Azure Monitor Private Link Scope.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// MonitorPrivateLinkScopeSpec defines the desired state of MonitorPrivateLinkScope
type MonitorPrivateLinkScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorPrivateLinkScopeParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MonitorPrivateLinkScopeInitParameters `json:"initProvider,omitempty"`
}

// MonitorPrivateLinkScopeStatus defines the observed state of MonitorPrivateLinkScope.
type MonitorPrivateLinkScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorPrivateLinkScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MonitorPrivateLinkScope is the Schema for the MonitorPrivateLinkScopes API. Manages an Azure Monitor Private Link Scope
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorPrivateLinkScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorPrivateLinkScopeSpec   `json:"spec"`
	Status            MonitorPrivateLinkScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorPrivateLinkScopeList contains a list of MonitorPrivateLinkScopes
type MonitorPrivateLinkScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorPrivateLinkScope `json:"items"`
}

// Repository type metadata.
var (
	MonitorPrivateLinkScope_Kind             = "MonitorPrivateLinkScope"
	MonitorPrivateLinkScope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorPrivateLinkScope_Kind}.String()
	MonitorPrivateLinkScope_KindAPIVersion   = MonitorPrivateLinkScope_Kind + "." + CRDGroupVersion.String()
	MonitorPrivateLinkScope_GroupVersionKind = CRDGroupVersion.WithKind(MonitorPrivateLinkScope_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorPrivateLinkScope{}, &MonitorPrivateLinkScopeList{})
}
