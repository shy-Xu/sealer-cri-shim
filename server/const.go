package server

const (
	// ContainerdSock is the CRI socket containerd listens on.
	ContainerdSock = "/var/run/containerd/containerd.sock"
	// SealerShimSock is the CRI socket the shim listens on.
	SealerShimSock = "/var/run/cri-resmgr/cri-resmgr.sock"
	// DirPermissions is the permissions to create the directory for sockets with.
	DirPermissions = 0711
)
