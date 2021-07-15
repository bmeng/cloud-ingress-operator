// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIScheme) DeepCopyInto(out *APIScheme) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIScheme.
func (in *APIScheme) DeepCopy() *APIScheme {
	if in == nil {
		return nil
	}
	out := new(APIScheme)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIScheme) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISchemeCondition) DeepCopyInto(out *APISchemeCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	if in.AllowedCIDRBlocks != nil {
		in, out := &in.AllowedCIDRBlocks, &out.AllowedCIDRBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISchemeCondition.
func (in *APISchemeCondition) DeepCopy() *APISchemeCondition {
	if in == nil {
		return nil
	}
	out := new(APISchemeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISchemeList) DeepCopyInto(out *APISchemeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIScheme, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISchemeList.
func (in *APISchemeList) DeepCopy() *APISchemeList {
	if in == nil {
		return nil
	}
	out := new(APISchemeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APISchemeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISchemeSpec) DeepCopyInto(out *APISchemeSpec) {
	*out = *in
	in.ManagementAPIServerIngress.DeepCopyInto(&out.ManagementAPIServerIngress)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISchemeSpec.
func (in *APISchemeSpec) DeepCopy() *APISchemeSpec {
	if in == nil {
		return nil
	}
	out := new(APISchemeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISchemeStatus) DeepCopyInto(out *APISchemeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]APISchemeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISchemeStatus.
func (in *APISchemeStatus) DeepCopy() *APISchemeStatus {
	if in == nil {
		return nil
	}
	out := new(APISchemeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationIngress) DeepCopyInto(out *ApplicationIngress) {
	*out = *in
	out.Certificate = in.Certificate
	in.RouteSelector.DeepCopyInto(&out.RouteSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationIngress.
func (in *ApplicationIngress) DeepCopy() *ApplicationIngress {
	if in == nil {
		return nil
	}
	out := new(ApplicationIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultAPIServerIngress) DeepCopyInto(out *DefaultAPIServerIngress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultAPIServerIngress.
func (in *DefaultAPIServerIngress) DeepCopy() *DefaultAPIServerIngress {
	if in == nil {
		return nil
	}
	out := new(DefaultAPIServerIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementAPIServerIngress) DeepCopyInto(out *ManagementAPIServerIngress) {
	*out = *in
	if in.AllowedCIDRBlocks != nil {
		in, out := &in.AllowedCIDRBlocks, &out.AllowedCIDRBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementAPIServerIngress.
func (in *ManagementAPIServerIngress) DeepCopy() *ManagementAPIServerIngress {
	if in == nil {
		return nil
	}
	out := new(ManagementAPIServerIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishingStrategy) DeepCopyInto(out *PublishingStrategy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishingStrategy.
func (in *PublishingStrategy) DeepCopy() *PublishingStrategy {
	if in == nil {
		return nil
	}
	out := new(PublishingStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublishingStrategy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishingStrategyList) DeepCopyInto(out *PublishingStrategyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PublishingStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishingStrategyList.
func (in *PublishingStrategyList) DeepCopy() *PublishingStrategyList {
	if in == nil {
		return nil
	}
	out := new(PublishingStrategyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublishingStrategyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishingStrategySpec) DeepCopyInto(out *PublishingStrategySpec) {
	*out = *in
	out.DefaultAPIServerIngress = in.DefaultAPIServerIngress
	if in.ApplicationIngress != nil {
		in, out := &in.ApplicationIngress, &out.ApplicationIngress
		*out = make([]ApplicationIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishingStrategySpec.
func (in *PublishingStrategySpec) DeepCopy() *PublishingStrategySpec {
	if in == nil {
		return nil
	}
	out := new(PublishingStrategySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishingStrategyStatus) DeepCopyInto(out *PublishingStrategyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishingStrategyStatus.
func (in *PublishingStrategyStatus) DeepCopy() *PublishingStrategyStatus {
	if in == nil {
		return nil
	}
	out := new(PublishingStrategyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHD) DeepCopyInto(out *SSHD) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHD.
func (in *SSHD) DeepCopy() *SSHD {
	if in == nil {
		return nil
	}
	out := new(SSHD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHD) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHDList) DeepCopyInto(out *SSHDList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SSHD, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHDList.
func (in *SSHDList) DeepCopy() *SSHDList {
	if in == nil {
		return nil
	}
	out := new(SSHDList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHDList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHDSpec) DeepCopyInto(out *SSHDSpec) {
	*out = *in
	if in.AllowedCIDRBlocks != nil {
		in, out := &in.AllowedCIDRBlocks, &out.AllowedCIDRBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ConfigMapSelector.DeepCopyInto(&out.ConfigMapSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHDSpec.
func (in *SSHDSpec) DeepCopy() *SSHDSpec {
	if in == nil {
		return nil
	}
	out := new(SSHDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHDStatus) DeepCopyInto(out *SSHDStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHDStatus.
func (in *SSHDStatus) DeepCopy() *SSHDStatus {
	if in == nil {
		return nil
	}
	out := new(SSHDStatus)
	in.DeepCopyInto(out)
	return out
}
