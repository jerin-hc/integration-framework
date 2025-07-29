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
	Type     string
	Identity string
}

type Response struct {
	Resources []ResourceChange
	Comment   Comment
}

type Request struct {
	Event     Event
	Resources []Resource
}

type Ctx context.Context

type Comment struct {
	Pass    bool
	Message string
}

type ResourceChange struct {
	Identity   string
	Mutate     map[string]interface{}
	Annotation map[string]interface{}
	Comment    Comment
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
