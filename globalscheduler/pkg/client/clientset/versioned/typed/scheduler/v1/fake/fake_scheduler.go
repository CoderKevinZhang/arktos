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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	schedulerv1 "k8s.io/kubernetes/globalscheduler/pkg/apis/scheduler/v1"
)

// FakeSchedulers implements SchedulerInterface
type FakeSchedulers struct {
	Fake *FakeGlobalschedulerV1
	ns   string
	te   string
}

var schedulersResource = schema.GroupVersionResource{Group: "globalscheduler.com", Version: "v1", Resource: "schedulers"}

var schedulersKind = schema.GroupVersionKind{Group: "globalscheduler.com", Version: "v1", Kind: "Scheduler"}

// Get takes name of the scheduler, and returns the corresponding scheduler object, and an error if there is any.
func (c *FakeSchedulers) Get(name string, options v1.GetOptions) (result *schedulerv1.Scheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithMultiTenancy(schedulersResource, c.ns, name, c.te), &schedulerv1.Scheduler{})

	if obj == nil {
		return nil, err
	}

	return obj.(*schedulerv1.Scheduler), err
}

// List takes label and field selectors, and returns the list of Schedulers that match those selectors.
func (c *FakeSchedulers) List(opts v1.ListOptions) (result *schedulerv1.SchedulerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithMultiTenancy(schedulersResource, schedulersKind, c.ns, opts, c.te), &schedulerv1.SchedulerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &schedulerv1.SchedulerList{ListMeta: obj.(*schedulerv1.SchedulerList).ListMeta}
	for _, item := range obj.(*schedulerv1.SchedulerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.AggregatedWatchInterface that watches the requested schedulers.
func (c *FakeSchedulers) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	aggWatch := watch.NewAggregatedWatcher()
	watcher, err := c.Fake.
		InvokesWatch(testing.NewWatchActionWithMultiTenancy(schedulersResource, c.ns, opts, c.te))

	aggWatch.AddWatchInterface(watcher, err)
	return aggWatch
}

// Create takes the representation of a scheduler and creates it.  Returns the server's representation of the scheduler, and an error, if there is any.
func (c *FakeSchedulers) Create(scheduler *schedulerv1.Scheduler) (result *schedulerv1.Scheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithMultiTenancy(schedulersResource, c.ns, scheduler, c.te), &schedulerv1.Scheduler{})

	if obj == nil {
		return nil, err
	}

	return obj.(*schedulerv1.Scheduler), err
}

// Update takes the representation of a scheduler and updates it. Returns the server's representation of the scheduler, and an error, if there is any.
func (c *FakeSchedulers) Update(scheduler *schedulerv1.Scheduler) (result *schedulerv1.Scheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithMultiTenancy(schedulersResource, c.ns, scheduler, c.te), &schedulerv1.Scheduler{})

	if obj == nil {
		return nil, err
	}

	return obj.(*schedulerv1.Scheduler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSchedulers) UpdateStatus(scheduler *schedulerv1.Scheduler) (*schedulerv1.Scheduler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithMultiTenancy(schedulersResource, "status", c.ns, scheduler, c.te), &schedulerv1.Scheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*schedulerv1.Scheduler), err
}

// Delete takes name of the scheduler and deletes it. Returns an error if one occurs.
func (c *FakeSchedulers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithMultiTenancy(schedulersResource, c.ns, name, c.te), &schedulerv1.Scheduler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSchedulers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithMultiTenancy(schedulersResource, c.ns, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &schedulerv1.SchedulerList{})
	return err
}

// Patch applies the patch and returns the patched scheduler.
func (c *FakeSchedulers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *schedulerv1.Scheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithMultiTenancy(schedulersResource, c.te, c.ns, name, pt, data, subresources...), &schedulerv1.Scheduler{})

	if obj == nil {
		return nil, err
	}

	return obj.(*schedulerv1.Scheduler), err
}
