// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generator information:
//- Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
//- ARM URI:
///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}
type SqlDatabaseContainerTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersTriggers_SPEC `json:"spec,omitempty"`
	Status            SqlTrigger_Status                                   `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerTrigger{}

// GetConditions returns the conditions of the resource
func (trigger *SqlDatabaseContainerTrigger) GetConditions() conditions.Conditions {
	return trigger.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (trigger *SqlDatabaseContainerTrigger) SetConditions(conditions conditions.Conditions) {
	trigger.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerTrigger{}

// ConvertFrom populates our SqlDatabaseContainerTrigger from the provided hub SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20210515storage.SqlDatabaseContainerTrigger)
	if !ok {
		return fmt.Errorf("expected storage:documentdb/v1alpha1api20210515storage/SqlDatabaseContainerTrigger but received %T instead", hub)
	}

	return trigger.AssignPropertiesFromSqlDatabaseContainerTrigger(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerTrigger from our SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20210515storage.SqlDatabaseContainerTrigger)
	if !ok {
		return fmt.Errorf("expected storage:documentdb/v1alpha1api20210515storage/SqlDatabaseContainerTrigger but received %T instead", hub)
	}

	return trigger.AssignPropertiesToSqlDatabaseContainerTrigger(destination)
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainertrigger,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainertriggers,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabasecontainertriggers.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabaseContainerTrigger{}

// Default applies defaults to the SqlDatabaseContainerTrigger resource
func (trigger *SqlDatabaseContainerTrigger) Default() {
	trigger.defaultImpl()
	var temp interface{} = trigger
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (trigger *SqlDatabaseContainerTrigger) defaultAzureName() {
	if trigger.Spec.AzureName == "" {
		trigger.Spec.AzureName = trigger.Name
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseContainerTrigger resource
func (trigger *SqlDatabaseContainerTrigger) defaultImpl() { trigger.defaultAzureName() }

var _ genruntime.KubernetesResource = &SqlDatabaseContainerTrigger{}

// AzureName returns the Azure name of the resource
func (trigger *SqlDatabaseContainerTrigger) AzureName() string {
	return trigger.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (trigger SqlDatabaseContainerTrigger) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (trigger *SqlDatabaseContainerTrigger) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (trigger *SqlDatabaseContainerTrigger) GetSpec() genruntime.ConvertibleSpec {
	return &trigger.Spec
}

// GetStatus returns the status of this resource
func (trigger *SqlDatabaseContainerTrigger) GetStatus() genruntime.ConvertibleStatus {
	return &trigger.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (trigger *SqlDatabaseContainerTrigger) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (trigger *SqlDatabaseContainerTrigger) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlTrigger_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (trigger *SqlDatabaseContainerTrigger) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(trigger.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  trigger.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (trigger *SqlDatabaseContainerTrigger) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlTrigger_Status); ok {
		trigger.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlTrigger_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	trigger.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainertrigger,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainertriggers,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabasecontainertriggers.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabaseContainerTrigger{}

// ValidateCreate validates the creation of the resource
func (trigger *SqlDatabaseContainerTrigger) ValidateCreate() error {
	validations := trigger.createValidations()
	var temp interface{} = trigger
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (trigger *SqlDatabaseContainerTrigger) ValidateDelete() error {
	validations := trigger.deleteValidations()
	var temp interface{} = trigger
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (trigger *SqlDatabaseContainerTrigger) ValidateUpdate(old runtime.Object) error {
	validations := trigger.updateValidations()
	var temp interface{} = trigger
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (trigger *SqlDatabaseContainerTrigger) createValidations() []func() error {
	return []func() error{trigger.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (trigger *SqlDatabaseContainerTrigger) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (trigger *SqlDatabaseContainerTrigger) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return trigger.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (trigger *SqlDatabaseContainerTrigger) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&trigger.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromSqlDatabaseContainerTrigger populates our SqlDatabaseContainerTrigger from the provided source SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) AssignPropertiesFromSqlDatabaseContainerTrigger(source *v1alpha1api20210515storage.SqlDatabaseContainerTrigger) error {

	// ObjectMeta
	trigger.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersTriggers_SPEC
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggers_SPEC(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggers_SPEC() to populate field Spec")
	}
	trigger.Spec = spec

	// Status
	var status SqlTrigger_Status
	err = status.AssignPropertiesFromSqlTrigger_Status(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSqlTrigger_Status() to populate field Status")
	}
	trigger.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerTrigger populates the provided destination SqlDatabaseContainerTrigger from our SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) AssignPropertiesToSqlDatabaseContainerTrigger(destination *v1alpha1api20210515storage.SqlDatabaseContainerTrigger) error {

	// ObjectMeta
	destination.ObjectMeta = *trigger.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersTriggers_SPEC
	err := trigger.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggers_SPEC(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggers_SPEC() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210515storage.SqlTrigger_Status
	err = trigger.Status.AssignPropertiesToSqlTrigger_Status(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlTrigger_Status() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (trigger *SqlDatabaseContainerTrigger) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: trigger.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseContainerTrigger",
	}
}

// +kubebuilder:object:root=true
//Generator information:
//- Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
//- ARM URI:
///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}
type SqlDatabaseContainerTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerTrigger `json:"items"`
}

type DatabaseAccountsSqlDatabasesContainersTriggers_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Options: A key-value pair of options to be applied for the request. This
	//corresponds to the headers sent with the request.
	Options *CreateUpdateOptions_Spec `json:"options,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	//Resource: The standard JSON format of a trigger
	Resource SqlTriggerResource_Spec `json:"resource"`
	Tags     map[string]string       `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabasesContainersTriggers_SPEC{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (spec *DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	var result DatabaseAccountsSqlDatabasesContainersTriggers_SPECARM

	// Set property ‘AzureName’:
	result.AzureName = spec.AzureName

	// Set property ‘Location’:
	if spec.Location != nil {
		location := *spec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if spec.Options != nil {
		optionsARM, err := (*spec.Options).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		options := optionsARM.(CreateUpdateOptions_SpecARM)
		result.Properties.Options = &options
	}
	resourceARM, err := spec.Resource.ConvertToARM(resolved)
	if err != nil {
		return nil, err
	}
	result.Properties.Resource = resourceARM.(SqlTriggerResource_SpecARM)

	// Set property ‘Tags’:
	if spec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range spec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (spec *DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabasesContainersTriggers_SPECARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (spec *DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabasesContainersTriggers_SPECARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabasesContainersTriggers_SPECARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	spec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		spec.Location = &location
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties.Options != nil {
		var options1 CreateUpdateOptions_Spec
		err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
		if err != nil {
			return err
		}
		options := options1
		spec.Options = &options
	}

	// Set property ‘Owner’:
	spec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	var resource SqlTriggerResource_Spec
	err := resource.PopulateFromARM(owner, typedInput.Properties.Resource)
	if err != nil {
		return err
	}
	spec.Resource = resource

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		spec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			spec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersTriggers_SPEC{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersTriggers_SPEC from the provided source
func (spec *DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersTriggers_SPEC)
	if ok {
		// Populate our instance from source
		return spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggers_SPEC(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersTriggers_SPEC{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggers_SPEC(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersTriggers_SPEC
func (spec *DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersTriggers_SPEC)
	if ok {
		// Populate destination from our instance
		return spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggers_SPEC(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersTriggers_SPEC{}
	err := spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggers_SPEC(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggers_SPEC populates our DatabaseAccountsSqlDatabasesContainersTriggers_SPEC from the provided source DatabaseAccountsSqlDatabasesContainersTriggers_SPEC
func (spec *DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggers_SPEC(source *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) error {

	// AzureName
	spec.AzureName = source.AzureName

	// Location
	spec.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions_Spec
		err := option.AssignPropertiesFromCreateUpdateOptions_Spec(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions_Spec() to populate field Options")
		}
		spec.Options = &option
	} else {
		spec.Options = nil
	}

	// Owner
	spec.Owner = source.Owner.Copy()

	// Resource
	if source.Resource != nil {
		var resource SqlTriggerResource_Spec
		err := resource.AssignPropertiesFromSqlTriggerResource_Spec(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlTriggerResource_Spec() to populate field Resource")
		}
		spec.Resource = resource
	} else {
		spec.Resource = SqlTriggerResource_Spec{}
	}

	// Tags
	spec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggers_SPEC populates the provided destination DatabaseAccountsSqlDatabasesContainersTriggers_SPEC from our DatabaseAccountsSqlDatabasesContainersTriggers_SPEC
func (spec *DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggers_SPEC(destination *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = spec.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(spec.Location)

	// Options
	if spec.Options != nil {
		var option v1alpha1api20210515storage.CreateUpdateOptions_Spec
		err := spec.Options.AssignPropertiesToCreateUpdateOptions_Spec(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions_Spec() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = spec.OriginalVersion()

	// Owner
	destination.Owner = spec.Owner.Copy()

	// Resource
	var resource v1alpha1api20210515storage.SqlTriggerResource_Spec
	err := spec.Resource.AssignPropertiesToSqlTriggerResource_Spec(&resource)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlTriggerResource_Spec() to populate field Resource")
	}
	destination.Resource = &resource

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(spec.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (spec *DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (spec *DatabaseAccountsSqlDatabasesContainersTriggers_SPEC) SetAzureName(azureName string) {
	spec.AzureName = azureName
}

type SqlTrigger_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	//Options: A key-value pair of options to be applied for the request. This
	//corresponds to the headers sent with the request.
	Options *CreateUpdateOptions_Status `json:"options,omitempty"`

	//Resource: The standard JSON format of a trigger
	Resource *SqlTriggerResource_Status `json:"resource,omitempty"`
	Tags     map[string]string          `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlTrigger_Status{}

// ConvertStatusFrom populates our SqlTrigger_Status from the provided source
func (trigger *SqlTrigger_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210515storage.SqlTrigger_Status)
	if ok {
		// Populate our instance from source
		return trigger.AssignPropertiesFromSqlTrigger_Status(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.SqlTrigger_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = trigger.AssignPropertiesFromSqlTrigger_Status(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlTrigger_Status
func (trigger *SqlTrigger_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210515storage.SqlTrigger_Status)
	if ok {
		// Populate destination from our instance
		return trigger.AssignPropertiesToSqlTrigger_Status(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.SqlTrigger_Status{}
	err := trigger.AssignPropertiesToSqlTrigger_Status(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &SqlTrigger_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (trigger *SqlTrigger_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlTrigger_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (trigger *SqlTrigger_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlTrigger_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlTrigger_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		trigger.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		trigger.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		trigger.Name = &name
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Options != nil {
			var options1 CreateUpdateOptions_Status
			err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
			if err != nil {
				return err
			}
			options := options1
			trigger.Options = &options
		}
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		var temp SqlTriggerResource_Status
		var temp1 SqlTriggerResource_Status
		err := temp1.PopulateFromARM(owner, typedInput.Properties.Resource)
		if err != nil {
			return err
		}
		temp = temp1
		trigger.Resource = &temp
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		trigger.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			trigger.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		trigger.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlTrigger_Status populates our SqlTrigger_Status from the provided source SqlTrigger_Status
func (trigger *SqlTrigger_Status) AssignPropertiesFromSqlTrigger_Status(source *v1alpha1api20210515storage.SqlTrigger_Status) error {

	// Conditions
	trigger.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	trigger.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	trigger.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	trigger.Name = genruntime.ClonePointerToString(source.Name)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions_Status
		err := option.AssignPropertiesFromCreateUpdateOptions_Status(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions_Status() to populate field Options")
		}
		trigger.Options = &option
	} else {
		trigger.Options = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlTriggerResource_Status
		err := resource.AssignPropertiesFromSqlTriggerResource_Status(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlTriggerResource_Status() to populate field Resource")
		}
		trigger.Resource = &resource
	} else {
		trigger.Resource = nil
	}

	// Tags
	trigger.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	trigger.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToSqlTrigger_Status populates the provided destination SqlTrigger_Status from our SqlTrigger_Status
func (trigger *SqlTrigger_Status) AssignPropertiesToSqlTrigger_Status(destination *v1alpha1api20210515storage.SqlTrigger_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(trigger.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(trigger.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(trigger.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(trigger.Name)

	// Options
	if trigger.Options != nil {
		var option v1alpha1api20210515storage.CreateUpdateOptions_Status
		err := trigger.Options.AssignPropertiesToCreateUpdateOptions_Status(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions_Status() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// Resource
	if trigger.Resource != nil {
		var resource v1alpha1api20210515storage.SqlTriggerResource_Status
		err := trigger.Resource.AssignPropertiesToSqlTriggerResource_Status(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlTriggerResource_Status() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(trigger.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(trigger.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type SqlTriggerResource_Spec struct {
	//Body: Body of the Trigger
	Body *string `json:"body,omitempty"`

	// +kubebuilder:validation:Required
	//Id: Name of the Cosmos DB SQL trigger
	Id string `json:"id"`

	//TriggerOperation: The operation the trigger is associated with
	TriggerOperation *SqlTriggerResource_TriggerOperation_Spec `json:"triggerOperation,omitempty"`

	//TriggerType: Type of the Trigger
	TriggerType *SqlTriggerResource_TriggerType_Spec `json:"triggerType,omitempty"`
}

var _ genruntime.ARMTransformer = &SqlTriggerResource_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (resource *SqlTriggerResource_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if resource == nil {
		return nil, nil
	}
	var result SqlTriggerResource_SpecARM

	// Set property ‘Body’:
	if resource.Body != nil {
		body := *resource.Body
		result.Body = &body
	}

	// Set property ‘Id’:
	result.Id = resource.Id

	// Set property ‘TriggerOperation’:
	if resource.TriggerOperation != nil {
		triggerOperation := *resource.TriggerOperation
		result.TriggerOperation = &triggerOperation
	}

	// Set property ‘TriggerType’:
	if resource.TriggerType != nil {
		triggerType := *resource.TriggerType
		result.TriggerType = &triggerType
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlTriggerResource_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlTriggerResource_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlTriggerResource_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlTriggerResource_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlTriggerResource_SpecARM, got %T", armInput)
	}

	// Set property ‘Body’:
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
	}

	// Set property ‘Id’:
	resource.Id = typedInput.Id

	// Set property ‘TriggerOperation’:
	if typedInput.TriggerOperation != nil {
		triggerOperation := *typedInput.TriggerOperation
		resource.TriggerOperation = &triggerOperation
	}

	// Set property ‘TriggerType’:
	if typedInput.TriggerType != nil {
		triggerType := *typedInput.TriggerType
		resource.TriggerType = &triggerType
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlTriggerResource_Spec populates our SqlTriggerResource_Spec from the provided source SqlTriggerResource_Spec
func (resource *SqlTriggerResource_Spec) AssignPropertiesFromSqlTriggerResource_Spec(source *v1alpha1api20210515storage.SqlTriggerResource_Spec) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.GetOptionalStringValue(source.Id)

	// TriggerOperation
	if source.TriggerOperation != nil {
		triggerOperation := SqlTriggerResource_TriggerOperation_Spec(*source.TriggerOperation)
		resource.TriggerOperation = &triggerOperation
	} else {
		resource.TriggerOperation = nil
	}

	// TriggerType
	if source.TriggerType != nil {
		triggerType := SqlTriggerResource_TriggerType_Spec(*source.TriggerType)
		resource.TriggerType = &triggerType
	} else {
		resource.TriggerType = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlTriggerResource_Spec populates the provided destination SqlTriggerResource_Spec from our SqlTriggerResource_Spec
func (resource *SqlTriggerResource_Spec) AssignPropertiesToSqlTriggerResource_Spec(destination *v1alpha1api20210515storage.SqlTriggerResource_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Id
	id := resource.Id
	destination.Id = &id

	// TriggerOperation
	if resource.TriggerOperation != nil {
		triggerOperation := string(*resource.TriggerOperation)
		destination.TriggerOperation = &triggerOperation
	} else {
		destination.TriggerOperation = nil
	}

	// TriggerType
	if resource.TriggerType != nil {
		triggerType := string(*resource.TriggerType)
		destination.TriggerType = &triggerType
	} else {
		destination.TriggerType = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type SqlTriggerResource_Status struct {
	//Body: Body of the Trigger
	Body *string `json:"body,omitempty"`

	// +kubebuilder:validation:Required
	//Id: Name of the Cosmos DB SQL trigger
	Id string `json:"id"`

	//TriggerOperation: The operation the trigger is associated with
	TriggerOperation *SqlTriggerResource_TriggerOperation_Status `json:"triggerOperation,omitempty"`

	//TriggerType: Type of the Trigger
	TriggerType *SqlTriggerResource_TriggerType_Status `json:"triggerType,omitempty"`
}

var _ genruntime.FromARMConverter = &SqlTriggerResource_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlTriggerResource_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlTriggerResource_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlTriggerResource_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlTriggerResource_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlTriggerResource_StatusARM, got %T", armInput)
	}

	// Set property ‘Body’:
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
	}

	// Set property ‘Id’:
	resource.Id = typedInput.Id

	// Set property ‘TriggerOperation’:
	if typedInput.TriggerOperation != nil {
		triggerOperation := *typedInput.TriggerOperation
		resource.TriggerOperation = &triggerOperation
	}

	// Set property ‘TriggerType’:
	if typedInput.TriggerType != nil {
		triggerType := *typedInput.TriggerType
		resource.TriggerType = &triggerType
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlTriggerResource_Status populates our SqlTriggerResource_Status from the provided source SqlTriggerResource_Status
func (resource *SqlTriggerResource_Status) AssignPropertiesFromSqlTriggerResource_Status(source *v1alpha1api20210515storage.SqlTriggerResource_Status) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.GetOptionalStringValue(source.Id)

	// TriggerOperation
	if source.TriggerOperation != nil {
		triggerOperation := SqlTriggerResource_TriggerOperation_Status(*source.TriggerOperation)
		resource.TriggerOperation = &triggerOperation
	} else {
		resource.TriggerOperation = nil
	}

	// TriggerType
	if source.TriggerType != nil {
		triggerType := SqlTriggerResource_TriggerType_Status(*source.TriggerType)
		resource.TriggerType = &triggerType
	} else {
		resource.TriggerType = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlTriggerResource_Status populates the provided destination SqlTriggerResource_Status from our SqlTriggerResource_Status
func (resource *SqlTriggerResource_Status) AssignPropertiesToSqlTriggerResource_Status(destination *v1alpha1api20210515storage.SqlTriggerResource_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Id
	id := resource.Id
	destination.Id = &id

	// TriggerOperation
	if resource.TriggerOperation != nil {
		triggerOperation := string(*resource.TriggerOperation)
		destination.TriggerOperation = &triggerOperation
	} else {
		destination.TriggerOperation = nil
	}

	// TriggerType
	if resource.TriggerType != nil {
		triggerType := string(*resource.TriggerType)
		destination.TriggerType = &triggerType
	} else {
		destination.TriggerType = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type SqlTriggerResource_TriggerOperation_Status string

const (
	SqlTriggerResource_TriggerOperation_StatusAll     = SqlTriggerResource_TriggerOperation_Status("All")
	SqlTriggerResource_TriggerOperation_StatusCreate  = SqlTriggerResource_TriggerOperation_Status("Create")
	SqlTriggerResource_TriggerOperation_StatusDelete  = SqlTriggerResource_TriggerOperation_Status("Delete")
	SqlTriggerResource_TriggerOperation_StatusReplace = SqlTriggerResource_TriggerOperation_Status("Replace")
	SqlTriggerResource_TriggerOperation_StatusUpdate  = SqlTriggerResource_TriggerOperation_Status("Update")
)

type SqlTriggerResource_TriggerType_Status string

const (
	SqlTriggerResource_TriggerType_StatusPost = SqlTriggerResource_TriggerType_Status("Post")
	SqlTriggerResource_TriggerType_StatusPre  = SqlTriggerResource_TriggerType_Status("Pre")
)

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerTrigger{}, &SqlDatabaseContainerTriggerList{})
}
