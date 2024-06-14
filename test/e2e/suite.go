package e2e

import (
	"github.com/openkruise/kruise/test/e2e/framework"
	k8sframework "k8s.io/kubernetes/test/e2e/framework"
)

// AfterSuiteActions are actions that are run on ginkgo's SynchronizedAfterSuite
func AfterSuiteActions() {
	// Run only Ginkgo on node 1
	framework.Logf("Running AfterSuite actions on node 1")
	if framework.TestContext.ReportDir != "" {
		k8sframework.CoreDump(framework.TestContext.ReportDir)
	}
}
