package main

import (
	"os"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func main() {
	// Set up a logger
	logger := zap.New(zap.UseDevMode(true))

	// Create a new manager
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{
		Logger: logger,
	})
	if err != nil {
		logger.Error(err, "unable to set up overall controller manager")
		os.Exit(1)
	}

	// Start the manager
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		logger.Error(err, "problem running manager")
		os.Exit(1)
	}
}
