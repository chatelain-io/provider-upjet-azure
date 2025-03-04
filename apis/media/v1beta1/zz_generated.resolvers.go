// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	rconfig "github.com/upbound/provider-azure/apis/rconfig"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Asset.
	apisresolver "github.com/upbound/provider-azure/internal/apis"
)

func (mg *Asset) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta2", "ServicesAccount", "ServicesAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
			Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AssetFilter.
func (mg *AssetFilter) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Asset", "AssetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AssetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.AssetIDRef,
			Selector:     mg.Spec.ForProvider.AssetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AssetID")
	}
	mg.Spec.ForProvider.AssetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AssetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ContentKeyPolicy.
func (mg *ContentKeyPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "ServicesAccount", "ServicesAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
			Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Job.
func (mg *Job) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.InputAsset); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Asset", "AssetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InputAsset[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.InputAsset[i3].NameRef,
				Selector:     mg.Spec.ForProvider.InputAsset[i3].NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.InputAsset[i3].Name")
		}
		mg.Spec.ForProvider.InputAsset[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.InputAsset[i3].NameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "ServicesAccount", "ServicesAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
			Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.OutputAsset); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Asset", "AssetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OutputAsset[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.OutputAsset[i3].NameRef,
				Selector:     mg.Spec.ForProvider.OutputAsset[i3].NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.OutputAsset[i3].Name")
		}
		mg.Spec.ForProvider.OutputAsset[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.OutputAsset[i3].NameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Transform", "TransformList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransformName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TransformNameRef,
			Selector:     mg.Spec.ForProvider.TransformNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransformName")
	}
	mg.Spec.ForProvider.TransformName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransformNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.InputAsset); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Asset", "AssetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InputAsset[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.InputAsset[i3].NameRef,
				Selector:     mg.Spec.InitProvider.InputAsset[i3].NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.InputAsset[i3].Name")
		}
		mg.Spec.InitProvider.InputAsset[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.InputAsset[i3].NameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.OutputAsset); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Asset", "AssetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.OutputAsset[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.OutputAsset[i3].NameRef,
				Selector:     mg.Spec.InitProvider.OutputAsset[i3].NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.OutputAsset[i3].Name")
		}
		mg.Spec.InitProvider.OutputAsset[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.OutputAsset[i3].NameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this LiveEvent.
func (mg *LiveEvent) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "ServicesAccount", "ServicesAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
			Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LiveEventOutput.
func (mg *LiveEventOutput) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Asset", "AssetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AssetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.AssetNameRef,
			Selector:     mg.Spec.ForProvider.AssetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AssetName")
	}
	mg.Spec.ForProvider.AssetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AssetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta2", "LiveEvent", "LiveEventList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LiveEventID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.LiveEventIDRef,
			Selector:     mg.Spec.ForProvider.LiveEventIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LiveEventID")
	}
	mg.Spec.ForProvider.LiveEventID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LiveEventIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Asset", "AssetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AssetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.AssetNameRef,
			Selector:     mg.Spec.InitProvider.AssetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AssetName")
	}
	mg.Spec.InitProvider.AssetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AssetNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServicesAccount.
func (mg *ServicesAccount) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageAccount); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Account", "AccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccount[i3].ID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.StorageAccount[i3].IDRef,
				Selector:     mg.Spec.ForProvider.StorageAccount[i3].IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccount[i3].ID")
		}
		mg.Spec.ForProvider.StorageAccount[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageAccount[i3].IDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageAccount); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Account", "AccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageAccount[i3].ID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.StorageAccount[i3].IDRef,
				Selector:     mg.Spec.InitProvider.StorageAccount[i3].IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageAccount[i3].ID")
		}
		mg.Spec.InitProvider.StorageAccount[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageAccount[i3].IDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ServicesAccountFilter.
func (mg *ServicesAccountFilter) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "ServicesAccount", "ServicesAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
			Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StreamingEndpoint.
func (mg *StreamingEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "ServicesAccount", "ServicesAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
			Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StreamingLocator.
func (mg *StreamingLocator) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Asset", "AssetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AssetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.AssetNameRef,
			Selector:     mg.Spec.ForProvider.AssetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AssetName")
	}
	mg.Spec.ForProvider.AssetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AssetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta2", "ServicesAccountFilter", "ServicesAccountFilterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.FilterNames),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.FilterNamesRefs,
			Selector:      mg.Spec.ForProvider.FilterNamesSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FilterNames")
	}
	mg.Spec.ForProvider.FilterNames = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.FilterNamesRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta2", "ServicesAccount", "ServicesAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
			Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "Asset", "AssetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AssetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.AssetNameRef,
			Selector:     mg.Spec.InitProvider.AssetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AssetName")
	}
	mg.Spec.InitProvider.AssetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AssetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta2", "ServicesAccountFilter", "ServicesAccountFilterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.FilterNames),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.FilterNamesRefs,
			Selector:      mg.Spec.InitProvider.FilterNamesSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FilterNames")
	}
	mg.Spec.InitProvider.FilterNames = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.FilterNamesRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this StreamingPolicy.
func (mg *StreamingPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CommonEncryptionCenc); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.CommonEncryptionCenc[i3].DefaultContentKey); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "ContentKeyPolicy", "ContentKeyPolicyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyName),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyNameRef,
					Selector:     mg.Spec.ForProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyName")
			}
			mg.Spec.ForProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyNameRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "ServicesAccount", "ServicesAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
			Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.CommonEncryptionCenc); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.CommonEncryptionCenc[i3].DefaultContentKey); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "ContentKeyPolicy", "ContentKeyPolicyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyName),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyNameRef,
					Selector:     mg.Spec.InitProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyName")
			}
			mg.Spec.InitProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.CommonEncryptionCenc[i3].DefaultContentKey[i4].PolicyNameRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Transform.
func (mg *Transform) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("media.azure.upbound.io", "v1beta1", "ServicesAccount", "ServicesAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MediaServicesAccountName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MediaServicesAccountNameRef,
			Selector:     mg.Spec.ForProvider.MediaServicesAccountNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MediaServicesAccountName")
	}
	mg.Spec.ForProvider.MediaServicesAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MediaServicesAccountNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
