package biz

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type ImageSummary struct {
	Containers int64 `json:"containers"`

	Created int64 `json:"created"`

	ID string `json:"id"`

	Labels map[string]string `json:"labels"`

	ParentID string `json:"parent_id"`

	RepoDigests []string `json:"repo_digests"`

	RepoTags []string `json:"repo_tags"`

	SharedSize int64 `json:"shared_size"`

	Size int64 `json:"size"`

	VirtualSize int64 `json:"virtual_size"`
}

type ImageUsecase struct {
	cli *client.Client
	log *zap.Logger
}

func NewImageUsecase(cli *client.Client, logger *zap.Logger) *ImageUsecase {
	return &ImageUsecase{cli: cli, log: logger}
}

// SearchByReference search image by name and version
func (u *ImageUsecase) SearchByReference(ctx context.Context, name, version string) ([]ImageSummary, error) {
	var imgs []types.ImageSummary
	var err error
	if name == "" {
		imgs, err = u.cli.ImageList(ctx, types.ImageListOptions{All: true})

	} else {
		imgs, err = u.cli.ImageList(ctx, types.ImageListOptions{All: true,
			Filters: filters.NewArgs(filters.Arg("reference", "*"+name+"*:"+version+"*"))})
	}

	if err != nil {
		return nil, err
	}
	var imageSummaries []ImageSummary
	err = copier.Copy(&imageSummaries, imgs)
	if err != nil {
		return nil, err
	}
	return imageSummaries, nil
}
func (u *ImageUsecase) DanglingImages() {

}
