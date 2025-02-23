/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Networking_NetworkSecurityGroup_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Network Security Group
	nsg := &network.NetworkSecurityGroup{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("nsg")),
		Spec: network.NetworkSecurityGroups_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(nsg)

	tc.Expect(nsg.Status.Id).ToNot(BeNil())
	armId := *nsg.Status.Id

	// Perform a simple patch
	old := nsg.DeepCopy()
	nsg.Spec.Tags = map[string]string{
		"foo": "bar",
	}
	tc.PatchResourceAndWait(old, nsg)
	tc.Expect(nsg.Status.Tags).To(HaveKey("foo"))

	// Run sub tests
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SecurityRules CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				NetworkSecurityGroup_SecurityRules_CRUD(testContext, nsg)
			},
		},
	)

	tc.DeleteResourceAndWait(nsg)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.NetworkSecurityGroupsSecurityRulesSpecAPIVersion20201101))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func NetworkSecurityGroup_SecurityRules_CRUD(tc *testcommon.KubePerTestContext, nsg client.Object) {
	rule1 := &network.NetworkSecurityGroupsSecurityRule{
		ObjectMeta: tc.MakeObjectMeta("rule1"),
		Spec: network.NetworkSecurityGroupsSecurityRules_Spec{
			Owner:                    testcommon.AsOwner(nsg),
			Protocol:                 network.SecurityRulePropertiesFormatProtocolTcp,
			SourcePortRange:          to.StringPtr("23-45"),
			DestinationPortRange:     to.StringPtr("46-56"),
			SourceAddressPrefix:      to.StringPtr("*"),
			DestinationAddressPrefix: to.StringPtr("*"),
			Access:                   network.SecurityRulePropertiesFormatAccessAllow,
			Priority:                 123,
			Direction:                network.SecurityRulePropertiesFormatDirectionInbound,
			Description:              to.StringPtr("The first rule of networking is don't talk about networking"),
		},
	}

	rule2 := &network.NetworkSecurityGroupsSecurityRule{
		ObjectMeta: tc.MakeObjectMeta("rule2"),
		Spec: network.NetworkSecurityGroupsSecurityRules_Spec{
			Owner:    testcommon.AsOwner(nsg),
			Protocol: network.SecurityRulePropertiesFormatProtocolTcp,
			SourcePortRanges: []string{
				"23-45",
				"5000-5100",
			},
			DestinationPortRange:     to.StringPtr("*"),
			SourceAddressPrefix:      to.StringPtr("*"),
			DestinationAddressPrefix: to.StringPtr("*"),
			Access:                   network.SecurityRulePropertiesFormatAccessDeny,
			Priority:                 124,
			Direction:                network.SecurityRulePropertiesFormatDirectionInbound,
		},
	}

	tc.CreateResourcesAndWait(rule1, rule2)
	defer tc.DeleteResourcesAndWait(rule1, rule2)

	tc.Expect(rule1.Status.Id).ToNot(BeNil())
	tc.Expect(rule2.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(rule1.Status.SourcePortRange).ToNot(BeNil())
	tc.Expect(*rule1.Status.SourcePortRange).To(Equal("23-45"))

	tc.Expect(rule2.Status.SourcePortRanges).To(HaveLen(2))
	tc.Expect(rule2.Status.SourcePortRanges[0]).To(Equal("23-45"))
	tc.Expect(rule2.Status.SourcePortRanges[1]).To(Equal("5000-5100"))

	// Perform a simple patch
	old := rule1.DeepCopy()
	newPriority := 100
	rule1.Spec.Priority = newPriority
	tc.PatchResourceAndWait(old, rule1)
	tc.Expect(rule1.Status.Priority).To(Equal(&newPriority))
}
