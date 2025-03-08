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

package fake

import (
	"context"

	v1alpha1 "go.virtual-secrets.dev/apimachinery/apis/config/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSecretMetadatas implements SecretMetadataInterface
type FakeSecretMetadatas struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var secretmetadatasResource = v1alpha1.SchemeGroupVersion.WithResource("secretmetadatas")

var secretmetadatasKind = v1alpha1.SchemeGroupVersion.WithKind("SecretMetadata")

// Get takes name of the secretMetadata, and returns the corresponding secretMetadata object, and an error if there is any.
func (c *FakeSecretMetadatas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SecretMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretmetadatasResource, c.ns, name), &v1alpha1.SecretMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretMetadata), err
}

// List takes label and field selectors, and returns the list of SecretMetadatas that match those selectors.
func (c *FakeSecretMetadatas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecretMetadataList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretmetadatasResource, secretmetadatasKind, c.ns, opts), &v1alpha1.SecretMetadataList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecretMetadataList{ListMeta: obj.(*v1alpha1.SecretMetadataList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecretMetadataList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secretMetadatas.
func (c *FakeSecretMetadatas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretmetadatasResource, c.ns, opts))

}

// Create takes the representation of a secretMetadata and creates it.  Returns the server's representation of the secretMetadata, and an error, if there is any.
func (c *FakeSecretMetadatas) Create(ctx context.Context, secretMetadata *v1alpha1.SecretMetadata, opts v1.CreateOptions) (result *v1alpha1.SecretMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretmetadatasResource, c.ns, secretMetadata), &v1alpha1.SecretMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretMetadata), err
}

// Update takes the representation of a secretMetadata and updates it. Returns the server's representation of the secretMetadata, and an error, if there is any.
func (c *FakeSecretMetadatas) Update(ctx context.Context, secretMetadata *v1alpha1.SecretMetadata, opts v1.UpdateOptions) (result *v1alpha1.SecretMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretmetadatasResource, c.ns, secretMetadata), &v1alpha1.SecretMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretMetadata), err
}

// Delete takes name of the secretMetadata and deletes it. Returns an error if one occurs.
func (c *FakeSecretMetadatas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(secretmetadatasResource, c.ns, name, opts), &v1alpha1.SecretMetadata{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecretMetadatas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretmetadatasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecretMetadataList{})
	return err
}

// Patch applies the patch and returns the patched secretMetadata.
func (c *FakeSecretMetadatas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecretMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretmetadatasResource, c.ns, name, pt, data, subresources...), &v1alpha1.SecretMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretMetadata), err
}
