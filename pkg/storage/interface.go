//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package storage

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/storage/storage"
)

type Util interface {
	Key(ctx context.Context, pattern ...string) string
}

type Storage interface {
	Cluster() storage.Cluster
	Hook() storage.Hook
	Namespace() storage.Namespace
	Deployment() storage.Deployment
	Service() storage.Service
	Route() storage.Route
	System() storage.System
	Pod() storage.Pod
	Volume() storage.Volume
	Node() storage.Node
	Endpoint() storage.Endpoint
}
