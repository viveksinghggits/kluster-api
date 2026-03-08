package server

import (
	"fmt"
	"io"
	"net"

	"github.com/viveksinghggits/kluster-api/pkg/apis/kluster/v1alpha1"
	genopenapi "github.com/viveksinghggits/kluster-api/pkg/generated/openapi"
	"k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

const defaultEtcdPathPrefix = "/registry/viveksingh.dev"

type KlusterServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions
}

func NewKlusterServerOptions(out, errOut io.Writer) *KlusterServerOptions {
	return &KlusterServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			Codec.LegacyCodec(v1alpha1.SchemeGroupVersion),
		),
	}
}

func (k *KlusterServerOptions) Config() (*KlusterConfig, error) {
	err := k.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")})
	if err != nil {
		return nil, fmt.Errorf("errro creating self signed certificate %v\n", err)
	}

	serverConfig := genericapiserver.NewRecommendedConfig(Codec)
	// this is where the options are converted into config using ApplyTo
	err = k.RecommendedOptions.ApplyTo(serverConfig)
	if err != nil {
		return nil, err
	}

	serverConfig.OpenAPIV3Config = genericapiserver.DefaultOpenAPIV3Config(genopenapi.GetOpenAPIDefinitions, &openapi.DefinitionNamer{})

	return &KlusterConfig{
		GenericConfig: serverConfig,
	}, nil
}

func (k *KlusterServerOptions) Run(stopCh <-chan struct{}) error {
	config, err := k.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}
