package govpsie

import (
	"context"
	"fmt"
	"net/http"
)

var serverBasePath = "/apps/v2/vm"

type ServerService interface {
	ListServer(context.Context, *ListOptions, string) ([]VmData, error)
	List(context.Context, *ListOptions) ([]VmData, error)
	GetServerByIdentifier(context.Context, string) (*VmData, error)
	GetServerStatusByIdentifier(context.Context, string) (*Status, error)
	GetServerConsole(ctx context.Context, identifierId string) (*ServerConsole, error)
	CreateServer(context.Context, *CreateServerRequest) error
	DeleteServer(ctx context.Context, identifierId string) error
	StartServer(ctx context.Context, identifierId string) error
	StopServer(ctx context.Context, identifierId string) error
	RestartServer(ctx context.Context, identifierId string) error
	ChangePassword(ctx context.Context, identifierId string, newPassword string) error
	ChangeHostName(ctx context.Context, identifierId string, newHostname string) error
	AddVPC(ctx context.Context, request *VpcRequest) error
	MoveVPC(ctx context.Context, request *VpcRequest) error
	AddTags(ctx context.Context, identifierId string, tags []string) error
	ResizeServer(ctx context.Context, identifierId, cpu, ram string) error
	AddSsh(ctx context.Context, identifierId, sshKeyIdentifier string) error
	AddScript(ctx context.Context, identifierId, scriptIdentifier string) error
	Lock(ctx context.Context, identifierId string) error
	UnLock(ctx context.Context, identifierId string) error
	DoMultiActions(ctx context.Context, vmsIdentifiers []string, actionType, sshKeyIdentifier string) error
	EnableIpv6(ctx context.Context, identifierId string) error
	EnableIpv4(ctx context.Context, identifierId string) error
	AddFip(ctx context.Context, identifierId, dcIdentifier string) error
	Resume(ctx context.Context, resumeReq *ResumeReq) error
	ResetNetwork(ctx context.Context, vmIdentifier string) error
	EditTag(ctx context.Context, tags []string, vmIdentifer string) error
}

type serverServiceHandler struct {
	client *Client
}

var _ ServerService = &serverServiceHandler{}

type ListServerRoot struct {
	Error bool     `json:"error"`
	Data  []VmData `json:"data"`
	Total int      `json:"total"`
}

type ListServerByIdentifierRoot struct {
	Error bool `json:"error"`
	Data  struct {
		VmData          VmData            `json:"vmData"`
		ImageCategories []ImageCategories `json:"imageCategories"`
		VmTags          []VmTags          `json:"vmTags"`
		PrivateIpData   []PrivateIpData   `json:"privateIpData"`
		FloatingIpData  []FloatingIpData  `json:"floatingIpData"`
	} `json:"data"`
	Total int `json:"total"`
}

type ImageCategories struct {
}
type VmTags struct {
}
type PrivateIpData struct {
}
type FloatingIpData struct {
}

type VmData struct {
	ID                  int     `json:"id"`
	UserID              int     `json:"user_id"`
	BoxSizeID           int     `json:"boxsize_id"`
	BoxImageID          int     `json:"boximage_id"`
	DataCenterID        int     `json:"datacenter_id"`
	NodeID              int     `json:"node_id"`
	BoxdIsCountID       *int    `json:"boxdiscount_id"`
	Hostname            string  `json:"hostname"`
	DefaultIP           string  `json:"default_ip"`
	DefaultIPv6         string  `json:"default_ipv6"`
	PrivateIP           string  `json:"private_ip"`
	IsAutoBackup        int     `json:"is_autobackup"`
	BoxVirtualization   string  `json:"box_virtualization_id"`
	Ram                 int     `json:"ram"`
	Cpu                 int     `json:"cpu"`
	Ssd                 int     `json:"ssd"`
	Traffic             int     `json:"traffic"`
	AddedIpAddresses    *string `json:"added_ip_addresses"`
	InitialPassword     string  `json:"initial_password"`
	Notes               *string `json:"notes"`
	CreatedOn           string  `json:"created_on"`
	LastUpdated         string  `json:"last_updated"`
	DroppedOn           *string `json:"dropped_on"`
	IsActive            int     `json:"is_active"`
	IsDeleted           int     `json:"is_deleted"`
	Identifier          string  `json:"identifier"`
	Power               int     `json:"power"`
	ProjectID           int     `json:"project_id"`
	IsCustom            int     `json:"is_custom"`
	NrAddedIps          int     `json:"nr_added_ips"`
	InPcs               int     `json:"in_pcs"`
	CustomPrice         *int    `json:"custom_price"`
	PayableLicense      int     `json:"payable_license"`
	LastLicensePay      *string `json:"last_license_pay"`
	ScriptID            *string `json:"script_id"`
	SshKeyID            *string `json:"sshkey_id"`
	IsLocked            int     `json:"is_locked"`
	IsWorkWithNew       int     `json:"is_work_with_new_version"`
	IsSuspended         int     `json:"is_suspended"`
	IsTerminated        int     `json:"is_terminated"`
	OldID               int     `json:"old_id"`
	CustomIsoID         *int    `json:"custom_iso_id"`
	IsIsoImageBootAble  int     `json:"is_iso_image_bootable"`
	HasSsl              int     `json:"has_ssl"`
	LastActionDate      *string `json:"last_action_date,omitempty"`
	IsCreatedFromLegacy int     `json:"is_created_from_legacy"`
	IsSmtpAllowed       int     `json:"is_smtp_allowed"`
	WeeklyBackup        int     `json:"weekly_backup"`
	MonthlyBackup       int     `json:"monthly_backup"`
	LibIsoID            *int    `json:"lib_iso_id,omitempty"`
	DailySnapshot       int     `json:"daily_snapshot"`
	WeeklySnapshot      int     `json:"weekly_snapshot"`
	MonthlySnapshot     int     `json:"monthly_snap"`
	LastActionInMin     int     `json:"last_action_in_min"`
	FirstName           string  `json:"firstname"`
	LastName            string  `json:"lastname"`
	Username            string  `json:"username"`
	State               string  `json:"state"`
	IsFipAvailable      int     `json:"is_fip_available"`
	IsBucketAvailable   int     `json:"is_bucket_available"`
	DcIdentifier        string  `json:"dcIdentifier"`
	Category            string  `json:"category"`
	FullName            string  `json:"fullname"`
	VmDescription       string  `json:"vmDescription"`
	BoxesSuspended      int     `json:"boxes_suspended"`
	IsSataAvailable     int     `json:"is_sata_available"`
	IsSsdAvailable      int     `json:"is_ssd_available"`
	PublicIp            *string `json:"publicIp,omitempty"`
}

type Status struct {
	Cpu            int    `json:"cpu"`
	Ballon         int    `json:"ballon"`
	Uptime         int    `json:"uptime"`
	Pid            string `json:"pid"`
	Disk           int    `json:"disk"`
	RunningMachine string `json:"running-machine"`
	RunningQemu    string `json:"running-qemu"`
	Status         string `json:"status"`
	DiskRead       string `json:"diskread"`
	DiskWrite      string `json:"diskwrite"`
	Fullname       string `json:"fullname"`
}

type GetStatusRoot struct {
	Error  bool   `json:"error"`
	Status Status `json:"data"`
}

type ServerConsoleRoot struct {
	Error         bool          `json:"error"`
	ServerConsole ServerConsole `json:"data"`
}

type ServerConsole struct {
	Upid     string `json:"upid"`
	Ticket   string `json:"ticket"`
	Cert     string `json:"cert"`
	User     string `json:"user"`
	Port     string `json:"port"`
	Token    string `json:"token"`
	OS       string `json:"os"`
	Fullname string `json:"fullname"`
}

type CreateServerRequest struct {
	ResourceIdentifier string    `json:"resourceIdentifier"`
	OsIdentifier       string    `json:"osIdentifier"`
	DcIdentifier       string    `json:"dcIdentifier"`
	Hostname           string    `json:"hostname"`
	Notes              *string   `json:"notes,omitempty"`
	BackupEnabled      *int      `json:"backupEnabled,omitempty"`
	AddPublicIpV4      *int      `json:"addPublicIpV4,omitempty"`
	AddPublicIpV6      *int      `json:"addPublicIpV6,omitempty"`
	AddPrivateIp       *int      `json:"addPrivateIp,omitempty"`
	SshKeyIdentifier   *string   `json:"sshKeyIdentifier,omitempty"`
	ProjectID          int       `json:"projectId"`
	Tags               []*string `json:"tags,omitempty"`
	ScriptIdentifier   *string   `json:"scriptIdentifier,omitempty"`
}

type ActionRequest struct {
	VmIdentifier string `json:"vmIdentifier"`
}

type VpcRequest struct {
	VmIdentifier string `json:"vmIdentifier"`
	VpcId        string `json:"vpcId"`
	DcIdentifier string `json:"dcIdentifier"`
}

type ResumeReq struct {
	VmIdentifier     string `json:"vmIdentifier"`
	OsIdentifier     string `json:"osIdentifier"`
	ScriptIdentifier string `json:"scriptIdentifier"`
	SshKeyIdentifier string `json:"sshKeyIdentifier"`
	HostName         string `json:"hostname"`
	Password         string `json:"password"`
	IsOnPremise      bool   `json:"isOnPremise"`
}

func (v *serverServiceHandler) ListServer(ctx context.Context, options *ListOptions, projectId string) ([]VmData, error) {
	path := fmt.Sprintf("%s?projectId=%s", serverBasePath, projectId)
	req, err := v.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	servers := new(ListServerRoot)
	if err = v.client.Do(ctx, req, servers); err != nil {
		return nil, err
	}

	return servers.Data, nil
}

func (v *serverServiceHandler) List(ctx context.Context, options *ListOptions) ([]VmData, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, serverBasePath, nil)
	if err != nil {
		return nil, err
	}

	Servers := new(ListServerRoot)
	if err = v.client.Do(ctx, req, Servers); err != nil {
		return nil, err
	}

	return Servers.Data, nil
}

func (v *serverServiceHandler) GetServerByIdentifier(ctx context.Context, identifierId string) (*VmData, error) {
	path := fmt.Sprintf("%s/%s", serverBasePath, identifierId)
	req, err := v.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	Servers := new(ListServerByIdentifierRoot)
	if err = v.client.Do(ctx, req, Servers); err != nil {
		return nil, err
	}

	return &Servers.Data.VmData, nil
}

func (v *serverServiceHandler) GetServerStatusByIdentifier(ctx context.Context, identifierId string) (*Status, error) {
	path := fmt.Sprintf("%s/status/ %s", serverBasePath, identifierId)
	req, err := v.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	status := new(GetStatusRoot)
	if err = v.client.Do(ctx, req, status); err != nil {
		return nil, err
	}

	return &status.Status, nil
}

func (v *serverServiceHandler) GetServerConsole(ctx context.Context, identifierId string) (*ServerConsole, error) {
	path := fmt.Sprintf("%s/console/%s", serverBasePath, identifierId)
	req, err := v.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	console := new(ServerConsoleRoot)
	if err = v.client.Do(ctx, req, console); err != nil {
		return nil, err
	}

	return &console.ServerConsole, nil
}

func (v *serverServiceHandler) CreateServer(ctx context.Context, server *CreateServerRequest) error {
	req, err := v.client.NewRequest(ctx, http.MethodPost, serverBasePath, server)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) DeleteServer(ctx context.Context, identifierId string) error {
	vmIdentifier := &ActionRequest{
		VmIdentifier: identifierId,
	}
	req, err := v.client.NewRequest(ctx, http.MethodDelete, serverBasePath, vmIdentifier)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) StartServer(ctx context.Context, identifierId string) error {
	vmIdentifier := &ActionRequest{
		VmIdentifier: identifierId,
	}

	path := fmt.Sprintf("%s/start", serverBasePath)
	req, err := v.client.NewRequest(ctx, http.MethodPost, path, vmIdentifier)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) StopServer(ctx context.Context, identifierId string) error {
	vmIdentifier := &ActionRequest{
		VmIdentifier: identifierId,
	}

	path := fmt.Sprintf("%s/stop", serverBasePath)
	req, err := v.client.NewRequest(ctx, http.MethodPost, path, vmIdentifier)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) RestartServer(ctx context.Context, identifierId string) error {
	vmIdentifier := &ActionRequest{
		VmIdentifier: identifierId,
	}
	path := fmt.Sprintf("%s/restart", serverBasePath)
	req, err := v.client.NewRequest(ctx, http.MethodPost, path, vmIdentifier)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) ChangePassword(ctx context.Context, identifierId string, newPassword string) error {
	changePassReq := struct {
		VmIdentifier string `json:"vmIdentifier"`
		NewPassword  string `json:"newpassword"`
	}{
		VmIdentifier: identifierId,
		NewPassword:  newPassword,
	}
	path := fmt.Sprintf("%s/changepass", serverBasePath)
	req, err := v.client.NewRequest(ctx, http.MethodPost, path, changePassReq)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) ChangeHostName(ctx context.Context, identifierId string, newHostname string) error {
	changeHostNameReq := struct {
		VmIdentifier string `json:"vmIdentifier"`
		Hostname     string `json:"hostname"`
	}{
		VmIdentifier: identifierId,
		Hostname:     newHostname,
	}
	path := fmt.Sprintf("%s/changehostname", serverBasePath)
	req, err := v.client.NewRequest(ctx, http.MethodPost, path, changeHostNameReq)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) AddVPC(ctx context.Context, request *VpcRequest) error {
	path := fmt.Sprintf("%s/add/vpc", serverBasePath)

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, request)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) MoveVPC(ctx context.Context, request *VpcRequest) error {
	path := fmt.Sprintf("%s/vpc/move", serverBasePath)

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, request)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) ResizeServer(ctx context.Context, identifierId, cpu, ram string) error {
	path := fmt.Sprintf("%s/resize", serverBasePath)

	resizeServer := struct {
		VmIdentifier string `json:"vmIdentifier"`
		Ram          string `json:"ram"`
		Cpu          string `json:"cup"`
	}{
		VmIdentifier: identifierId,
		Ram:          ram,
		Cpu:          cpu,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, resizeServer)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) AddTags(ctx context.Context, identifierId string, tags []string) error {
	path := fmt.Sprintf("%s/addtags", serverBasePath)

	addTagsRequest := struct {
		VmIdentifier string   `json:"vmIdentifier"`
		Tags         []string `json:"tags"`
	}{
		VmIdentifier: identifierId,
		Tags:         tags,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, addTagsRequest)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) AddSsh(ctx context.Context, identifierId, sshKeyIdentifier string) error {
	path := fmt.Sprintf("%s/sshkey", serverBasePath)
	addSshReq := struct {
		VmIdentifier     string `json:"vmIdentifier"`
		SshKeyIdentifier string `json:"sshKeyIdentifier"`
	}{
		VmIdentifier:     identifierId,
		SshKeyIdentifier: sshKeyIdentifier,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, addSshReq)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) AddScript(ctx context.Context, identifierId, scriptIdentifier string) error {
	path := fmt.Sprintf("%s/script", serverBasePath)
	addScriptReq := struct {
		VmIdentifier     string `json:"vmIdentifier"`
		ScriptIdentifier string `json:"scriptIdentifier"`
	}{
		VmIdentifier:     identifierId,
		ScriptIdentifier: scriptIdentifier,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, addScriptReq)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) Lock(ctx context.Context, identifierId string) error {
	path := fmt.Sprintf("%s/lock", serverBasePath)
	lockReq := struct {
		VmIdentifier string `json:"vmIdentifier"`
	}{
		VmIdentifier: identifierId,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, lockReq)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) UnLock(ctx context.Context, identifierId string) error {
	path := fmt.Sprintf("%s/unlock", serverBasePath)
	unLockReq := struct {
		VmIdentifier string `json:"vmIdentifier"`
	}{
		VmIdentifier: identifierId,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, unLockReq)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) DoMultiActions(ctx context.Context, vmsIdentifiers []string, actionType, sshKeyIdentifier string) error {
	path := fmt.Sprintf("%s/actions", serverBasePath)
	doMultiActionsReq := struct {
		VmsIdentifiers   []string `json:"vmsIdentifiers"`
		ActionType       string   `json:"actionType"`
		SshKeyIdentifier string   `json:"sshKeyIdentifier"`
	}{
		VmsIdentifiers:   vmsIdentifiers,
		ActionType:       actionType,
		SshKeyIdentifier: sshKeyIdentifier,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, doMultiActionsReq)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) AddFip(ctx context.Context, identifierId, dcIdentifier string) error {
	path := fmt.Sprintf("%s/fip/add", serverBasePath)
	addFipReq := struct {
		VmIdentifier string `json:"vmIdentifier"`
		DcIdentifier string `json:"dcIdentifier"`
	}{
		VmIdentifier: identifierId,
		DcIdentifier: dcIdentifier,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, addFipReq)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) EnableIpv6(ctx context.Context, identifierId string) error {
	path := fmt.Sprintf("%s/enable/ipv6", serverBasePath)
	enableIpv6Req := struct {
		VmIdentifier string `json:"vmIdentifier"`
	}{
		VmIdentifier: identifierId,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, enableIpv6Req)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) EnableIpv4(ctx context.Context, identifierId string) error {
	path := fmt.Sprintf("%s/enable/ipv4", serverBasePath)
	enableIpv4Req := struct {
		VmIdentifier string `json:"vmIdentifier"`
	}{
		VmIdentifier: identifierId,
	}

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, enableIpv4Req)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) Resume(ctx context.Context, resumeReq *ResumeReq) error {
	path := fmt.Sprintf("%s/resume", serverBasePath)

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, resumeReq)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) ResetNetwork(ctx context.Context, vmIdentifier string) error {
	path := fmt.Sprintf("/apps/v2/refresh/vm/ips/%s", vmIdentifier)

	req, err := v.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (v *serverServiceHandler) EditTag(ctx context.Context, tags []string, vmIdentifer string) error {
	path := fmt.Sprintf("%s/tags/edit", serverBasePath)

	req, err := v.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}

	if err = v.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
