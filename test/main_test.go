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
		exec.Command("rm", "main").Run()
	})

	Describe("version", func() {
		It("prints version output", func() {
			session, err := gexec.Start(exec.Command("./main", "version"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			output := string(session.Wait().Out.Contents())
			Expect(output).To(Equal("| app configuration\n\t| default-command\n\t\t| parameter\n\t\t\t[ info ] param name \"filePath\"\n\t\t\t[ info ] param description \"path to file\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"port\"\n\t\t\t[ info ] param description \"port to listen\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"verbose\"\n\t\t\t[ info ] param description \"verbose mode in which more detailed output is presented\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\n\t| command \"echo\"\n\t\t| parameter\n\t\t\t[ info ] param name \"message\"\n\t\t\t[ info ] param description \"returns your message back\"\n\t\t\t[ info ] param checker specified\n\t| command \"hello\"\n\t\t| parameter\n\t\t\t[ info ] param name \"name\"\n\t\t\t[ info ] param description \"name for hello\"\n\t\t\t[ info ] param checker specified\n\t\t\t| dependent parameter\n\t\t\t\t[ info ] param name \"lastname\"\n\t\t\t\t[ info ] param description \"lastname for hello\"\n\t\t\t\t[ info ] param default value \"\"\n\t\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"upcase\"\n\t\t\t[ info ] param description \"if passed upcaes the text\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\n0.0.1 (test)\n| app initialization\n\t[ info ] called with params: \"[]string{\"./main\", \"version\"}\"\n\t[ info ] command \"version and build info\" picked\n\t| command version and build info execution\n\t\t[ info ] command \"version\" called with: \"map[string]string{}\"\n"))
		})
	})
	Describe("help", func() {
		It("prints help output", func() {
			session, err := gexec.Start(exec.Command("./main", "help"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			output := string(session.Wait().Out.Contents())
			Expect(output).To(Equal("| app configuration\n\t| default-command\n\t\t| parameter\n\t\t\t[ info ] param name \"filePath\"\n\t\t\t[ info ] param description \"path to file\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"port\"\n\t\t\t[ info ] param description \"port to listen\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"verbose\"\n\t\t\t[ info ] param description \"verbose mode in which more detailed output is presented\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\n\t| command \"echo\"\n\t\t| parameter\n\t\t\t[ info ] param name \"message\"\n\t\t\t[ info ] param description \"returns your message back\"\n\t\t\t[ info ] param checker specified\n\t| command \"hello\"\n\t\t| parameter\n\t\t\t[ info ] param name \"name\"\n\t\t\t[ info ] param description \"name for hello\"\n\t\t\t[ info ] param checker specified\n\t\t\t| dependent parameter\n\t\t\t\t[ info ] param name \"lastname\"\n\t\t\t\t[ info ] param description \"lastname for hello\"\n\t\t\t\t[ info ] param default value \"\"\n\t\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"upcase\"\n\t\t\t[ info ] param description \"if passed upcaes the text\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\ntestapp (0.0.1 (test))\n\ttestapp is a test application which is an example of use of github.com/Contra-Culture/cli library.\n\n\t>> commands <<\n\n\t[default] - parameters\n\t\tprints its params\n\n\t\t>> parameters <<\n\n\t\t-filePath\n\t\t\tpath to file\n\n\t\t-port\n\t\t\tport to listen\n\n\t\t-verbose (optional, default: y)\n\t\t\tverbose mode in which more detailed output is presented\n\n\techo - echo\n\t\tprints your text\n\n\t\t>> parameters <<\n\n\t\t-message\n\t\t\treturns your message back\n\n\thello - hello\n\t\tprints hello message\n\n\t\t>> parameters <<\n\n\t\t-name\n\t\t\tname for hello\n\n\t\t-upcase\n\t\t\tif passed upcaes the text\n\n\thelp - help info\n\t\thelp shows help information.\n\n\tversion - version and build info\n\t\tversion shows the application version, build information and credits.\n| app initialization\n\t[ info ] called with params: \"[]string{\"./main\", \"help\"}\"\n\t[ info ] command \"help info\" picked\n\t| command help info execution\n\t\t[ info ] command \"help\" called with: \"map[string]string{}\"\n"))
		})
	})

	Describe("custom commands", func() {
		Describe("[default]", func() {
			It("prints [default] output", func() {
				session, err := gexec.Start(exec.Command("./main", "-filePath=/", "-port=100"), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				output := string(session.Wait().Out.Contents())
				Expect(output).To(Equal("| app configuration\n\t| default-command\n\t\t| parameter\n\t\t\t[ info ] param name \"filePath\"\n\t\t\t[ info ] param description \"path to file\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"port\"\n\t\t\t[ info ] param description \"port to listen\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"verbose\"\n\t\t\t[ info ] param description \"verbose mode in which more detailed output is presented\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\n\t| command \"echo\"\n\t\t| parameter\n\t\t\t[ info ] param name \"message\"\n\t\t\t[ info ] param description \"returns your message back\"\n\t\t\t[ info ] param checker specified\n\t| command \"hello\"\n\t\t| parameter\n\t\t\t[ info ] param name \"name\"\n\t\t\t[ info ] param description \"name for hello\"\n\t\t\t[ info ] param checker specified\n\t\t\t| dependent parameter\n\t\t\t\t[ info ] param name \"lastname\"\n\t\t\t\t[ info ] param description \"lastname for hello\"\n\t\t\t\t[ info ] param default value \"\"\n\t\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"upcase\"\n\t\t\t[ info ] param description \"if passed upcaes the text\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\n\n\nparams: map[string]string{\"filePath\":\"/\", \"port\":\"100\"}\n| app initialization\n\t[ info ] called with params: \"[]string{\"./main\", \"-filePath=/\", \"-port=100\"}\"\n\t[ info ] default command picked\n\t| command parameters execution\n\t\t[ info ] command \"[default]\" called with: \"map[string]string{\"filePath\":\"/\", \"port\":\"100\"}\"\n\t\t| prepare param \"filePath\"\n\t\t\t| check param \"filePath\"=\"/\"\n\t\t| prepare param \"port\"\n\t\t\t| check param \"port\"=\"100\"\n\t\t| prepare param \"verbose\"\n\t\t\t[ info ] trying to get param \"verbose\" from env variable\n\t\t\t[ info ] env variable does not provide \"verbose\" parameter value\n\t\t\t[ info ] default value \"y\" for \"verbose\" parameter picked\n"))
			})
		})
		Describe("echo", func() {
			It("prints echo output", func() {
				session, err := gexec.Start(exec.Command("./main", "echo", "-message=hi"), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				output := string(session.Wait().Out.Contents())
				Expect(output).To(Equal("| app configuration\n\t| default-command\n\t\t| parameter\n\t\t\t[ info ] param name \"filePath\"\n\t\t\t[ info ] param description \"path to file\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"port\"\n\t\t\t[ info ] param description \"port to listen\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"verbose\"\n\t\t\t[ info ] param description \"verbose mode in which more detailed output is presented\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\n\t| command \"echo\"\n\t\t| parameter\n\t\t\t[ info ] param name \"message\"\n\t\t\t[ info ] param description \"returns your message back\"\n\t\t\t[ info ] param checker specified\n\t| command \"hello\"\n\t\t| parameter\n\t\t\t[ info ] param name \"name\"\n\t\t\t[ info ] param description \"name for hello\"\n\t\t\t[ info ] param checker specified\n\t\t\t| dependent parameter\n\t\t\t\t[ info ] param name \"lastname\"\n\t\t\t\t[ info ] param description \"lastname for hello\"\n\t\t\t\t[ info ] param default value \"\"\n\t\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"upcase\"\n\t\t\t[ info ] param description \"if passed upcaes the text\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\n\necho > hi\n| app initialization\n\t[ info ] called with params: \"[]string{\"./main\", \"echo\", \"-message=hi\"}\"\n\t[ info ] command \"echo\" picked\n\t| command echo execution\n\t\t[ info ] command \"echo\" called with: \"map[string]string{\"message\":\"hi\"}\"\n\t\t| prepare param \"message\"\n\t\t\t| check param \"message\"=\"hi\"\n"))
			})
		})
		Describe("hello", func() {
			It("prints hello output", func() {
				session, err := gexec.Start(exec.Command("./main", "hello", "-name=Johny"), GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				output := string(session.Wait().Out.Contents())
				Expect(output).To(Equal("| app configuration\n\t| default-command\n\t\t| parameter\n\t\t\t[ info ] param name \"filePath\"\n\t\t\t[ info ] param description \"path to file\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"port\"\n\t\t\t[ info ] param description \"port to listen\"\n\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"verbose\"\n\t\t\t[ info ] param description \"verbose mode in which more detailed output is presented\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\n\t| command \"echo\"\n\t\t| parameter\n\t\t\t[ info ] param name \"message\"\n\t\t\t[ info ] param description \"returns your message back\"\n\t\t\t[ info ] param checker specified\n\t| command \"hello\"\n\t\t| parameter\n\t\t\t[ info ] param name \"name\"\n\t\t\t[ info ] param description \"name for hello\"\n\t\t\t[ info ] param checker specified\n\t\t\t| dependent parameter\n\t\t\t\t[ info ] param name \"lastname\"\n\t\t\t\t[ info ] param description \"lastname for hello\"\n\t\t\t\t[ info ] param default value \"\"\n\t\t\t\t[ info ] param checker specified\n\t\t| parameter\n\t\t\t[ info ] param name \"upcase\"\n\t\t\t[ info ] param description \"if passed upcaes the text\"\n\t\t\t[ info ] param default value \"y\"\n\t\t\t[ info ] param checker specified\n| app initialization\n\t[ info ] called with params: \"[]string{\"./main\", \"hello\", \"-name=Johny\"}\"\n\t[ info ] command \"hello\" picked\n\t| command hello execution\n\t\t[ info ] command \"hello\" called with: \"map[string]string{\"name\":\"Johny\"}\"\n\t\t| prepare param \"name\"\n\t\t\t| check param \"name\"=\"Johny\"\n\t\t\t| prepare param \"name\"\n\t\t\t\t[ info ] trying to get param \"name.\" from env variable\n\t\t\t\t[ info ] env variable does not provide \"name.\" parameter value\n\t\t\t\t[ info ] default value \"\" for \"name.\" parameter picked\n\t\t\t\t[ error ] praram \"name.\" required\n"))
			})
		})
	})
})
