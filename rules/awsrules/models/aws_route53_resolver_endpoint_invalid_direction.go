// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsRoute53ResolverEndpointInvalidDirectionRule checks the pattern is valid
type AwsRoute53ResolverEndpointInvalidDirectionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsRoute53ResolverEndpointInvalidDirectionRule returns new rule with default attributes
func NewAwsRoute53ResolverEndpointInvalidDirectionRule() *AwsRoute53ResolverEndpointInvalidDirectionRule {
	return &AwsRoute53ResolverEndpointInvalidDirectionRule{
		resourceType:  "aws_route53_resolver_endpoint",
		attributeName: "direction",
		enum: []string{
			"INBOUND",
			"OUTBOUND",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverEndpointInvalidDirectionRule) Name() string {
	return "aws_route53_resolver_endpoint_invalid_direction"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverEndpointInvalidDirectionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverEndpointInvalidDirectionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverEndpointInvalidDirectionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverEndpointInvalidDirectionRule) Check(runner *tflint.Runner) error {
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
					`direction is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
