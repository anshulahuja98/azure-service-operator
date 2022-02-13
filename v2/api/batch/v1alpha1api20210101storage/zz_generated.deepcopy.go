//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20210101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStorageBaseProperties_Spec) DeepCopyInto(out *AutoStorageBaseProperties_Spec) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.StorageAccountReference = in.StorageAccountReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStorageBaseProperties_Spec.
func (in *AutoStorageBaseProperties_Spec) DeepCopy() *AutoStorageBaseProperties_Spec {
	if in == nil {
		return nil
	}
	out := new(AutoStorageBaseProperties_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoStorageBaseProperties_Status) DeepCopyInto(out *AutoStorageBaseProperties_Status) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StorageAccountId != nil {
		in, out := &in.StorageAccountId, &out.StorageAccountId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoStorageBaseProperties_Status.
func (in *AutoStorageBaseProperties_Status) DeepCopy() *AutoStorageBaseProperties_Status {
	if in == nil {
		return nil
	}
	out := new(AutoStorageBaseProperties_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccount) DeepCopyInto(out *BatchAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccount.
func (in *BatchAccount) DeepCopy() *BatchAccount {
	if in == nil {
		return nil
	}
	out := new(BatchAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountCreateParameters_Status) DeepCopyInto(out *BatchAccountCreateParameters_Status) {
	*out = *in
	if in.AutoStorage != nil {
		in, out := &in.AutoStorage, &out.AutoStorage
		*out = new(AutoStorageBaseProperties_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentity_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReference_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountCreateParameters_Status.
func (in *BatchAccountCreateParameters_Status) DeepCopy() *BatchAccountCreateParameters_Status {
	if in == nil {
		return nil
	}
	out := new(BatchAccountCreateParameters_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_Spec) DeepCopyInto(out *BatchAccountIdentity_Spec) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_Spec.
func (in *BatchAccountIdentity_Spec) DeepCopy() *BatchAccountIdentity_Spec {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_Status) DeepCopyInto(out *BatchAccountIdentity_Status) {
	*out = *in
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]BatchAccountIdentity_UserAssignedIdentities_Status, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_Status.
func (in *BatchAccountIdentity_Status) DeepCopy() *BatchAccountIdentity_Status {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountIdentity_UserAssignedIdentities_Status) DeepCopyInto(out *BatchAccountIdentity_UserAssignedIdentities_Status) {
	*out = *in
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountIdentity_UserAssignedIdentities_Status.
func (in *BatchAccountIdentity_UserAssignedIdentities_Status) DeepCopy() *BatchAccountIdentity_UserAssignedIdentities_Status {
	if in == nil {
		return nil
	}
	out := new(BatchAccountIdentity_UserAssignedIdentities_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccountList) DeepCopyInto(out *BatchAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BatchAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccountList.
func (in *BatchAccountList) DeepCopy() *BatchAccountList {
	if in == nil {
		return nil
	}
	out := new(BatchAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchAccounts_SPEC) DeepCopyInto(out *BatchAccounts_SPEC) {
	*out = *in
	if in.AutoStorage != nil {
		in, out := &in.AutoStorage, &out.AutoStorage
		*out = new(AutoStorageBaseProperties_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionProperties_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(BatchAccountIdentity_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(KeyVaultReference_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	out.Owner = in.Owner
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchAccounts_SPEC.
func (in *BatchAccounts_SPEC) DeepCopy() *BatchAccounts_SPEC {
	if in == nil {
		return nil
	}
	out := new(BatchAccounts_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProperties_Spec) DeepCopyInto(out *EncryptionProperties_Spec) {
	*out = *in
	if in.KeySource != nil {
		in, out := &in.KeySource, &out.KeySource
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultProperties_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProperties_Spec.
func (in *EncryptionProperties_Spec) DeepCopy() *EncryptionProperties_Spec {
	if in == nil {
		return nil
	}
	out := new(EncryptionProperties_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionProperties_Status) DeepCopyInto(out *EncryptionProperties_Status) {
	*out = *in
	if in.KeySource != nil {
		in, out := &in.KeySource, &out.KeySource
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultProperties != nil {
		in, out := &in.KeyVaultProperties, &out.KeyVaultProperties
		*out = new(KeyVaultProperties_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionProperties_Status.
func (in *EncryptionProperties_Status) DeepCopy() *EncryptionProperties_Status {
	if in == nil {
		return nil
	}
	out := new(EncryptionProperties_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultProperties_Spec) DeepCopyInto(out *KeyVaultProperties_Spec) {
	*out = *in
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultProperties_Spec.
func (in *KeyVaultProperties_Spec) DeepCopy() *KeyVaultProperties_Spec {
	if in == nil {
		return nil
	}
	out := new(KeyVaultProperties_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultProperties_Status) DeepCopyInto(out *KeyVaultProperties_Status) {
	*out = *in
	if in.KeyIdentifier != nil {
		in, out := &in.KeyIdentifier, &out.KeyIdentifier
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultProperties_Status.
func (in *KeyVaultProperties_Status) DeepCopy() *KeyVaultProperties_Status {
	if in == nil {
		return nil
	}
	out := new(KeyVaultProperties_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultReference_Spec) DeepCopyInto(out *KeyVaultReference_Spec) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Reference = in.Reference
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultReference_Spec.
func (in *KeyVaultReference_Spec) DeepCopy() *KeyVaultReference_Spec {
	if in == nil {
		return nil
	}
	out := new(KeyVaultReference_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultReference_Status) DeepCopyInto(out *KeyVaultReference_Status) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultReference_Status.
func (in *KeyVaultReference_Status) DeepCopy() *KeyVaultReference_Status {
	if in == nil {
		return nil
	}
	out := new(KeyVaultReference_Status)
	in.DeepCopyInto(out)
	return out
}
