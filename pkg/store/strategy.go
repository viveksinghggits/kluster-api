package store

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
)

type KlusterStrategy struct {
	runtime.ObjectTyper
	// The name generator is used when the standard GenerateName field is set.
	// The NameGenerator will be invoked prior to validation.
	names.NameGenerator
}

func NewKlusterStrategy(scheme runtime.ObjectTyper) KlusterStrategy {
	return KlusterStrategy{
		ObjectTyper:   scheme,
		NameGenerator: names.SimpleNameGenerator,
	}
}

// NamespaceScoped returns true if the object must be within a namespace.
func (k KlusterStrategy) NamespaceScoped() bool {
	return true
}

// PrepareForCreate is invoked on create before validation to normalize
// the object.  For example: remove fields that are not to be persisted,
// sort order-insensitive list fields, etc.  This should not remove fields
// whose presence would be considered a validation error.
//
// Often implemented as a type check and an initailization or clearing of
// status. Clear the status because status changes are internal. External
// callers of an api (users) should not be setting an initial status on
// newly created objects.
func (k KlusterStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {

}

// Validate returns an ErrorList with validation errors or nil.  Validate
// is invoked after default fields in the object have been filled in
// before the object is persisted.  This method should not mutate the
// object.
func (k KlusterStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return nil
}

// WarningsOnCreate returns warnings to the client performing a create.
// WarningsOnCreate is invoked after default fields in the object have been filled in
// and after Validate has passed, before Canonicalize is called, and the object is persisted.
// This method must not mutate the object.
//
// Be brief; limit warnings to 120 characters if possible.
// Don't include a "Warning:" prefix in the message (that is added by clients on output).
// Warnings returned about a specific field should be formatted as "path.to.field: message".
// For example: `spec.imagePullSecrets[0].name: invalid empty name ""`
//
// Use warning messages to describe problems the client making the API request should correct or be aware of.
// For example:
//   - use of deprecated fields/labels/annotations that will stop working in a future release
//   - use of obsolete fields/labels/annotations that are non-functional
//   - malformed or invalid specifications that prevent successful handling of the submitted object,
//     but are not rejected by validation for compatibility reasons
//
// Warnings should not be returned for fields which cannot be resolved by the caller.
// For example, do not warn about spec fields in a subresource creation request.
func (k KlusterStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

// Canonicalize allows an object to be mutated into a canonical form. This
// ensures that code that operates on these objects can rely on the common
// form for things like comparison.  Canonicalize is invoked after
// validation has succeeded but before the object has been persisted.
// This method may mutate the object. Often implemented as a type check or
// empty method.
func (k KlusterStrategy) Canonicalize(obj runtime.Object) {

}

func (k KlusterStrategy) AllowCreateOnUpdate() bool {
	return false
}

// PrepareForUpdate is invoked on update before validation to normalize
// the object.  For example: remove fields that are not to be persisted,
// sort order-insensitive list fields, etc.  This should not remove fields
// whose presence would be considered a validation error.
func (k KlusterStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {

}

// ValidateUpdate is invoked after default fields in the object have been
// filled in before the object is persisted.  This method should not mutate
// the object.
func (k KlusterStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return nil
}

// WarningsOnUpdate returns warnings to the client performing the update.
// WarningsOnUpdate is invoked after default fields in the object have been filled in
// and after ValidateUpdate has passed, before Canonicalize is called, and before the object is persisted.
// This method must not mutate either object.
//
// Be brief; limit warnings to 120 characters if possible.
// Don't include a "Warning:" prefix in the message (that is added by clients on output).
// Warnings returned about a specific field should be formatted as "path.to.field: message".
// For example: `spec.imagePullSecrets[0].name: invalid empty name ""`
//
// Use warning messages to describe problems the client making the API request should correct or be aware of.
// For example:
//   - use of deprecated fields/labels/annotations that will stop working in a future release
//   - use of obsolete fields/labels/annotations that are non-functional
//   - malformed or invalid specifications that prevent successful handling of the submitted object,
//     but are not rejected by validation for compatibility reasons
//
// Warnings should not be returned for fields which cannot be resolved by the caller.
// For example, do not warn about spec fields in a status update.
func (k KlusterStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

// AllowUnconditionalUpdate returns true if the object can be updated
// unconditionally (irrespective of the latest resource version), when
// there is no resource version specified in the object.
func (k KlusterStrategy) AllowUnconditionalUpdate() bool {
	return false
}
