package basic

import (
	"testing"
	"time"

	"e2e/internal/testhelpers"

	"github.com/giantswarm/apptest-framework/v2/pkg/state"
	"github.com/giantswarm/apptest-framework/v2/pkg/suite"
	"github.com/giantswarm/clustertest/v2/pkg/failurehandler"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	isUpgrade = false
)

func TestBasic(t *testing.T) {
	suite.New().
		WithInCluster(true).
		WithInstallNamespace("").
		WithIsUpgrade(isUpgrade).
		WithValuesFile("./values.yaml").
		Tests(func() {
			It("should have the HelmRelease ready on the management cluster", func() {
				mcClient := state.GetFramework().MC()
				clusterName := state.GetCluster().Name
				orgName := state.GetCluster().Organization.Name

				Eventually(func() (bool, error) {
					ready, err := testhelpers.HelmReleaseIsReady(*mcClient, clusterName, orgName)
					if err != nil {
						GinkgoLogr.Info("HelmRelease check failed", "error", err.Error())
					} else if !ready {
						GinkgoLogr.Info("HelmRelease not ready yet", "name", clusterName+"-aws-node-termination-handler")
					} else {
						GinkgoLogr.Info("HelmRelease is ready", "name", clusterName+"-aws-node-termination-handler")
					}
					return ready, err
				}).
					WithTimeout(15 * time.Minute).
					WithPolling(10 * time.Second).
					Should(BeTrue(), failurehandler.LLMPrompt(state.GetFramework(), state.GetCluster(), "Investigate HelmRelease not ready for aws-node-termination-handler"))
			})

			It("should have the aws-node-termination-handler deployment running", func() {
				wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
				Expect(err).Should(Succeed())

				Eventually(func() error {
					var dp appsv1.Deployment
					err := wcClient.Get(state.GetContext(), types.NamespacedName{
						Namespace: "kube-system",
						Name:      "aws-node-termination-handler",
					}, &dp)
					if err != nil {
						GinkgoLogr.Info("aws-node-termination-handler not found yet", "error", err.Error())
					} else {
						GinkgoLogr.Info("aws-node-termination-handler found",
							"replicas", dp.Status.Replicas,
							"ready", dp.Status.ReadyReplicas,
							"available", dp.Status.AvailableReplicas,
						)
					}
					return err
				}).
					WithTimeout(10 * time.Minute).
					WithPolling(5 * time.Second).
					ShouldNot(HaveOccurred(), failurehandler.LLMPrompt(state.GetFramework(), state.GetCluster(), "Investigate aws-node-termination-handler deployment not found or not running"))
			})

			It("should have at least one ready replica", func() {
				wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
				Expect(err).Should(Succeed())

				Eventually(func() (int32, error) {
					var dp appsv1.Deployment
					err := wcClient.Get(state.GetContext(), types.NamespacedName{
						Namespace: "kube-system",
						Name:      "aws-node-termination-handler",
					}, &dp)
					if err != nil {
						return 0, err
					}
					GinkgoLogr.Info("deployment replicas", "ready", dp.Status.ReadyReplicas, "desired", *dp.Spec.Replicas)
					return dp.Status.ReadyReplicas, nil
				}).
					WithTimeout(10 * time.Minute).
					WithPolling(5 * time.Second).
					Should(BeNumerically(">=", int32(1)), failurehandler.LLMPrompt(state.GetFramework(), state.GetCluster(), "Investigate aws-node-termination-handler has no ready replicas"))
			})
		}).
		Run(t, "Node Termination Handler Basic")
}
