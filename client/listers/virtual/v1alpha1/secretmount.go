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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "go.virtual-secrets.dev/apimachinery/apis/virtual/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SecretMountLister helps list SecretMounts.
// All objects returned here must be treated as read-only.
type SecretMountLister interface {
	// List lists all SecretMounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SecretMount, err error)
	// SecretMounts returns an object that can list and get SecretMounts.
	SecretMounts(namespace string) SecretMountNamespaceLister
	SecretMountListerExpansion
}

// secretMountLister implements the SecretMountLister interface.
type secretMountLister struct {
	indexer cache.Indexer
}

// NewSecretMountLister returns a new SecretMountLister.
func NewSecretMountLister(indexer cache.Indexer) SecretMountLister {
	return &secretMountLister{indexer: indexer}
}

// List lists all SecretMounts in the indexer.
func (s *secretMountLister) List(selector labels.Selector) (ret []*v1alpha1.SecretMount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretMount))
	})
	return ret, err
}

// SecretMounts returns an object that can list and get SecretMounts.
func (s *secretMountLister) SecretMounts(namespace string) SecretMountNamespaceLister {
	return secretMountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecretMountNamespaceLister helps list and get SecretMounts.
// All objects returned here must be treated as read-only.
type SecretMountNamespaceLister interface {
	// List lists all SecretMounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SecretMount, err error)
	// Get retrieves the SecretMount from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SecretMount, error)
	SecretMountNamespaceListerExpansion
}

// secretMountNamespaceLister implements the SecretMountNamespaceLister
// interface.
type secretMountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecretMounts in the indexer for a given namespace.
func (s secretMountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecretMount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretMount))
	})
	return ret, err
}

// Get retrieves the SecretMount from the indexer for a given namespace and name.
func (s secretMountNamespaceLister) Get(name string) (*v1alpha1.SecretMount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("secretmount"), name)
	}
	return obj.(*v1alpha1.SecretMount), nil
}
