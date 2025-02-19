// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package talent

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newProfileClientHook clientHook

// ProfileCallOptions contains the retry settings for each method of ProfileClient.
type ProfileCallOptions struct {
	ListProfiles   []gax.CallOption
	CreateProfile  []gax.CallOption
	GetProfile     []gax.CallOption
	UpdateProfile  []gax.CallOption
	DeleteProfile  []gax.CallOption
	SearchProfiles []gax.CallOption
}

func defaultProfileGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("jobs.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("jobs.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://jobs.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultProfileCallOptions() *ProfileCallOptions {
	return &ProfileCallOptions{
		ListProfiles: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateProfile: []gax.CallOption{},
		GetProfile: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateProfile: []gax.CallOption{},
		DeleteProfile: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchProfiles: []gax.CallOption{},
	}
}

// internalProfileClient is an interface that defines the methods availaible from Cloud Talent Solution API.
type internalProfileClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListProfiles(context.Context, *talentpb.ListProfilesRequest, ...gax.CallOption) *ProfileIterator
	CreateProfile(context.Context, *talentpb.CreateProfileRequest, ...gax.CallOption) (*talentpb.Profile, error)
	GetProfile(context.Context, *talentpb.GetProfileRequest, ...gax.CallOption) (*talentpb.Profile, error)
	UpdateProfile(context.Context, *talentpb.UpdateProfileRequest, ...gax.CallOption) (*talentpb.Profile, error)
	DeleteProfile(context.Context, *talentpb.DeleteProfileRequest, ...gax.CallOption) error
	SearchProfiles(context.Context, *talentpb.SearchProfilesRequest, ...gax.CallOption) (*talentpb.SearchProfilesResponse, error)
}

// ProfileClient is a client for interacting with Cloud Talent Solution API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service that handles profile management, including profile CRUD,
// enumeration and search.
type ProfileClient struct {
	// The internal transport-dependent client.
	internalClient internalProfileClient

	// The call options for this service.
	CallOptions *ProfileCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ProfileClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ProfileClient) setGoogleClientInfo(...string) {
	c.internalClient.setGoogleClientInfo()
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ProfileClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListProfiles lists profiles by filter. The order is unspecified.
func (c *ProfileClient) ListProfiles(ctx context.Context, req *talentpb.ListProfilesRequest, opts ...gax.CallOption) *ProfileIterator {
	return c.internalClient.ListProfiles(ctx, req, opts...)
}

// CreateProfile creates and returns a new profile.
func (c *ProfileClient) CreateProfile(ctx context.Context, req *talentpb.CreateProfileRequest, opts ...gax.CallOption) (*talentpb.Profile, error) {
	return c.internalClient.CreateProfile(ctx, req, opts...)
}

// GetProfile gets the specified profile.
func (c *ProfileClient) GetProfile(ctx context.Context, req *talentpb.GetProfileRequest, opts ...gax.CallOption) (*talentpb.Profile, error) {
	return c.internalClient.GetProfile(ctx, req, opts...)
}

// UpdateProfile updates the specified profile and returns the updated result.
func (c *ProfileClient) UpdateProfile(ctx context.Context, req *talentpb.UpdateProfileRequest, opts ...gax.CallOption) (*talentpb.Profile, error) {
	return c.internalClient.UpdateProfile(ctx, req, opts...)
}

// DeleteProfile deletes the specified profile.
// Prerequisite: The profile has no associated applications or assignments
// associated.
func (c *ProfileClient) DeleteProfile(ctx context.Context, req *talentpb.DeleteProfileRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteProfile(ctx, req, opts...)
}

// SearchProfiles searches for profiles within a tenant.
//
// For example, search by raw queries “software engineer in Mountain View” or
// search by structured filters (location filter, education filter, etc.).
//
// See SearchProfilesRequest for more information.
func (c *ProfileClient) SearchProfiles(ctx context.Context, req *talentpb.SearchProfilesRequest, opts ...gax.CallOption) (*talentpb.SearchProfilesResponse, error) {
	return c.internalClient.SearchProfiles(ctx, req, opts...)
}

// profileGRPCClient is a client for interacting with Cloud Talent Solution API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type profileGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ProfileClient
	CallOptions **ProfileCallOptions

	// The gRPC API client.
	profileClient talentpb.ProfileServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewProfileClient creates a new profile service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service that handles profile management, including profile CRUD,
// enumeration and search.
func NewProfileClient(ctx context.Context, opts ...option.ClientOption) (*ProfileClient, error) {
	clientOpts := defaultProfileGRPCClientOptions()
	if newProfileClientHook != nil {
		hookOpts, err := newProfileClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ProfileClient{CallOptions: defaultProfileCallOptions()}

	c := &profileGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		profileClient:    talentpb.NewProfileServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *profileGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *profileGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *profileGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *profileGRPCClient) ListProfiles(ctx context.Context, req *talentpb.ListProfilesRequest, opts ...gax.CallOption) *ProfileIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListProfiles[0:len((*c.CallOptions).ListProfiles):len((*c.CallOptions).ListProfiles)], opts...)
	it := &ProfileIterator{}
	req = proto.Clone(req).(*talentpb.ListProfilesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.Profile, string, error) {
		var resp *talentpb.ListProfilesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.profileClient.ListProfiles(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetProfiles(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

func (c *profileGRPCClient) CreateProfile(ctx context.Context, req *talentpb.CreateProfileRequest, opts ...gax.CallOption) (*talentpb.Profile, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateProfile[0:len((*c.CallOptions).CreateProfile):len((*c.CallOptions).CreateProfile)], opts...)
	var resp *talentpb.Profile
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.profileClient.CreateProfile(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *profileGRPCClient) GetProfile(ctx context.Context, req *talentpb.GetProfileRequest, opts ...gax.CallOption) (*talentpb.Profile, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetProfile[0:len((*c.CallOptions).GetProfile):len((*c.CallOptions).GetProfile)], opts...)
	var resp *talentpb.Profile
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.profileClient.GetProfile(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *profileGRPCClient) UpdateProfile(ctx context.Context, req *talentpb.UpdateProfileRequest, opts ...gax.CallOption) (*talentpb.Profile, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "profile.name", url.QueryEscape(req.GetProfile().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateProfile[0:len((*c.CallOptions).UpdateProfile):len((*c.CallOptions).UpdateProfile)], opts...)
	var resp *talentpb.Profile
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.profileClient.UpdateProfile(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *profileGRPCClient) DeleteProfile(ctx context.Context, req *talentpb.DeleteProfileRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteProfile[0:len((*c.CallOptions).DeleteProfile):len((*c.CallOptions).DeleteProfile)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.profileClient.DeleteProfile(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *profileGRPCClient) SearchProfiles(ctx context.Context, req *talentpb.SearchProfilesRequest, opts ...gax.CallOption) (*talentpb.SearchProfilesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SearchProfiles[0:len((*c.CallOptions).SearchProfiles):len((*c.CallOptions).SearchProfiles)], opts...)
	var resp *talentpb.SearchProfilesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.profileClient.SearchProfiles(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ProfileIterator manages a stream of *talentpb.Profile.
type ProfileIterator struct {
	items    []*talentpb.Profile
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.Profile, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ProfileIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ProfileIterator) Next() (*talentpb.Profile, error) {
	var item *talentpb.Profile
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ProfileIterator) bufLen() int {
	return len(it.items)
}

func (it *ProfileIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
