package cli_test

import (
	. "github.com/Contra-Culture/cli"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cli", func() {
	Describe("creation", func() {
		Describe("New()", func() {
			It("returns new app", func() {
				app := New(func(*AppCfgr) {

				})
				Expect(app).NotTo(BeNil())
			})
		})
	})
})
