package main

import (
	"cri-proxy/shim"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"k8s.io/klog/v2"
)

// test 1

//test 2
func main() {
	options := shim.Options{
		ShimSocket:  "/var/run/cri-resmgr/cri-resmgr.sock",
		ImageSocket: "/var/run/dockershim.sock",
	}

	_shim, err := shim.NewShim(options)
	if err != nil {
		klog.Fatalf("failed to new _shim, %s", err)
	}

	err = _shim.Setup()
	if err != nil {
		klog.Fatalf("failed to setup sealer _shim, %s", err)
	}

	err = _shim.Start()
	if err != nil {
		klog.Fatalf(fmt.Sprintf("failed to start sealer _shim, %s", err))
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	stopCh := make(chan struct{}, 1)
	select {
	case <-signalCh:
		close(stopCh)
	case <-stopCh:

	}
	klog.Infof("Shutting donw the sealer _shim")
}
