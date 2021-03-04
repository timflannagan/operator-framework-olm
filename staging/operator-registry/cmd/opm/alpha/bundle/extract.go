package bundle

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/openshift/operator-framework-olm/staging/operator-registry/pkg/configmap"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extracts the data in a bundle directory via ConfigMap",
	Long:  "Extract takes as input a directory containing manifests and writes the per file contents to a ConfipMap",

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if debug, _ := cmd.Flags().GetBool("debug"); debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	},

	RunE: runExtractCmd,
}

func init() {
	extractCmd.Flags().Bool("debug", false, "enable debug logging")
	extractCmd.Flags().StringP("kubeconfig", "k", "", "absolute path to kubeconfig file")
	extractCmd.Flags().StringP("manifestsdir", "m", "/", "path to directory containing manifests")
	extractCmd.Flags().StringP("configmapname", "c", "", "name of configmap to write bundle data")
	extractCmd.Flags().StringP("namespace", "n", "openshift-operator-lifecycle-manager", "namespace to write configmap data")
	extractCmd.Flags().Uint64P("datalimit", "l", 1<<20, "maximum limit in bytes for total bundle data")
	extractCmd.MarkPersistentFlagRequired("configmapname")
}

func runExtractCmd(cmd *cobra.Command, args []string) error {
	manifestsDir, err := cmd.Flags().GetString("manifestsdir")
	if err != nil {
		return err
	}
	kubeconfig, err := cmd.Flags().GetString("kubeconfig")
	if err != nil {
		return err
	}
	configmapName, err := cmd.Flags().GetString("configmapname")
	if err != nil {
		return err
	}
	namespace, err := cmd.Flags().GetString("namespace")
	if err != nil {
		return err
	}
	datalimit, err := cmd.Flags().GetUint64("datalimit")
	if err != nil {
		return err
	}

	loader := configmap.NewConfigMapLoaderForDirectory(configmapName, namespace, manifestsDir, kubeconfig)
	if err := loader.Populate(datalimit); err != nil {
		return fmt.Errorf("error loading manifests from directory: %s", err)
	}

	return nil
}
