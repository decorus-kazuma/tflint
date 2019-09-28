// This file generated by `tools/api-rule-gen/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsDBInstanceInvalidDBSubnetGroupRule checks whether attribute value actually exists
type AwsDBInstanceInvalidDBSubnetGroupRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsDBInstanceInvalidDBSubnetGroupRule returns new rule with default attributes
func NewAwsDBInstanceInvalidDBSubnetGroupRule() *AwsDBInstanceInvalidDBSubnetGroupRule {
	return &AwsDBInstanceInvalidDBSubnetGroupRule{
		resourceType:  "aws_db_instance",
		attributeName: "db_subnet_group_name",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsDBInstanceInvalidDBSubnetGroupRule) Name() string {
	return "aws_db_instance_invalid_db_subnet_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDBInstanceInvalidDBSubnetGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDBInstanceInvalidDBSubnetGroupRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDBInstanceInvalidDBSubnetGroupRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeDBSubnetGroups
func (r *AwsDBInstanceInvalidDBSubnetGroupRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeDBSubnetGroups")
			var err error
			r.data, err = runner.AwsClient.DescribeDBSubnetGroups()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeDBSubnetGroups",
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
					fmt.Sprintf(`"%s" is invalid DB subnet group name.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
