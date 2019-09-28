// This file generated by `tools/api-rule-gen/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsRouteInvalidInstanceRule checks whether attribute value actually exists
type AwsRouteInvalidInstanceRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsRouteInvalidInstanceRule returns new rule with default attributes
func NewAwsRouteInvalidInstanceRule() *AwsRouteInvalidInstanceRule {
	return &AwsRouteInvalidInstanceRule{
		resourceType:  "aws_route",
		attributeName: "instance_id",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsRouteInvalidInstanceRule) Name() string {
	return "aws_route_invalid_instance"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRouteInvalidInstanceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRouteInvalidInstanceRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRouteInvalidInstanceRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeInstances
func (r *AwsRouteInvalidInstanceRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeInstances")
			var err error
			r.data, err = runner.AwsClient.DescribeInstances()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeInstances",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid instance ID.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
