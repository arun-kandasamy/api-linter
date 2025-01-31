package aiphcp

import (
	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/rules/internal/utils"
	"github.com/jhump/protoreflect/desc"
)

// AddRules accepts a register function and registers each of
// this AIP's rules to it.
func AddRules(r lint.RuleRegistry) error {
	return r.Register(
		9001,
		rule3,
	)
}

func hasMethodSignatures(m *desc.MethodDescriptor) bool {
	sigs := utils.GetMethodSignatures(m)
	return len(sigs) > 0
}
