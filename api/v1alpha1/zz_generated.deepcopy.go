//go:build !ignore_autogenerated

/*
Copyright 2023 zncdata-labs.

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
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alluxio) DeepCopyInto(out *Alluxio) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alluxio.
func (in *Alluxio) DeepCopy() *Alluxio {
	if in == nil {
		return nil
	}
	out := new(Alluxio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alluxio) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioList) DeepCopyInto(out *AlluxioList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alluxio, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioList.
func (in *AlluxioList) DeepCopy() *AlluxioList {
	if in == nil {
		return nil
	}
	out := new(AlluxioList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlluxioList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioSpec) DeepCopyInto(out *AlluxioSpec) {
	*out = *in
	if in.ClusterConfig != nil {
		in, out := &in.ClusterConfig, &out.ClusterConfig
		*out = new(ClusterConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = new(MasterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = new(WorkerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioSpec.
func (in *AlluxioSpec) DeepCopy() *AlluxioSpec {
	if in == nil {
		return nil
	}
	out := new(AlluxioSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioStatus) DeepCopyInto(out *AlluxioStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URLs != nil {
		in, out := &in.URLs, &out.URLs
		*out = make([]StatusURL, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioStatus.
func (in *AlluxioStatus) DeepCopy() *AlluxioStatus {
	if in == nil {
		return nil
	}
	out := new(AlluxioStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigSpec) DeepCopyInto(out *ClusterConfigSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = new(v1.Toleration)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MasterPorts != nil {
		in, out := &in.MasterPorts, &out.MasterPorts
		*out = new(MasterPortsSpec)
		**out = **in
	}
	if in.WorkerPorts != nil {
		in, out := &in.WorkerPorts, &out.WorkerPorts
		*out = new(WorkerPortsSpec)
		**out = **in
	}
	if in.ExtraContainers != nil {
		in, out := &in.ExtraContainers, &out.ExtraContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.JobMaster != nil {
		in, out := &in.JobMaster, &out.JobMaster
		*out = new(JobMasterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JobWorker != nil {
		in, out := &in.JobWorker, &out.JobWorker
		*out = new(JobWorkerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TieredStore != nil {
		in, out := &in.TieredStore, &out.TieredStore
		*out = make([]*TieredStore, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TieredStore)
				**out = **in
			}
		}
	}
	if in.ShortCircuit != nil {
		in, out := &in.ShortCircuit, &out.ShortCircuit
		*out = new(ShortCircuitSpec)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Journal != nil {
		in, out := &in.Journal, &out.Journal
		*out = new(JournalSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigSpec.
func (in *ClusterConfigSpec) DeepCopy() *ClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobMasterPortsSpec) DeepCopyInto(out *JobMasterPortsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobMasterPortsSpec.
func (in *JobMasterPortsSpec) DeepCopy() *JobMasterPortsSpec {
	if in == nil {
		return nil
	}
	out := new(JobMasterPortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobMasterSpec) DeepCopyInto(out *JobMasterSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = new(JobMasterPortsSpec)
		**out = **in
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobMasterSpec.
func (in *JobMasterSpec) DeepCopy() *JobMasterSpec {
	if in == nil {
		return nil
	}
	out := new(JobMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobWorkerPortsSpec) DeepCopyInto(out *JobWorkerPortsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobWorkerPortsSpec.
func (in *JobWorkerPortsSpec) DeepCopy() *JobWorkerPortsSpec {
	if in == nil {
		return nil
	}
	out := new(JobWorkerPortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobWorkerSpec) DeepCopyInto(out *JobWorkerSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = new(JobWorkerPortsSpec)
		**out = **in
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobWorkerSpec.
func (in *JobWorkerSpec) DeepCopy() *JobWorkerSpec {
	if in == nil {
		return nil
	}
	out := new(JobWorkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JournalSpec) DeepCopyInto(out *JournalSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JournalSpec.
func (in *JournalSpec) DeepCopy() *JournalSpec {
	if in == nil {
		return nil
	}
	out := new(JournalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterPortsSpec) DeepCopyInto(out *MasterPortsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterPortsSpec.
func (in *MasterPortsSpec) DeepCopy() *MasterPortsSpec {
	if in == nil {
		return nil
	}
	out := new(MasterPortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterSpec) DeepCopyInto(out *MasterSpec) {
	*out = *in
	if in.RoleConfig != nil {
		in, out := &in.RoleConfig, &out.RoleConfig
		*out = new(RoleGroupMasterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*RoleGroupMasterSpec, len(*in))
		for key, val := range *in {
			var outVal *RoleGroupMasterSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(RoleGroupMasterSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterSpec.
func (in *MasterSpec) DeepCopy() *MasterSpec {
	if in == nil {
		return nil
	}
	out := new(MasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleConfigMasterSpec) DeepCopyInto(out *RoleConfigMasterSpec) {
	*out = *in
	if in.JobMaster != nil {
		in, out := &in.JobMaster, &out.JobMaster
		*out = new(JobMasterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleConfigMasterSpec.
func (in *RoleConfigMasterSpec) DeepCopy() *RoleConfigMasterSpec {
	if in == nil {
		return nil
	}
	out := new(RoleConfigMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleGroupMasterSpec) DeepCopyInto(out *RoleGroupMasterSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = new(v1.Toleration)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = new(MasterPortsSpec)
		**out = **in
	}
	if in.ExtraContainers != nil {
		in, out := &in.ExtraContainers, &out.ExtraContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.JobMaster != nil {
		in, out := &in.JobMaster, &out.JobMaster
		*out = new(JobMasterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleGroupMasterSpec.
func (in *RoleGroupMasterSpec) DeepCopy() *RoleGroupMasterSpec {
	if in == nil {
		return nil
	}
	out := new(RoleGroupMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleGroupWorkerSpec) DeepCopyInto(out *RoleGroupWorkerSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = new(v1.Toleration)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = new(WorkerPortsSpec)
		**out = **in
	}
	if in.ExtraContainers != nil {
		in, out := &in.ExtraContainers, &out.ExtraContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.JobWorker != nil {
		in, out := &in.JobWorker, &out.JobWorker
		*out = new(JobWorkerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostPID != nil {
		in, out := &in.HostPID, &out.HostPID
		*out = new(bool)
		**out = **in
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.ShareProcessNamespace != nil {
		in, out := &in.ShareProcessNamespace, &out.ShareProcessNamespace
		*out = new(bool)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleGroupWorkerSpec.
func (in *RoleGroupWorkerSpec) DeepCopy() *RoleGroupWorkerSpec {
	if in == nil {
		return nil
	}
	out := new(RoleGroupWorkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShortCircuitSpec) DeepCopyInto(out *ShortCircuitSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShortCircuitSpec.
func (in *ShortCircuitSpec) DeepCopy() *ShortCircuitSpec {
	if in == nil {
		return nil
	}
	out := new(ShortCircuitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusURL) DeepCopyInto(out *StatusURL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusURL.
func (in *StatusURL) DeepCopy() *StatusURL {
	if in == nil {
		return nil
	}
	out := new(StatusURL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TieredStore) DeepCopyInto(out *TieredStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TieredStore.
func (in *TieredStore) DeepCopy() *TieredStore {
	if in == nil {
		return nil
	}
	out := new(TieredStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPortsSpec) DeepCopyInto(out *WorkerPortsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPortsSpec.
func (in *WorkerPortsSpec) DeepCopy() *WorkerPortsSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerPortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSpec) DeepCopyInto(out *WorkerSpec) {
	*out = *in
	if in.RoleConfig != nil {
		in, out := &in.RoleConfig, &out.RoleConfig
		*out = new(RoleGroupWorkerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*RoleGroupWorkerSpec, len(*in))
		for key, val := range *in {
			var outVal *RoleGroupWorkerSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(RoleGroupWorkerSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSpec.
func (in *WorkerSpec) DeepCopy() *WorkerSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerSpec)
	in.DeepCopyInto(out)
	return out
}
