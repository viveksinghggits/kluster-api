package server

import (
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

	return &KlusterServer{
		GenericAPIServer: genericServer,
	}, nil
}
