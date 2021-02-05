/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReplicaSets implements ReplicaSetInterface
type FakeReplicaSets struct {
	Fake *FakeAppsV1
	ns   string
	te   string
}

var replicasetsResource = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "replicasets"}

var replicasetsKind = schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "ReplicaSet"}

// Get takes name of the replicaSet, and returns the corresponding replicaSet object, and an error if there is any.
func (c *FakeReplicaSets) Get(name string, options v1.GetOptions) (result *appsv1.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithMultiTenancy(replicasetsResource, c.ns, name, c.te), &appsv1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}

	return obj.(*appsv1.ReplicaSet), err
}

// List takes label and field selectors, and returns the list of ReplicaSets that match those selectors.
func (c *FakeReplicaSets) List(opts v1.ListOptions) (result *appsv1.ReplicaSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithMultiTenancy(replicasetsResource, replicasetsKind, c.ns, opts, c.te), &appsv1.ReplicaSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1.ReplicaSetList{ListMeta: obj.(*appsv1.ReplicaSetList).ListMeta}
	for _, item := range obj.(*appsv1.ReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.AggregatedWatchInterface that watches the requested replicaSets.
func (c *FakeReplicaSets) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	aggWatch := watch.NewAggregatedWatcher()
	watcher, err := c.Fake.
		InvokesWatch(testing.NewWatchActionWithMultiTenancy(replicasetsResource, c.ns, opts, c.te))

	aggWatch.AddWatchInterface(watcher, err)
	return aggWatch
}

// Create takes the representation of a replicaSet and creates it.  Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *FakeReplicaSets) Create(replicaSet *appsv1.ReplicaSet) (result *appsv1.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithMultiTenancy(replicasetsResource, c.ns, replicaSet, c.te), &appsv1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}

	return obj.(*appsv1.ReplicaSet), err
}

// Update takes the representation of a replicaSet and updates it. Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *FakeReplicaSets) Update(replicaSet *appsv1.ReplicaSet) (result *appsv1.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithMultiTenancy(replicasetsResource, c.ns, replicaSet, c.te), &appsv1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}

	return obj.(*appsv1.ReplicaSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicaSets) UpdateStatus(replicaSet *appsv1.ReplicaSet) (*appsv1.ReplicaSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithMultiTenancy(replicasetsResource, "status", c.ns, replicaSet, c.te), &appsv1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.ReplicaSet), err
}

// Delete takes name of the replicaSet and deletes it. Returns an error if one occurs.
func (c *FakeReplicaSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithMultiTenancy(replicasetsResource, c.ns, name, c.te), &appsv1.ReplicaSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicaSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithMultiTenancy(replicasetsResource, c.ns, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &appsv1.ReplicaSetList{})
	return err
}

// Patch applies the patch and returns the patched replicaSet.
func (c *FakeReplicaSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *appsv1.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithMultiTenancy(replicasetsResource, c.te, c.ns, name, pt, data, subresources...), &appsv1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}

	return obj.(*appsv1.ReplicaSet), err
}

// GetScale takes name of the replicaSet, and returns the corresponding scale object, and an error if there is any.
func (c *FakeReplicaSets) GetScale(replicaSetName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithMultiTenancy(replicasetsResource, c.ns, "scale", replicaSetName, c.te), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}

	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeReplicaSets) UpdateScale(replicaSetName string, scale *autoscalingv1.Scale) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithMultiTenancy(replicasetsResource, "scale", c.ns, scale, c.te), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}

	return obj.(*autoscalingv1.Scale), err
}
