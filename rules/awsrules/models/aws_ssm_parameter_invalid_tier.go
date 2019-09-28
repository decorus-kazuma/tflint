// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmParameterInvalidTierRule checks the pattern is valid
type AwsSsmParameterInvalidTierRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmParameterInvalidTierRule returns new rule with default attributes
func NewAwsSsmParameterInvalidTierRule() *AwsSsmParameterInvalidTierRule {
	return &AwsSsmParameterInvalidTierRule{
		resourceType:  "aws_ssm_parameter",
		attributeName: "tier",
		enum: []string{
			"Standard",
			"Advanced",
			"Intelligent-Tiering",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmParameterInvalidTierRule) Name() string {
	return "aws_ssm_parameter_invalid_tier"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmParameterInvalidTierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmParameterInvalidTierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmParameterInvalidTierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmParameterInvalidTierRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`tier is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
