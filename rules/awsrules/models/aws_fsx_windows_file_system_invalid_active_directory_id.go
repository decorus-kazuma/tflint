// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule checks the pattern is valid
type AwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule returns new rule with default attributes
func NewAwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule() *AwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule {
	return &AwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule{
		resourceType:  "aws_fsx_windows_file_system",
		attributeName: "active_directory_id",
		max:           12,
		min:           12,
		pattern:       regexp.MustCompile(`^d-[0-9a-f]{10}$`),
	}
}

// Name returns the rule name
func (r *AwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule) Name() string {
	return "aws_fsx_windows_file_system_invalid_active_directory_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"active_directory_id must be 12 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"active_directory_id must be 12 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`active_directory_id does not match valid pattern ^d-[0-9a-f]{10}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
