package server

import (
	"github.com/spf13/cobra"
)

func NewCommandStartKlusterServer(
	defaults *KlusterServerOptions,
	stopCh <-chan struct{},
) *cobra.Command {
	o := *defaults
	cmd := &cobra.Command{
		Short: "Run Kluster agg api server",
		RunE: func(cmd *cobra.Command, args []string) error {
			// validate options
			// completed them

			if err := o.Run(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.RecommendedOptions.AddFlags(flags)

	return cmd
}
