package schema

type Task struct {
	Name        string `hcl:"name,label"`
	Description string `hcl:"description,optional"`
	Enabled     bool   `hcl:"enabled,optional"`

	On struct {
		PostPlan               bool   `hcl:"post_plan,optional"`
		PostApply              bool   `hcl:"post_apply,optional"`
		PostResourceAssessment bool   `hcl:"post_resource_assessment,optional"`
		Trigger                string `hcl:"trigger,optional"`
	} `hcl:"on,optional"`

	Runner struct {
		Type       string   `hcl:"type,optional"`
		Source     string   `hcl:"source,optional"`
		Name       string   `hcl:"name,optional"`
		Entrypoint []string `hcl:"entrypoint,optional"`
	} `hcl:"runner,optional"`

	Env map[string]string `hcl:"env,optional"`

	Inputs struct {
		ReportingCurrency string `hcl:"reporting_currency,optional"`
		OptimizationLevel string `hcl:"optimization_level,optional"`
	} `hcl:"inputs,optional"`

	Script      string `hcl:"script,optional"`
	Enforcement string `hcl:"enforcement_level,optional"`

	Notification struct {
		Slack struct {
			WebhookURL string `hcl:"webhook_url,optional"`
			Channel    string `hcl:"channel,optional"`
			OnFailure  bool   `hcl:"on_failure,optional"`
			OnSuccess  bool   `hcl:"on_success,optional"`
		} `hcl:"slack,optional"`
	} `hcl:"notification,optional"`
}
