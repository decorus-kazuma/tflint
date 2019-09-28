// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsAppsyncFunctionInvalidRequestMappingTemplateRule checks the pattern is valid
type AwsAppsyncFunctionInvalidRequestMappingTemplateRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppsyncFunctionInvalidRequestMappingTemplateRule returns new rule with default attributes
func NewAwsAppsyncFunctionInvalidRequestMappingTemplateRule() *AwsAppsyncFunctionInvalidRequestMappingTemplateRule {
	return &AwsAppsyncFunctionInvalidRequestMappingTemplateRule{
		resourceType:  "aws_appsync_function",
		attributeName: "request_mapping_template",
		max:           65536,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppsyncFunctionInvalidRequestMappingTemplateRule) Name() string {
	return "aws_appsync_function_invalid_request_mapping_template"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncFunctionInvalidRequestMappingTemplateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncFunctionInvalidRequestMappingTemplateRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncFunctionInvalidRequestMappingTemplateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncFunctionInvalidRequestMappingTemplateRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"request_mapping_template must be 65536 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"request_mapping_template must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
