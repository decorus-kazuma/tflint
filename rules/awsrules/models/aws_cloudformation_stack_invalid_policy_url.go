// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudformationStackInvalidPolicyURLRule checks the pattern is valid
type AwsCloudformationStackInvalidPolicyURLRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudformationStackInvalidPolicyURLRule returns new rule with default attributes
func NewAwsCloudformationStackInvalidPolicyURLRule() *AwsCloudformationStackInvalidPolicyURLRule {
	return &AwsCloudformationStackInvalidPolicyURLRule{
		resourceType:  "aws_cloudformation_stack",
		attributeName: "policy_url",
		max:           1350,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackInvalidPolicyURLRule) Name() string {
	return "aws_cloudformation_stack_invalid_policy_url"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackInvalidPolicyURLRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackInvalidPolicyURLRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackInvalidPolicyURLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackInvalidPolicyURLRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"policy_url must be 1350 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"policy_url must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
