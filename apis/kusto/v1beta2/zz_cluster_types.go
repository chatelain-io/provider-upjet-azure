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

type ClusterInitParameters struct {

	// List of allowed FQDNs(Fully Qualified Domain Name) for egress from Cluster.
	AllowedFqdns []*string `json:"allowedFqdns,omitempty" tf:"allowed_fqdns,omitempty"`

	// The list of ips in the format of CIDR allowed to connect to the cluster.
	AllowedIPRanges []*string `json:"allowedIpRanges,omitempty" tf:"allowed_ip_ranges,omitempty"`

	// Specifies if the cluster could be automatically stopped (due to lack of data or no activity for many days). Defaults to true.
	AutoStopEnabled *bool `json:"autoStopEnabled,omitempty" tf:"auto_stop_enabled,omitempty"`

	// Specifies if the cluster's disks are encrypted.
	DiskEncryptionEnabled *bool `json:"diskEncryptionEnabled,omitempty" tf:"disk_encryption_enabled,omitempty"`

	// Is the cluster's double encryption enabled? Changing this forces a new resource to be created.
	DoubleEncryptionEnabled *bool `json:"doubleEncryptionEnabled,omitempty" tf:"double_encryption_enabled,omitempty"`

	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// An identity block as defined below.
	Identity *IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// An list of language_extensions to enable. Valid values are: PYTHON, PYTHON_3.10.8 and R. PYTHON is used to specify Python 3.6.5 image and PYTHON_3.10.8 is used to specify Python 3.10.8 image. Note that PYTHON_3.10.8 is only available in skus which support nested virtualization.
	// +listType=set
	LanguageExtensions []*string `json:"languageExtensions,omitempty" tf:"language_extensions,omitempty"`

	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// An optimized_auto_scale block as defined below.
	OptimizedAutoScale *OptimizedAutoScaleInitParameters `json:"optimizedAutoScale,omitempty" tf:"optimized_auto_scale,omitempty"`

	// Whether to restrict outbound network access. Value is optional but if passed in, must be true or false, default is false.
	OutboundNetworkAccessRestricted *bool `json:"outboundNetworkAccessRestricted,omitempty" tf:"outbound_network_access_restricted,omitempty"`

	// Indicates what public IP type to create - IPv4 (default), or DualStack (both IPv4 and IPv6). Defaults to IPv4.
	PublicIPType *string `json:"publicIpType,omitempty" tf:"public_ip_type,omitempty"`

	// Is the public network access enabled? Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies if the purge operations are enabled.
	PurgeEnabled *bool `json:"purgeEnabled,omitempty" tf:"purge_enabled,omitempty"`

	// A sku block as defined below.
	Sku *SkuInitParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// Specifies if the streaming ingest is enabled.
	StreamingIngestionEnabled *bool `json:"streamingIngestionEnabled,omitempty" tf:"streaming_ingestion_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of tenant IDs that are trusted by the cluster. Default setting trusts all other tenants. Use trusted_external_tenants = ["*"] to explicitly allow all other tenants, trusted_external_tenants = ["MyTenantOnly"] for only your tenant or trusted_external_tenants = ["<tenantId1>", "<tenantIdx>"] to allow specific other tenants.
	TrustedExternalTenants []*string `json:"trustedExternalTenants,omitempty" tf:"trusted_external_tenants,omitempty"`

	// A virtual_network_configuration block as defined below.
	VirtualNetworkConfiguration *VirtualNetworkConfigurationInitParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`

	// Specifies a list of Availability Zones in which this Kusto Cluster should be located. Changing this forces a new Kusto Cluster to be created.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ClusterObservation struct {

	// List of allowed FQDNs(Fully Qualified Domain Name) for egress from Cluster.
	AllowedFqdns []*string `json:"allowedFqdns,omitempty" tf:"allowed_fqdns,omitempty"`

	// The list of ips in the format of CIDR allowed to connect to the cluster.
	AllowedIPRanges []*string `json:"allowedIpRanges,omitempty" tf:"allowed_ip_ranges,omitempty"`

	// Specifies if the cluster could be automatically stopped (due to lack of data or no activity for many days). Defaults to true.
	AutoStopEnabled *bool `json:"autoStopEnabled,omitempty" tf:"auto_stop_enabled,omitempty"`

	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionURI *string `json:"dataIngestionUri,omitempty" tf:"data_ingestion_uri,omitempty"`

	// Specifies if the cluster's disks are encrypted.
	DiskEncryptionEnabled *bool `json:"diskEncryptionEnabled,omitempty" tf:"disk_encryption_enabled,omitempty"`

	// Is the cluster's double encryption enabled? Changing this forces a new resource to be created.
	DoubleEncryptionEnabled *bool `json:"doubleEncryptionEnabled,omitempty" tf:"double_encryption_enabled,omitempty"`

	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The Kusto Cluster ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity *IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// An list of language_extensions to enable. Valid values are: PYTHON, PYTHON_3.10.8 and R. PYTHON is used to specify Python 3.6.5 image and PYTHON_3.10.8 is used to specify Python 3.10.8 image. Note that PYTHON_3.10.8 is only available in skus which support nested virtualization.
	// +listType=set
	LanguageExtensions []*string `json:"languageExtensions,omitempty" tf:"language_extensions,omitempty"`

	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// An optimized_auto_scale block as defined below.
	OptimizedAutoScale *OptimizedAutoScaleObservation `json:"optimizedAutoScale,omitempty" tf:"optimized_auto_scale,omitempty"`

	// Whether to restrict outbound network access. Value is optional but if passed in, must be true or false, default is false.
	OutboundNetworkAccessRestricted *bool `json:"outboundNetworkAccessRestricted,omitempty" tf:"outbound_network_access_restricted,omitempty"`

	// Indicates what public IP type to create - IPv4 (default), or DualStack (both IPv4 and IPv6). Defaults to IPv4.
	PublicIPType *string `json:"publicIpType,omitempty" tf:"public_ip_type,omitempty"`

	// Is the public network access enabled? Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies if the purge operations are enabled.
	PurgeEnabled *bool `json:"purgeEnabled,omitempty" tf:"purge_enabled,omitempty"`

	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A sku block as defined below.
	Sku *SkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// Specifies if the streaming ingest is enabled.
	StreamingIngestionEnabled *bool `json:"streamingIngestionEnabled,omitempty" tf:"streaming_ingestion_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of tenant IDs that are trusted by the cluster. Default setting trusts all other tenants. Use trusted_external_tenants = ["*"] to explicitly allow all other tenants, trusted_external_tenants = ["MyTenantOnly"] for only your tenant or trusted_external_tenants = ["<tenantId1>", "<tenantIdx>"] to allow specific other tenants.
	TrustedExternalTenants []*string `json:"trustedExternalTenants,omitempty" tf:"trusted_external_tenants,omitempty"`

	// The FQDN of the Azure Kusto Cluster.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// A virtual_network_configuration block as defined below.
	VirtualNetworkConfiguration *VirtualNetworkConfigurationObservation `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`

	// Specifies a list of Availability Zones in which this Kusto Cluster should be located. Changing this forces a new Kusto Cluster to be created.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ClusterParameters struct {

	// List of allowed FQDNs(Fully Qualified Domain Name) for egress from Cluster.
	// +kubebuilder:validation:Optional
	AllowedFqdns []*string `json:"allowedFqdns,omitempty" tf:"allowed_fqdns,omitempty"`

	// The list of ips in the format of CIDR allowed to connect to the cluster.
	// +kubebuilder:validation:Optional
	AllowedIPRanges []*string `json:"allowedIpRanges,omitempty" tf:"allowed_ip_ranges,omitempty"`

	// Specifies if the cluster could be automatically stopped (due to lack of data or no activity for many days). Defaults to true.
	// +kubebuilder:validation:Optional
	AutoStopEnabled *bool `json:"autoStopEnabled,omitempty" tf:"auto_stop_enabled,omitempty"`

	// Specifies if the cluster's disks are encrypted.
	// +kubebuilder:validation:Optional
	DiskEncryptionEnabled *bool `json:"diskEncryptionEnabled,omitempty" tf:"disk_encryption_enabled,omitempty"`

	// Is the cluster's double encryption enabled? Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DoubleEncryptionEnabled *bool `json:"doubleEncryptionEnabled,omitempty" tf:"double_encryption_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity *IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// An list of language_extensions to enable. Valid values are: PYTHON, PYTHON_3.10.8 and R. PYTHON is used to specify Python 3.6.5 image and PYTHON_3.10.8 is used to specify Python 3.10.8 image. Note that PYTHON_3.10.8 is only available in skus which support nested virtualization.
	// +kubebuilder:validation:Optional
	// +listType=set
	LanguageExtensions []*string `json:"languageExtensions,omitempty" tf:"language_extensions,omitempty"`

	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// An optimized_auto_scale block as defined below.
	// +kubebuilder:validation:Optional
	OptimizedAutoScale *OptimizedAutoScaleParameters `json:"optimizedAutoScale,omitempty" tf:"optimized_auto_scale,omitempty"`

	// Whether to restrict outbound network access. Value is optional but if passed in, must be true or false, default is false.
	// +kubebuilder:validation:Optional
	OutboundNetworkAccessRestricted *bool `json:"outboundNetworkAccessRestricted,omitempty" tf:"outbound_network_access_restricted,omitempty"`

	// Indicates what public IP type to create - IPv4 (default), or DualStack (both IPv4 and IPv6). Defaults to IPv4.
	// +kubebuilder:validation:Optional
	PublicIPType *string `json:"publicIpType,omitempty" tf:"public_ip_type,omitempty"`

	// Is the public network access enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies if the purge operations are enabled.
	// +kubebuilder:validation:Optional
	PurgeEnabled *bool `json:"purgeEnabled,omitempty" tf:"purge_enabled,omitempty"`

	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A sku block as defined below.
	// +kubebuilder:validation:Optional
	Sku *SkuParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// Specifies if the streaming ingest is enabled.
	// +kubebuilder:validation:Optional
	StreamingIngestionEnabled *bool `json:"streamingIngestionEnabled,omitempty" tf:"streaming_ingestion_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of tenant IDs that are trusted by the cluster. Default setting trusts all other tenants. Use trusted_external_tenants = ["*"] to explicitly allow all other tenants, trusted_external_tenants = ["MyTenantOnly"] for only your tenant or trusted_external_tenants = ["<tenantId1>", "<tenantIdx>"] to allow specific other tenants.
	// +kubebuilder:validation:Optional
	TrustedExternalTenants []*string `json:"trustedExternalTenants,omitempty" tf:"trusted_external_tenants,omitempty"`

	// A virtual_network_configuration block as defined below.
	// +kubebuilder:validation:Optional
	VirtualNetworkConfiguration *VirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`

	// Specifies a list of Availability Zones in which this Kusto Cluster should be located. Changing this forces a new Kusto Cluster to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Kusto Cluster.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: SystemAssigned, UserAssigned and SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Kusto Cluster.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this System Assigned Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this System Assigned Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: SystemAssigned, UserAssigned and SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Kusto Cluster.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: SystemAssigned, UserAssigned and SystemAssigned, UserAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type OptimizedAutoScaleInitParameters struct {

	// The maximum number of allowed instances. Must between 0 and 1000.
	MaximumInstances *float64 `json:"maximumInstances,omitempty" tf:"maximum_instances,omitempty"`

	// The minimum number of allowed instances. Must between 0 and 1000.
	MinimumInstances *float64 `json:"minimumInstances,omitempty" tf:"minimum_instances,omitempty"`
}

type OptimizedAutoScaleObservation struct {

	// The maximum number of allowed instances. Must between 0 and 1000.
	MaximumInstances *float64 `json:"maximumInstances,omitempty" tf:"maximum_instances,omitempty"`

	// The minimum number of allowed instances. Must between 0 and 1000.
	MinimumInstances *float64 `json:"minimumInstances,omitempty" tf:"minimum_instances,omitempty"`
}

type OptimizedAutoScaleParameters struct {

	// The maximum number of allowed instances. Must between 0 and 1000.
	// +kubebuilder:validation:Optional
	MaximumInstances *float64 `json:"maximumInstances" tf:"maximum_instances,omitempty"`

	// The minimum number of allowed instances. Must between 0 and 1000.
	// +kubebuilder:validation:Optional
	MinimumInstances *float64 `json:"minimumInstances" tf:"minimum_instances,omitempty"`
}

type SkuInitParameters struct {

	// Specifies the node count for the cluster. Boundaries depend on the SKU name.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the SKU. Possible values are Dev(No SLA)_Standard_D11_v2, Dev(No SLA)_Standard_E2a_v4, Standard_D14_v2, Standard_D11_v2, Standard_D16d_v5, Standard_D13_v2, Standard_D12_v2, Standard_DS14_v2+4TB_PS, Standard_DS14_v2+3TB_PS, Standard_DS13_v2+1TB_PS, Standard_DS13_v2+2TB_PS, Standard_D32d_v5, Standard_D32d_v4, Standard_EC8ads_v5, Standard_EC8as_v5+1TB_PS, Standard_EC8as_v5+2TB_PS, Standard_EC16ads_v5, Standard_EC16as_v5+4TB_PS, Standard_EC16as_v5+3TB_PS, Standard_E80ids_v4, Standard_E8a_v4, Standard_E8ads_v5, Standard_E8as_v5+1TB_PS, Standard_E8as_v5+2TB_PS, Standard_E8as_v4+1TB_PS, Standard_E8as_v4+2TB_PS, Standard_E8d_v5, Standard_E8d_v4, Standard_E8s_v5+1TB_PS, Standard_E8s_v5+2TB_PS, Standard_E8s_v4+1TB_PS, Standard_E8s_v4+2TB_PS, Standard_E4a_v4, Standard_E4ads_v5, Standard_E4d_v5, Standard_E4d_v4, Standard_E16a_v4, Standard_E16ads_v5, Standard_E16as_v5+4TB_PS, Standard_E16as_v5+3TB_PS, Standard_E16as_v4+4TB_PS, Standard_E16as_v4+3TB_PS, Standard_E16d_v5, Standard_E16d_v4, Standard_E16s_v5+4TB_PS, Standard_E16s_v5+3TB_PS, Standard_E16s_v4+4TB_PS, Standard_E16s_v4+3TB_PS, Standard_E64i_v3, Standard_E2a_v4, Standard_E2ads_v5, Standard_E2d_v5, Standard_E2d_v4, Standard_L8as_v3, Standard_L8s, Standard_L8s_v3, Standard_L8s_v2, Standard_L4s, Standard_L16as_v3, Standard_L16s, Standard_L16s_v3, Standard_L16s_v2, Standard_L32as_v3 and Standard_L32s_v3.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SkuObservation struct {

	// Specifies the node count for the cluster. Boundaries depend on the SKU name.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the SKU. Possible values are Dev(No SLA)_Standard_D11_v2, Dev(No SLA)_Standard_E2a_v4, Standard_D14_v2, Standard_D11_v2, Standard_D16d_v5, Standard_D13_v2, Standard_D12_v2, Standard_DS14_v2+4TB_PS, Standard_DS14_v2+3TB_PS, Standard_DS13_v2+1TB_PS, Standard_DS13_v2+2TB_PS, Standard_D32d_v5, Standard_D32d_v4, Standard_EC8ads_v5, Standard_EC8as_v5+1TB_PS, Standard_EC8as_v5+2TB_PS, Standard_EC16ads_v5, Standard_EC16as_v5+4TB_PS, Standard_EC16as_v5+3TB_PS, Standard_E80ids_v4, Standard_E8a_v4, Standard_E8ads_v5, Standard_E8as_v5+1TB_PS, Standard_E8as_v5+2TB_PS, Standard_E8as_v4+1TB_PS, Standard_E8as_v4+2TB_PS, Standard_E8d_v5, Standard_E8d_v4, Standard_E8s_v5+1TB_PS, Standard_E8s_v5+2TB_PS, Standard_E8s_v4+1TB_PS, Standard_E8s_v4+2TB_PS, Standard_E4a_v4, Standard_E4ads_v5, Standard_E4d_v5, Standard_E4d_v4, Standard_E16a_v4, Standard_E16ads_v5, Standard_E16as_v5+4TB_PS, Standard_E16as_v5+3TB_PS, Standard_E16as_v4+4TB_PS, Standard_E16as_v4+3TB_PS, Standard_E16d_v5, Standard_E16d_v4, Standard_E16s_v5+4TB_PS, Standard_E16s_v5+3TB_PS, Standard_E16s_v4+4TB_PS, Standard_E16s_v4+3TB_PS, Standard_E64i_v3, Standard_E2a_v4, Standard_E2ads_v5, Standard_E2d_v5, Standard_E2d_v4, Standard_L8as_v3, Standard_L8s, Standard_L8s_v3, Standard_L8s_v2, Standard_L4s, Standard_L16as_v3, Standard_L16s, Standard_L16s_v3, Standard_L16s_v2, Standard_L32as_v3 and Standard_L32s_v3.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SkuParameters struct {

	// Specifies the node count for the cluster. Boundaries depend on the SKU name.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the SKU. Possible values are Dev(No SLA)_Standard_D11_v2, Dev(No SLA)_Standard_E2a_v4, Standard_D14_v2, Standard_D11_v2, Standard_D16d_v5, Standard_D13_v2, Standard_D12_v2, Standard_DS14_v2+4TB_PS, Standard_DS14_v2+3TB_PS, Standard_DS13_v2+1TB_PS, Standard_DS13_v2+2TB_PS, Standard_D32d_v5, Standard_D32d_v4, Standard_EC8ads_v5, Standard_EC8as_v5+1TB_PS, Standard_EC8as_v5+2TB_PS, Standard_EC16ads_v5, Standard_EC16as_v5+4TB_PS, Standard_EC16as_v5+3TB_PS, Standard_E80ids_v4, Standard_E8a_v4, Standard_E8ads_v5, Standard_E8as_v5+1TB_PS, Standard_E8as_v5+2TB_PS, Standard_E8as_v4+1TB_PS, Standard_E8as_v4+2TB_PS, Standard_E8d_v5, Standard_E8d_v4, Standard_E8s_v5+1TB_PS, Standard_E8s_v5+2TB_PS, Standard_E8s_v4+1TB_PS, Standard_E8s_v4+2TB_PS, Standard_E4a_v4, Standard_E4ads_v5, Standard_E4d_v5, Standard_E4d_v4, Standard_E16a_v4, Standard_E16ads_v5, Standard_E16as_v5+4TB_PS, Standard_E16as_v5+3TB_PS, Standard_E16as_v4+4TB_PS, Standard_E16as_v4+3TB_PS, Standard_E16d_v5, Standard_E16d_v4, Standard_E16s_v5+4TB_PS, Standard_E16s_v5+3TB_PS, Standard_E16s_v4+4TB_PS, Standard_E16s_v4+3TB_PS, Standard_E64i_v3, Standard_E2a_v4, Standard_E2ads_v5, Standard_E2d_v5, Standard_E2d_v4, Standard_L8as_v3, Standard_L8s, Standard_L8s_v3, Standard_L8s_v2, Standard_L4s, Standard_L16as_v3, Standard_L16s, Standard_L16s_v3, Standard_L16s_v2, Standard_L32as_v3 and Standard_L32s_v3.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type VirtualNetworkConfigurationInitParameters struct {

	// Data management's service public IP address resource id.
	DataManagementPublicIPID *string `json:"dataManagementPublicIpId,omitempty" tf:"data_management_public_ip_id,omitempty"`

	// Engine service's public IP address resource id.
	EnginePublicIPID *string `json:"enginePublicIpId,omitempty" tf:"engine_public_ip_id,omitempty"`

	// The subnet resource id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type VirtualNetworkConfigurationObservation struct {

	// Data management's service public IP address resource id.
	DataManagementPublicIPID *string `json:"dataManagementPublicIpId,omitempty" tf:"data_management_public_ip_id,omitempty"`

	// Engine service's public IP address resource id.
	EnginePublicIPID *string `json:"enginePublicIpId,omitempty" tf:"engine_public_ip_id,omitempty"`

	// The subnet resource id.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type VirtualNetworkConfigurationParameters struct {

	// Data management's service public IP address resource id.
	// +kubebuilder:validation:Optional
	DataManagementPublicIPID *string `json:"dataManagementPublicIpId" tf:"data_management_public_ip_id,omitempty"`

	// Engine service's public IP address resource id.
	// +kubebuilder:validation:Optional
	EnginePublicIPID *string `json:"enginePublicIpId" tf:"engine_public_ip_id,omitempty"`

	// The subnet resource id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Cluster is the Schema for the Clusters API. Manages Kusto (also known as Azure Data Explorer) Cluster
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
