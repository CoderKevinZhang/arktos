// +build !ignore_autogenerated

/*
Copyright 2020 Authors of Arktos.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	clusterv1 "k8s.io/kubernetes/globalscheduler/pkg/apis/cluster/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUnion) DeepCopyInto(out *ClusterUnion) {
	*out = *in
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GeoLocation != nil {
		in, out := &in.GeoLocation, &out.GeoLocation
		*out = make([]*clusterv1.GeolocationInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(clusterv1.GeolocationInfo)
				**out = **in
			}
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = make([]*clusterv1.RegionInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(clusterv1.RegionInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = make([]*clusterv1.OperatorInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(clusterv1.OperatorInfo)
				**out = **in
			}
		}
	}
	if in.Flavors != nil {
		in, out := &in.Flavors, &out.Flavors
		*out = make([][]*clusterv1.FlavorInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]*clusterv1.FlavorInfo, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(clusterv1.FlavorInfo)
						**out = **in
					}
				}
			}
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([][]*clusterv1.StorageSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]*clusterv1.StorageSpec, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(clusterv1.StorageSpec)
						**out = **in
					}
				}
			}
		}
	}
	if in.EipCapacity != nil {
		in, out := &in.EipCapacity, &out.EipCapacity
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.CPUCapacity != nil {
		in, out := &in.CPUCapacity, &out.CPUCapacity
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.MemCapacity != nil {
		in, out := &in.MemCapacity, &out.MemCapacity
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.ServerPrice != nil {
		in, out := &in.ServerPrice, &out.ServerPrice
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.HomeScheduler != nil {
		in, out := &in.HomeScheduler, &out.HomeScheduler
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUnion.
func (in *ClusterUnion) DeepCopy() *ClusterUnion {
	if in == nil {
		return nil
	}
	out := new(ClusterUnion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeolocationInfo) DeepCopyInto(out *GeolocationInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeolocationInfo.
func (in *GeolocationInfo) DeepCopy() *GeolocationInfo {
	if in == nil {
		return nil
	}
	out := new(GeolocationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scheduler) DeepCopyInto(out *Scheduler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler.
func (in *Scheduler) DeepCopy() *Scheduler {
	if in == nil {
		return nil
	}
	out := new(Scheduler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Scheduler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerList) DeepCopyInto(out *SchedulerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Scheduler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerList.
func (in *SchedulerList) DeepCopy() *SchedulerList {
	if in == nil {
		return nil
	}
	out := new(SchedulerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerSpec) DeepCopyInto(out *SchedulerSpec) {
	*out = *in
	out.Location = in.Location
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = make([]*clusterv1.Cluster, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(clusterv1.Cluster)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.Union.DeepCopyInto(&out.Union)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerSpec.
func (in *SchedulerSpec) DeepCopy() *SchedulerSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerStatus) DeepCopyInto(out *SchedulerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerStatus.
func (in *SchedulerStatus) DeepCopy() *SchedulerStatus {
	if in == nil {
		return nil
	}
	out := new(SchedulerStatus)
	in.DeepCopyInto(out)
	return out
}
