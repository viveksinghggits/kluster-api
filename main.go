package main

import (
	"flag"
	"os"

	"github.com/viveksinghggits/kluster-api/pkg/server"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	stopCh := genericapiserver.SetupSignalHandler()
	options := server.NewKlusterServerOptions(os.Stdout, os.Stderr)

	cmd := server.NewCommandStartKlusterServer(options, stopCh)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
