/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Storage_StorageAccount_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Custom namer because storage accounts have strict names
	namer := tc.Namer.WithSeparator("")

	// Create a storage account
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
			// TODO: They mark this property as optional but actually it is required
			AccessTier: &accessTier,
		},
	}

	tc.CreateResourceAndWait(acct)

	tc.Expect(acct.Status.Location).To(Equal(&tc.AzureRegion))
	expectedKind := storage.StorageAccountStatusKindStorageV2
	tc.Expect(acct.Status.Kind).To(Equal(&expectedKind))
	tc.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	// Run sub-tests on storage account
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Blob Services CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				StorageAccount_BlobServices_CRUD(tc, acct)
			},
		},
		testcommon.Subtest{
			Name: "Queue Services CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				StorageAccount_QueueServices_CRUD(tc, acct)
			},
		},
	)

	tc.DeleteResourceAndWait(acct)

	// Ensure that the account was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(storage.StorageAccountsSpecAPIVersion20210401))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func StorageAccount_BlobServices_CRUD(tc *testcommon.KubePerTestContext, storageAccount client.Object) {
	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccountsBlobServices_Spec{
			Owner: testcommon.AsOwner(storageAccount),
		},
	}

	tc.CreateResourceAndWait(blobService)
	// no DELETE, this is not a real resource

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Container CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				StorageAccount_BlobServices_Container_CRUD(testContext, blobService)
			},
		})
}

func StorageAccount_BlobServices_Container_CRUD(tc *testcommon.KubePerTestContext, blobService *storage.StorageAccountsBlobService) {
	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMeta("container"),
		Spec: storage.StorageAccountsBlobServicesContainers_Spec{
			Owner: testcommon.AsOwner(blobService),
		},
	}

	tc.CreateResourceAndWait(blobContainer)
	defer tc.DeleteResourceAndWait(blobContainer)
}

func StorageAccount_QueueServices_CRUD(tc *testcommon.KubePerTestContext, storageAccount client.Object) {
	queueService := &storage.StorageAccountsQueueService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccountsQueueServices_Spec{
			Owner: testcommon.AsOwner(storageAccount),
		},
	}

	tc.CreateResourceAndWait(queueService)
	// cannot delete - not a real resource

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Queue CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				StorageAccount_QueueServices_Queue_CRUD(tc, queueService)
			},
		},
	)
}

func StorageAccount_QueueServices_Queue_CRUD(tc *testcommon.KubePerTestContext, queueService *storage.StorageAccountsQueueService) {
	queue := &storage.StorageAccountsQueueServicesQueue{
		ObjectMeta: tc.MakeObjectMeta("queue"),
		Spec: storage.StorageAccountsQueueServicesQueues_Spec{
			Owner: testcommon.AsOwner(queueService),
		},
	}

	tc.CreateResourceAndWait(queue)
	defer tc.DeleteResourceAndWait(queue)
}

func Test_Storage_StorageAccount_SecretsFromAzure(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Custom namer because storage accounts have strict names
	namer := tc.Namer.WithSeparator("")

	// Create a storage account
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
			// Initially with no OperatorSpec.Secrets, to ensure no secrets are created
		},
	}

	tc.CreateResourceAndWait(acct)

	tc.Expect(acct.Status.Location).To(Equal(&tc.AzureRegion))
	expectedKind := storage.StorageAccountStatusKindStorageV2
	tc.Expect(acct.Status.Kind).To(Equal(&expectedKind))

	// There should be no secrets at this point
	list := &v1.SecretList{}
	tc.ListResources(list, client.InNamespace(tc.Namespace))
	tc.Expect(list.Items).To(HaveLen(0))

	// Run sub-tests on storage account
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				StorageAccount_SecretsWrittenToSameKubeSecret(tc, acct)
			},
		},
		testcommon.Subtest{
			Name: "SecretsWrittenToDifferentKubeSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				StorageAccount_SecretsWrittenToDifferentKubeSecrets(tc, acct)
			},
		},
	)
}

func StorageAccount_SecretsWrittenToSameKubeSecret(tc *testcommon.KubePerTestContext, acct *storage.StorageAccount) {
	old := acct.DeepCopy()
	storageKeysSecret := "storagekeys"
	acct.Spec.OperatorSpec = &storage.StorageAccountOperatorSpec{
		Secrets: &storage.StorageAccountOperatorSecrets{
			Key1:         &genruntime.SecretDestination{Name: storageKeysSecret, Key: "key1"},
			BlobEndpoint: &genruntime.SecretDestination{Name: storageKeysSecret, Key: "blob"},
		},
	}
	tc.PatchResourceAndWait(old, acct)

	tc.ExpectSecretHasKeys(storageKeysSecret, "key1", "blob")
}

func StorageAccount_SecretsWrittenToDifferentKubeSecrets(tc *testcommon.KubePerTestContext, acct *storage.StorageAccount) {
	old := acct.DeepCopy()
	key1Secret := "secret1"
	key2Secret := "secret2"
	blobSecret := "secret3"
	queueSecret := "secret4"
	tableSecret := "secret5"
	fileSecret := "secret6"
	webSecret := "secret7"
	dfsSecret := "secret8"

	// Writing 8 secrets from a single resource is a degenerate case (it's not likely people are
	// going to do this regularly), but best we make sure it works...
	acct.Spec.OperatorSpec = &storage.StorageAccountOperatorSpec{
		Secrets: &storage.StorageAccountOperatorSecrets{
			Key1:          &genruntime.SecretDestination{Name: key1Secret, Key: "key1"},
			Key2:          &genruntime.SecretDestination{Name: key2Secret, Key: "key2"},
			BlobEndpoint:  &genruntime.SecretDestination{Name: blobSecret, Key: "blob"},
			QueueEndpoint: &genruntime.SecretDestination{Name: queueSecret, Key: "queue"},
			TableEndpoint: &genruntime.SecretDestination{Name: tableSecret, Key: "table"},
			FileEndpoint:  &genruntime.SecretDestination{Name: fileSecret, Key: "file"},
			WebEndpoint:   &genruntime.SecretDestination{Name: webSecret, Key: "web"},
			DfsEndpoint:   &genruntime.SecretDestination{Name: dfsSecret, Key: "dfs"},
		},
	}
	tc.PatchResourceAndWait(old, acct)

	tc.ExpectSecretHasKeys(key1Secret, "key1")
	tc.ExpectSecretHasKeys(key2Secret, "key2")
	tc.ExpectSecretHasKeys(blobSecret, "blob")
	tc.ExpectSecretHasKeys(queueSecret, "queue")
	tc.ExpectSecretHasKeys(tableSecret, "table")
	tc.ExpectSecretHasKeys(fileSecret, "file")
	tc.ExpectSecretHasKeys(webSecret, "web")
	tc.ExpectSecretHasKeys(dfsSecret, "dfs")
}
