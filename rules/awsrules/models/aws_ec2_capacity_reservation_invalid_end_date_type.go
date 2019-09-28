// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsEc2CapacityReservationInvalidEndDateTypeRule checks the pattern is valid
type AwsEc2CapacityReservationInvalidEndDateTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2CapacityReservationInvalidEndDateTypeRule returns new rule with default attributes
func NewAwsEc2CapacityReservationInvalidEndDateTypeRule() *AwsEc2CapacityReservationInvalidEndDateTypeRule {
	return &AwsEc2CapacityReservationInvalidEndDateTypeRule{
		resourceType:  "aws_ec2_capacity_reservation",
		attributeName: "end_date_type",
		enum: []string{
			"unlimited",
			"limited",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2CapacityReservationInvalidEndDateTypeRule) Name() string {
	return "aws_ec2_capacity_reservation_invalid_end_date_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2CapacityReservationInvalidEndDateTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2CapacityReservationInvalidEndDateTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2CapacityReservationInvalidEndDateTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2CapacityReservationInvalidEndDateTypeRule) Check(runner *tflint.Runner) error {
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
					`end_date_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
