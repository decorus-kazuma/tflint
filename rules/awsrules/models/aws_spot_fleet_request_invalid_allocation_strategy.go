// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsSpotFleetRequestInvalidAllocationStrategyRule checks the pattern is valid
type AwsSpotFleetRequestInvalidAllocationStrategyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSpotFleetRequestInvalidAllocationStrategyRule returns new rule with default attributes
func NewAwsSpotFleetRequestInvalidAllocationStrategyRule() *AwsSpotFleetRequestInvalidAllocationStrategyRule {
	return &AwsSpotFleetRequestInvalidAllocationStrategyRule{
		resourceType:  "aws_spot_fleet_request",
		attributeName: "allocation_strategy",
		enum: []string{
			"lowestPrice",
			"diversified",
			"capacityOptimized",
		},
	}
}

// Name returns the rule name
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Name() string {
	return "aws_spot_fleet_request_invalid_allocation_strategy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Check(runner *tflint.Runner) error {
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
					`allocation_strategy is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
