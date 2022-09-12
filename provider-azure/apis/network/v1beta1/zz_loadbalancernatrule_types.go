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

type LoadBalancerNatRuleObservation struct {

	// The ID of the Load Balancer NAT Rule.
	BackendIPConfigurationID *string `json:"backendIpConfigurationId,omitempty" tf:"backend_ip_configuration_id,omitempty"`

	// The ID of the Load Balancer NAT Rule.
	FrontendIPConfigurationID *string `json:"frontendIpConfigurationId,omitempty" tf:"frontend_ip_configuration_id,omitempty"`

	// The ID of the Load Balancer NAT Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LoadBalancerNatRuleParameters struct {

	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive.
	// +kubebuilder:validation:Required
	BackendPort *float64 `json:"backendPort" tf:"backend_port,omitempty"`

	// Are the Floating IPs enabled for this Load Balancer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// The name of the frontend IP configuration exposing this rule.
	// +kubebuilder:validation:Required
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName" tf:"frontend_ip_configuration_name,omitempty"`

	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive.
	// +kubebuilder:validation:Required
	FrontendPort *float64 `json:"frontendPort" tf:"frontend_port,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 30 minutes. Defaults to 4 minutes.
	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// The ID of the Load Balancer in which to create the NAT Rule.
	// +crossplane:generate:reference:type=LoadBalancer
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// The transport protocol for the external endpoint. Possible values are Udp, Tcp or All.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// The name of the resource group in which to create the resource.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// LoadBalancerNatRuleSpec defines the desired state of LoadBalancerNatRule
type LoadBalancerNatRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerNatRuleParameters `json:"forProvider"`
}

// LoadBalancerNatRuleStatus defines the observed state of LoadBalancerNatRule.
type LoadBalancerNatRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerNatRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNatRule is the Schema for the LoadBalancerNatRules API. Manages a Load Balancer NAT Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LoadBalancerNatRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerNatRuleSpec   `json:"spec"`
	Status            LoadBalancerNatRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNatRuleList contains a list of LoadBalancerNatRules
type LoadBalancerNatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerNatRule `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerNatRule_Kind             = "LoadBalancerNatRule"
	LoadBalancerNatRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerNatRule_Kind}.String()
	LoadBalancerNatRule_KindAPIVersion   = LoadBalancerNatRule_Kind + "." + CRDGroupVersion.String()
	LoadBalancerNatRule_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerNatRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerNatRule{}, &LoadBalancerNatRuleList{})
}
