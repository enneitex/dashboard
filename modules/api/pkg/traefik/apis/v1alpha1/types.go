/*
Copyright 2017 The Kubernetes Authors.
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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:generate controller-gen object paths=$GOFILE

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressRoute struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// spec is the desired state of the IngressRoute.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec IngressRouteSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IngressRouteList is a collection of IngressRoute.
type IngressRouteList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of IngressRoute.
	Items []IngressRoute `json:"items" protobuf:"bytes,2,opt,name=items"`
}

// IngressRouteSpec describes the IngressRoute the user wishes to exist.
type IngressRouteSpec struct {
	// EntryPoints defines the list of entry point names to bind to. Entry points
	// have to be configured in the static configuration. More info:
	// https://doc.traefik.io/traefik/v2.9/routing/entrypoints/ Default: all.
	EntryPoints []string `json:"entryPoints,omitempty" protobuf:"bytes,1,rep,name=entryPoints"`

	// Routes defines the list of routes.
	Routes []IngressRouteRoutes `json:"defaultBackend,omitempty" protobuf:"bytes,1,rep,name=routes"`

	// TLS defines the TLS configuration. More info:
	// https://doc.traefik.io/traefik/v2.9/routing/routers/#tls
	TLS *IngressRouteTLS `json:"tls,omitempty" protobuf:"bytes,2,opt,name=tls"`

}

type IngressRouteRoutes struct {

	Kind string `json:"kind,omitempty" protobuf:"bytes,2,opt,name=kind"`

	Match string `json:"match,omitempty" protobuf:"bytes,2,opt,name=match"`

	Middlewares []IngressRouteRoutesMiddlewares `json:"middlewares,omitempty" protobuf:"bytes,2,rep,name=middlewares"`

	Priority int32 `json:"priority,omitempty" protobuf:"bytes,2,opt,name=priority"`

	Services []IngressRouteRoutesServices `json:"services,omitempty" protobuf:"bytes,2,rep,name=services"`


}

type IngressRouteRoutesMiddlewares struct {

	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`

	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`

}

type IngressRouteRoutesServices struct {

	Kind string `json:"kind,omitempty" protobuf:"bytes,2,opt,name=kind"`

	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`

	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`

	PassHostHeader bool `json:"passHostHeader,omitempty" protobuf:"bytes,2,opt,name=passHostHeader"`

	Port string `json:"port,omitempty" protobuf:"bytes,2,opt,name=port"`

	ResponseForwarding *IngressRouteRoutesServicesResponseForwarding `json:"responseForwarding,omitempty" protobuf:"bytes,2,opt,name=responseForwarding"`

	Scheme string `json:"scheme,omitempty" protobuf:"bytes,2,opt,name=scheme"`

	ServersTransport string `json:"serversTransport,omitempty" protobuf:"bytes,2,opt,name=serversTransport"`

	Sticky *IngressRouteRoutesServicesSticky `json:"sticky,omitempty" protobuf:"bytes,2,opt,name=sticky"`

	Strategy string `json:"strategy,omitempty" protobuf:"bytes,2,opt,name=strategy"`

	Weight int32 `json:"weight,omitempty" protobuf:"bytes,2,opt,name=weight"`


}

type IngressRouteRoutesServicesResponseForwarding struct {

	FlushInterval string `json:"flushInterval,omitempty" protobuf:"bytes,2,opt,name=flushInterval"`

}

type IngressRouteRoutesServicesSticky struct {

	Cookie *IngressRouteRoutesServicesStickyCookie `json:"cookie,omitempty" protobuf:"bytes,2,opt,name=cookie"`

}

type IngressRouteRoutesServicesStickyCookie struct {

	HttpOnly bool `json:"httpOnly,omitempty" protobuf:"bytes,2,opt,name=httpOnly"`

	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`

	SameSite string `json:"sameSite,omitempty" protobuf:"bytes,2,opt,name=sameSite"`

	Secure bool `json:"secure,omitempty" protobuf:"bytes,2,opt,name=secure"`

}

// TLS defines the TLS configuration. More info:
// https://doc.traefik.io/traefik/v2.9/routing/routers/#tls
type IngressRouteTLS struct {

	// CertResolver defines the name of the certificate resolver to use. Cert
    // resolvers have to be configured in the static configuration. More info:
    // https://doc.traefik.io/traefik/v2.9/https/acme/#certificate-resolvers
	CertResolver string `json:"certResolver,omitempty" protobuf:"bytes,2,opt,name=certResolver"`

	// Domains defines the list of domains that will be used to issue
    // certificates. More info:
    // https://doc.traefik.io/traefik/v2.9/routing/routers/#domains
	Domains []IngressRouteTLSDomains `json:"domains,omitempty" protobuf:"bytes,2,rep,name=domains"`

	// Options defines the reference to a TLSOption, that specifies the parameters
    // of the TLS connection. If not defined, the `default` TLSOption is used.
    // More info: https://doc.traefik.io/traefik/v2.9/https/tls/#tls-options
	Options *IngressRouteTLSOptions `json:"options,omitempty" protobuf:"bytes,2,opt,name=options"`

	// SecretName is the name of the referenced Kubernetes Secret to specify the
	// certificate details.
	SecretName string `json:"secretName,omitempty" protobuf:"bytes,2,opt,name=secretName"`

	// Store defines the reference to the TLSStore, that will be used to store
    // certificates. Please note that only `default` TLSStore can be used.
	Store *IngressRouteTLSStore `json:"store,omitempty" protobuf:"bytes,2,opt,name=store"`

}

type IngressRouteTLSStore struct {

	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`

	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`

}

type IngressRouteTLSDomains struct {

	Main string `json:"main,omitempty" protobuf:"bytes,2,opt,name=main"`

	Sans []string `json:"sans,omitempty" protobuf:"bytes,2,rep,name=sans"`

}

type IngressRouteTLSOptions struct {

	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`

	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`

}