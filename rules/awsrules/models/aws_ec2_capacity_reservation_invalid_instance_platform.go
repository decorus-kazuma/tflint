// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsEc2CapacityReservationInvalidInstancePlatformRule checks the pattern is valid
type AwsEc2CapacityReservationInvalidInstancePlatformRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2CapacityReservationInvalidInstancePlatformRule returns new rule with default attributes
func NewAwsEc2CapacityReservationInvalidInstancePlatformRule() *AwsEc2CapacityReservationInvalidInstancePlatformRule {
	return &AwsEc2CapacityReservationInvalidInstancePlatformRule{
		resourceType:  "aws_ec2_capacity_reservation",
		attributeName: "instance_platform",
		enum: []string{
			"Linux/UNIX",
			"Red Hat Enterprise Linux",
			"SUSE Linux",
			"Windows",
			"Windows with SQL Server",
			"Windows with SQL Server Enterprise",
			"Windows with SQL Server Standard",
			"Windows with SQL Server Web",
			"Linux with SQL Server Standard",
			"Linux with SQL Server Web",
			"Linux with SQL Server Enterprise",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2CapacityReservationInvalidInstancePlatformRule) Name() string {
	return "aws_ec2_capacity_reservation_invalid_instance_platform"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2CapacityReservationInvalidInstancePlatformRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2CapacityReservationInvalidInstancePlatformRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2CapacityReservationInvalidInstancePlatformRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2CapacityReservationInvalidInstancePlatformRule) Check(runner *tflint.Runner) error {
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
					`instance_platform is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
