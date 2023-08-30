package govpsie

import (
	"context"
	"fmt"
	"net/http"
)

var scriptsBasePath = "/apps/v2"

type ScriptsService interface {
	GetScripts(ctx context.Context) ([]Script, error)
	GetScript(ctx context.Context, scriptId string) (Script, error)
	CreateScript(ctx context.Context, createScriptRequest *CreateScriptRequest) error
}

type ScriptsServiceHandler struct {
	client *Client
}

var _ ScriptsService = &ScriptsServiceHandler{}

type Script struct {
	Identifier string `json:"identifier"`
	ScriptName string `json:"script_name"`
	Script     string `json:"script"`
}

type CreateScriptRequest struct {
	Name          string   `json:"name"`
	ScriptContent string   `json:"scriptContent"`
	ScriptType    string   `json:"scriptType"`
	Tags          []string `json:"tags"`
}

type ListScriptRoot struct {
	Error bool     `json:"error"`
	Data  []Script `json:"data"`
	Total int      `json:"total"`
}

type ScriptRoot struct {
	Error bool   `json:"error"`
	Data  Script `json:"data"`
}

func (s *ScriptsServiceHandler) GetScripts(ctx context.Context) ([]Script, error) {
	path := fmt.Sprintf("%s/scripts", scriptsBasePath)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	scripts := new(ListScriptRoot)
	if err := s.client.Do(ctx, req, scripts); err != nil {
		return nil, err
	}

	return scripts.Data, nil
}

func (s *ScriptsServiceHandler) GetScript(ctx context.Context, scriptId string) (Script, error) {
	path := fmt.Sprintf("%s/script/%s", scriptsBasePath, scriptId)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return Script{}, err
	}
	script := new(ScriptRoot)
	if err := s.client.Do(ctx, req, script); err != nil {
		return Script{}, err
	}

	return script.Data, nil
}

func (s *ScriptsServiceHandler) CreateScript(ctx context.Context, createScriptRequest *CreateScriptRequest) error {
	path := fmt.Sprintf("%s/script/add", scriptsBasePath)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createScriptRequest)
	if err != nil {
		return err
	}

	if err := s.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
