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

type AccessPolicyObservation struct {

	// The object ID of an Application in Azure Active Directory.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// List of certificate permissions, must be one or more from the following: Backup, Create, Delete, DeleteIssuers, Get, GetIssuers, Import, List, ListIssuers, ManageContacts, ManageIssuers, Purge, Recover, Restore, SetIssuers and Update.
	CertificatePermissions []*string `json:"certificatePermissions,omitempty" tf:"certificate_permissions,omitempty"`

	// List of key permissions, must be one or more from the following: Backup, Create, Decrypt, Delete, Encrypt, Get, Import, List, Purge, Recover, Restore, Sign, UnwrapKey, Update, Verify and WrapKey.
	KeyPermissions []*string `json:"keyPermissions,omitempty" tf:"key_permissions,omitempty"`

	// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// List of secret permissions, must be one or more from the following: Backup, Delete, Get, List, Purge, Recover, Restore and Set.
	SecretPermissions []*string `json:"secretPermissions,omitempty" tf:"secret_permissions,omitempty"`

	// List of storage permissions, must be one or more from the following: Backup, Delete, DeleteSAS, Get, GetSAS, List, ListSAS, Purge, Recover, RegenerateKey, Restore, Set, SetSAS and Update.
	StoragePermissions []*string `json:"storagePermissions,omitempty" tf:"storage_permissions,omitempty"`

	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault. Must match the tenant_id used above.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AccessPolicyParameters struct {
}

type ContactObservation struct {
}

type ContactParameters struct {

	// E-mail address of the contact.
	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`

	// Name of the contact.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Phone number of the contact.
	// +kubebuilder:validation:Optional
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`
}

type NetworkAclsObservation struct {
}

type NetworkAclsParameters struct {

	// Specifies which traffic can bypass the network rules. Possible values are AzureServices and None.
	// +kubebuilder:validation:Required
	Bypass *string `json:"bypass" tf:"bypass,omitempty"`

	// The Default Action to use when no rules match from ip_rules / virtual_network_subnet_ids. Possible values are Allow and Deny.
	// +kubebuilder:validation:Required
	DefaultAction *string `json:"defaultAction" tf:"default_action,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the Key Vault.
	// +kubebuilder:validation:Optional
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// One or more Subnet IDs which should be able to access this Key Vault.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIds []*string `json:"virtualNetworkSubnetIds,omitempty" tf:"virtual_network_subnet_ids,omitempty"`
}

type VaultObservation struct {

	// A list of up to 16 objects describing access policies, as described below.
	AccessPolicy []AccessPolicyObservation `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The ID of the Key Vault.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URI of the Key Vault, used for performing operations on keys and secrets.
	VaultURI *string `json:"vaultUri,omitempty" tf:"vault_uri,omitempty"`
}

type VaultParameters struct {

	// One or more contact block as defined below.
	// +kubebuilder:validation:Optional
	Contact []ContactParameters `json:"contact,omitempty" tf:"contact,omitempty"`

	// Boolean flag to specify whether Azure Key Vault uses Role Based Access Control (RBAC) for authorization of data actions. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableRbacAuthorization *bool `json:"enableRbacAuthorization,omitempty" tf:"enable_rbac_authorization,omitempty"`

	// Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to false.
	// +kubebuilder:validation:Optional
	EnabledForDeployment *bool `json:"enabledForDeployment,omitempty" tf:"enabled_for_deployment,omitempty"`

	// Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to false.
	// +kubebuilder:validation:Optional
	EnabledForDiskEncryption *bool `json:"enabledForDiskEncryption,omitempty" tf:"enabled_for_disk_encryption,omitempty"`

	// Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to false.
	// +kubebuilder:validation:Optional
	EnabledForTemplateDeployment *bool `json:"enabledForTemplateDeployment,omitempty" tf:"enabled_for_template_deployment,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// A network_acls block as defined below.
	// +kubebuilder:validation:Optional
	NetworkAcls []NetworkAclsParameters `json:"networkAcls,omitempty" tf:"network_acls,omitempty"`

	// Is Purge Protection enabled for this Key Vault? Defaults to false.
	// +kubebuilder:validation:Optional
	PurgeProtectionEnabled *bool `json:"purgeProtectionEnabled,omitempty" tf:"purge_protection_enabled,omitempty"`

	// The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Name of the SKU used for this Key Vault. Possible values are standard and premium.
	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// The number of days that items should be retained for once soft-deleted. This value can be between 7 and 90 (the default) days.
	// +kubebuilder:validation:Optional
	SoftDeleteRetentionDays *float64 `json:"softDeleteRetentionDays,omitempty" tf:"soft_delete_retention_days,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

// VaultSpec defines the desired state of Vault
type VaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultParameters `json:"forProvider"`
}

// VaultStatus defines the observed state of Vault.
type VaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vault is the Schema for the Vaults API. Manages a Key Vault.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Vault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultSpec   `json:"spec"`
	Status            VaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultList contains a list of Vaults
type VaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vault `json:"items"`
}

// Repository type metadata.
var (
	Vault_Kind             = "Vault"
	Vault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vault_Kind}.String()
	Vault_KindAPIVersion   = Vault_Kind + "." + CRDGroupVersion.String()
	Vault_GroupVersionKind = CRDGroupVersion.WithKind(Vault_Kind)
)

func init() {
	SchemeBuilder.Register(&Vault{}, &VaultList{})
}
