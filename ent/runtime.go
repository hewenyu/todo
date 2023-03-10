// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/hewenyu/todo/ent/schema"
	"github.com/hewenyu/todo/ent/temple"
	"github.com/hewenyu/todo/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	templeMixin := schema.Temple{}.Mixin()
	templeMixinFields0 := templeMixin[0].Fields()
	_ = templeMixinFields0
	templeFields := schema.Temple{}.Fields()
	_ = templeFields
	// templeDescCreatedAt is the schema descriptor for created_at field.
	templeDescCreatedAt := templeMixinFields0[1].Descriptor()
	// temple.DefaultCreatedAt holds the default value on creation for the created_at field.
	temple.DefaultCreatedAt = templeDescCreatedAt.Default.(func() time.Time)
	// templeDescUpdatedAt is the schema descriptor for updated_at field.
	templeDescUpdatedAt := templeMixinFields0[2].Descriptor()
	// temple.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	temple.DefaultUpdatedAt = templeDescUpdatedAt.Default.(func() time.Time)
	// temple.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	temple.UpdateDefaultUpdatedAt = templeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// templeDescText is the schema descriptor for text field.
	templeDescText := templeFields[0].Descriptor()
	// temple.TextValidator is a validator for the "text" field. It is called by the builders before save.
	temple.TextValidator = templeDescText.Validators[0].(func(string) error)
	// templeDescPriority is the schema descriptor for priority field.
	templeDescPriority := templeFields[1].Descriptor()
	// temple.DefaultPriority holds the default value on creation for the priority field.
	temple.DefaultPriority = templeDescPriority.Default.(int)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
