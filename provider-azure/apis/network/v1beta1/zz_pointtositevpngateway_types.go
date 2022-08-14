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

type ConnectionConfigurationObservation struct {
}

type ConnectionConfigurationParameters struct {

	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to false.
	// +kubebuilder:validation:Optional
	InternetSecurityEnabled *bool `json:"internetSecurityEnabled,omitempty" tf:"internet_security_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A route block as defined below.
	// +kubebuilder:validation:Optional
	Route []RouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// A vpn_client_address_pool block as defined below.
	// +kubebuilder:validation:Required
	VPNClientAddressPool []VPNClientAddressPoolParameters `json:"vpnClientAddressPool" tf:"vpn_client_address_pool,omitempty"`
}

type PointToSiteVPNGatewayObservation struct {

	// The ID of the Point-to-Site VPN Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PointToSiteVPNGatewayParameters struct {

	// A connection_configuration block as defined below.
	// +kubebuilder:validation:Required
	ConnectionConfiguration []ConnectionConfigurationParameters `json:"connectionConfiguration" tf:"connection_configuration,omitempty"`

	// A list of IP Addresses of DNS Servers for the Point-to-Site VPN Gateway.
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group in which to create the Point-to-Site VPN Gateway. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Scale Unit for this Point-to-Site VPN Gateway.
	// +kubebuilder:validation:Required
	ScaleUnit *float64 `json:"scaleUnit" tf:"scale_unit,omitempty"`

	// A mapping of tags to assign to the Point-to-Site VPN Gateway.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the VPN Server Configuration which this Point-to-Site VPN Gateway should use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=VPNServerConfiguration
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPNServerConfigurationID *string `json:"vpnServerConfigurationId,omitempty" tf:"vpn_server_configuration_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPNServerConfigurationIDRef *v1.Reference `json:"vpnServerConfigurationIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPNServerConfigurationIDSelector *v1.Selector `json:"vpnServerConfigurationIdSelector,omitempty" tf:"-"`

	// The ID of the Virtual Hub where this Point-to-Site VPN Gateway should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=VirtualHub
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualHubIDRef *v1.Reference `json:"virtualHubIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VirtualHubIDSelector *v1.Selector `json:"virtualHubIdSelector,omitempty" tf:"-"`
}

type RouteObservation struct {
}

type RouteParameters struct {

	// The Virtual Hub Route Table resource id associated with this Routing Configuration.
	// +kubebuilder:validation:Required
	AssociatedRouteTableID *string `json:"associatedRouteTableId" tf:"associated_route_table_id,omitempty"`

	// A propagated_route_table block as defined below.
	// +kubebuilder:validation:Optional
	PropagatedRouteTable []RoutePropagatedRouteTableParameters `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

type RoutePropagatedRouteTableObservation struct {
}

type RoutePropagatedRouteTableParameters struct {

	// The list of Virtual Hub Route Table resource id which the routes will be propagated to.
	// +kubebuilder:validation:Required
	Ids []*string `json:"ids" tf:"ids,omitempty"`

	// The list of labels to logically group Virtual Hub Route Tables which the routes will be propagated to.
	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type VPNClientAddressPoolObservation struct {
}

type VPNClientAddressPoolParameters struct {

	// A list of CIDR Ranges which should be used as Address Prefixes.
	// +kubebuilder:validation:Required
	AddressPrefixes []*string `json:"addressPrefixes" tf:"address_prefixes,omitempty"`
}

// PointToSiteVPNGatewaySpec defines the desired state of PointToSiteVPNGateway
type PointToSiteVPNGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PointToSiteVPNGatewayParameters `json:"forProvider"`
}

// PointToSiteVPNGatewayStatus defines the observed state of PointToSiteVPNGateway.
type PointToSiteVPNGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PointToSiteVPNGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PointToSiteVPNGateway is the Schema for the PointToSiteVPNGateways API. Manages a Point-to-Site VPN Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PointToSiteVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PointToSiteVPNGatewaySpec   `json:"spec"`
	Status            PointToSiteVPNGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PointToSiteVPNGatewayList contains a list of PointToSiteVPNGateways
type PointToSiteVPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PointToSiteVPNGateway `json:"items"`
}

// Repository type metadata.
var (
	PointToSiteVPNGateway_Kind             = "PointToSiteVPNGateway"
	PointToSiteVPNGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PointToSiteVPNGateway_Kind}.String()
	PointToSiteVPNGateway_KindAPIVersion   = PointToSiteVPNGateway_Kind + "." + CRDGroupVersion.String()
	PointToSiteVPNGateway_GroupVersionKind = CRDGroupVersion.WithKind(PointToSiteVPNGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&PointToSiteVPNGateway{}, &PointToSiteVPNGatewayList{})
}
