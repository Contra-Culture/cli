package cli_test

import (
	. "github.com/Contra-Culture/cli"
	"github.com/Contra-Culture/report"
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
					Expect(report.ToString(r)).To(Equal("| app configuration\n\t[ error ] no app title specified\n\t[ error ] no app description specified\n"))
				})
			})
		})
	})
})
