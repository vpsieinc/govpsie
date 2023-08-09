package govpsie

import (
	"context"
	"fmt"
	"net/http"
)

var imagesPath = "/apps/v2/custom"

type ImagesService interface {
	DeleteImage(ctx context.Context, imageIdentifier string) error
	List(ctx context.Context, options *ListOptions) ([]CustomImage, error)
	CreateImages(ctx context.Context, dcIdentifier, imageName, imageUrl string) error
	CreateVPSieByImage(ctx context.Context, createVpsieReq *CreateVpsieRequest) error
}

type imagesServiceHandler struct {
	client *Client
}

var _ ImagesService = &imagesServiceHandler{}

type ListCustomImageRoot struct {
	Error bool          `json:"error"`
	Data  []CustomImage `json:"data"`
	Total int           `json:"total "`
}

type CustomImage struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	DataCenterID   int    `json:"datacenter_id"`
	ImageSize      int    `json:"image_size"`
	OriginalName   string `json:"original_name"`
	FetchedFromUrl string `json:"fetched_from_url"`
	ImageHash      string `json:"image_hash"`
	ImageLabel     string `json:"image_label"`
	CreatedOn      string `json:"created_on"`
	Deleted        int    `json:"deleted"`
	Identifier     string `json:"identifier"`
	DcName         string `json:"dc_name"`
	DcIdentifier   string `json:"dcIdentifier"`
	CreatedBy      string `json:"createdBy"`
}

func (i *imagesServiceHandler) List(ctx context.Context, options *ListOptions) ([]CustomImage, error) {
	path := fmt.Sprintf("%s/images", imagesPath)

	req, err := i.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	images := new(ListCustomImageRoot)
	if err = i.client.Do(ctx, req, images); err != nil {
		return nil, err
	}

	return images.Data, nil
}

func (i *imagesServiceHandler) CreateImages(ctx context.Context, dcIdentifier, imageName, imageUrl string) error {
	path := fmt.Sprintf("%s/images", imagesPath)

	createReq := struct {
		DcIdentifier string `json:"dcIdentifier"`
		ImageName    string `json:"imageName"`
		ImageUrl     string `json:"imageUrl"`
	}{
		DcIdentifier: dcIdentifier,
		ImageName:    imageName,
		ImageUrl:     imageUrl,
	}
	req, err := i.client.NewRequest(ctx, http.MethodPost, path, createReq)
	if err != nil {
		return err
	}

	if err = i.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (i *imagesServiceHandler) DeleteImage(ctx context.Context, imageIdentifier string) error {
	path := fmt.Sprintf("%s/images/%s", imagesPath, imageIdentifier)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	if err = i.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (i *imagesServiceHandler) CreateVPSieByImage(ctx context.Context, createVpsieReq *CreateVpsieRequest) error {
	path := fmt.Sprintf("%s/vm", imagesPath)
	req, err := i.client.NewRequest(ctx, http.MethodPost, path, createVpsieReq)
	if err != nil {
		return err
	}

	if err = i.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
