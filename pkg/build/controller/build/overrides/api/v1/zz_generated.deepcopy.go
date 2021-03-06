// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	build_v1 "github.com/openshift/origin/pkg/build/apis/build/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildOverridesConfig, InType: reflect.TypeOf(&BuildOverridesConfig{})},
	)
}

// DeepCopy_v1_BuildOverridesConfig is an autogenerated deepcopy function.
func DeepCopy_v1_BuildOverridesConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildOverridesConfig)
		out := out.(*BuildOverridesConfig)
		*out = *in
		if in.ImageLabels != nil {
			in, out := &in.ImageLabels, &out.ImageLabels
			*out = make([]build_v1.ImageLabel, len(*in))
			copy(*out, *in)
		}
		if in.NodeSelector != nil {
			in, out := &in.NodeSelector, &out.NodeSelector
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.Annotations != nil {
			in, out := &in.Annotations, &out.Annotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}
