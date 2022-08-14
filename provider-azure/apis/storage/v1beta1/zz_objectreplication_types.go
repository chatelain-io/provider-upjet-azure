/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ObjectReplicationObservation struct {

	// The ID of the Object Replication in the destination storage account.
	DestinationObjectReplicationID *string `json:"destinationObjectReplicationId,omitempty" tf:"destination_object_replication_id,omitempty"`

	// The ID of the Storage Object Replication in the destination storage account. It's composed as format source_object_replication_id;destination_object_replication_id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more rules blocks as defined below.
	// +kubebuilder:validation:Required
	Rules []ObjectReplicationRulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// The ID of the Object Replication in the source storage account.
	SourceObjectReplicationID *string `json:"sourceObjectReplicationId,omitempty" tf:"source_object_replication_id,omitempty"`
}

type ObjectReplicationParameters struct {

	// The ID of the destination storage account. Changing this forces a new Storage Object Replication to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DestinationStorageAccountID *string `json:"destinationStorageAccountId,omitempty" tf:"destination_storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationStorageAccountIDRef *v1.Reference `json:"destinationStorageAccountIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DestinationStorageAccountIDSelector *v1.Selector `json:"destinationStorageAccountIdSelector,omitempty" tf:"-"`

	// One or more rules blocks as defined below.
	// +kubebuilder:validation:Required
	Rules []ObjectReplicationRulesParameters `json:"rules" tf:"rules,omitempty"`

	// The ID of the source storage account. Changing this forces a new Storage Object Replication to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceStorageAccountID *string `json:"sourceStorageAccountId,omitempty" tf:"source_storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceStorageAccountIDRef *v1.Reference `json:"sourceStorageAccountIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceStorageAccountIDSelector *v1.Selector `json:"sourceStorageAccountIdSelector,omitempty" tf:"-"`
}

type ObjectReplicationRulesObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ObjectReplicationRulesParameters struct {

	// The time after which the Block Blobs created will be copies to the destination. Possible values are OnlyNewObjects, Everything and time in RFC3339 format: 2006-01-02T15:04:00Z.
	// +kubebuilder:validation:Optional
	CopyBlobsCreatedAfter *string `json:"copyBlobsCreatedAfter,omitempty" tf:"copy_blobs_created_after,omitempty"`

	// The destination storage container name. Changing this forces a new Storage Object Replication to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	DestinationContainerName *string `json:"destinationContainerName,omitempty" tf:"destination_container_name,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationContainerNameRef *v1.Reference `json:"destinationContainerNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DestinationContainerNameSelector *v1.Selector `json:"destinationContainerNameSelector,omitempty" tf:"-"`

	// Specifies a list of filters prefixes, the blobs whose names begin with which will be replicated.
	// +kubebuilder:validation:Optional
	FilterOutBlobsWithPrefix []*string `json:"filterOutBlobsWithPrefix,omitempty" tf:"filter_out_blobs_with_prefix,omitempty"`

	// The source storage container name. Changing this forces a new Storage Object Replication to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	SourceContainerName *string `json:"sourceContainerName,omitempty" tf:"source_container_name,omitempty"`

	// +kubebuilder:validation:Optional
	SourceContainerNameRef *v1.Reference `json:"sourceContainerNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceContainerNameSelector *v1.Selector `json:"sourceContainerNameSelector,omitempty" tf:"-"`
}

// ObjectReplicationSpec defines the desired state of ObjectReplication
type ObjectReplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectReplicationParameters `json:"forProvider"`
}

// ObjectReplicationStatus defines the observed state of ObjectReplication.
type ObjectReplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectReplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectReplication is the Schema for the ObjectReplications API. Copy Block Blobs between a source storage account and a destination account
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ObjectReplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ObjectReplicationSpec   `json:"spec"`
	Status            ObjectReplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectReplicationList contains a list of ObjectReplications
type ObjectReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectReplication `json:"items"`
}

// Repository type metadata.
var (
	ObjectReplication_Kind             = "ObjectReplication"
	ObjectReplication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectReplication_Kind}.String()
	ObjectReplication_KindAPIVersion   = ObjectReplication_Kind + "." + CRDGroupVersion.String()
	ObjectReplication_GroupVersionKind = CRDGroupVersion.WithKind(ObjectReplication_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectReplication{}, &ObjectReplicationList{})
}
