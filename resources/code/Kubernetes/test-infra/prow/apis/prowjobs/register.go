/*
Copyright 2018 The Kubernetes Authors.

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

package prowjobs

const (
	// GroupName defines the API group identifier for prow types.
	GroupName = "prow.k8s.io"

	// Kind defines the API kind identifier.
	Kind = "ProwJob"
)

var (
	// KindPlural is the plural form of Kind
	KindPlural = Kind + "s"
)
