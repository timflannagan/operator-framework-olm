package main

import (
	"os"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/openshift/operator-framework-olm/staging/operator-registry/cmd/opm/root"
	registrylib "github.com/openshift/operator-framework-olm/staging/operator-registry/pkg/registry"
)

func main() {
	cmd := root.NewCmd()
	if err := cmd.Execute(); err != nil {
		agg, ok := err.(utilerrors.Aggregate)
		if !ok {
			os.Exit(1)
		}
		for _, e := range agg.Errors() {
			if _, ok := e.(registrylib.BundleImageAlreadyAddedErr); ok {
				os.Exit(2)
			}
			if _, ok := e.(registrylib.PackageVersionAlreadyAddedErr); ok {
				os.Exit(3)
			}
		}
		os.Exit(1)
	}
}
