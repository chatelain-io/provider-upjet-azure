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

type APNSCredentialObservation struct {
}

type APNSCredentialParameters struct {

	// The Application Mode which defines which server the APNS Messages should be sent to. Possible values are Production and Sandbox.
	// +kubebuilder:validation:Required
	ApplicationMode *string `json:"applicationMode" tf:"application_mode,omitempty"`

	// The Bundle ID of the iOS/macOS application to send push notifications for, such as com.hashicorp.example.
	// +kubebuilder:validation:Required
	BundleID *string `json:"bundleId" tf:"bundle_id,omitempty"`

	// The Apple Push Notifications Service (APNS) Key.
	// +kubebuilder:validation:Required
	KeyID *string `json:"keyId" tf:"key_id,omitempty"`

	// The ID of the team the Token.
	// +kubebuilder:validation:Required
	TeamID *string `json:"teamId" tf:"team_id,omitempty"`

	// The Push Token associated with the Apple Developer Account. This is the contents of the key downloaded from the Apple Developer Portal between the -----BEGIN PRIVATE KEY----- and -----END PRIVATE KEY----- blocks.
	// +kubebuilder:validation:Required
	TokenSecretRef v1.SecretKeySelector `json:"tokenSecretRef" tf:"-"`
}

type GCMCredentialObservation struct {
}

type GCMCredentialParameters struct {

	// The API Key associated with the Google Cloud Messaging service.
	// +kubebuilder:validation:Required
	APIKeySecretRef v1.SecretKeySelector `json:"apiKeySecretRef" tf:"-"`
}

type NotificationHubObservation struct {

	// The ID of the Notification Hub.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NotificationHubParameters struct {

	// A apns_credential block as defined below.
	// +kubebuilder:validation:Optional
	APNSCredential []APNSCredentialParameters `json:"apnsCredential,omitempty" tf:"apns_credential,omitempty"`

	// A gcm_credential block as defined below.
	// +kubebuilder:validation:Optional
	GCMCredential []GCMCredentialParameters `json:"gcmCredential,omitempty" tf:"gcm_credential,omitempty"`

	// The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	NamespaceName *string `json:"namespaceName" tf:"namespace_name,omitempty"`

	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// NotificationHubSpec defines the desired state of NotificationHub
type NotificationHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationHubParameters `json:"forProvider"`
}

// NotificationHubStatus defines the observed state of NotificationHub.
type NotificationHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationHub is the Schema for the NotificationHubs API. Manages a Notification Hub within a Notification Hub Namespace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NotificationHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationHubSpec   `json:"spec"`
	Status            NotificationHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationHubList contains a list of NotificationHubs
type NotificationHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationHub `json:"items"`
}

// Repository type metadata.
var (
	NotificationHub_Kind             = "NotificationHub"
	NotificationHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationHub_Kind}.String()
	NotificationHub_KindAPIVersion   = NotificationHub_Kind + "." + CRDGroupVersion.String()
	NotificationHub_GroupVersionKind = CRDGroupVersion.WithKind(NotificationHub_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationHub{}, &NotificationHubList{})
}
