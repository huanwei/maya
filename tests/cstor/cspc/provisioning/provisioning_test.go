/*
Copyright 2019 The OpenEBS Authors
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

package provisioning

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cspc_v1alpha1 "github.com/openebs/maya/pkg/cstor/poolcluster/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var skipPositiveCaseIfRequired = false

var _ = Describe("[CSPC] CSTOR STRIPE POOL PROVISIONING AND RECONCILIATION ", func() {
	provisioningAndReconciliationTest(createCSPCObjectForStripe)
})
var _ = Describe("[CSPC] CSTOR MIRROR POOL PROVISIONING AND RECONCILIATION ", func() {
	provisioningAndReconciliationTest(createCSPCObjectForMirror)
})
var _ = Describe("[CSPC] CSTOR RAIDZ POOL PROVISIONING AND RECONCILIATION ", func() {
	provisioningAndReconciliationTest(createCSPCObjectForRaidz)
})
var _ = Describe("[CSPC] CSTOR RAIDZ2 POOL PROVISIONING AND RECONCILIATION ", func() {
	provisioningAndReconciliationTest(createCSPCObjectForRaidz2)
})
var _ = Describe("[CSPC] CSTOR STRIPE POOL PROVISIONING AND DownScaling With BDC cleanup ", func() {
	provisioningAndCleanupTest(createCSPCObjectForStripe)
})

func provisioningAndReconciliationTest(createCSPCObject func()) {
	When("A CSPC Is Created", func() {
		It("cStor Pools Should be Provisioned ", func() {
			SkipTest(skipPositiveCaseIfRequired)
			By("Preparing A CSPC Object, No Error Should Occur", createCSPCObject)

			By("Creating A CSPC Object, Desired Number of CSPIs Should Be Created", verifyDesiredCSPICount)
		})
	})

	When("The CSPC Finalizer Is Removed From CSPC", func() {
		It("The Finalizer Should Come Back As Part Of Reconcilation", func() {
			SkipTest(skipPositiveCaseIfRequired)
			err := Cspc.RemoveFinalizer(cspc_v1alpha1.CSPCFinalizer)
			Expect(err).To(BeNil())
			Expect(ops.IsCSPCFinalizerExistsOnCSPC(cspcObj.Name, cspc_v1alpha1.CSPCFinalizer)).To(BeTrue())
		})
	})
	// TODO : Add test case for pool import
	// TODO : Remove hardcoded expected count and pass via flag to enable test to run on single node also.
	for i := 1; i <= 3; i++ {
		When(strconv.Itoa(i)+" CSPI Is Deleted", func() {
			// pin it
			i := i
			It(strconv.Itoa(i)+" New CSPI Should Come Up Again", func() {
				SkipTest(skipPositiveCaseIfRequired)
				ops.DeleteCSPI(cspcObj.Name, i)
				// We expect 3 cstorPool objects.
				cspCount := ops.GetHealthyCSPICount(cspcObj.Name, 3)
				Expect(cspCount).To(Equal(3))
			})
		})
	}
	// cleanup cspcObj after each test
	cleanCSPCObject()
}

func provisioningAndCleanupTest(createCSPCObject func()) {
	When("A CSPC Is Created", func() {
		It("cStor Pools Should be Provisioned ", func() {
			SkipTest(skipPositiveCaseIfRequired)
			By("Preparing A CSPC Object, No Error Should Occur", createCSPCObject)

			By("Creating A CSPC Object, Desired Number of CSPIs Should Be Created", verifyDesiredCSPICount)
		})
	})

	When("The CSPC Is DownScaled ", func() {
		It("Extra BDCs Should Be Removed", func() {
			SkipTest(skipPositiveCaseIfRequired)
			downScaleCSPCObject()
			listOpts := metav1.ListOptions{
				LabelSelector: getLabelSelector(cspcObj),
			}
			Expect(ops.GetBDCCountEventually(listOpts, 2, ops.NameSpace)).Should(Equal(2))
		})
	})
	// cleanup cspcObj after each test
	cleanCSPCObject()
}
