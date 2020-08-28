/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	monitor "tkestack.io/tke/api/monitor"
)

// ClusterOverviewsGetter has a method to return a ClusterOverviewInterface.
// A group's client should implement this interface.
type ClusterOverviewsGetter interface {
	ClusterOverviews() ClusterOverviewInterface
}

// ClusterOverviewInterface has methods to work with ClusterOverview resources.
type ClusterOverviewInterface interface {
	Create(ctx context.Context, clusterOverview *monitor.ClusterOverview, opts v1.CreateOptions) (*monitor.ClusterOverview, error)
	ClusterOverviewExpansion
}

// clusterOverviews implements ClusterOverviewInterface
type clusterOverviews struct {
	client rest.Interface
}

// newClusterOverviews returns a ClusterOverviews
func newClusterOverviews(c *MonitorClient) *clusterOverviews {
	return &clusterOverviews{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a clusterOverview and creates it.  Returns the server's representation of the clusterOverview, and an error, if there is any.
func (c *clusterOverviews) Create(ctx context.Context, clusterOverview *monitor.ClusterOverview, opts v1.CreateOptions) (result *monitor.ClusterOverview, err error) {
	result = &monitor.ClusterOverview{}
	err = c.client.Post().
		Resource("clusteroverviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterOverview).
		Do(ctx).
		Into(result)
	return
}