package schema

import (
	"context"
)

type Event string

const (
	PrePlan                Event = "pre-plan"
	PostPlan               Event = "post-plan"
	PreApply               Event = "pre-apply"
	PostApply              Event = "post-apply"
	Test                   Event = "test"
	PostResourceAssessment Event = "post-resource-assessment"
)

type Resource struct {
	Type     string `json:"type"`
	Identity string `json:"identity"`
}

type Response struct {
	Resources []ResourceChange `json:"resources"`
	Comment   Comment          `json:"comment"`
}

type Request struct {
	Event     Event      `json:"event"`
	Resources []Resource `json:"resources"`
}

type Ctx context.Context

type Comment struct {
	Pass    bool   `json:"pass"`
	Message string `json:"message"`
}

type ResourceChange struct {
	Identity   string                 `json:"identity"`
	Mutate     map[string]interface{} `json:"mutate"`
	Annotation map[string]interface{} `json:"annotation"`
	Comment    Comment                `json:"comment"`
}

type RunTaskHandler struct {
	HandlePrePlan   func(resources []Resource, ctx Ctx) (*Response, error)
	HandlePostPlan  func(resources []Resource, ctx Ctx) (*Response, error)
	HandlePreApply  func(resources []Resource, ctx Ctx) (*Response, error)
	HandlePostApply func(resources []Resource, ctx Ctx) (*Response, error)
	HandleTest      func(resources []Resource, ctx Ctx) (*Response, error)
	Trigger         func(resources []Resource, ctx Ctx) (*Response, error)
}

type IntegrationServer interface {
	HandlePrePlan(resources []Resource, ctx Ctx) (*Response, error)
	HandlePostPlan(resources []Resource, ctx Ctx) (*Response, error)
	HandlePreApply(resources []Resource, ctx Ctx) (*Response, error)
	HandlePostApply(resources []Resource, ctx Ctx) (*Response, error)
	HandleTest(resources []Resource, ctx Ctx) (*Response, error)
	Trigger(resources []Resource, ctx Ctx) (*Response, error)
}
