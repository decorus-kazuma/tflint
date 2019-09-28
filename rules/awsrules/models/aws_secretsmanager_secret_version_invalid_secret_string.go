// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsSecretsmanagerSecretVersionInvalidSecretStringRule checks the pattern is valid
type AwsSecretsmanagerSecretVersionInvalidSecretStringRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsSecretsmanagerSecretVersionInvalidSecretStringRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretVersionInvalidSecretStringRule() *AwsSecretsmanagerSecretVersionInvalidSecretStringRule {
	return &AwsSecretsmanagerSecretVersionInvalidSecretStringRule{
		resourceType:  "aws_secretsmanager_secret_version",
		attributeName: "secret_string",
		max:           10240,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Name() string {
	return "aws_secretsmanager_secret_version_invalid_secret_string"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretVersionInvalidSecretStringRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"secret_string must be 10240 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
