// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/codelite7/momentum/api/ent/agent"
	"github.com/codelite7/momentum/api/ent/bookmark"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/response"
	"github.com/codelite7/momentum/api/ent/schema"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/codelite7/momentum/api/ent/tenant"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	agentMixin := schema.Agent{}.Mixin()
	agentMixinFields0 := agentMixin[0].Fields()
	_ = agentMixinFields0
	agentMixinFields1 := agentMixin[1].Fields()
	_ = agentMixinFields1
	agentFields := schema.Agent{}.Fields()
	_ = agentFields
	// agentDescCreatedAt is the schema descriptor for created_at field.
	agentDescCreatedAt := agentMixinFields0[0].Descriptor()
	// agent.DefaultCreatedAt holds the default value on creation for the created_at field.
	agent.DefaultCreatedAt = agentDescCreatedAt.Default.(func() time.Time)
	// agentDescUpdatedAt is the schema descriptor for updated_at field.
	agentDescUpdatedAt := agentMixinFields0[1].Descriptor()
	// agent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	agent.DefaultUpdatedAt = agentDescUpdatedAt.Default.(func() time.Time)
	// agentDescID is the schema descriptor for id field.
	agentDescID := agentMixinFields1[0].Descriptor()
	// agent.DefaultID holds the default value on creation for the id field.
	agent.DefaultID = agentDescID.Default.(func() pulid.ID)
	bookmarkMixin := schema.Bookmark{}.Mixin()
	bookmarkMixinFields0 := bookmarkMixin[0].Fields()
	_ = bookmarkMixinFields0
	bookmarkMixinFields1 := bookmarkMixin[1].Fields()
	_ = bookmarkMixinFields1
	bookmarkFields := schema.Bookmark{}.Fields()
	_ = bookmarkFields
	// bookmarkDescCreatedAt is the schema descriptor for created_at field.
	bookmarkDescCreatedAt := bookmarkMixinFields0[0].Descriptor()
	// bookmark.DefaultCreatedAt holds the default value on creation for the created_at field.
	bookmark.DefaultCreatedAt = bookmarkDescCreatedAt.Default.(func() time.Time)
	// bookmarkDescUpdatedAt is the schema descriptor for updated_at field.
	bookmarkDescUpdatedAt := bookmarkMixinFields0[1].Descriptor()
	// bookmark.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	bookmark.DefaultUpdatedAt = bookmarkDescUpdatedAt.Default.(func() time.Time)
	// bookmarkDescID is the schema descriptor for id field.
	bookmarkDescID := bookmarkMixinFields1[0].Descriptor()
	// bookmark.DefaultID holds the default value on creation for the id field.
	bookmark.DefaultID = bookmarkDescID.Default.(func() pulid.ID)
	messageMixin := schema.Message{}.Mixin()
	messageMixinFields0 := messageMixin[0].Fields()
	_ = messageMixinFields0
	messageMixinFields1 := messageMixin[1].Fields()
	_ = messageMixinFields1
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageMixinFields0[0].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	// messageDescUpdatedAt is the schema descriptor for updated_at field.
	messageDescUpdatedAt := messageMixinFields0[1].Descriptor()
	// message.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	message.DefaultUpdatedAt = messageDescUpdatedAt.Default.(func() time.Time)
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageMixinFields1[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() pulid.ID)
	responseMixin := schema.Response{}.Mixin()
	responseMixinFields0 := responseMixin[0].Fields()
	_ = responseMixinFields0
	responseMixinFields1 := responseMixin[1].Fields()
	_ = responseMixinFields1
	responseFields := schema.Response{}.Fields()
	_ = responseFields
	// responseDescCreatedAt is the schema descriptor for created_at field.
	responseDescCreatedAt := responseMixinFields0[0].Descriptor()
	// response.DefaultCreatedAt holds the default value on creation for the created_at field.
	response.DefaultCreatedAt = responseDescCreatedAt.Default.(func() time.Time)
	// responseDescUpdatedAt is the schema descriptor for updated_at field.
	responseDescUpdatedAt := responseMixinFields0[1].Descriptor()
	// response.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	response.DefaultUpdatedAt = responseDescUpdatedAt.Default.(func() time.Time)
	// responseDescID is the schema descriptor for id field.
	responseDescID := responseMixinFields1[0].Descriptor()
	// response.DefaultID holds the default value on creation for the id field.
	response.DefaultID = responseDescID.Default.(func() pulid.ID)
	tenantMixin := schema.Tenant{}.Mixin()
	tenantMixinFields0 := tenantMixin[0].Fields()
	_ = tenantMixinFields0
	tenantFields := schema.Tenant{}.Fields()
	_ = tenantFields
	// tenantDescCreatedAt is the schema descriptor for created_at field.
	tenantDescCreatedAt := tenantFields[0].Descriptor()
	// tenant.DefaultCreatedAt holds the default value on creation for the created_at field.
	tenant.DefaultCreatedAt = tenantDescCreatedAt.Default.(func() time.Time)
	// tenantDescUpdatedAt is the schema descriptor for updated_at field.
	tenantDescUpdatedAt := tenantFields[1].Descriptor()
	// tenant.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tenant.DefaultUpdatedAt = tenantDescUpdatedAt.Default.(func() time.Time)
	// tenantDescID is the schema descriptor for id field.
	tenantDescID := tenantMixinFields0[0].Descriptor()
	// tenant.DefaultID holds the default value on creation for the id field.
	tenant.DefaultID = tenantDescID.Default.(func() pulid.ID)
	threadMixin := schema.Thread{}.Mixin()
	threadMixinFields0 := threadMixin[0].Fields()
	_ = threadMixinFields0
	threadMixinFields1 := threadMixin[1].Fields()
	_ = threadMixinFields1
	threadFields := schema.Thread{}.Fields()
	_ = threadFields
	// threadDescCreatedAt is the schema descriptor for created_at field.
	threadDescCreatedAt := threadMixinFields0[0].Descriptor()
	// thread.DefaultCreatedAt holds the default value on creation for the created_at field.
	thread.DefaultCreatedAt = threadDescCreatedAt.Default.(func() time.Time)
	// threadDescUpdatedAt is the schema descriptor for updated_at field.
	threadDescUpdatedAt := threadMixinFields0[1].Descriptor()
	// thread.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	thread.DefaultUpdatedAt = threadDescUpdatedAt.Default.(func() time.Time)
	// threadDescID is the schema descriptor for id field.
	threadDescID := threadMixinFields1[0].Descriptor()
	// thread.DefaultID holds the default value on creation for the id field.
	thread.DefaultID = threadDescID.Default.(func() pulid.ID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields1[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() pulid.ID)
}
