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
	v1alpha1 "go.virtual-secrets.dev/apimachinery/apis/config/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SecretSourceLister helps list SecretSources.
// All objects returned here must be treated as read-only.
type SecretSourceLister interface {
	// List lists all SecretSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SecretSource, err error)
	// Get retrieves the SecretSource from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SecretSource, error)
	SecretSourceListerExpansion
}

// secretSourceLister implements the SecretSourceLister interface.
type secretSourceLister struct {
	indexer cache.Indexer
}

// NewSecretSourceLister returns a new SecretSourceLister.
func NewSecretSourceLister(indexer cache.Indexer) SecretSourceLister {
	return &secretSourceLister{indexer: indexer}
}

// List lists all SecretSources in the indexer.
func (s *secretSourceLister) List(selector labels.Selector) (ret []*v1alpha1.SecretSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretSource))
	})
	return ret, err
}

// Get retrieves the SecretSource from the index for a given name.
func (s *secretSourceLister) Get(name string) (*v1alpha1.SecretSource, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("secretsource"), name)
	}
	return obj.(*v1alpha1.SecretSource), nil
}
