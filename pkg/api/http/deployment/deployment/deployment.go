//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2019] Last.Backend LLC
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

package deployment

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/api/envs"
	"github.com/lastbackend/lastbackend/pkg/api/types/v1/request"
	"github.com/lastbackend/lastbackend/pkg/distribution"
	"github.com/lastbackend/lastbackend/pkg/distribution/errors"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
	"github.com/lastbackend/lastbackend/pkg/log"
	"net/http"
)

const (
	logPrefix = "api:handler:deployment"
	logLevel  = 3
)

func Fetch(ctx context.Context, namespace, service, name string) (*types.Deployment, *errors.Err) {

	nm := distribution.NewDeploymentModel(ctx, envs.Get().GetStorage())
	dep, err := nm.Get(namespace, service, name)

	if err != nil {
		log.V(logLevel).Errorf("%s:fetch:> err: %s", logPrefix, err.Error())
		return nil, errors.New("deployment").InternalServerError(err)
	}

	if dep == nil {
		err := errors.New("deployment not found")
		log.V(logLevel).Errorf("%s:fetch:> err: %s", logPrefix, err.Error())
		return nil, errors.New("deployment").NotFound()
	}

	return dep, nil
}

func Apply(ctx context.Context, ns *types.Namespace, svc *types.Service, mf *request.DeploymentManifest, opts *request.DeploymentUpdateOptions) (*types.Deployment, *errors.Err) {

	if mf.Meta.Name == nil {
		return nil, errors.New("service").BadParameter("meta.name")
	}

	dep, err := Fetch(ctx, ns.Meta.Name, svc.Meta.Name, *mf.Meta.Name)
	if err != nil {
		if err.Code != http.StatusText(http.StatusNotFound) {
			return nil, errors.New("service").InternalServerError()
		}
	}

	if dep == nil {
		return Create(ctx, ns, svc, mf)
	}

	return Update(ctx, ns, svc, dep, mf, opts)
}

func Create(ctx context.Context, ns *types.Namespace, svc *types.Service, mf *request.DeploymentManifest) (*types.Deployment, *errors.Err) {

	dm := distribution.NewDeploymentModel(ctx, envs.Get().GetStorage())

	if mf.Meta.Name != nil {

		dep, err := dm.Get(ns.Meta.Name, svc.Meta.Name, *mf.Meta.Name)
		if err != nil {
			log.V(logLevel).Errorf("%s:create:> get deployment by name `%s` in namespace `%s` err: %s", logPrefix, mf.Meta.Name, ns.Meta.Name, err.Error())
			return nil, errors.New("deployment").InternalServerError()

		}

		if dep != nil {
			log.V(logLevel).Warnf("%s:create:> deployment name `%s` in namespace `%s` not unique", logPrefix, mf.Meta.Name, ns.Meta.Name)
			return nil, errors.New("deployment").NotUnique("name")

		}
	}

	dep := new(types.Deployment)
	mf.SetDeploymentMeta(dep)

	dep.Meta.SelfLink = *types.NewDeploymentSelfLink(ns.Meta.Name, svc.Meta.Name, *mf.Meta.Name)
	dep.Meta.Namespace = ns.Meta.Name
	dep.Meta.Endpoint = svc.Meta.Endpoint

	if err := mf.SetDeploymentSpec(dep); err != nil {
		return nil, errors.New("deployment").BadRequest(err.Error())
	}

	err := dm.Insert(dep)
	if err != nil {
		log.V(logLevel).Errorf("%s:create:> create deployment err: %s", logPrefix, err.Error())
		return nil, errors.New("deployment").InternalServerError()
	}

	return dep, nil
}

func Update(ctx context.Context, ns *types.Namespace, svc *types.Service, dep *types.Deployment, mf *request.DeploymentManifest, opts *request.DeploymentUpdateOptions) (*types.Deployment, *errors.Err) {

	dm := distribution.NewDeploymentModel(ctx, envs.Get().GetStorage())
	mf.SetDeploymentMeta(dep)

	dep.Meta.Endpoint = svc.Meta.Endpoint

	if err := mf.SetDeploymentSpec(dep); err != nil {
		return nil, errors.New("service").BadRequest(err.Error())
	}

	if err := dm.Update(dep); err != nil {
		log.V(logLevel).Errorf("%s:update:> update service err: %s", logPrefix, err.Error())
		return nil, errors.New("service").InternalServerError()
	}

	return dep, nil
}
