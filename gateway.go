package govpsie

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var gatewayPath = "/apps/v2/gateways"

type GatewayService interface {
	List(ctx context.Context, options *ListOptions) ([]Gateway, error)
	Delete(ctx context.Context, ipId int) error
	Create(ctx context.Context, createReq *CreateGatewayReq) error
}

type gatewayServiceHandler struct {
	client *Client
}

var _ GatewayService = &gatewayServiceHandler{}

type ListGatewayRoot struct {
	Error bool `json:"error"`
	Data  struct {
		Rows                    []Gateway `json:"rows"`
		IsThereUserGatewayLevel bool      `json:"isThereUserGatewayLevel"`
		Count                   int       `json:"count"`
	} `json:"data"`
}

type Gateway struct {
	ID                   int         `json:"id"`
	DatacenterID         int         `json:"datacenter_id"`
	IPPropertiesID       int         `json:"ip_properties_id"`
	IP                   string      `json:"ip"`
	IsReserved           int         `json:"is_reserved"`
	IPVersion            string      `json:"ip_version"`
	BoxID                interface{} `json:"box_id"`
	IsPrimary            int         `json:"is_primary"`
	Notes                interface{} `json:"notes"`
	UserID               int         `json:"user_id"`
	UpdatedAt            time.Time   `json:"updated_at"`
	IsGatewayReserved    int         `json:"is_gateway_reserved"`
	IsUserAccountGateway int         `json:"is_user_account_gateway"`
	DatacenterName       string      `json:"datacenterName"`
	State                string      `json:"state"`
	DcIdentifier         string      `json:"dcIdentifier"`
	CreatedBy            string      `json:"created_by"`
	AttachedVms          []string    `json:"attachedVms"`
}

type CreateGatewayReq struct {
	IPType       string        `json:"ipType"`
	Tags         []string `json:"tags"`
	DcIdentifier string        `json:"dcIdentifier"`
}

func (s *gatewayServiceHandler) List(ctx context.Context, options *ListOptions) ([]Gateway, error) {
	path := fmt.Sprintf("%s/ips", gatewayPath)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var root ListGatewayRoot
	if err := s.client.Do(ctx, req, &root); err != nil {
		return nil, err
	}

	return root.Data.Rows, nil
}

func (s *gatewayServiceHandler) Create(ctx context.Context, createReq *CreateGatewayReq) error {
	path := fmt.Sprintf("%s/add/ip", gatewayPath)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createReq)
	if err != nil {
		return err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *gatewayServiceHandler) Delete(ctx context.Context, ipId int) error {
	path := fmt.Sprintf("%s/delete/ip", gatewayPath)

	delReq := struct {
		IPID int `json:"ipId"`
	}{
		IPID: ipId,
	}

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, delReq)
	if err != nil {
		return err
	}

	return s.client.Do(ctx, req, nil)
}
