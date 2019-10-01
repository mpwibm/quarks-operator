package environment

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega/gexec"
)

// StartDependencies starts dependencies of cf-operator, like quarks-job
func (e *Environment) StartDependencies() error {
	cliPath, err := gexec.Build("code.cloudfoundry.org/quarks-job/cmd")
	cmd := exec.Command(cliPath, "-n "+e.Namespace)
	_, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

	return err
}
