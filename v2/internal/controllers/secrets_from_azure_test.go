/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1alpha1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func Test_WhenObjectDeleted_SecretIsDeleted(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a storage account
	storageKeysSecret := "storagekeys"
	acct := makeSimpleStorageAccountWithOperatorSpecSecrets(tc, rg, storageKeysSecret, "key1")

	tc.CreateResourceAndWait(acct)

	// Expect that the secret exists and has an ownerReference set. We can't actually ensure
	// that the resource is deleted in envtest because the controller that performs ownerReference
	// resource deletion is not running in envtest (so resources do not get deleted when they would
	// in a real cluster).
	secretName := types.NamespacedName{Namespace: tc.Namespace, Name: storageKeysSecret}
	var secret v1.Secret
	tc.GetResource(secretName, &secret)
	tc.Expect(secret.OwnerReferences).To(HaveLen(1))
	tc.Expect(secret.OwnerReferences[0].UID).To(Equal(acct.GetUID()))

	// Delete the resource
	tc.DeleteResourceAndWait(acct)
}

func Test_WhenObjectPullSecretsAndSecretAlreadyExists_WarningConditionIsSet(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a storage account
	storageKeysSecret := "storagekeys"
	acct := makeSimpleStorageAccountWithOperatorSpecSecrets(tc, rg, storageKeysSecret, "key1")

	collidingSecret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMetaWithName(storageKeysSecret),
		StringData: map[string]string{
			"val1": "abc",
		},
	}
	tc.CreateResource(collidingSecret)

	tc.CreateResourceAndWaitForState(acct, metav1.ConditionFalse, conditions.ConditionSeverityWarning)

	// Expect that the ARM ID in status is set. This indicates that status is filled out and the resource
	// has been created in Azure
	tc.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	// We expect the ready condition to include details of the error
	tc.Expect(acct.Status.Conditions[0].Reason).To(Equal("FailedWritingSecret"))
	tc.Expect(acct.Status.Conditions[0].Message).To(MatchRegexp("cannot overwrite secret.*which is not owned by"))

	// Delete the resource, it should be able to proceed and delete the underlying Azure resource
	tc.DeleteResourceAndWait(acct)

	// Ensure that the account was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(storage.StorageAccountsSpecAPIVersion20210401))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Test_TwoObjectsWriteSameSecret_WarningConditionIsSetOnSecond(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a storage account
	storageKeysSecret := "k1"
	acct1 := makeSimpleStorageAccountWithOperatorSpecSecrets(tc, rg, storageKeysSecret, "key1")
	acct2 := makeSimpleStorageAccountWithOperatorSpecSecrets(tc, rg, storageKeysSecret, "key2")

	tc.CreateResourceAndWait(acct1)
	tc.CreateResourceAndWaitForState(acct2, metav1.ConditionFalse, conditions.ConditionSeverityWarning)

	// We expect the ready condition to include details of the error
	tc.Expect(acct2.Status.Conditions[0].Reason).To(Equal("FailedWritingSecret"))
	tc.Expect(acct2.Status.Conditions[0].Message).To(MatchRegexp("cannot overwrite secret.*which is not owned by"))
}

func makeSimpleStorageAccountWithOperatorSpecSecrets(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, secretName string, secretKey string) *storage.StorageAccount {
	namer := tc.Namer.WithSeparator("")
	accessTier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("stor")),
		Spec: storage.StorageAccounts_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     storage.StorageAccountsSpecKindStorageV2,
			Sku: storage.Sku{
				Name: storage.SkuNameStandardLRS,
			},
			AccessTier: &accessTier,
			OperatorSpec: &storage.StorageAccountOperatorSpec{
				Secrets: &storage.StorageAccountOperatorSecrets{
					Key1: &genruntime.SecretDestination{
						Name: secretName,
						Key:  secretKey,
					},
				},
			},
		},
	}

	return acct
}
