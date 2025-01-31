package aiphcp

import (
	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/locations"
	"github.com/jhump/protoreflect/desc"
	apb "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
)

var rule3 = &lint.MethodRule{
	Name:       lint.NewRuleName(9001, "no-additional-bindings"),
	LintMethod: rule3LintMethod,
}

func rule3LintMethod(m *desc.MethodDescriptor) []lint.Problem {
	methodOptions := m.GetMethodOptions()

	if x := proto.GetExtension(methodOptions, apb.E_Http); x != nil {
		httpRule, ok := proto.GetExtension(methodOptions, apb.E_Http).(*apb.HttpRule)
		if ok {
			additionalBindingsCount := len(httpRule.GetAdditionalBindings())
			if additionalBindingsCount > 0 {
				return []lint.Problem{{
					Message:    "The additional_bindings MUST NOT be used",
					Descriptor: m,
					Location:   locations.MethodHTTPRule(m),
				}}
			}
		}
	}

	return nil
}
