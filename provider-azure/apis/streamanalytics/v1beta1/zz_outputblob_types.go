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

type OutputBlobObservation struct {

	// The ID of the Stream Analytics Output Blob Storage.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OutputBlobParameters struct {

	// The maximum wait time per batch in hh:mm:ss e.g. 00:02:00 for two minutes.
	// +kubebuilder:validation:Optional
	BatchMaxWaitTime *string `json:"batchMaxWaitTime,omitempty" tf:"batch_max_wait_time,omitempty"`

	// The minimum number of rows per batch .
	// +kubebuilder:validation:Optional
	BatchMinRows *float64 `json:"batchMinRows,omitempty" tf:"batch_min_rows,omitempty"`

	// The date format. Wherever {date} appears in path_pattern, the value of this property is used as the date format instead.
	// +kubebuilder:validation:Required
	DateFormat *string `json:"dateFormat" tf:"date_format,omitempty"`

	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	// +kubebuilder:validation:Required
	PathPattern *string `json:"pathPattern" tf:"path_pattern,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A serialization block as defined below.
	// +kubebuilder:validation:Required
	Serialization []SerializationParameters `json:"serialization" tf:"serialization,omitempty"`

	// The Access Key which should be used to connect to this Storage Account.
	// +kubebuilder:validation:Required
	StorageAccountKeySecretRef v1.SecretKeySelector `json:"storageAccountKeySecretRef" tf:"-"`

	// The name of the Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`

	// The name of the Container within the Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageContainerNameRef *v1.Reference `json:"storageContainerNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageContainerNameSelector *v1.Selector `json:"storageContainerNameSelector,omitempty" tf:"-"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name,omitempty"`

	// The time format. Wherever {time} appears in path_pattern, the value of this property is used as the time format instead.
	// +kubebuilder:validation:Required
	TimeFormat *string `json:"timeFormat" tf:"time_format,omitempty"`
}

type SerializationObservation struct {
}

type SerializationParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value  records. Possible values are   , , ,     , |  and ;.
	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// Specifies the format of the JSON the output will be written in. Possible values are Array and LineSeparated.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The serialization format used for outgoing data streams. Possible values are Avro, Csv, Json and Parquet.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// OutputBlobSpec defines the desired state of OutputBlob
type OutputBlobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputBlobParameters `json:"forProvider"`
}

// OutputBlobStatus defines the observed state of OutputBlob.
type OutputBlobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputBlobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputBlob is the Schema for the OutputBlobs API. Manages a Stream Analytics Output to Blob Storage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputBlob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OutputBlobSpec   `json:"spec"`
	Status            OutputBlobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputBlobList contains a list of OutputBlobs
type OutputBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputBlob `json:"items"`
}

// Repository type metadata.
var (
	OutputBlob_Kind             = "OutputBlob"
	OutputBlob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputBlob_Kind}.String()
	OutputBlob_KindAPIVersion   = OutputBlob_Kind + "." + CRDGroupVersion.String()
	OutputBlob_GroupVersionKind = CRDGroupVersion.WithKind(OutputBlob_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputBlob{}, &OutputBlobList{})
}
