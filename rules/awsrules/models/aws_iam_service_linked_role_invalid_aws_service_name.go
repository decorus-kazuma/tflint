// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsIAMServiceLinkedRoleInvalidAwsServiceNameRule checks the pattern is valid
type AwsIAMServiceLinkedRoleInvalidAwsServiceNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMServiceLinkedRoleInvalidAwsServiceNameRule returns new rule with default attributes
func NewAwsIAMServiceLinkedRoleInvalidAwsServiceNameRule() *AwsIAMServiceLinkedRoleInvalidAwsServiceNameRule {
	return &AwsIAMServiceLinkedRoleInvalidAwsServiceNameRule{
		resourceType:  "aws_iam_service_linked_role",
		attributeName: "aws_service_name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMServiceLinkedRoleInvalidAwsServiceNameRule) Name() string {
	return "aws_iam_service_linked_role_invalid_aws_service_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMServiceLinkedRoleInvalidAwsServiceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMServiceLinkedRoleInvalidAwsServiceNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMServiceLinkedRoleInvalidAwsServiceNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMServiceLinkedRoleInvalidAwsServiceNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"aws_service_name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"aws_service_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`aws_service_name does not match valid pattern ^[\w+=,.@-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
