/*
Copyright 2019 The Kruise Authors.
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"os"
	"testing"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/openkruise/kruise/test/e2e/framework"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	k8sframework "k8s.io/kubernetes/test/e2e/framework"
	e2ereporters "k8s.io/kubernetes/test/e2e/reporters"
)

const (
	// namespaceCleanupTimeout is how long to wait for the namespace to be deleted.
	// If there are any orphaned namespaces to clean up, this test is running
	// on a long lived cluster. A long wait here is preferably to spurious test
	// failures caused by leaked resources from a previous test run.
	namespaceCleanupTimeout = 15 * time.Minute
)

var progressReporter = &e2ereporters.ProgressReporter{}

var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	// Reference common test to make the import valid.
	//commontest.CurrentSuite = commontest.E2E
	progressReporter.SetStartMsg()
	return nil
}, func(data []byte) {
	// Run on all Ginkgo nodes
})

var _ = ginkgo.SynchronizedAfterSuite(func() {
	progressReporter.SetEndMsg()
}, func() {
	AfterSuiteActions()
	framework.Logf("Running AfterSuite actions on all nodes")
	framework.RunCleanupActions()
})

// RunE2ETests checks configuration parameters (specified through flags) and then runs
// E2E tests using the Ginkgo runner.
// If a "report directory" is specified, one or more JUnit test reports will be
// generated in this directory, and cluster logs will also be saved.
// This function is called on each Ginkgo node in parallel mode.
// RunE2ETests checks configuration parameters (specified through flags) and then runs
// E2E tests using the Ginkgo runner.
// If a "report directory" is specified, one or more JUnit test reports will be
// generated in this directory, and cluster logs will also be saved.
// This function is called on each Ginkgo node in parallel mode.
func RunE2ETests(t *testing.T) {
	// InitLogs disables contextual logging, without a way to enable it again
	// in the E2E test suite because it has no feature gates. It used to have a
	// misleading --feature-gates parameter but that didn't do what users
	// and developers expected (define which features the cluster supports)
	// and therefore got removed.
	//
	// Because contextual logging is useful and should get tested, it gets
	// re-enabled here unconditionally.
	logs.InitLogs()
	defer logs.FlushLogs()
	klog.EnableContextualLogging(true)

	progressReporter = e2ereporters.NewProgressReporter(framework.TestContext.ProgressReportURL)
	gomega.RegisterFailHandler(k8sframework.Fail)

	// Run tests through the Ginkgo runner with output to console + JUnit for Jenkins
	if framework.TestContext.ReportDir != "" {
		// TODO: we should probably only be trying to create this directory once
		// rather than once-per-Ginkgo-node.
		// NOTE: junit report can be simply created by executing your tests with the new --junit-report flags instead.
		if err := os.MkdirAll(framework.TestContext.ReportDir, 0755); err != nil {
			klog.Errorf("Failed creating report directory: %v", err)
		}
	}

	suiteConfig, reporterConfig := k8sframework.CreateGinkgoConfig()
	klog.Infof("Starting e2e run %q on Ginkgo node %d", framework.RunID, suiteConfig.ParallelProcess)
	ginkgo.RunSpecs(t, "Kubernetes e2e suite", suiteConfig, reporterConfig)
}
