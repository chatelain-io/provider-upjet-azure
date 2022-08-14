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

type AccessRuleObservation struct {
}

type AccessRuleParameters struct {

	// The access level for this rule. Possible values are: rw, ro, no.
	// +kubebuilder:validation:Required
	Access *string `json:"access" tf:"access,omitempty"`

	// The anonymous GID used when root_squash_enabled is true.
	// +kubebuilder:validation:Optional
	AnonymousGID *float64 `json:"anonymousGid,omitempty" tf:"anonymous_gid,omitempty"`

	// The anonymous UID used when root_squash_enabled is true.
	// +kubebuilder:validation:Optional
	AnonymousUID *float64 `json:"anonymousUid,omitempty" tf:"anonymous_uid,omitempty"`

	// The filter applied to the scope for this rule. The filter's format depends on its scope: default scope matches all clients and has no filter value; network scope takes a CIDR format; host takes an IP address or fully qualified domain name. If a client does not match any filter rule and there is no default rule, access is denied.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Whether to enable root squash? Defaults to false.
	// +kubebuilder:validation:Optional
	RootSquashEnabled *bool `json:"rootSquashEnabled,omitempty" tf:"root_squash_enabled,omitempty"`

	// The scope of this rule. The scope and  the filter determine which clients match the rule. Possible values are: default, network, host.
	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`

	// Whether allow access to subdirectories under the root export? Defaults to false.
	// +kubebuilder:validation:Optional
	SubmountAccessEnabled *bool `json:"submountAccessEnabled,omitempty" tf:"submount_access_enabled,omitempty"`

	// Whether SUID is allowed? Defaults to false.
	// +kubebuilder:validation:Optional
	SuidEnabled *bool `json:"suidEnabled,omitempty" tf:"suid_enabled,omitempty"`
}

type BindObservation struct {
}

type BindParameters struct {

	// The Bind Distinguished Name  identity to be used in the secure LDAP connection.
	// +kubebuilder:validation:Required
	Dn *string `json:"dn" tf:"dn,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`
}

type DNSObservation struct {
}

type DNSParameters struct {

	// The DNS search domain for the HPC Cache.
	// +kubebuilder:validation:Optional
	SearchDomain *string `json:"searchDomain,omitempty" tf:"search_domain,omitempty"`

	// A list of DNS servers for the HPC Cache. At most three IP are allowed to set.
	// +kubebuilder:validation:Required
	Servers []*string `json:"servers" tf:"servers,omitempty"`
}

type DefaultAccessPolicyObservation struct {
}

type DefaultAccessPolicyParameters struct {

	// One to three access_rule blocks as defined above.
	// +kubebuilder:validation:Required
	AccessRule []AccessRuleParameters `json:"accessRule" tf:"access_rule,omitempty"`
}

type DirectoryActiveDirectoryObservation struct {
}

type DirectoryActiveDirectoryParameters struct {

	// The NetBIOS name to assign to the HPC Cache when it joins the Active Directory domain as a server.
	// +kubebuilder:validation:Required
	CacheNetbiosName *string `json:"cacheNetbiosName" tf:"cache_netbios_name,omitempty"`

	// The primary DNS IP address used to resolve the Active Directory domain controller's FQDN.
	// +kubebuilder:validation:Required
	DNSPrimaryIP *string `json:"dnsPrimaryIp" tf:"dns_primary_ip,omitempty"`

	// The secondary DNS IP address used to resolve the Active Directory domain controller's FQDN.
	// +kubebuilder:validation:Optional
	DNSSecondaryIP *string `json:"dnsSecondaryIp,omitempty" tf:"dns_secondary_ip,omitempty"`

	// The fully qualified domain name of the Active Directory domain controller.
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// The Active Directory domain's NetBIOS name.
	// +kubebuilder:validation:Required
	DomainNetbiosName *string `json:"domainNetbiosName" tf:"domain_netbios_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username of the Active Directory domain administrator.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type DirectoryFlatFileObservation struct {
}

type DirectoryFlatFileParameters struct {

	// The URI of the file containing group information .
	// +kubebuilder:validation:Required
	GroupFileURI *string `json:"groupFileUri" tf:"group_file_uri,omitempty"`

	// The URI of the file containing user information .
	// +kubebuilder:validation:Required
	PasswordFileURI *string `json:"passwordFileUri" tf:"password_file_uri,omitempty"`
}

type DirectoryLdapObservation struct {
}

type DirectoryLdapParameters struct {

	// The base distinguished name  for the LDAP domain.
	// +kubebuilder:validation:Required
	BaseDn *string `json:"baseDn" tf:"base_dn,omitempty"`

	// A bind block as defined above.
	// +kubebuilder:validation:Optional
	Bind []BindParameters `json:"bind,omitempty" tf:"bind,omitempty"`

	// The URI of the CA certificate to validate the LDAP secure connection.
	// +kubebuilder:validation:Optional
	CertificateValidationURI *string `json:"certificateValidationUri,omitempty" tf:"certificate_validation_uri,omitempty"`

	// Whether the certificate should be automatically downloaded. This can be set to true only when certificate_validation_uri is provided. Defaults to false.
	// +kubebuilder:validation:Optional
	DownloadCertificateAutomatically *bool `json:"downloadCertificateAutomatically,omitempty" tf:"download_certificate_automatically,omitempty"`

	// Whether the LDAP connection should be encrypted? Defaults to false.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The FQDN or IP address of the LDAP server.
	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`
}

type HPCCacheObservation struct {

	// The id of the HPC Cache.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of IP Addresses where the HPC Cache can be mounted.
	MountAddresses []*string `json:"mountAddresses,omitempty" tf:"mount_addresses,omitempty"`
}

type HPCCacheParameters struct {

	// The size of the HPC Cache, in GB. Possible values are 3072, 6144, 12288, 21623, 24576, 43246, 49152 and 86491. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	CacheSizeInGb *float64 `json:"cacheSizeInGb" tf:"cache_size_in_gb,omitempty"`

	// A dns block as defined below.
	// +kubebuilder:validation:Optional
	DNS []DNSParameters `json:"dns,omitempty" tf:"dns,omitempty"`

	// A default_access_policy block as defined below.
	// +kubebuilder:validation:Optional
	DefaultAccessPolicy []DefaultAccessPolicyParameters `json:"defaultAccessPolicy,omitempty" tf:"default_access_policy,omitempty"`

	// A directory_active_directory block as defined below.
	// +kubebuilder:validation:Optional
	DirectoryActiveDirectory []DirectoryActiveDirectoryParameters `json:"directoryActiveDirectory,omitempty" tf:"directory_active_directory,omitempty"`

	// A directory_flat_file block as defined below.
	// +kubebuilder:validation:Optional
	DirectoryFlatFile []DirectoryFlatFileParameters `json:"directoryFlatFile,omitempty" tf:"directory_flat_file,omitempty"`

	// A directory_ldap block as defined below.
	// +kubebuilder:validation:Optional
	DirectoryLdap []DirectoryLdapParameters `json:"directoryLdap,omitempty" tf:"directory_ldap,omitempty"`

	// Specifies the supported Azure Region where the HPC Cache should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The IPv4 maximum transmission unit configured for the subnet of the HPC Cache. Possible values range from 576 - 1500. Defaults to 1500.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The NTP server IP Address or FQDN for the HPC Cache. Defaults to time.windows.com.
	// +kubebuilder:validation:Optional
	NtpServer *string `json:"ntpServer,omitempty" tf:"ntp_server,omitempty"`

	// The name of the Resource Group in which to create the HPC Cache. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU of HPC Cache to use. Possible values are  - Standard_2G, Standard_4G Standard_8G or  - Standard_L4_5G, Standard_L9G, and Standard_L16G. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// The ID of the Subnet for the HPC Cache. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the HPC Cache.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// HPCCacheSpec defines the desired state of HPCCache
type HPCCacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HPCCacheParameters `json:"forProvider"`
}

// HPCCacheStatus defines the observed state of HPCCache.
type HPCCacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HPCCacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCache is the Schema for the HPCCaches API. Manages a HPC Cache.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HPCCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HPCCacheSpec   `json:"spec"`
	Status            HPCCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheList contains a list of HPCCaches
type HPCCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HPCCache `json:"items"`
}

// Repository type metadata.
var (
	HPCCache_Kind             = "HPCCache"
	HPCCache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HPCCache_Kind}.String()
	HPCCache_KindAPIVersion   = HPCCache_Kind + "." + CRDGroupVersion.String()
	HPCCache_GroupVersionKind = CRDGroupVersion.WithKind(HPCCache_Kind)
)

func init() {
	SchemeBuilder.Register(&HPCCache{}, &HPCCacheList{})
}
