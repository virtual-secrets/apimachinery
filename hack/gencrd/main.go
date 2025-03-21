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

package main

import (
	"os"
	"path/filepath"

	configinstall "go.virtual-secrets.dev/apimachinery/apis/config/install"
	configapi "go.virtual-secrets.dev/apimachinery/apis/config/v1alpha1"
	vsainstall "go.virtual-secrets.dev/apimachinery/apis/virtual/install"
	vsapi "go.virtual-secrets.dev/apimachinery/apis/virtual/v1alpha1"

	gort "gomodules.xyz/runtime"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/klog/v2"
	"k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/validation/spec"
	"kmodules.xyz/client-go/openapi"
)

func generateSwaggerJson() {
	var (
		Scheme = runtime.NewScheme()
		Codecs = serializer.NewCodecFactory(Scheme)
	)

	configinstall.Install(Scheme)
	vsainstall.Install(Scheme)

	apispec, err := openapi.RenderOpenAPISpec(openapi.Config{
		Scheme: Scheme,
		Codecs: Codecs,
		Info: spec.InfoProps{
			Title:   "KubeDB",
			Version: "v0",
			Contact: &spec.ContactInfo{
				Name:  "AppsCode Inc.",
				URL:   "https://appscode.com",
				Email: "hello@appscode.com",
			},
			License: &spec.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
			},
		},
		OpenAPIDefinitions: []common.GetOpenAPIDefinitions{
			configapi.GetOpenAPIDefinitions,
			vsapi.GetOpenAPIDefinitions,
		},
		//nolint:govet
		Resources: []openapi.TypeInfo{
			{configapi.SchemeGroupVersion, configapi.ResourceSecretMetadatas, configapi.ResourceKindSecretMetadata, true},
			{configapi.SchemeGroupVersion, configapi.ResourceSecretStores, configapi.ResourceKindSecretStore, false},

			{vsapi.SchemeGroupVersion, vsapi.ResourceSecrets, vsapi.ResourceKindSecret, true},
		},
	})
	if err != nil {
		klog.Fatal(err)
	}

	filename := gort.GOPath() + "/src/go.virtual-secrets.dev/apimachinery/openapi/swagger.json"
	err = os.MkdirAll(filepath.Dir(filename), 0o755)
	if err != nil {
		klog.Fatal(err)
	}
	err = os.WriteFile(filename, []byte(apispec), 0o644)
	if err != nil {
		klog.Fatal(err)
	}
}

func main() {
	generateSwaggerJson()
}
