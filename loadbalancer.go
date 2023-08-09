package govpsie

import (
	"context"
	"fmt"
	"net/http"
)

var lbPath = "/api/v1/lb"

type LBsService interface {
	ListLBs(ctx context.Context, options *ListOptions) ([]LB, error)
	ListLBDataCenters(ctx context.Context, options *ListOptions) ([]LBDataCenter, error)
	ListOffers(ctx context.Context, dcIdentifier string) ([]LBOffers, error)
	GetLB(ctx context.Context, lbID string) (*LB, error)
	CreateLB(ctx context.Context, createLBReq *CreateLBReq) error
	DeleteLB(ctx context.Context, lbID string) error
	DeleteLBRule(ctx context.Context, ruleID string) error
	DeleteLBDomain(ctx context.Context, domainID string) error
	DeleteLBBackend(ctx context.Context, lbBackendID string) error
}

type lbsServiceHandler struct {
	client *Client
}

var _ LBsService = &lbsServiceHandler{}

type ListLBsRoot struct {
	Error bool `json:"error"`
	Data  []LB `json:"data"`
	Total int  `json:"total"`
}

type GetLBRoot struct {
	Error bool `json:"error"`
	Data  LB   `json:"data"`
}

type ListLBDataCentersRoot struct {
	Error bool           `json:"error"`
	Data  []LBDataCenter `json:"data"`
	Total int            `json:"total"`
}

type LB struct {
	Cpu        int    `json:"cpu"`
	Ssd        int    `json:"ssd"`
	Ram        int    `json:"ram"`
	LBName     string `json:"lbName"`
	Traffic    int    `json:"traffic"`
	BoxsizeID  int    `json:"boxsize_id"`
	DefaultIP  string `json:"default_ip"`
	DCName     string `json:"dc_name"`
	Identifier string `json:"identifier"`
	CreatedOn  string `json:"created_on"`
	UpdatedAt  string `json:"updated_at"`
	Package    string `json:"package"`
	CreatedBy  string `json:"created_by"`
	UserID     int    `json:"user_id"`
}

type CreateLBReq struct {
	Rules              []Rule `json:"rules"`
	Algorithm          string `json:"algorithm"`
	CookieName         string `json:"cookieName"`
	HealthCheckPath    string `json:"healthCheckPath"`
	CookieCheck        bool   `json:"cookieCheck"`
	RedirectHTTP       bool   `json:"redirectHTTP"`
	LBName             string `json:"lbName"`
	ResourceIdentifier string `json:"resourceIdentifier"`
	DcIdentifier       string `json:"dcIdentifier"`
	CheckInterval      int    `json:"checkInterval"`
	FastInterval       int    `json:"fastInterval"`
	Rise               int    `json:"rise"`
	Fall               int    `json:"fall"`
}

type Rule struct {
	Scheme     string     `json:"scheme"`
	FrontPort  string     `json:"frontPort"`
	Domains    []LBDomain `json:"domains"`
	Backends   []Backend  `json:"backends"`
	BackPort   string     `json:"backPort"`
	DomainName string     `json:"domainName"`
}
type LBDomain struct {
	DomainID      string    `json:"domainId"`
	Backends      []Backend `json:"backends"`
	BackPort      string    `json:"backPort"`
	BackendScheme string    `json:"backendScheme"`
	DomainName    string    `json:"domainName"`
}
type Backend struct {
	Ip           string `json:"ip"`
	VmIdentifier string `json:"vmIdentifier"`
}

type LBDataCenter struct {
	DcName     string `json:"dc_name"`
	DcImage    string `json:"dc_image"`
	State      string `json:"state"`
	Country    string `json:"country"`
	Identifier string `json:"identifier"`
	IsActive   int    `json:"is_active"`
	IsDeleted  int    `json:"is_deleted"`
}

type LBOffers struct {
	Cpu         int    `json:"cpu"`
	Ram         int    `json:"ram"`
	Ssd         int    `json:"ssd"`
	Traffic     int    `json:"traffic"`
	Price       string `json:"price"`
	NickName    string `json:"nickname"`
	Identifier  string `json:"identifier"`
	Color       string `json:"color"`
	NetSpeed    int    `json:"net_speed"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type ListOffersRoot struct {
	Error bool       `json:"error"`
	Data  []LBOffers `json:"Data"`
}

func (l *lbsServiceHandler) ListLBs(ctx context.Context, options *ListOptions) ([]LB, error) {
	path := fmt.Sprintf("%s/all?sortField=created_on&sortDirection=DESC", lbPath)

	req, err := l.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	listLbsRoot := new(ListLBsRoot)
	if err := l.client.Do(ctx, req, listLbsRoot); err != nil {
		return nil, err
	}

	return listLbsRoot.Data, nil
}

func (l *lbsServiceHandler) GetLB(ctx context.Context, lbID string) (*LB, error) {
	path := fmt.Sprintf("%s/%s", lbPath, lbID)

	req, err := l.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	lb := new(GetLBRoot)
	if err := l.client.Do(ctx, req, &lb); err != nil {
		return nil, err
	}

	return &lb.Data, nil
}

func (l *lbsServiceHandler) ListLBDataCenters(ctx context.Context, options *ListOptions) ([]LBDataCenter, error) {
	path := fmt.Sprintf("%s/datacenter", lbPath)

	req, err := l.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	lbDataCenters := new(ListLBDataCentersRoot)
	if err := l.client.Do(ctx, req, &lbDataCenters); err != nil {
		return nil, err
	}

	return lbDataCenters.Data, nil
}

func (l *lbsServiceHandler) CreateLB(ctx context.Context, createLBReq *CreateLBReq) error {
	path := fmt.Sprintf("%s/create", lbPath)

	req, err := l.client.NewRequest(ctx, http.MethodPost, path, createLBReq)
	if err != nil {
		return err
	}

	return l.client.Do(ctx, req, nil)
}

func (l *lbsServiceHandler) DeleteLB(ctx context.Context, lbID string) error {
	path := fmt.Sprintf("%s/%s", lbPath, lbID)

	req, err := l.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return l.client.Do(ctx, req, nil)
}

func (l *lbsServiceHandler) DeleteLBRule(ctx context.Context, ruleID string) error {
	path := fmt.Sprintf("%s/delete/rule", lbPath)
	delReq := struct {
		RuleID string `json:"ruleId"`
	}{
		RuleID: ruleID,
	}

	req, err := l.client.NewRequest(ctx, http.MethodDelete, path, delReq)
	if err != nil {
		return err
	}

	return l.client.Do(ctx, req, nil)
}

func (l *lbsServiceHandler) DeleteLBDomain(ctx context.Context, domainID string) error {
	path := fmt.Sprintf("%s/delete/domain", lbPath)

	delReq := struct {
		DomainID string `json:"domainId"`
	}{
		DomainID: domainID,
	}

	req, err := l.client.NewRequest(ctx, http.MethodDelete, path, delReq)
	if err != nil {
		return err
	}

	return l.client.Do(ctx, req, nil)
}

func (l *lbsServiceHandler) DeleteLBBackend(ctx context.Context, lbBackendID string) error {
	path := fmt.Sprintf("%s/delete/backend", lbPath)

	delReq := struct {
		BackendID string `json:"backendId"`
	}{
		BackendID: lbBackendID,
	}

	req, err := l.client.NewRequest(ctx, http.MethodDelete, path, delReq)
	if err != nil {
		return err
	}

	return l.client.Do(ctx, req, nil)
}

func (l *lbsServiceHandler) ListOffers(ctx context.Context, dcIdentifier string) ([]LBOffers, error) {
	path := fmt.Sprintf("%s/offers", lbPath)

	offerReq := struct {
		DcIdentifier string `json:"dcIdentifier"`
	}{
		DcIdentifier: dcIdentifier,
	}

	req, err := l.client.NewRequest(ctx, http.MethodPost, path, &offerReq)
	if err != nil {
		return nil, err
	}

	offersRoot := new(ListOffersRoot)
	if err := l.client.Do(ctx, req, &offersRoot); err != nil {
		return nil, err
	}

	return offersRoot.Data, nil
}

func (l *lbsServiceHandler) AddRuleToLB(ctx context.Context, lbID string, rule *Rule) error {
	return nil
}

func (l *lbsServiceHandler) AddDomainToRule(ctx context.Context, lbID string, domain *LBDomain) error {
	return nil
}
