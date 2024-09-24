// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EncryptionSettingsDiskEncryptionKeyInitParameters struct {

	// The URL to the Key Vault Secret used as the Disk Encryption Key. This can be found as id on the azurerm_key_vault_secret resource.
	SecretURL *string `json:"secretUrl,omitempty" tf:"secret_url,omitempty"`

	// The ID of the source Key Vault. This can be found as id on the azurerm_key_vault resource.
	SourceVaultID *string `json:"sourceVaultId,omitempty" tf:"source_vault_id,omitempty"`
}

type EncryptionSettingsDiskEncryptionKeyObservation struct {

	// The URL to the Key Vault Secret used as the Disk Encryption Key. This can be found as id on the azurerm_key_vault_secret resource.
	SecretURL *string `json:"secretUrl,omitempty" tf:"secret_url,omitempty"`

	// The ID of the source Key Vault. This can be found as id on the azurerm_key_vault resource.
	SourceVaultID *string `json:"sourceVaultId,omitempty" tf:"source_vault_id,omitempty"`
}

type EncryptionSettingsDiskEncryptionKeyParameters struct {

	// The URL to the Key Vault Secret used as the Disk Encryption Key. This can be found as id on the azurerm_key_vault_secret resource.
	// +kubebuilder:validation:Optional
	SecretURL *string `json:"secretUrl" tf:"secret_url,omitempty"`

	// The ID of the source Key Vault. This can be found as id on the azurerm_key_vault resource.
	// +kubebuilder:validation:Optional
	SourceVaultID *string `json:"sourceVaultId" tf:"source_vault_id,omitempty"`
}

type EncryptionSettingsKeyEncryptionKeyInitParameters struct {

	// The URL to the Key Vault Key used as the Key Encryption Key. This can be found as id on the azurerm_key_vault_key resource.
	KeyURL *string `json:"keyUrl,omitempty" tf:"key_url,omitempty"`

	// The ID of the source Key Vault. This can be found as id on the azurerm_key_vault resource.
	SourceVaultID *string `json:"sourceVaultId,omitempty" tf:"source_vault_id,omitempty"`
}

type EncryptionSettingsKeyEncryptionKeyObservation struct {

	// The URL to the Key Vault Key used as the Key Encryption Key. This can be found as id on the azurerm_key_vault_key resource.
	KeyURL *string `json:"keyUrl,omitempty" tf:"key_url,omitempty"`

	// The ID of the source Key Vault. This can be found as id on the azurerm_key_vault resource.
	SourceVaultID *string `json:"sourceVaultId,omitempty" tf:"source_vault_id,omitempty"`
}

type EncryptionSettingsKeyEncryptionKeyParameters struct {

	// The URL to the Key Vault Key used as the Key Encryption Key. This can be found as id on the azurerm_key_vault_key resource.
	// +kubebuilder:validation:Optional
	KeyURL *string `json:"keyUrl" tf:"key_url,omitempty"`

	// The ID of the source Key Vault. This can be found as id on the azurerm_key_vault resource.
	// +kubebuilder:validation:Optional
	SourceVaultID *string `json:"sourceVaultId" tf:"source_vault_id,omitempty"`
}

type SnapshotEncryptionSettingsInitParameters struct {

	// A disk_encryption_key block as defined below.
	DiskEncryptionKey *EncryptionSettingsDiskEncryptionKeyInitParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A key_encryption_key block as defined below.
	KeyEncryptionKey *EncryptionSettingsKeyEncryptionKeyInitParameters `json:"keyEncryptionKey,omitempty" tf:"key_encryption_key,omitempty"`
}

type SnapshotEncryptionSettingsObservation struct {

	// A disk_encryption_key block as defined below.
	DiskEncryptionKey *EncryptionSettingsDiskEncryptionKeyObservation `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A key_encryption_key block as defined below.
	KeyEncryptionKey *EncryptionSettingsKeyEncryptionKeyObservation `json:"keyEncryptionKey,omitempty" tf:"key_encryption_key,omitempty"`
}

type SnapshotEncryptionSettingsParameters struct {

	// A disk_encryption_key block as defined below.
	// +kubebuilder:validation:Optional
	DiskEncryptionKey *EncryptionSettingsDiskEncryptionKeyParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A key_encryption_key block as defined below.
	// +kubebuilder:validation:Optional
	KeyEncryptionKey *EncryptionSettingsKeyEncryptionKeyParameters `json:"keyEncryptionKey,omitempty" tf:"key_encryption_key,omitempty"`
}

type SnapshotInitParameters struct {

	// Indicates how the snapshot is to be created. Possible values are Copy or Import.
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// Specifies the ID of the Disk Access which should be used for this Snapshot. This is used in conjunction with setting network_access_policy to AllowPrivate.
	DiskAccessID *string `json:"diskAccessId,omitempty" tf:"disk_access_id,omitempty"`

	// The size of the Snapshotted Disk in GB.
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// A encryption_settings block as defined below.
	EncryptionSettings *SnapshotEncryptionSettingsInitParameters `json:"encryptionSettings,omitempty" tf:"encryption_settings,omitempty"`

	// Specifies if the Snapshot is incremental. Changing this forces a new resource to be created.
	IncrementalEnabled *bool `json:"incrementalEnabled,omitempty" tf:"incremental_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Policy for accessing the disk via network. Possible values are AllowAll, AllowPrivate, or DenyAll. Defaults to AllowAll.
	NetworkAccessPolicy *string `json:"networkAccessPolicy,omitempty" tf:"network_access_policy,omitempty"`

	// Policy for controlling export on the disk. Possible values are true or false. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies a reference to an existing snapshot, when create_option is Copy. Changing this forces a new resource to be created.
	SourceResourceID *string `json:"sourceResourceId,omitempty" tf:"source_resource_id,omitempty"`

	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta2.ManagedDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// Reference to a ManagedDisk in compute to populate sourceUri.
	// +kubebuilder:validation:Optional
	SourceURIRef *v1.Reference `json:"sourceUriRef,omitempty" tf:"-"`

	// Selector for a ManagedDisk in compute to populate sourceUri.
	// +kubebuilder:validation:Optional
	SourceURISelector *v1.Selector `json:"sourceUriSelector,omitempty" tf:"-"`

	// Specifies the ID of an storage account. Used with source_uri to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SnapshotObservation struct {

	// Indicates how the snapshot is to be created. Possible values are Copy or Import.
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// Specifies the ID of the Disk Access which should be used for this Snapshot. This is used in conjunction with setting network_access_policy to AllowPrivate.
	DiskAccessID *string `json:"diskAccessId,omitempty" tf:"disk_access_id,omitempty"`

	// The size of the Snapshotted Disk in GB.
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// A encryption_settings block as defined below.
	EncryptionSettings *SnapshotEncryptionSettingsObservation `json:"encryptionSettings,omitempty" tf:"encryption_settings,omitempty"`

	// The Snapshot ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies if the Snapshot is incremental. Changing this forces a new resource to be created.
	IncrementalEnabled *bool `json:"incrementalEnabled,omitempty" tf:"incremental_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Policy for accessing the disk via network. Possible values are AllowAll, AllowPrivate, or DenyAll. Defaults to AllowAll.
	NetworkAccessPolicy *string `json:"networkAccessPolicy,omitempty" tf:"network_access_policy,omitempty"`

	// Policy for controlling export on the disk. Possible values are true or false. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies a reference to an existing snapshot, when create_option is Copy. Changing this forces a new resource to be created.
	SourceResourceID *string `json:"sourceResourceId,omitempty" tf:"source_resource_id,omitempty"`

	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// Specifies the ID of an storage account. Used with source_uri to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether Trusted Launch is enabled for the Snapshot.
	TrustedLaunchEnabled *bool `json:"trustedLaunchEnabled,omitempty" tf:"trusted_launch_enabled,omitempty"`
}

type SnapshotParameters struct {

	// Indicates how the snapshot is to be created. Possible values are Copy or Import.
	// +kubebuilder:validation:Optional
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// Specifies the ID of the Disk Access which should be used for this Snapshot. This is used in conjunction with setting network_access_policy to AllowPrivate.
	// +kubebuilder:validation:Optional
	DiskAccessID *string `json:"diskAccessId,omitempty" tf:"disk_access_id,omitempty"`

	// The size of the Snapshotted Disk in GB.
	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// A encryption_settings block as defined below.
	// +kubebuilder:validation:Optional
	EncryptionSettings *SnapshotEncryptionSettingsParameters `json:"encryptionSettings,omitempty" tf:"encryption_settings,omitempty"`

	// Specifies if the Snapshot is incremental. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IncrementalEnabled *bool `json:"incrementalEnabled,omitempty" tf:"incremental_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Policy for accessing the disk via network. Possible values are AllowAll, AllowPrivate, or DenyAll. Defaults to AllowAll.
	// +kubebuilder:validation:Optional
	NetworkAccessPolicy *string `json:"networkAccessPolicy,omitempty" tf:"network_access_policy,omitempty"`

	// Policy for controlling export on the disk. Possible values are true or false. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies a reference to an existing snapshot, when create_option is Copy. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceResourceID *string `json:"sourceResourceId,omitempty" tf:"source_resource_id,omitempty"`

	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta2.ManagedDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// Reference to a ManagedDisk in compute to populate sourceUri.
	// +kubebuilder:validation:Optional
	SourceURIRef *v1.Reference `json:"sourceUriRef,omitempty" tf:"-"`

	// Selector for a ManagedDisk in compute to populate sourceUri.
	// +kubebuilder:validation:Optional
	SourceURISelector *v1.Selector `json:"sourceUriSelector,omitempty" tf:"-"`

	// Specifies the ID of an storage account. Used with source_uri to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
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
	InitProvider SnapshotInitParameters `json:"initProvider,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Snapshot is the Schema for the Snapshots API. Manages a Disk Snapshot.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.createOption) || (has(self.initProvider) && has(self.initProvider.createOption))",message="spec.forProvider.createOption is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   SnapshotSpec   `json:"spec"`
	Status SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
