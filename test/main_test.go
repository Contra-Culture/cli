package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("example app", func() {
	BeforeSuite(func() {
		buildSession, err := gexec.Start(exec.Command("go", "build", "-v", "./main.go"), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(string(buildSession.Wait().Out.Contents())).Should(Equal(""))
	})
	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	Describe("version", func() {
		It("prints version output", func() {
			session, err := gexec.Start(exec.Command("./main", "version"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			output := string(session.Wait().Out.Contents())
			Expect(output).To(Equal("0.0.1 (test)\n"))
		})
	})
	Describe("help", func() {
		It("prints help output", func() {
			session, err := gexec.Start(exec.Command("./main", "help"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			output := string(session.Wait().Out.Contents())
			Expect(output).To(Equal("testapp (0.0.1 (test))\n\ttestapp is a test application which is an example of use of github.com/Contra-Culture/cli library.\n\n\t>> commands <<\n\n\t[default] - no-title\n\t\tprint its params\n\n\t\t>> parameters <<\n\n\t\t-filePath\n\t\t\tpath to file\n\n\t\t-port\n\t\t\tport to listen\n\n\t\t-verbose\n\t\t\tverbose mode in which more detailed output is presented\n\n\techo - echo\n\t\tprints your text\n\n\t\t>> parameters <<\n\n\t\t-message\n\t\t\treturns your message back\n\n\thello - hello\n\t\tprints hello message\n\n\t\t>> parameters <<\n\n\t\t-name\n\t\t\tname for welcome\n\n\t\t-upcase\n\t\t\tif passed upcaes the text\n\n\thelp - help info\n\t\thelp shows help information.\n\n\tversion - version and build info\n\t\tversion shows the application version, build information and credits.\n"))
		})
	})

	Describe("custom commands", func() {
		Describe("[default]", func() {
			It("prints [default] output", func() {
				session, err := gexec.Start(exec.Command("./main"), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				output := string(session.Wait().Out.Contents())
				Expect(output).To(Equal(""))
			})
		})
		Describe("echo", func() {
			It("prints echo output", func() {
				session, err := gexec.Start(exec.Command("./main", "echo", "-message=hi"), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				output := string(session.Wait().Out.Contents())
				Expect(output).To(Equal(""))
			})
		})
		Describe("hello", func() {

		})
	})
})
