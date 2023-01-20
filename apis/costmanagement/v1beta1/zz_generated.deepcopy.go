//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportDataOptionsObservation) DeepCopyInto(out *ExportDataOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportDataOptionsObservation.
func (in *ExportDataOptionsObservation) DeepCopy() *ExportDataOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(ExportDataOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportDataOptionsParameters) DeepCopyInto(out *ExportDataOptionsParameters) {
	*out = *in
	if in.TimeFrame != nil {
		in, out := &in.TimeFrame, &out.TimeFrame
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportDataOptionsParameters.
func (in *ExportDataOptionsParameters) DeepCopy() *ExportDataOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(ExportDataOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportDataStorageLocationObservation) DeepCopyInto(out *ExportDataStorageLocationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportDataStorageLocationObservation.
func (in *ExportDataStorageLocationObservation) DeepCopy() *ExportDataStorageLocationObservation {
	if in == nil {
		return nil
	}
	out := new(ExportDataStorageLocationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportDataStorageLocationParameters) DeepCopyInto(out *ExportDataStorageLocationParameters) {
	*out = *in
	if in.ContainerID != nil {
		in, out := &in.ContainerID, &out.ContainerID
		*out = new(string)
		**out = **in
	}
	if in.ContainerIDRef != nil {
		in, out := &in.ContainerIDRef, &out.ContainerIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerIDSelector != nil {
		in, out := &in.ContainerIDSelector, &out.ContainerIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RootFolderPath != nil {
		in, out := &in.RootFolderPath, &out.RootFolderPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportDataStorageLocationParameters.
func (in *ExportDataStorageLocationParameters) DeepCopy() *ExportDataStorageLocationParameters {
	if in == nil {
		return nil
	}
	out := new(ExportDataStorageLocationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupCostManagementExport) DeepCopyInto(out *ResourceGroupCostManagementExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupCostManagementExport.
func (in *ResourceGroupCostManagementExport) DeepCopy() *ResourceGroupCostManagementExport {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupCostManagementExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceGroupCostManagementExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupCostManagementExportList) DeepCopyInto(out *ResourceGroupCostManagementExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceGroupCostManagementExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupCostManagementExportList.
func (in *ResourceGroupCostManagementExportList) DeepCopy() *ResourceGroupCostManagementExportList {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupCostManagementExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceGroupCostManagementExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupCostManagementExportObservation) DeepCopyInto(out *ResourceGroupCostManagementExportObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupCostManagementExportObservation.
func (in *ResourceGroupCostManagementExportObservation) DeepCopy() *ResourceGroupCostManagementExportObservation {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupCostManagementExportObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupCostManagementExportParameters) DeepCopyInto(out *ResourceGroupCostManagementExportParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.ExportDataOptions != nil {
		in, out := &in.ExportDataOptions, &out.ExportDataOptions
		*out = make([]ExportDataOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExportDataStorageLocation != nil {
		in, out := &in.ExportDataStorageLocation, &out.ExportDataStorageLocation
		*out = make([]ExportDataStorageLocationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecurrencePeriodEndDate != nil {
		in, out := &in.RecurrencePeriodEndDate, &out.RecurrencePeriodEndDate
		*out = new(string)
		**out = **in
	}
	if in.RecurrencePeriodStartDate != nil {
		in, out := &in.RecurrencePeriodStartDate, &out.RecurrencePeriodStartDate
		*out = new(string)
		**out = **in
	}
	if in.RecurrenceType != nil {
		in, out := &in.RecurrenceType, &out.RecurrenceType
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupIDRef != nil {
		in, out := &in.ResourceGroupIDRef, &out.ResourceGroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupIDSelector != nil {
		in, out := &in.ResourceGroupIDSelector, &out.ResourceGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupCostManagementExportParameters.
func (in *ResourceGroupCostManagementExportParameters) DeepCopy() *ResourceGroupCostManagementExportParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupCostManagementExportParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupCostManagementExportSpec) DeepCopyInto(out *ResourceGroupCostManagementExportSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupCostManagementExportSpec.
func (in *ResourceGroupCostManagementExportSpec) DeepCopy() *ResourceGroupCostManagementExportSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupCostManagementExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupCostManagementExportStatus) DeepCopyInto(out *ResourceGroupCostManagementExportStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupCostManagementExportStatus.
func (in *ResourceGroupCostManagementExportStatus) DeepCopy() *ResourceGroupCostManagementExportStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupCostManagementExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExport) DeepCopyInto(out *SubscriptionCostManagementExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExport.
func (in *SubscriptionCostManagementExport) DeepCopy() *SubscriptionCostManagementExport {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscriptionCostManagementExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExportExportDataOptionsObservation) DeepCopyInto(out *SubscriptionCostManagementExportExportDataOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExportExportDataOptionsObservation.
func (in *SubscriptionCostManagementExportExportDataOptionsObservation) DeepCopy() *SubscriptionCostManagementExportExportDataOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExportExportDataOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExportExportDataOptionsParameters) DeepCopyInto(out *SubscriptionCostManagementExportExportDataOptionsParameters) {
	*out = *in
	if in.TimeFrame != nil {
		in, out := &in.TimeFrame, &out.TimeFrame
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExportExportDataOptionsParameters.
func (in *SubscriptionCostManagementExportExportDataOptionsParameters) DeepCopy() *SubscriptionCostManagementExportExportDataOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExportExportDataOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExportExportDataStorageLocationObservation) DeepCopyInto(out *SubscriptionCostManagementExportExportDataStorageLocationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExportExportDataStorageLocationObservation.
func (in *SubscriptionCostManagementExportExportDataStorageLocationObservation) DeepCopy() *SubscriptionCostManagementExportExportDataStorageLocationObservation {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExportExportDataStorageLocationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExportExportDataStorageLocationParameters) DeepCopyInto(out *SubscriptionCostManagementExportExportDataStorageLocationParameters) {
	*out = *in
	if in.ContainerID != nil {
		in, out := &in.ContainerID, &out.ContainerID
		*out = new(string)
		**out = **in
	}
	if in.ContainerIDRef != nil {
		in, out := &in.ContainerIDRef, &out.ContainerIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerIDSelector != nil {
		in, out := &in.ContainerIDSelector, &out.ContainerIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RootFolderPath != nil {
		in, out := &in.RootFolderPath, &out.RootFolderPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExportExportDataStorageLocationParameters.
func (in *SubscriptionCostManagementExportExportDataStorageLocationParameters) DeepCopy() *SubscriptionCostManagementExportExportDataStorageLocationParameters {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExportExportDataStorageLocationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExportList) DeepCopyInto(out *SubscriptionCostManagementExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SubscriptionCostManagementExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExportList.
func (in *SubscriptionCostManagementExportList) DeepCopy() *SubscriptionCostManagementExportList {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscriptionCostManagementExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExportObservation) DeepCopyInto(out *SubscriptionCostManagementExportObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExportObservation.
func (in *SubscriptionCostManagementExportObservation) DeepCopy() *SubscriptionCostManagementExportObservation {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExportObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExportParameters) DeepCopyInto(out *SubscriptionCostManagementExportParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.ExportDataOptions != nil {
		in, out := &in.ExportDataOptions, &out.ExportDataOptions
		*out = make([]SubscriptionCostManagementExportExportDataOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExportDataStorageLocation != nil {
		in, out := &in.ExportDataStorageLocation, &out.ExportDataStorageLocation
		*out = make([]SubscriptionCostManagementExportExportDataStorageLocationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecurrencePeriodEndDate != nil {
		in, out := &in.RecurrencePeriodEndDate, &out.RecurrencePeriodEndDate
		*out = new(string)
		**out = **in
	}
	if in.RecurrencePeriodStartDate != nil {
		in, out := &in.RecurrencePeriodStartDate, &out.RecurrencePeriodStartDate
		*out = new(string)
		**out = **in
	}
	if in.RecurrenceType != nil {
		in, out := &in.RecurrenceType, &out.RecurrenceType
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionID != nil {
		in, out := &in.SubscriptionID, &out.SubscriptionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExportParameters.
func (in *SubscriptionCostManagementExportParameters) DeepCopy() *SubscriptionCostManagementExportParameters {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExportParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExportSpec) DeepCopyInto(out *SubscriptionCostManagementExportSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExportSpec.
func (in *SubscriptionCostManagementExportSpec) DeepCopy() *SubscriptionCostManagementExportSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCostManagementExportStatus) DeepCopyInto(out *SubscriptionCostManagementExportStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCostManagementExportStatus.
func (in *SubscriptionCostManagementExportStatus) DeepCopy() *SubscriptionCostManagementExportStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCostManagementExportStatus)
	in.DeepCopyInto(out)
	return out
}