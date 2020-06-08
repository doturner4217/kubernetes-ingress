// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(ActionRedirect)
		**out = **in
	}
	if in.Return != nil {
		in, out := &in.Return, &out.Return
		*out = new(ActionReturn)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(ActionProxy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionProxy) DeepCopyInto(out *ActionProxy) {
	*out = *in
	if in.RequestHeaders != nil {
		in, out := &in.RequestHeaders, &out.RequestHeaders
		*out = new(ProxyRequestHeaders)
		(*in).DeepCopyInto(*out)
	}
	if in.ResponseHeaders != nil {
		in, out := &in.ResponseHeaders, &out.ResponseHeaders
		*out = new(ProxyResponseHeaders)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionProxy.
func (in *ActionProxy) DeepCopy() *ActionProxy {
	if in == nil {
		return nil
	}
	out := new(ActionProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionRedirect) DeepCopyInto(out *ActionRedirect) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionRedirect.
func (in *ActionRedirect) DeepCopy() *ActionRedirect {
	if in == nil {
		return nil
	}
	out := new(ActionRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionReturn) DeepCopyInto(out *ActionReturn) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionReturn.
func (in *ActionReturn) DeepCopy() *ActionReturn {
	if in == nil {
		return nil
	}
	out := new(ActionReturn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddHeader) DeepCopyInto(out *AddHeader) {
	*out = *in
	out.Header = in.Header
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddHeader.
func (in *AddHeader) DeepCopy() *AddHeader {
	if in == nil {
		return nil
	}
	out := new(AddHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorPage) DeepCopyInto(out *ErrorPage) {
	*out = *in
	if in.Codes != nil {
		in, out := &in.Codes, &out.Codes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Return != nil {
		in, out := &in.Return, &out.Return
		*out = new(ErrorPageReturn)
		(*in).DeepCopyInto(*out)
	}
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(ErrorPageRedirect)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorPage.
func (in *ErrorPage) DeepCopy() *ErrorPage {
	if in == nil {
		return nil
	}
	out := new(ErrorPage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorPageRedirect) DeepCopyInto(out *ErrorPageRedirect) {
	*out = *in
	out.ActionRedirect = in.ActionRedirect
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorPageRedirect.
func (in *ErrorPageRedirect) DeepCopy() *ErrorPageRedirect {
	if in == nil {
		return nil
	}
	out := new(ErrorPageRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorPageReturn) DeepCopyInto(out *ErrorPageReturn) {
	*out = *in
	out.ActionReturn = in.ActionReturn
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]Header, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorPageReturn.
func (in *ErrorPageReturn) DeepCopy() *ErrorPageReturn {
	if in == nil {
		return nil
	}
	out := new(ErrorPageReturn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalEndpoint) DeepCopyInto(out *ExternalEndpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalEndpoint.
func (in *ExternalEndpoint) DeepCopy() *ExternalEndpoint {
	if in == nil {
		return nil
	}
	out := new(ExternalEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Header) DeepCopyInto(out *Header) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Header.
func (in *Header) DeepCopy() *Header {
	if in == nil {
		return nil
	}
	out := new(Header)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(UpstreamTLS)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]Header, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Match) DeepCopyInto(out *Match) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(Action)
		(*in).DeepCopyInto(*out)
	}
	if in.Splits != nil {
		in, out := &in.Splits, &out.Splits
		*out = make([]Split, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Match.
func (in *Match) DeepCopy() *Match {
	if in == nil {
		return nil
	}
	out := new(Match)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyRequestHeaders) DeepCopyInto(out *ProxyRequestHeaders) {
	*out = *in
	if in.Pass != nil {
		in, out := &in.Pass, &out.Pass
		*out = new(bool)
		**out = **in
	}
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make([]Header, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyRequestHeaders.
func (in *ProxyRequestHeaders) DeepCopy() *ProxyRequestHeaders {
	if in == nil {
		return nil
	}
	out := new(ProxyRequestHeaders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyResponseHeaders) DeepCopyInto(out *ProxyResponseHeaders) {
	*out = *in
	if in.Hide != nil {
		in, out := &in.Hide, &out.Hide
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pass != nil {
		in, out := &in.Pass, &out.Pass
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ignore != nil {
		in, out := &in.Ignore, &out.Ignore
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make([]AddHeader, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyResponseHeaders.
func (in *ProxyResponseHeaders) DeepCopy() *ProxyResponseHeaders {
	if in == nil {
		return nil
	}
	out := new(ProxyResponseHeaders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(Action)
		(*in).DeepCopyInto(*out)
	}
	if in.Splits != nil {
		in, out := &in.Splits, &out.Splits
		*out = make([]Split, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Matches != nil {
		in, out := &in.Matches, &out.Matches
		*out = make([]Match, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ErrorPages != nil {
		in, out := &in.ErrorPages, &out.ErrorPages
		*out = make([]ErrorPage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionCookie) DeepCopyInto(out *SessionCookie) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionCookie.
func (in *SessionCookie) DeepCopy() *SessionCookie {
	if in == nil {
		return nil
	}
	out := new(SessionCookie)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Split) DeepCopyInto(out *Split) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(Action)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Split.
func (in *Split) DeepCopy() *Split {
	if in == nil {
		return nil
	}
	out := new(Split)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLS) DeepCopyInto(out *TLS) {
	*out = *in
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(TLSRedirect)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLS.
func (in *TLS) DeepCopy() *TLS {
	if in == nil {
		return nil
	}
	out := new(TLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSRedirect) DeepCopyInto(out *TLSRedirect) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSRedirect.
func (in *TLSRedirect) DeepCopy() *TLSRedirect {
	if in == nil {
		return nil
	}
	out := new(TLSRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Upstream) DeepCopyInto(out *Upstream) {
	*out = *in
	if in.Subselector != nil {
		in, out := &in.Subselector, &out.Subselector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MaxFails != nil {
		in, out := &in.MaxFails, &out.MaxFails
		*out = new(int)
		**out = **in
	}
	if in.MaxConns != nil {
		in, out := &in.MaxConns, &out.MaxConns
		*out = new(int)
		**out = **in
	}
	if in.Keepalive != nil {
		in, out := &in.Keepalive, &out.Keepalive
		*out = new(int)
		**out = **in
	}
	if in.ProxyBuffering != nil {
		in, out := &in.ProxyBuffering, &out.ProxyBuffering
		*out = new(bool)
		**out = **in
	}
	if in.ProxyBuffers != nil {
		in, out := &in.ProxyBuffers, &out.ProxyBuffers
		*out = new(UpstreamBuffers)
		**out = **in
	}
	out.TLS = in.TLS
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Queue != nil {
		in, out := &in.Queue, &out.Queue
		*out = new(UpstreamQueue)
		**out = **in
	}
	if in.SessionCookie != nil {
		in, out := &in.SessionCookie, &out.SessionCookie
		*out = new(SessionCookie)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Upstream.
func (in *Upstream) DeepCopy() *Upstream {
	if in == nil {
		return nil
	}
	out := new(Upstream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamBuffers) DeepCopyInto(out *UpstreamBuffers) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamBuffers.
func (in *UpstreamBuffers) DeepCopy() *UpstreamBuffers {
	if in == nil {
		return nil
	}
	out := new(UpstreamBuffers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamQueue) DeepCopyInto(out *UpstreamQueue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamQueue.
func (in *UpstreamQueue) DeepCopy() *UpstreamQueue {
	if in == nil {
		return nil
	}
	out := new(UpstreamQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamTLS) DeepCopyInto(out *UpstreamTLS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamTLS.
func (in *UpstreamTLS) DeepCopy() *UpstreamTLS {
	if in == nil {
		return nil
	}
	out := new(UpstreamTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServer) DeepCopyInto(out *VirtualServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServer.
func (in *VirtualServer) DeepCopy() *VirtualServer {
	if in == nil {
		return nil
	}
	out := new(VirtualServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerList) DeepCopyInto(out *VirtualServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerList.
func (in *VirtualServerList) DeepCopy() *VirtualServerList {
	if in == nil {
		return nil
	}
	out := new(VirtualServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerRoute) DeepCopyInto(out *VirtualServerRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerRoute.
func (in *VirtualServerRoute) DeepCopy() *VirtualServerRoute {
	if in == nil {
		return nil
	}
	out := new(VirtualServerRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServerRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerRouteList) DeepCopyInto(out *VirtualServerRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualServerRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerRouteList.
func (in *VirtualServerRouteList) DeepCopy() *VirtualServerRouteList {
	if in == nil {
		return nil
	}
	out := new(VirtualServerRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServerRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerRouteSpec) DeepCopyInto(out *VirtualServerRouteSpec) {
	*out = *in
	if in.Upstreams != nil {
		in, out := &in.Upstreams, &out.Upstreams
		*out = make([]Upstream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subroutes != nil {
		in, out := &in.Subroutes, &out.Subroutes
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerRouteSpec.
func (in *VirtualServerRouteSpec) DeepCopy() *VirtualServerRouteSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualServerRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerRouteStatus) DeepCopyInto(out *VirtualServerRouteStatus) {
	*out = *in
	if in.ExternalEndpoints != nil {
		in, out := &in.ExternalEndpoints, &out.ExternalEndpoints
		*out = make([]ExternalEndpoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerRouteStatus.
func (in *VirtualServerRouteStatus) DeepCopy() *VirtualServerRouteStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualServerRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerSpec) DeepCopyInto(out *VirtualServerSpec) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.Upstreams != nil {
		in, out := &in.Upstreams, &out.Upstreams
		*out = make([]Upstream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerSpec.
func (in *VirtualServerSpec) DeepCopy() *VirtualServerSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerStatus) DeepCopyInto(out *VirtualServerStatus) {
	*out = *in
	if in.ExternalEndpoints != nil {
		in, out := &in.ExternalEndpoints, &out.ExternalEndpoints
		*out = make([]ExternalEndpoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerStatus.
func (in *VirtualServerStatus) DeepCopy() *VirtualServerStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualServerStatus)
	in.DeepCopyInto(out)
	return out
}
