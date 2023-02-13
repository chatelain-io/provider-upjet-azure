/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/compute/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualMachineID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualMachineIDRef,
		Selector:     mg.Spec.ForProvider.VirtualMachineIDSelector,
		To: reference.To{
			List:    &v1beta1.WindowsVirtualMachineList{},
			Managed: &v1beta1.WindowsVirtualMachine{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualMachineID")
	}
	mg.Spec.ForProvider.VirtualMachineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualMachineIDRef = rsp.ResolvedReference

	return nil
}
