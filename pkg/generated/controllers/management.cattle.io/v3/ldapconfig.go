/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type LdapConfigHandler func(string, *v3.LdapConfig) (*v3.LdapConfig, error)

type LdapConfigController interface {
	generic.ControllerMeta
	LdapConfigClient

	OnChange(ctx context.Context, name string, sync LdapConfigHandler)
	OnRemove(ctx context.Context, name string, sync LdapConfigHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() LdapConfigCache
}

type LdapConfigClient interface {
	Create(*v3.LdapConfig) (*v3.LdapConfig, error)
	Update(*v3.LdapConfig) (*v3.LdapConfig, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.LdapConfig, error)
	List(opts metav1.ListOptions) (*v3.LdapConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.LdapConfig, err error)
}

type LdapConfigCache interface {
	Get(name string) (*v3.LdapConfig, error)
	List(selector labels.Selector) ([]*v3.LdapConfig, error)

	AddIndexer(indexName string, indexer LdapConfigIndexer)
	GetByIndex(indexName, key string) ([]*v3.LdapConfig, error)
}

type LdapConfigIndexer func(obj *v3.LdapConfig) ([]string, error)

type ldapConfigController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewLdapConfigController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) LdapConfigController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &ldapConfigController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromLdapConfigHandlerToHandler(sync LdapConfigHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.LdapConfig
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.LdapConfig))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *ldapConfigController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.LdapConfig))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateLdapConfigDeepCopyOnChange(client LdapConfigClient, obj *v3.LdapConfig, handler func(obj *v3.LdapConfig) (*v3.LdapConfig, error)) (*v3.LdapConfig, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *ldapConfigController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *ldapConfigController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *ldapConfigController) OnChange(ctx context.Context, name string, sync LdapConfigHandler) {
	c.AddGenericHandler(ctx, name, FromLdapConfigHandlerToHandler(sync))
}

func (c *ldapConfigController) OnRemove(ctx context.Context, name string, sync LdapConfigHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromLdapConfigHandlerToHandler(sync)))
}

func (c *ldapConfigController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *ldapConfigController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *ldapConfigController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *ldapConfigController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *ldapConfigController) Cache() LdapConfigCache {
	return &ldapConfigCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *ldapConfigController) Create(obj *v3.LdapConfig) (*v3.LdapConfig, error) {
	result := &v3.LdapConfig{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *ldapConfigController) Update(obj *v3.LdapConfig) (*v3.LdapConfig, error) {
	result := &v3.LdapConfig{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *ldapConfigController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *ldapConfigController) Get(name string, options metav1.GetOptions) (*v3.LdapConfig, error) {
	result := &v3.LdapConfig{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *ldapConfigController) List(opts metav1.ListOptions) (*v3.LdapConfigList, error) {
	result := &v3.LdapConfigList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *ldapConfigController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *ldapConfigController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.LdapConfig, error) {
	result := &v3.LdapConfig{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type ldapConfigCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *ldapConfigCache) Get(name string) (*v3.LdapConfig, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.LdapConfig), nil
}

func (c *ldapConfigCache) List(selector labels.Selector) (ret []*v3.LdapConfig, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.LdapConfig))
	})

	return ret, err
}

func (c *ldapConfigCache) AddIndexer(indexName string, indexer LdapConfigIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.LdapConfig))
		},
	}))
}

func (c *ldapConfigCache) GetByIndex(indexName, key string) (result []*v3.LdapConfig, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.LdapConfig, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.LdapConfig))
	}
	return result, nil
}