//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CTlog) DeepCopyInto(out *CTlog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CTlog.
func (in *CTlog) DeepCopy() *CTlog {
	if in == nil {
		return nil
	}
	out := new(CTlog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CTlog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CTlogList) DeepCopyInto(out *CTlogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CTlog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CTlogList.
func (in *CTlogList) DeepCopy() *CTlogList {
	if in == nil {
		return nil
	}
	out := new(CTlogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CTlogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CTlogSpec) DeepCopyInto(out *CTlogSpec) {
	*out = *in
	if in.TreeID != nil {
		in, out := &in.TreeID, &out.TreeID
		*out = new(int64)
		**out = **in
	}
	if in.PrivateKeyRef != nil {
		in, out := &in.PrivateKeyRef, &out.PrivateKeyRef
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.PrivateKeyPasswordRef != nil {
		in, out := &in.PrivateKeyPasswordRef, &out.PrivateKeyPasswordRef
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.PublicKeyRef != nil {
		in, out := &in.PublicKeyRef, &out.PublicKeyRef
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.RootCertificates != nil {
		in, out := &in.RootCertificates, &out.RootCertificates
		*out = make([]SecretKeySelector, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CTlogSpec.
func (in *CTlogSpec) DeepCopy() *CTlogSpec {
	if in == nil {
		return nil
	}
	out := new(CTlogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CTlogStatus) DeepCopyInto(out *CTlogStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CTlogStatus.
func (in *CTlogStatus) DeepCopy() *CTlogStatus {
	if in == nil {
		return nil
	}
	out := new(CTlogStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAccess) DeepCopyInto(out *ExternalAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAccess.
func (in *ExternalAccess) DeepCopy() *ExternalAccess {
	if in == nil {
		return nil
	}
	out := new(ExternalAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fulcio) DeepCopyInto(out *Fulcio) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fulcio.
func (in *Fulcio) DeepCopy() *Fulcio {
	if in == nil {
		return nil
	}
	out := new(Fulcio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fulcio) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FulcioCert) DeepCopyInto(out *FulcioCert) {
	*out = *in
	if in.PrivateKeyRef != nil {
		in, out := &in.PrivateKeyRef, &out.PrivateKeyRef
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.PrivateKeyPasswordRef != nil {
		in, out := &in.PrivateKeyPasswordRef, &out.PrivateKeyPasswordRef
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.CARef != nil {
		in, out := &in.CARef, &out.CARef
		*out = new(SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FulcioCert.
func (in *FulcioCert) DeepCopy() *FulcioCert {
	if in == nil {
		return nil
	}
	out := new(FulcioCert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FulcioConfig) DeepCopyInto(out *FulcioConfig) {
	*out = *in
	if in.OIDCIssuers != nil {
		in, out := &in.OIDCIssuers, &out.OIDCIssuers
		*out = make(map[string]OIDCIssuer, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MetaIssuers != nil {
		in, out := &in.MetaIssuers, &out.MetaIssuers
		*out = make(map[string]OIDCIssuer, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FulcioConfig.
func (in *FulcioConfig) DeepCopy() *FulcioConfig {
	if in == nil {
		return nil
	}
	out := new(FulcioConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FulcioList) DeepCopyInto(out *FulcioList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fulcio, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FulcioList.
func (in *FulcioList) DeepCopy() *FulcioList {
	if in == nil {
		return nil
	}
	out := new(FulcioList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FulcioList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FulcioSpec) DeepCopyInto(out *FulcioSpec) {
	*out = *in
	out.ExternalAccess = in.ExternalAccess
	in.Config.DeepCopyInto(&out.Config)
	in.Certificate.DeepCopyInto(&out.Certificate)
	out.Monitoring = in.Monitoring
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FulcioSpec.
func (in *FulcioSpec) DeepCopy() *FulcioSpec {
	if in == nil {
		return nil
	}
	out := new(FulcioSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FulcioStatus) DeepCopyInto(out *FulcioStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FulcioStatus.
func (in *FulcioStatus) DeepCopy() *FulcioStatus {
	if in == nil {
		return nil
	}
	out := new(FulcioStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfig) DeepCopyInto(out *MonitoringConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfig.
func (in *MonitoringConfig) DeepCopy() *MonitoringConfig {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCIssuer) DeepCopyInto(out *OIDCIssuer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCIssuer.
func (in *OIDCIssuer) DeepCopy() *OIDCIssuer {
	if in == nil {
		return nil
	}
	out := new(OIDCIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pvc) DeepCopyInto(out *Pvc) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pvc.
func (in *Pvc) DeepCopy() *Pvc {
	if in == nil {
		return nil
	}
	out := new(Pvc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rekor) DeepCopyInto(out *Rekor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rekor.
func (in *Rekor) DeepCopy() *Rekor {
	if in == nil {
		return nil
	}
	out := new(Rekor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Rekor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RekorList) DeepCopyInto(out *RekorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rekor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RekorList.
func (in *RekorList) DeepCopy() *RekorList {
	if in == nil {
		return nil
	}
	out := new(RekorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RekorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RekorSearchUI) DeepCopyInto(out *RekorSearchUI) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RekorSearchUI.
func (in *RekorSearchUI) DeepCopy() *RekorSearchUI {
	if in == nil {
		return nil
	}
	out := new(RekorSearchUI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RekorSigner) DeepCopyInto(out *RekorSigner) {
	*out = *in
	if in.PasswordRef != nil {
		in, out := &in.PasswordRef, &out.PasswordRef
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.KeyRef != nil {
		in, out := &in.KeyRef, &out.KeyRef
		*out = new(SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RekorSigner.
func (in *RekorSigner) DeepCopy() *RekorSigner {
	if in == nil {
		return nil
	}
	out := new(RekorSigner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RekorSpec) DeepCopyInto(out *RekorSpec) {
	*out = *in
	if in.TreeID != nil {
		in, out := &in.TreeID, &out.TreeID
		*out = new(int64)
		**out = **in
	}
	out.ExternalAccess = in.ExternalAccess
	out.Monitoring = in.Monitoring
	out.RekorSearchUI = in.RekorSearchUI
	in.Signer.DeepCopyInto(&out.Signer)
	out.Pvc = in.Pvc
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RekorSpec.
func (in *RekorSpec) DeepCopy() *RekorSpec {
	if in == nil {
		return nil
	}
	out := new(RekorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RekorStatus) DeepCopyInto(out *RekorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RekorStatus.
func (in *RekorStatus) DeepCopy() *RekorStatus {
	if in == nil {
		return nil
	}
	out := new(RekorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Securesign) DeepCopyInto(out *Securesign) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Securesign.
func (in *Securesign) DeepCopy() *Securesign {
	if in == nil {
		return nil
	}
	out := new(Securesign)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Securesign) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuresignList) DeepCopyInto(out *SecuresignList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Securesign, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuresignList.
func (in *SecuresignList) DeepCopy() *SecuresignList {
	if in == nil {
		return nil
	}
	out := new(SecuresignList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecuresignList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuresignSpec) DeepCopyInto(out *SecuresignSpec) {
	*out = *in
	in.Rekor.DeepCopyInto(&out.Rekor)
	in.Fulcio.DeepCopyInto(&out.Fulcio)
	in.Trillian.DeepCopyInto(&out.Trillian)
	in.Tuf.DeepCopyInto(&out.Tuf)
	in.Ctlog.DeepCopyInto(&out.Ctlog)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuresignSpec.
func (in *SecuresignSpec) DeepCopy() *SecuresignSpec {
	if in == nil {
		return nil
	}
	out := new(SecuresignSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuresignStatus) DeepCopyInto(out *SecuresignStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuresignStatus.
func (in *SecuresignStatus) DeepCopy() *SecuresignStatus {
	if in == nil {
		return nil
	}
	out := new(SecuresignStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trillian) DeepCopyInto(out *Trillian) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trillian.
func (in *Trillian) DeepCopy() *Trillian {
	if in == nil {
		return nil
	}
	out := new(Trillian)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trillian) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrillianDB) DeepCopyInto(out *TrillianDB) {
	*out = *in
	if in.DatabaseSecretRef != nil {
		in, out := &in.DatabaseSecretRef, &out.DatabaseSecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	out.Pvc = in.Pvc
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrillianDB.
func (in *TrillianDB) DeepCopy() *TrillianDB {
	if in == nil {
		return nil
	}
	out := new(TrillianDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrillianList) DeepCopyInto(out *TrillianList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trillian, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrillianList.
func (in *TrillianList) DeepCopy() *TrillianList {
	if in == nil {
		return nil
	}
	out := new(TrillianList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrillianList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrillianSpec) DeepCopyInto(out *TrillianSpec) {
	*out = *in
	in.Db.DeepCopyInto(&out.Db)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrillianSpec.
func (in *TrillianSpec) DeepCopy() *TrillianSpec {
	if in == nil {
		return nil
	}
	out := new(TrillianSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrillianStatus) DeepCopyInto(out *TrillianStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrillianStatus.
func (in *TrillianStatus) DeepCopy() *TrillianStatus {
	if in == nil {
		return nil
	}
	out := new(TrillianStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tuf) DeepCopyInto(out *Tuf) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tuf.
func (in *Tuf) DeepCopy() *Tuf {
	if in == nil {
		return nil
	}
	out := new(Tuf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tuf) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TufKey) DeepCopyInto(out *TufKey) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TufKey.
func (in *TufKey) DeepCopy() *TufKey {
	if in == nil {
		return nil
	}
	out := new(TufKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TufList) DeepCopyInto(out *TufList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tuf, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TufList.
func (in *TufList) DeepCopy() *TufList {
	if in == nil {
		return nil
	}
	out := new(TufList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TufList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TufSpec) DeepCopyInto(out *TufSpec) {
	*out = *in
	out.ExternalAccess = in.ExternalAccess
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]TufKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TufSpec.
func (in *TufSpec) DeepCopy() *TufSpec {
	if in == nil {
		return nil
	}
	out := new(TufSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TufStatus) DeepCopyInto(out *TufStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TufStatus.
func (in *TufStatus) DeepCopy() *TufStatus {
	if in == nil {
		return nil
	}
	out := new(TufStatus)
	in.DeepCopyInto(out)
	return out
}
