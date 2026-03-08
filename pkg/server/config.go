package server

import (
	"github.com/viveksinghggits/kluster-api/pkg/apis/kluster/v1alpha1"
	"github.com/viveksinghggits/kluster-api/pkg/store"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/util/compatibility"
)

type KlusterConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
}

type completedConfig struct {
	// add new fields if required
	GenericConfig genericapiserver.CompletedConfig
}

type KlusterServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

func (k *KlusterConfig) Complete() completedConfig {
	k.GenericConfig.EffectiveVersion = compatibility.DefaultBuildEffectiveVersion()
	return completedConfig{
		GenericConfig: k.GenericConfig.Complete(),
	}
}

func (c completedConfig) New() (*KlusterServer, error) {
	genericServer, err := c.GenericConfig.New("kluster-server", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(v1alpha1.SchemeGroupVersion.Group, Scheme, metav1.ParameterCodec, Codec)
	v1alpha1Storage := map[string]rest.Storage{}
	store, err := store.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, err
	}
	v1alpha1Storage["klusters"] = store

	apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1Storage

	if err := genericServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	return &KlusterServer{
		GenericAPIServer: genericServer,
	}, nil
}
