package server

import (
	"context"
	"fmt"
	"strings"

	api "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

const (
	legacyDefaultDomain = "index.docker.io"
	defaultDomain       = "docker.io"
	officialRepoName    = "library"
	defaultTag          = "latest"

	sealerHub = "sea.hub:5000"
)

func (s *server) ListImages(ctx context.Context,
	req *api.ListImagesRequest) (*api.ListImagesResponse, error) {
	rsp, err := (*s.imageService).ListImages(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (s *server) ImageStatus(ctx context.Context,
	req *api.ImageStatusRequest) (*api.ImageStatusResponse, error) {
	rsp, err := (*s.imageService).ImageStatus(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (s *server) PullImage(ctx context.Context,
	req *api.PullImageRequest) (*api.PullImageResponse, error) {
	// TODO we can change the image name of req, and make the cri pull the image we need.
	// for example:
	// req.Image.Image = "sealer.hub/library/nginx:1.1.1"
	// and the cri will pull "sealer.hub/library/nginx:1.1.1", and save it as "sealer.hub/library/nginx:1.1.1"
	// note:
	// but kubelet sometimes will invoke imageService.RemoveImage() or something else. The req.Image.Image will the original name.
	// so we'd better tag "sealer.hub/library/nginx:1.1.1" with original name "req.Image.Image" After "rsp, err := (*s.imageService).PullImage(ctx, req)".
	domain, _ := splitDockerDomain(req.Image.Image)
	if domain == "" {
		req.Image.Image = sealerHub + "/" + req.Image.Image
	}
	fmt.Println(req.Image.Image)
	rsp, err := (*s.imageService).PullImage(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (s *server) RemoveImage(ctx context.Context,
	req *api.RemoveImageRequest) (*api.RemoveImageResponse, error) {
	rsp, err := (*s.imageService).RemoveImage(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (s *server) ImageFsInfo(ctx context.Context,
	req *api.ImageFsInfoRequest) (*api.ImageFsInfoResponse, error) {
	rsp, err := (*s.imageService).ImageFsInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, err
}

func splitDockerDomain(name string) (domain, remainder string) {
	i := strings.IndexRune(name, '/')
	if i == -1 || (!strings.ContainsAny(name[:i], ".:") && name[:i] != "localhost" && strings.ToLower(name[:i]) == name[:i]) {
		domain, remainder = "", name
	} else {
		domain, remainder = name[:i], name[i+1:]
	}
	if domain == legacyDefaultDomain {
		domain = defaultDomain
	}
	if domain == defaultDomain && !strings.ContainsRune(remainder, '/') {
		remainder = officialRepoName + "/" + remainder
	}
	return
}
