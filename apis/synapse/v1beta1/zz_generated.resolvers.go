/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta13 "github.com/upbound/provider-azure/apis/keyvault/v1beta1"
	v1beta12 "github.com/upbound/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	v1beta1 "github.com/upbound/provider-azure/apis/storage/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this FirewallRule.
func (mg *FirewallRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LinkedService.
func (mg *LinkedService) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.IntegrationRuntime); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IntegrationRuntime[i3].Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.IntegrationRuntime[i3].NameRef,
			Selector:     mg.Spec.ForProvider.IntegrationRuntime[i3].NameSelector,
			To: reference.To{
				List:    &IntegrationRuntimeAzureList{},
				Managed: &IntegrationRuntimeAzure{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.IntegrationRuntime[i3].Name")
		}
		mg.Spec.ForProvider.IntegrationRuntime[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.IntegrationRuntime[i3].NameRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TargetResourceIDRef,
		Selector:     mg.Spec.ForProvider.TargetResourceIDSelector,
		To: reference.To{
			List:    &v1beta1.AccountList{},
			Managed: &v1beta1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetResourceID")
	}
	mg.Spec.ForProvider.TargetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PrivateLinkHub.
func (mg *PrivateLinkHub) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoleAssignment.
func (mg *RoleAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SQLPool.
func (mg *SQLPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SQLPoolIDRef,
		Selector:     mg.Spec.ForProvider.SQLPoolIDSelector,
		To: reference.To{
			List:    &SQLPoolList{},
			Managed: &SQLPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SQLPoolID")
	}
	mg.Spec.ForProvider.SQLPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SQLPoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageEndpoint),
		Extract:      resource.ExtractParamPath("primary_blob_endpoint", true),
		Reference:    mg.Spec.ForProvider.StorageEndpointRef,
		Selector:     mg.Spec.ForProvider.StorageEndpointSelector,
		To: reference.To{
			List:    &v1beta1.AccountList{},
			Managed: &v1beta1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageEndpoint")
	}
	mg.Spec.ForProvider.StorageEndpoint = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageEndpointRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SQLPoolIDRef,
		Selector:     mg.Spec.ForProvider.SQLPoolIDSelector,
		To: reference.To{
			List:    &SQLPoolList{},
			Managed: &SQLPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SQLPoolID")
	}
	mg.Spec.ForProvider.SQLPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SQLPoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageEndpoint),
		Extract:      resource.ExtractParamPath("primary_blob_endpoint", true),
		Reference:    mg.Spec.ForProvider.StorageEndpointRef,
		Selector:     mg.Spec.ForProvider.StorageEndpointSelector,
		To: reference.To{
			List:    &v1beta1.AccountList{},
			Managed: &v1beta1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageEndpoint")
	}
	mg.Spec.ForProvider.StorageEndpoint = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageEndpointRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkloadGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WorkloadGroupIDRef,
		Selector:     mg.Spec.ForProvider.WorkloadGroupIDSelector,
		To: reference.To{
			List:    &SQLPoolWorkloadGroupList{},
			Managed: &SQLPoolWorkloadGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkloadGroupID")
	}
	mg.Spec.ForProvider.WorkloadGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkloadGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SQLPoolIDRef,
		Selector:     mg.Spec.ForProvider.SQLPoolIDSelector,
		To: reference.To{
			List:    &SQLPoolList{},
			Managed: &SQLPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SQLPoolID")
	}
	mg.Spec.ForProvider.SQLPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SQLPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SparkPool.
func (mg *SparkPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workspace.
func (mg *Workspace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ComputeSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ComputeSubnetIDRef,
		Selector:     mg.Spec.ForProvider.ComputeSubnetIDSelector,
		To: reference.To{
			List:    &v1beta12.SubnetList{},
			Managed: &v1beta12.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ComputeSubnetID")
	}
	mg.Spec.ForProvider.ComputeSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ComputeSubnetIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CustomerManagedKey); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomerManagedKey[i3].KeyVersionlessID),
			Extract:      resource.ExtractParamPath("versionless_id", true),
			Reference:    mg.Spec.ForProvider.CustomerManagedKey[i3].KeyVersionlessIDRef,
			Selector:     mg.Spec.ForProvider.CustomerManagedKey[i3].KeyVersionlessIDSelector,
			To: reference.To{
				List:    &v1beta13.KeyList{},
				Managed: &v1beta13.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomerManagedKey[i3].KeyVersionlessID")
		}
		mg.Spec.ForProvider.CustomerManagedKey[i3].KeyVersionlessID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomerManagedKey[i3].KeyVersionlessIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagedResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ManagedResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ManagedResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManagedResourceGroupName")
	}
	mg.Spec.ForProvider.ManagedResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ManagedResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageDataLakeGen2FileSystemID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.StorageDataLakeGen2FileSystemIDRef,
		Selector:     mg.Spec.ForProvider.StorageDataLakeGen2FileSystemIDSelector,
		To: reference.To{
			List:    &v1beta1.DataLakeGen2FileSystemList{},
			Managed: &v1beta1.DataLakeGen2FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageDataLakeGen2FileSystemID")
	}
	mg.Spec.ForProvider.StorageDataLakeGen2FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageDataLakeGen2FileSystemIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageEndpoint),
		Extract:      resource.ExtractParamPath("primary_blob_endpoint", true),
		Reference:    mg.Spec.ForProvider.StorageEndpointRef,
		Selector:     mg.Spec.ForProvider.StorageEndpointSelector,
		To: reference.To{
			List:    &v1beta1.AccountList{},
			Managed: &v1beta1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageEndpoint")
	}
	mg.Spec.ForProvider.StorageEndpoint = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageEndpointRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageEndpoint),
		Extract:      resource.ExtractParamPath("primary_blob_endpoint", true),
		Reference:    mg.Spec.ForProvider.StorageEndpointRef,
		Selector:     mg.Spec.ForProvider.StorageEndpointSelector,
		To: reference.To{
			List:    &v1beta1.AccountList{},
			Managed: &v1beta1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageEndpoint")
	}
	mg.Spec.ForProvider.StorageEndpoint = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageEndpointRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseWorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SynapseWorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.SynapseWorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseWorkspaceID")
	}
	mg.Spec.ForProvider.SynapseWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseWorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WorkspaceVulnerabilityAssessment.
func (mg *WorkspaceVulnerabilityAssessment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceSecurityAlertPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WorkspaceSecurityAlertPolicyIDRef,
		Selector:     mg.Spec.ForProvider.WorkspaceSecurityAlertPolicyIDSelector,
		To: reference.To{
			List:    &WorkspaceSecurityAlertPolicyList{},
			Managed: &WorkspaceSecurityAlertPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceSecurityAlertPolicyID")
	}
	mg.Spec.ForProvider.WorkspaceSecurityAlertPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceSecurityAlertPolicyIDRef = rsp.ResolvedReference

	return nil
}
