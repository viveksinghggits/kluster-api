package store

import (
	"github.com/viveksinghggits/kluster-api/pkg/apis/kluster"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

type KlusterRegistry struct {
	genericregistry.Store
}

func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*KlusterRegistry, error) {
	strategy := NewKlusterStrategy(scheme)

	store := genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &kluster.Kluster{} },
		NewListFunc:              func() runtime.Object { return &kluster.KlusterList{} },
		DefaultQualifiedResource: kluster.Resource("klusters"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		SingularQualifiedResource: kluster.Resource("kluster"),

		TableConvertor: rest.NewDefaultTableConvertor(kluster.Resource("klusters")),
	}

	options := generic.StoreOptions{
		RESTOptions: optsGetter,
		AttrFunc: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			return nil, nil, nil
		},
	}

	if err := store.CompleteWithOptions(&options); err != nil {
		return nil, err
	}

	return &KlusterRegistry{
		Store: store,
	}, nil
}
