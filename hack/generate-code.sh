#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail


MODULE="github.com/viveksinghggits/kluster-api"
OUTPUT_DIR="pkg/generated"
OUTPUT_PKG="${MODULE}/pkg/generated"
BOILERPLATE="hack/boilerplate.go.txt"


rm -rf pkg/generated



echo "Generating clientset..."
go run k8s.io/code-generator/cmd/client-gen \
    --clientset-name versioned \
    --input-base "" \
    --input "${MODULE}/pkg/apis/kluster/v1alpha1" \
    --output-pkg "${OUTPUT_PKG}/clientset" \
    --output-dir "${OUTPUT_DIR}/clientset" \
    --go-header-file ${BOILERPLATE}

echo "Generating listers..."
go run k8s.io/code-generator/cmd/lister-gen \
    --output-pkg "${OUTPUT_PKG}/listers" \
    --output-dir "${OUTPUT_DIR}/listers" \
    --go-header-file ${BOILERPLATE} \
    "${MODULE}/pkg/apis/kluster/v1alpha1"

echo "Generating informers..."
go run k8s.io/code-generator/cmd/informer-gen \
    --versioned-clientset-package "${OUTPUT_PKG}/clientset/versioned" \
    --listers-package "${OUTPUT_PKG}/listers" \
    --output-pkg "${OUTPUT_PKG}/informers" \
    --output-dir "${OUTPUT_DIR}/informers" \
    --go-header-file ${BOILERPLATE} \
    "${MODULE}/pkg/apis/kluster/v1alpha1"


echo "Generating deepcopy..."
go run k8s.io/code-generator/cmd/deepcopy-gen \
    --output-file zz_generated.deepcopy.go \
    --go-header-file ${BOILERPLATE} \
    "${MODULE}/pkg/apis/kluster/v1alpha1" \
    "${MODULE}/pkg/apis/kluster"

echo "Generating conversion..."
go run k8s.io/code-generator/cmd/conversion-gen \
    --output-file zz_generated.conversion.go \
    --go-header-file ${BOILERPLATE} \
    "${MODULE}/pkg/apis/kluster/v1alpha1" \
    "${MODULE}/pkg/apis/kluster"


echo "Generating OpenAPI..."
go run k8s.io/kube-openapi/cmd/openapi-gen \
    --output-pkg "${OUTPUT_PKG}/openapi" \
    --output-dir "${OUTPUT_DIR}/openapi" \
    --report-filename pkg/generated/openapi_violations.report \
    --go-header-file ${BOILERPLATE} \
    "${MODULE}/pkg/apis/kluster/v1alpha1" \
    "k8s.io/apimachinery/pkg/apis/meta/v1" \
    "k8s.io/apimachinery/pkg/runtime" \
    "k8s.io/apimachinery/pkg/version"

echo "Done!"
