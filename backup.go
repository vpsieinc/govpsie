package govpsie

import (
	"context"
	"fmt"
	"net/http"
)

var backupsPath = "/apps/v2"

type BackupsService interface {
	List(ctx context.Context, options *ListOptions) ([]Backup, error)
	DeleteBackup(ctx context.Context, backupIdentifier string) error
	CreateBackups(ctx context.Context, vmIdentifier, name, notes string) error
	ListByVpsie(ctx context.Context, options *ListOptions, vmIdentifier string) ([]Backup, error)
	CreateVPSieByBackup(ctx context.Context, backupIdentifier string) error
}

type backupsServiceHandler struct {
	client *Client
}

var _ BackupsService = &backupsServiceHandler{}

type ListBackupsRoot struct {
	Error bool     `json:"error"`
	Data  []Backup `json:"data"`
	Total int      `json:"total "`
}

type Backup struct {
	HostName     string `json:"hostname"`
	Name         string `json:"name"`
	Identifier   string `json:"identifier"`
	Note         string `json:"note"`
	BackupKey    string `json:"backupKey"`
	State        string `json:"state"`
	DcIdentifier string `json:"dcIdentifier"`
	VMIdentifier string `json:"vmIdentifier"`
	BoxID        int    `json:"boxId"`
	BackupSHA1   string `json:"backupsha1"`
	OSFullName   string `json:"osFullName"`
	VMCategory   string `json:"vmCategory"`
	CreatedBy    string `json:"created_by"`
	CreatedOn    string `json:"created_on"`
}

func (b *backupsServiceHandler) List(ctx context.Context, options *ListOptions) ([]Backup, error) {
	path := fmt.Sprintf("%s/backups", backupsPath)

	req, err := b.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	backups := new(ListBackupsRoot)
	if err = b.client.Do(ctx, req, backups); err != nil {
		return nil, err
	}

	return backups.Data, nil
}

func (b *backupsServiceHandler) DeleteBackup(ctx context.Context, backupIdentifier string) error {
	path := fmt.Sprintf("%s/backup", backupsPath)

	deleteReq := struct {
		BackupIdentifier string `json:"backupIdentifier"`
	}{
		BackupIdentifier: backupIdentifier,
	}

	req, err := b.client.NewRequest(ctx, http.MethodDelete, path, &deleteReq)
	if err != nil {
		return err
	}

	return b.client.Do(ctx, req, nil)
}

func (b *backupsServiceHandler) CreateBackups(ctx context.Context, vmIdentifier, name, notes string) error {
	path := fmt.Sprintf("%s/backup/add", backupsPath)

	createBackupReq := struct {
		VmIdentifier string `json:"vmIdentifier"`
		Name         string `json:"name"`
	}{
		VmIdentifier: vmIdentifier,
		Name:         name,
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, path, &createBackupReq)

	if err != nil {
		return err
	}

	return b.client.Do(ctx, req, nil)

}

func (b *backupsServiceHandler) ListByVpsie(ctx context.Context, options *ListOptions, vpsieId string) ([]Backup, error) {
	path := fmt.Sprintf("%s/vm/backups/%s", backupsPath, vpsieId)

	req, err := b.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	backups := new(ListBackupsRoot)
	if err = b.client.Do(ctx, req, backups); err != nil {
		return nil, err
	}

	return backups.Data, nil
}

func (b *backupsServiceHandler) CreateVPSieByBackup(ctx context.Context, backupIdentifier string) error {
	path := fmt.Sprintf("%s/backups/create", backupsPath)

	createVpsieReq := struct {
		BackupIdentifier string `json:"backupIdentifier"`
	}{
		BackupIdentifier: backupIdentifier,
	}

	req, err := b.client.NewRequest(ctx, http.MethodPost, path, &createVpsieReq)
	if err != nil {
		return err
	}

	return b.client.Do(ctx, req, nil)
}
