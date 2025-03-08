/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "go.virtual-secrets.dev/apimachinery/apis/config/v1alpha1"
	scheme "go.virtual-secrets.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecretMetadatasGetter has a method to return a SecretMetadataInterface.
// A group's client should implement this interface.
type SecretMetadatasGetter interface {
	SecretMetadatas(namespace string) SecretMetadataInterface
}

// SecretMetadataInterface has methods to work with SecretMetadata resources.
type SecretMetadataInterface interface {
	Create(ctx context.Context, secretMetadata *v1alpha1.SecretMetadata, opts v1.CreateOptions) (*v1alpha1.SecretMetadata, error)
	Update(ctx context.Context, secretMetadata *v1alpha1.SecretMetadata, opts v1.UpdateOptions) (*v1alpha1.SecretMetadata, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SecretMetadata, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SecretMetadataList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecretMetadata, err error)
	SecretMetadataExpansion
}

// secretMetadatas implements SecretMetadataInterface
type secretMetadatas struct {
	client rest.Interface
	ns     string
}

// newSecretMetadatas returns a SecretMetadatas
func newSecretMetadatas(c *ConfigV1alpha1Client, namespace string) *secretMetadatas {
	return &secretMetadatas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secretMetadata, and returns the corresponding secretMetadata object, and an error if there is any.
func (c *secretMetadatas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SecretMetadata, err error) {
	result = &v1alpha1.SecretMetadata{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretmetadatas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecretMetadatas that match those selectors.
func (c *secretMetadatas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecretMetadataList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecretMetadataList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretmetadatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secretMetadatas.
func (c *secretMetadatas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("secretmetadatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a secretMetadata and creates it.  Returns the server's representation of the secretMetadata, and an error, if there is any.
func (c *secretMetadatas) Create(ctx context.Context, secretMetadata *v1alpha1.SecretMetadata, opts v1.CreateOptions) (result *v1alpha1.SecretMetadata, err error) {
	result = &v1alpha1.SecretMetadata{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("secretmetadatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secretMetadata).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a secretMetadata and updates it. Returns the server's representation of the secretMetadata, and an error, if there is any.
func (c *secretMetadatas) Update(ctx context.Context, secretMetadata *v1alpha1.SecretMetadata, opts v1.UpdateOptions) (result *v1alpha1.SecretMetadata, err error) {
	result = &v1alpha1.SecretMetadata{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secretmetadatas").
		Name(secretMetadata.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secretMetadata).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the secretMetadata and deletes it. Returns an error if one occurs.
func (c *secretMetadatas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretmetadatas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secretMetadatas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretmetadatas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched secretMetadata.
func (c *secretMetadatas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecretMetadata, err error) {
	result = &v1alpha1.SecretMetadata{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("secretmetadatas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
