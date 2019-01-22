/*
Copyright 2019 The OpenEBS Authors

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

package v1beta1

import (
	v1beta1 "github.com/openebs/maya/pkg/apis/openebs.io/runtask/v1beta1"
	scheme "github.com/openebs/maya/pkg/client/generated/openebs.io/runtask/v1beta1/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RunTasksGetter has a method to return a RunTaskInterface.
// A group's client should implement this interface.
type RunTasksGetter interface {
	RunTasks(namespace string) RunTaskInterface
}

// RunTaskInterface has methods to work with RunTask resources.
type RunTaskInterface interface {
	Create(*v1beta1.RunTask) (*v1beta1.RunTask, error)
	Update(*v1beta1.RunTask) (*v1beta1.RunTask, error)
	UpdateStatus(*v1beta1.RunTask) (*v1beta1.RunTask, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.RunTask, error)
	List(opts v1.ListOptions) (*v1beta1.RunTaskList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RunTask, err error)
	RunTaskExpansion
}

// runTasks implements RunTaskInterface
type runTasks struct {
	client rest.Interface
	ns     string
}

// newRunTasks returns a RunTasks
func newRunTasks(c *OpenebsV1beta1Client, namespace string) *runTasks {
	return &runTasks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the runTask, and returns the corresponding runTask object, and an error if there is any.
func (c *runTasks) Get(name string, options v1.GetOptions) (result *v1beta1.RunTask, err error) {
	result = &v1beta1.RunTask{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("runtasks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RunTasks that match those selectors.
func (c *runTasks) List(opts v1.ListOptions) (result *v1beta1.RunTaskList, err error) {
	result = &v1beta1.RunTaskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("runtasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested runTasks.
func (c *runTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("runtasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a runTask and creates it.  Returns the server's representation of the runTask, and an error, if there is any.
func (c *runTasks) Create(runTask *v1beta1.RunTask) (result *v1beta1.RunTask, err error) {
	result = &v1beta1.RunTask{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("runtasks").
		Body(runTask).
		Do().
		Into(result)
	return
}

// Update takes the representation of a runTask and updates it. Returns the server's representation of the runTask, and an error, if there is any.
func (c *runTasks) Update(runTask *v1beta1.RunTask) (result *v1beta1.RunTask, err error) {
	result = &v1beta1.RunTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("runtasks").
		Name(runTask.Name).
		Body(runTask).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *runTasks) UpdateStatus(runTask *v1beta1.RunTask) (result *v1beta1.RunTask, err error) {
	result = &v1beta1.RunTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("runtasks").
		Name(runTask.Name).
		SubResource("status").
		Body(runTask).
		Do().
		Into(result)
	return
}

// Delete takes name of the runTask and deletes it. Returns an error if one occurs.
func (c *runTasks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("runtasks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *runTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("runtasks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched runTask.
func (c *runTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RunTask, err error) {
	result = &v1beta1.RunTask{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("runtasks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
