package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = BeforeSuite(func() {
	buildSession, err := gexec.Start(exec.Command("go", "build", "-v", "./main.go"), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(string(buildSession.Wait().Out.Contents())).Should(Equal(""))
})
var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
	exec.Command("rm", "main").Run()
})
var _ = Describe("example app", func() {
	Describe("version", func() {
		It("prints version output", func() {
			session, err := gexec.Start(exec.Command("./main", "version"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			output := string(session.Wait().Out.Contents())
			Expect(output).To(Equal(""))
		})
	})
	Describe("help", func() {
		It("prints help output", func() {
			session, err := gexec.Start(exec.Command("./main", "help"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			output := string(session.Wait().Out.Contents())
			Expect(output).To(Equal(""))
		})
	})
	Describe("custom commands", func() {
		Describe("[default]", func() {
			It("prints [default] output", func() {
				session, err := gexec.Start(exec.Command("./main", "-filePath=/", "-port=100"), GinkgoWriter, GinkgoWriter)
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
			It("prints hello output", func() {
				session, err := gexec.Start(exec.Command("./main", "hello", "-name=Johny"), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				output := string(session.Wait().Out.Contents())
				Expect(output).To(Equal(""))
			})
		})
	})
})
