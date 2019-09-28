// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsDynamoDBTableItemInvalidTableNameRule checks the pattern is valid
type AwsDynamoDBTableItemInvalidTableNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsDynamoDBTableItemInvalidTableNameRule returns new rule with default attributes
func NewAwsDynamoDBTableItemInvalidTableNameRule() *AwsDynamoDBTableItemInvalidTableNameRule {
	return &AwsDynamoDBTableItemInvalidTableNameRule{
		resourceType:  "aws_dynamodb_table_item",
		attributeName: "table_name",
		max:           255,
		min:           3,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`),
	}
}

// Name returns the rule name
func (r *AwsDynamoDBTableItemInvalidTableNameRule) Name() string {
	return "aws_dynamodb_table_item_invalid_table_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBTableItemInvalidTableNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBTableItemInvalidTableNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBTableItemInvalidTableNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBTableItemInvalidTableNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"table_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"table_name must be 3 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`table_name does not match valid pattern ^[a-zA-Z0-9_.-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
