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

type AutomaticInstanceRepairObservation struct {
}

type AutomaticInstanceRepairParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Amount of time  for which automatic repairs will be delayed. The grace period starts right after the VM is found unhealthy. The time duration should be specified in ISO 8601 format.
	// +kubebuilder:validation:Optional
	GracePeriod *string `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`
}

type AutomaticOsUpgradePolicyObservation struct {
}

type AutomaticOsUpgradePolicyParameters struct {

	// Should automatic rollbacks be disabled?
	// +kubebuilder:validation:Required
	DisableAutomaticRollback *bool `json:"disableAutomaticRollback" tf:"disable_automatic_rollback,omitempty"`

	// Should OS Upgrades automatically be applied to Scale Set instances in a rolling fashion when a newer version of the OS Image becomes available?
	// +kubebuilder:validation:Required
	EnableAutomaticOsUpgrade *bool `json:"enableAutomaticOsUpgrade" tf:"enable_automatic_os_upgrade,omitempty"`
}

type ExtensionObservation struct {
}

type ExtensionParameters struct {

	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to true.
	// +kubebuilder:validation:Optional
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version,omitempty"`

	// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to false.
	// +kubebuilder:validation:Optional
	AutomaticUpgradeEnabled *bool `json:"automaticUpgradeEnabled,omitempty" tf:"automatic_upgrade_enabled,omitempty"`

	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	// +kubebuilder:validation:Optional
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A JSON String which specifies Sensitive Settings  for the Extension.
	// +kubebuilder:validation:Optional
	ProtectedSettingsSecretRef *v1.SecretKeySelector `json:"protectedSettingsSecretRef,omitempty" tf:"-"`

	// An ordered list of Extension names which this should be provisioned after.
	// +kubebuilder:validation:Optional
	ProvisionAfterExtensions []*string `json:"provisionAfterExtensions,omitempty" tf:"provision_after_extensions,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// A JSON String which specifies Settings for the Extension.
	// +kubebuilder:validation:Optional
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
	// +kubebuilder:validation:Required
	TypeHandlerVersion *string `json:"typeHandlerVersion" tf:"type_handler_version,omitempty"`
}

type IPConfigurationObservation struct {
}

type IPConfigurationParameters struct {

	// A list of Backend Address Pools ID's from a Application Gateway which this Virtual Machine Scale Set should be connected to.
	// +kubebuilder:validation:Optional
	ApplicationGatewayBackendAddressPoolIds []*string `json:"applicationGatewayBackendAddressPoolIds,omitempty" tf:"application_gateway_backend_address_pool_ids,omitempty"`

	// A list of Application Security Group ID's which this Virtual Machine Scale Set should be connected to.
	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupIds []*string `json:"applicationSecurityGroupIds,omitempty" tf:"application_security_group_ids,omitempty"`

	// A list of Backend Address Pools ID's from a Load Balancer which this Virtual Machine Scale Set should be connected to.
	// +kubebuilder:validation:Optional
	LoadBalancerBackendAddressPoolIds []*string `json:"loadBalancerBackendAddressPoolIds,omitempty" tf:"load_balancer_backend_address_pool_ids,omitempty"`

	// A list of NAT Rule ID's from a Load Balancer which this Virtual Machine Scale Set should be connected to.
	// +kubebuilder:validation:Optional
	LoadBalancerInboundNATRulesIds []*string `json:"loadBalancerInboundNatRulesIds,omitempty" tf:"load_balancer_inbound_nat_rules_ids,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// A public_ip_address block as defined below.
	// +kubebuilder:validation:Optional
	PublicIPAddress []PublicIPAddressParameters `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`

	// The ID of the Subnet which this IP Configuration should be connected to.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type IPTagObservation struct {
}

type IPTagParameters struct {

	// The IP Tag associated with the Public IP, such as SQL or Storage.
	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LinuxVirtualMachineScaleSetAdditionalCapabilitiesObservation struct {
}

type LinuxVirtualMachineScaleSetAdditionalCapabilitiesParameters struct {

	// Should the capacity to enable Data Disks of the UltraSSD_LRS storage account type be supported on this Virtual Machine Scale Set? Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled,omitempty"`
}

type LinuxVirtualMachineScaleSetAdminSSHKeyObservation struct {
}

type LinuxVirtualMachineScaleSetAdminSSHKeyParameters struct {

	// The Public Key which should be used for authentication, which needs to be at least 2048-bit and in ssh-rsa format.
	// +kubebuilder:validation:Required
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`

	// The Username for which this Public SSH Key should be configured.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type LinuxVirtualMachineScaleSetBootDiagnosticsObservation struct {
}

type LinuxVirtualMachineScaleSetBootDiagnosticsParameters struct {

	// The Primary/Secondary Endpoint for the Azure Storage Account which should be used to store Boot Diagnostics, including Console Output and Screenshots from the Hypervisor.
	// +kubebuilder:validation:Optional
	StorageAccountURI *string `json:"storageAccountUri,omitempty" tf:"storage_account_uri,omitempty"`
}

type LinuxVirtualMachineScaleSetDataDiskObservation struct {
}

type LinuxVirtualMachineScaleSetDataDiskParameters struct {

	// +kubebuilder:validation:Required
	Caching *string `json:"caching" tf:"caching,omitempty"`

	// The create option which should be used for this Data Disk. Possible values are Empty and FromImage. Defaults to Empty. .
	// +kubebuilder:validation:Optional
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Required
	DiskSizeGb *float64 `json:"diskSizeGb" tf:"disk_size_gb,omitempty"`

	// The Logical Unit Number of the Data Disk, which must be unique within the Virtual Machine.
	// +kubebuilder:validation:Required
	Lun *float64 `json:"lun" tf:"lun,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// +kubebuilder:validation:Optional
	UltraSsdDiskIopsReadWrite *float64 `json:"ultraSsdDiskIopsReadWrite,omitempty" tf:"ultra_ssd_disk_iops_read_write,omitempty"`

	// +kubebuilder:validation:Optional
	UltraSsdDiskMbpsReadWrite *float64 `json:"ultraSsdDiskMbpsReadWrite,omitempty" tf:"ultra_ssd_disk_mbps_read_write,omitempty"`

	// +kubebuilder:validation:Optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type LinuxVirtualMachineScaleSetIdentityObservation struct {

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type LinuxVirtualMachineScaleSetIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Linux Virtual Machine Scale Set.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LinuxVirtualMachineScaleSetObservation struct {

	// The ID of the Linux Virtual Machine Scale Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []LinuxVirtualMachineScaleSetIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Unique ID for this Linux Virtual Machine Scale Set.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type LinuxVirtualMachineScaleSetOsDiskObservation struct {
}

type LinuxVirtualMachineScaleSetOsDiskParameters struct {

	// +kubebuilder:validation:Required
	Caching *string `json:"caching" tf:"caching,omitempty"`

	// A diff_disk_settings block as defined above. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DiffDiskSettings []OsDiskDiffDiskSettingsParameters `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// The ID of the Disk Encryption Set which should be used to Encrypt the OS Disk when the Virtual Machine Scale Set is Confidential VMSS. Conflicts with disk_encryption_set_id. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SecureVMDiskEncryptionSetID *string `json:"secureVmDiskEncryptionSetId,omitempty" tf:"secure_vm_disk_encryption_set_id,omitempty"`

	// Encryption Type when the Virtual Machine Scale Set is Confidential VMSS. Possible values are VMGuestStateOnly and DiskWithVMGuestState. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SecurityEncryptionType *string `json:"securityEncryptionType,omitempty" tf:"security_encryption_type,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// +kubebuilder:validation:Optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type LinuxVirtualMachineScaleSetParameters struct {

	// A additional_capabilities block as defined below.
	// +kubebuilder:validation:Optional
	AdditionalCapabilities []LinuxVirtualMachineScaleSetAdditionalCapabilitiesParameters `json:"additionalCapabilities,omitempty" tf:"additional_capabilities,omitempty"`

	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AdminPasswordSecretRef *v1.SecretKeySelector `json:"adminPasswordSecretRef,omitempty" tf:"-"`

	// One or more admin_ssh_key blocks as defined below.
	// +kubebuilder:validation:Optional
	AdminSSHKey []LinuxVirtualMachineScaleSetAdminSSHKeyParameters `json:"adminSshKey,omitempty" tf:"admin_ssh_key,omitempty"`

	// The username of the local administrator on each Virtual Machine Scale Set instance. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// A automatic_instance_repair block as defined below. To enable the automatic instance repair, this Virtual Machine Scale Set must have a valid health_probe_id or an Application Health Extension.
	// +kubebuilder:validation:Optional
	AutomaticInstanceRepair []AutomaticInstanceRepairParameters `json:"automaticInstanceRepair,omitempty" tf:"automatic_instance_repair,omitempty"`

	// A automatic_os_upgrade_policy block as defined below. This can only be specified when upgrade_mode is set to Automatic.
	// +kubebuilder:validation:Optional
	AutomaticOsUpgradePolicy []AutomaticOsUpgradePolicyParameters `json:"automaticOsUpgradePolicy,omitempty" tf:"automatic_os_upgrade_policy,omitempty"`

	// A boot_diagnostics block as defined below.
	// +kubebuilder:validation:Optional
	BootDiagnostics []LinuxVirtualMachineScaleSetBootDiagnosticsParameters `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`

	// The prefix which should be used for the name of the Virtual Machines in this Scale Set. If unspecified this defaults to the value for the name field. If the value of the name field is not a valid computer_name_prefix, then you must specify computer_name_prefix.
	// +kubebuilder:validation:Optional
	ComputerNamePrefix *string `json:"computerNamePrefix,omitempty" tf:"computer_name_prefix,omitempty"`

	// The Base64-Encoded Custom Data which should be used for this Virtual Machine Scale Set.
	// +kubebuilder:validation:Optional
	CustomDataSecretRef *v1.SecretKeySelector `json:"customDataSecretRef,omitempty" tf:"-"`

	// One or more data_disk blocks as defined below.
	// +kubebuilder:validation:Optional
	DataDisk []LinuxVirtualMachineScaleSetDataDiskParameters `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`

	// Should Password Authentication be disabled on this Virtual Machine Scale Set? Defaults to true.
	// +kubebuilder:validation:Optional
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication,omitempty"`

	// Should Virtual Machine Extensions be run on Overprovisioned Virtual Machines in the Scale Set? Defaults to false.
	// +kubebuilder:validation:Optional
	DoNotRunExtensionsOnOverprovisionedMachines *bool `json:"doNotRunExtensionsOnOverprovisionedMachines,omitempty" tf:"do_not_run_extensions_on_overprovisioned_machines,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Linux Virtual Machine Scale Set should exist. Changing this forces a new Linux Virtual Machine Scale Set to be created.
	// +kubebuilder:validation:Optional
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// Should all of the disks  attached to this Virtual Machine be encrypted by enabling Encryption at Host?
	// +kubebuilder:validation:Optional
	EncryptionAtHostEnabled *bool `json:"encryptionAtHostEnabled,omitempty" tf:"encryption_at_host_enabled,omitempty"`

	// The Policy which should be used Virtual Machines are Evicted from the Scale Set. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// One or more extension blocks as defined below
	// +kubebuilder:validation:Optional
	Extension []ExtensionParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// Specifies the duration allocated for all extensions to start. The time duration should be between 15 minutes and 120 minutes  and should be specified in ISO 8601 format. Defaults to 90 minutes .
	// +kubebuilder:validation:Optional
	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty" tf:"extensions_time_budget,omitempty"`

	// The ID of a Load Balancer Probe which should be used to determine the health of an instance. This is Required and can only be specified when upgrade_mode is set to Automatic or Rolling.
	// +kubebuilder:validation:Optional
	HealthProbeID *string `json:"healthProbeId,omitempty" tf:"health_probe_id,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []LinuxVirtualMachineScaleSetIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The number of Virtual Machines in the Scale Set.
	// +kubebuilder:validation:Required
	Instances *float64 `json:"instances" tf:"instances,omitempty"`

	// The Azure location where the Linux Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The maximum price you're willing to pay for each Virtual Machine in this Scale Set, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machines in the Scale Set will be evicted using the eviction_policy. Defaults to -1, which means that each Virtual Machine in this Scale Set should not be evicted for price reasons.
	// +kubebuilder:validation:Optional
	MaxBidPrice *float64 `json:"maxBidPrice,omitempty" tf:"max_bid_price,omitempty"`

	// One or more network_interface blocks as defined below.
	// +kubebuilder:validation:Required
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface" tf:"network_interface,omitempty"`

	// An os_disk block as defined below.
	// +kubebuilder:validation:Required
	OsDisk []LinuxVirtualMachineScaleSetOsDiskParameters `json:"osDisk" tf:"os_disk,omitempty"`

	// Should Azure over-provision Virtual Machines in this Scale Set? This means that multiple Virtual Machines will be provisioned and Azure will keep the instances which become available first - which improves provisioning success rates and improves deployment time. You're not billed for these over-provisioned VM's and they don't count towards the Subscription Quota. Defaults to true.
	// +kubebuilder:validation:Optional
	Overprovision *bool `json:"overprovision,omitempty" tf:"overprovision,omitempty"`

	// A plan block as documented below.
	// +kubebuilder:validation:Optional
	Plan []LinuxVirtualMachineScaleSetPlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// Specifies the number of fault domains that are used by this Linux Virtual Machine Scale Set. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PlatformFaultDomainCount *float64 `json:"platformFaultDomainCount,omitempty" tf:"platform_fault_domain_count,omitempty"`

	// The Priority of this Virtual Machine Scale Set. Possible values are Regular and Spot. Defaults to Regular. Changing this value forces a new resource.
	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// Should the Azure VM Agent be provisioned on each Virtual Machine in the Scale Set? Defaults to true. Changing this value forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ProvisionVMAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`

	// The ID of the Proximity Placement Group in which the Virtual Machine Scale Set should be assigned to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id,omitempty"`

	// The name of the Resource Group in which the Linux Virtual Machine Scale Set should be exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A rolling_upgrade_policy block as defined below. This is Required and can only be specified when upgrade_mode is set to Automatic or Rolling.
	// +kubebuilder:validation:Optional
	RollingUpgradePolicy []RollingUpgradePolicyParameters `json:"rollingUpgradePolicy,omitempty" tf:"rolling_upgrade_policy,omitempty"`

	// The scale-in policy rule that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled in. Possible values for the scale-in policy rules are Default, NewestVM and OldestVM, defaults to Default. For more information about scale in policy, please refer to this doc.
	// +kubebuilder:validation:Optional
	ScaleInPolicy *string `json:"scaleInPolicy,omitempty" tf:"scale_in_policy,omitempty"`

	// One or more secret blocks as defined below.
	// +kubebuilder:validation:Optional
	Secret []LinuxVirtualMachineScaleSetSecretParameters `json:"secret,omitempty" tf:"secret,omitempty"`

	// Specifies whether secure boot should be enabled on the virtual machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty" tf:"secure_boot_enabled,omitempty"`

	// Should this Virtual Machine Scale Set be limited to a Single Placement Group, which means the number of instances will be capped at 100 Virtual Machines. Defaults to true.
	// +kubebuilder:validation:Optional
	SinglePlacementGroup *bool `json:"singlePlacementGroup,omitempty" tf:"single_placement_group,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// The ID of an Image which each Virtual Machine in this Scale Set should be based on.
	// +kubebuilder:validation:Optional
	SourceImageID *string `json:"sourceImageId,omitempty" tf:"source_image_id,omitempty"`

	// A source_image_reference block as defined below.
	// +kubebuilder:validation:Optional
	SourceImageReference []LinuxVirtualMachineScaleSetSourceImageReferenceParameters `json:"sourceImageReference,omitempty" tf:"source_image_reference,omitempty"`

	// A mapping of tags which should be assigned to this Virtual Machine Scale Set.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A terminate_notification block as defined below.
	// +kubebuilder:validation:Optional
	TerminateNotification []TerminateNotificationParameters `json:"terminateNotification,omitempty" tf:"terminate_notification,omitempty"`

	// A termination_notification block as defined below.
	// +kubebuilder:validation:Optional
	TerminationNotification []LinuxVirtualMachineScaleSetTerminationNotificationParameters `json:"terminationNotification,omitempty" tf:"termination_notification,omitempty"`

	// Specifies how Upgrades  should be performed to Virtual Machine Instances. Possible values are Automatic, Manual and Rolling. Defaults to Manual.
	// +kubebuilder:validation:Optional
	UpgradeMode *string `json:"upgradeMode,omitempty" tf:"upgrade_mode,omitempty"`

	// The Base64-Encoded User Data which should be used for this Virtual Machine Scale Set.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// Specifies whether vTPM should be enabled on the virtual machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	VtpmEnabled *bool `json:"vtpmEnabled,omitempty" tf:"vtpm_enabled,omitempty"`

	// Should the Virtual Machines in this Scale Set be strictly evenly distributed across Availability Zones? Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ZoneBalance *bool `json:"zoneBalance,omitempty" tf:"zone_balance,omitempty"`

	// Specifies a list of Availability Zones in which this Linux Virtual Machine Scale Set should be located. Changing this forces a new Linux Virtual Machine Scale Set to be created.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type LinuxVirtualMachineScaleSetPlanObservation struct {
}

type LinuxVirtualMachineScaleSetPlanParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the product of the image from the marketplace. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Product *string `json:"product" tf:"product,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`
}

type LinuxVirtualMachineScaleSetSecretObservation struct {
}

type LinuxVirtualMachineScaleSetSecretParameters struct {

	// One or more certificate blocks as defined above.
	// +kubebuilder:validation:Required
	Certificate []SecretCertificateParameters `json:"certificate" tf:"certificate,omitempty"`

	// The ID of the Key Vault from which all Secrets should be sourced.
	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`
}

type LinuxVirtualMachineScaleSetSourceImageReferenceObservation struct {
}

type LinuxVirtualMachineScaleSetSourceImageReferenceParameters struct {

	// Specifies the offer of the image used to create the virtual machines.
	// +kubebuilder:validation:Required
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type LinuxVirtualMachineScaleSetTerminationNotificationObservation struct {
}

type LinuxVirtualMachineScaleSetTerminationNotificationParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type NetworkInterfaceObservation struct {
}

type NetworkInterfaceParameters struct {

	// A list of IP Addresses of DNS Servers which should be assigned to the Network Interface.
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Does this Network Interface support Accelerated Networking? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking,omitempty"`

	// Does this Network Interface support IP Forwarding? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableIPForwarding *bool `json:"enableIpForwarding,omitempty" tf:"enable_ip_forwarding,omitempty"`

	// One or more ip_configuration blocks as defined above.
	// +kubebuilder:validation:Required
	IPConfiguration []IPConfigurationParameters `json:"ipConfiguration" tf:"ip_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of a Network Security Group which should be assigned to this Network Interface.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`
}

type OsDiskDiffDiskSettingsObservation struct {
}

type OsDiskDiffDiskSettingsParameters struct {

	// +kubebuilder:validation:Required
	Option *string `json:"option" tf:"option,omitempty"`
}

type PublicIPAddressObservation struct {
}

type PublicIPAddressParameters struct {

	// The Prefix which should be used for the Domain Name Label for each Virtual Machine Instance. Azure concatenates the Domain Name Label and Virtual Machine Index to create a unique Domain Name Label for each Virtual Machine.
	// +kubebuilder:validation:Optional
	DomainNameLabel *string `json:"domainNameLabel,omitempty" tf:"domain_name_label,omitempty"`

	// One or more ip_tag blocks as defined above.
	// +kubebuilder:validation:Optional
	IPTag []IPTagParameters `json:"ipTag,omitempty" tf:"ip_tag,omitempty"`

	// The Idle Timeout in Minutes for the Public IP Address. Possible values are in the range 4 to 32.
	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the Public IP Address Prefix from where Public IP Addresses should be allocated. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`
}

type RollingUpgradePolicyObservation struct {
}

type RollingUpgradePolicyParameters struct {

	// The maximum percent of total virtual machine instances that will be upgraded simultaneously by the rolling upgrade in one batch. As this is a maximum, unhealthy instances in previous or future batches can cause the percentage of instances in a batch to decrease to ensure higher reliability.
	// +kubebuilder:validation:Required
	MaxBatchInstancePercent *float64 `json:"maxBatchInstancePercent" tf:"max_batch_instance_percent,omitempty"`

	// The maximum percentage of the total virtual machine instances in the scale set that can be simultaneously unhealthy, either as a result of being upgraded, or by being found in an unhealthy state by the virtual machine health checks before the rolling upgrade aborts. This constraint will be checked prior to starting any batch.
	// +kubebuilder:validation:Required
	MaxUnhealthyInstancePercent *float64 `json:"maxUnhealthyInstancePercent" tf:"max_unhealthy_instance_percent,omitempty"`

	// The maximum percentage of upgraded virtual machine instances that can be found to be in an unhealthy state. This check will happen after each batch is upgraded. If this percentage is ever exceeded, the rolling update aborts.
	// +kubebuilder:validation:Required
	MaxUnhealthyUpgradedInstancePercent *float64 `json:"maxUnhealthyUpgradedInstancePercent" tf:"max_unhealthy_upgraded_instance_percent,omitempty"`

	// The wait time between completing the update for all virtual machines in one batch and starting the next batch. The time duration should be specified in ISO 8601 format.
	// +kubebuilder:validation:Required
	PauseTimeBetweenBatches *string `json:"pauseTimeBetweenBatches" tf:"pause_time_between_batches,omitempty"`
}

type SecretCertificateObservation struct {
}

type SecretCertificateParameters struct {

	// The Secret URL of a Key Vault Certificate.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type TerminateNotificationObservation struct {
}

type TerminateNotificationParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

// LinuxVirtualMachineScaleSetSpec defines the desired state of LinuxVirtualMachineScaleSet
type LinuxVirtualMachineScaleSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinuxVirtualMachineScaleSetParameters `json:"forProvider"`
}

// LinuxVirtualMachineScaleSetStatus defines the observed state of LinuxVirtualMachineScaleSet.
type LinuxVirtualMachineScaleSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinuxVirtualMachineScaleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinuxVirtualMachineScaleSet is the Schema for the LinuxVirtualMachineScaleSets API. Manages a Linux Virtual Machine Scale Set.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinuxVirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinuxVirtualMachineScaleSetSpec   `json:"spec"`
	Status            LinuxVirtualMachineScaleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinuxVirtualMachineScaleSetList contains a list of LinuxVirtualMachineScaleSets
type LinuxVirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinuxVirtualMachineScaleSet `json:"items"`
}

// Repository type metadata.
var (
	LinuxVirtualMachineScaleSet_Kind             = "LinuxVirtualMachineScaleSet"
	LinuxVirtualMachineScaleSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinuxVirtualMachineScaleSet_Kind}.String()
	LinuxVirtualMachineScaleSet_KindAPIVersion   = LinuxVirtualMachineScaleSet_Kind + "." + CRDGroupVersion.String()
	LinuxVirtualMachineScaleSet_GroupVersionKind = CRDGroupVersion.WithKind(LinuxVirtualMachineScaleSet_Kind)
)

func init() {
	SchemeBuilder.Register(&LinuxVirtualMachineScaleSet{}, &LinuxVirtualMachineScaleSetList{})
}
