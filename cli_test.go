package cli_test

import (
	. "github.com/Contra-Culture/cli"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cli", func() {
	Describe("creation", func() {
		Describe("New()", func() {
			Context("when wrong configuration", func() {
				It("returns new app", func() {
					app, r := New(
						func(*AppCfgr) {

						})
					Expect(app).NotTo(BeNil())
					Expect(r).NotTo(BeNil())
					Expect(r.String()).To(Equal("app configuration\n\t\t[ error ] no app title specified\n\t\t[ error ] no app description specified\n\t\t[ error ] no app default command specified\n"))
				})
			})
		})
	})
})
