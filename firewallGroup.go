package govpsie

import (
	"context"
	"fmt"
	"net/http"
)

var firewallGroupBasePath = "/apps/v2/firewall"

type FirewallGroupService interface {
	Create(ctx context.Context, groupName string) error
	List(ctx context.Context, options *ListOptions) ([]FirewallGroup, error)
	Get(ctx context.Context, fwGroupId string) (*FirewallGroup, error)
	Delete(ctx context.Context, fwGroupId string) error
	Update(ctx context.Context, fwGroupReq *FirewallUpdateReq, fwGroupId string) error
	AssignToVpsie(ctx context.Context, groupId string, vmId string) error
}

type firewallGroupServiceHandler struct {
	client *Client
}

var _ FirewallGroupService = &firewallGroupServiceHandler{}

type ListFirewallGroupsRoot struct {
	Error bool            `json:"error"`
	Data  []FirewallGroup `json:"data"`
	Total int             `json:"total"`
}

type GetFirewallGroupRoot struct {
	Error bool `json:"error"`
	Data  struct {
		Group *FirewallGroup `json:"group"`
	} `json:"data"`
}

type FirewallGroup struct {
	UserName      string `json:"user_name"`
	ID            int    `json:"id"`
	GroupName     string `json:"group_name"`
	Identifier    string `json:"identifier"`
	CreatedOn     string `json:"created_on"`
	UpdatedOn     string `json:"updated_on"`
	InboundCount  int    `json:"inbound_count"`
	OutboundCount int    `json:"outbound_count"`
	Vms           int    `json:"vms"`
}

type FirewallUpdateReq struct {
	Action     string     `json:"action"`
	Type       string     `json:"type"`
	Comment    string     `json:"comment"`
	Dest       []IpsetObj `json:"dest"`
	Dport      string     `json:"dport"`
	Source     []IpsetObj `json:"source"`
	Proto      string     `json:"proto"`
	Sport      string     `json:"sport"`
	Enable     int        `json:"enable"`
	Iface      string     `json:"iface"`
	Log        string     `json:"log"`
	Macro      string     `json:"macro"`
	Identifier string     `json:"identifier"`
}

type IpsetObj struct {
	Ipset string `json:"ipset"`
}

var _ FirewallGroupService = &firewallGroupServiceHandler{}

func (f *firewallGroupServiceHandler) Create(ctx context.Context, groupName string) error {
	fwGroupReq := struct {
		GroupName string `json:"groupName"`
	}{
		GroupName: groupName,
	}

	path := fmt.Sprintf("%s/create/group", firewallGroupBasePath)

	req, err := f.client.NewRequest(ctx, http.MethodPost, path, &fwGroupReq)
	if err != nil {
		return err
	}

	return f.client.Do(ctx, req, nil)
}

func (f *firewallGroupServiceHandler) List(ctx context.Context, options *ListOptions) ([]FirewallGroup, error) {
	path := fmt.Sprintf("%s/groups", firewallGroupBasePath)

	req, err := f.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	fwGroups := new(ListFirewallGroupsRoot)

	if err = f.client.Do(ctx, req, &fwGroups); err != nil {
		return nil, err
	}

	return fwGroups.Data, nil

}

func (f *firewallGroupServiceHandler) Get(ctx context.Context, fwGroupId string) (*FirewallGroup, error) {
	path := fmt.Sprintf("%s/group/%s", firewallGroupBasePath, fwGroupId)

	req, err := f.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	fwGroup := new(GetFirewallGroupRoot)

	if err = f.client.Do(ctx, req, fwGroup); err != nil {
		return nil, err
	}

	return fwGroup.Data.Group, nil
}

func (f *firewallGroupServiceHandler) Delete(ctx context.Context, fwGroupId string) error {
	path := fmt.Sprintf("%s/delete/group", firewallGroupBasePath)

	delReq := struct {
		GroupId string `json:"groupId"`
	}{
		GroupId: fwGroupId,
	}

	req, err := f.client.NewRequest(ctx, http.MethodDelete, path, &delReq)
	if err != nil {
		return err
	}

	return f.client.Do(ctx, req, nil)
}

func (f *firewallGroupServiceHandler) AssignToVpsie(ctx context.Context, groupId string, vmId string) error {
	path := fmt.Sprintf("%s/attach/group", firewallGroupBasePath)

	assignReq := struct {
		VmID    string `json:"vmId"`
		GroupID string `json:"groupId"`
	}{
		VmID:    vmId,
		GroupID: groupId,
	}

	req, err := f.client.NewRequest(ctx, http.MethodPost, path, assignReq)
	if err != nil {
		return err
	}

	return f.client.Do(ctx, req, nil)
}

func (f *firewallGroupServiceHandler) Update(ctx context.Context, fwGroupReq *FirewallUpdateReq, fwGroupId string) error {
	path := fmt.Sprintf("%s/groups/%s", firewallGroupBasePath, fwGroupId)

	req, err := f.client.NewRequest(ctx, http.MethodPost, path, fwGroupReq)
	if err != nil {
		return err
	}

	return f.client.Do(ctx, req, nil)
}
