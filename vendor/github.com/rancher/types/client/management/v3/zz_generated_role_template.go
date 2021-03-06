package client

import (
	"github.com/rancher/norman/types"
)

const (
	RoleTemplateType                 = "roleTemplate"
	RoleTemplateFieldAnnotations     = "annotations"
	RoleTemplateFieldBuiltin         = "builtin"
	RoleTemplateFieldContext         = "context"
	RoleTemplateFieldCreated         = "created"
	RoleTemplateFieldCreatorID       = "creatorId"
	RoleTemplateFieldDescription     = "description"
	RoleTemplateFieldExternal        = "external"
	RoleTemplateFieldHidden          = "hidden"
	RoleTemplateFieldLabels          = "labels"
	RoleTemplateFieldName            = "name"
	RoleTemplateFieldOwnerReferences = "ownerReferences"
	RoleTemplateFieldRemoved         = "removed"
	RoleTemplateFieldRoleTemplateIds = "roleTemplateIds"
	RoleTemplateFieldRules           = "rules"
	RoleTemplateFieldUuid            = "uuid"
)

type RoleTemplate struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Builtin         bool              `json:"builtin,omitempty" yaml:"builtin,omitempty"`
	Context         string            `json:"context,omitempty" yaml:"context,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Description     string            `json:"description,omitempty" yaml:"description,omitempty"`
	External        bool              `json:"external,omitempty" yaml:"external,omitempty"`
	Hidden          bool              `json:"hidden,omitempty" yaml:"hidden,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	RoleTemplateIds []string          `json:"roleTemplateIds,omitempty" yaml:"roleTemplateIds,omitempty"`
	Rules           []PolicyRule      `json:"rules,omitempty" yaml:"rules,omitempty"`
	Uuid            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type RoleTemplateCollection struct {
	types.Collection
	Data   []RoleTemplate `json:"data,omitempty"`
	client *RoleTemplateClient
}

type RoleTemplateClient struct {
	apiClient *Client
}

type RoleTemplateOperations interface {
	List(opts *types.ListOpts) (*RoleTemplateCollection, error)
	Create(opts *RoleTemplate) (*RoleTemplate, error)
	Update(existing *RoleTemplate, updates interface{}) (*RoleTemplate, error)
	ByID(id string) (*RoleTemplate, error)
	Delete(container *RoleTemplate) error
}

func newRoleTemplateClient(apiClient *Client) *RoleTemplateClient {
	return &RoleTemplateClient{
		apiClient: apiClient,
	}
}

func (c *RoleTemplateClient) Create(container *RoleTemplate) (*RoleTemplate, error) {
	resp := &RoleTemplate{}
	err := c.apiClient.Ops.DoCreate(RoleTemplateType, container, resp)
	return resp, err
}

func (c *RoleTemplateClient) Update(existing *RoleTemplate, updates interface{}) (*RoleTemplate, error) {
	resp := &RoleTemplate{}
	err := c.apiClient.Ops.DoUpdate(RoleTemplateType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RoleTemplateClient) List(opts *types.ListOpts) (*RoleTemplateCollection, error) {
	resp := &RoleTemplateCollection{}
	err := c.apiClient.Ops.DoList(RoleTemplateType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *RoleTemplateCollection) Next() (*RoleTemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &RoleTemplateCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *RoleTemplateClient) ByID(id string) (*RoleTemplate, error) {
	resp := &RoleTemplate{}
	err := c.apiClient.Ops.DoByID(RoleTemplateType, id, resp)
	return resp, err
}

func (c *RoleTemplateClient) Delete(container *RoleTemplate) error {
	return c.apiClient.Ops.DoResourceDelete(RoleTemplateType, &container.Resource)
}
