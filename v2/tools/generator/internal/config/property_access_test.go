/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
	. "github.com/onsi/gomega"
)

func Test_PropertyAccess_Lookup_ReturnsConfiguredValue_WhenPresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	ref := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	value := ExplicitCollections

	model := NewObjectModelConfiguration()
	g.Expect(
		model.ModifyProperty(
			ref,
			"Name",
			func(tc *PropertyConfiguration) error {
				tc.PayloadType.Set(value)
				return nil
			}),
	).To(Succeed())

	access := makePropertyAccess[PayloadType](
		model,
		func(t *PropertyConfiguration) *configurable[PayloadType] {
			return &t.PayloadType
		})

	// Act
	actual, ok := access.Lookup(ref, "Name")

	// Assert
	g.Expect(ok).To(BeTrue())
	g.Expect(actual).To(Equal(value))
}
