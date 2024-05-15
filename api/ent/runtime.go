// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/codelite7/momentum/api/ent/agent"
	"github.com/codelite7/momentum/api/ent/bookmark"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/response"
	"github.com/codelite7/momentum/api/ent/schema"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	agentMixin := schema.Agent{}.Mixin()
	agentMixinFields0 := agentMixin[0].Fields()
	_ = agentMixinFields0
	agentFields := schema.Agent{}.Fields()
	_ = agentFields
	// agentDescCreatedAt is the schema descriptor for created_at field.
	agentDescCreatedAt := agentMixinFields0[1].Descriptor()
	// agent.DefaultCreatedAt holds the default value on creation for the created_at field.
	agent.DefaultCreatedAt = agentDescCreatedAt.Default.(func() time.Time)
	// agentDescUpdatedAt is the schema descriptor for updated_at field.
	agentDescUpdatedAt := agentMixinFields0[2].Descriptor()
	// agent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	agent.DefaultUpdatedAt = agentDescUpdatedAt.Default.(func() time.Time)
	// agentDescID is the schema descriptor for id field.
	agentDescID := agentMixinFields0[0].Descriptor()
	// agent.DefaultID holds the default value on creation for the id field.
	agent.DefaultID = agentDescID.Default.(func() uuid.UUID)
	bookmarkMixin := schema.Bookmark{}.Mixin()
	bookmarkMixinFields0 := bookmarkMixin[0].Fields()
	_ = bookmarkMixinFields0
	bookmarkFields := schema.Bookmark{}.Fields()
	_ = bookmarkFields
	// bookmarkDescCreatedAt is the schema descriptor for created_at field.
	bookmarkDescCreatedAt := bookmarkMixinFields0[1].Descriptor()
	// bookmark.DefaultCreatedAt holds the default value on creation for the created_at field.
	bookmark.DefaultCreatedAt = bookmarkDescCreatedAt.Default.(func() time.Time)
	// bookmarkDescUpdatedAt is the schema descriptor for updated_at field.
	bookmarkDescUpdatedAt := bookmarkMixinFields0[2].Descriptor()
	// bookmark.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	bookmark.DefaultUpdatedAt = bookmarkDescUpdatedAt.Default.(func() time.Time)
	// bookmarkDescID is the schema descriptor for id field.
	bookmarkDescID := bookmarkMixinFields0[0].Descriptor()
	// bookmark.DefaultID holds the default value on creation for the id field.
	bookmark.DefaultID = bookmarkDescID.Default.(func() uuid.UUID)
	messageMixin := schema.Message{}.Mixin()
	messageMixinFields0 := messageMixin[0].Fields()
	_ = messageMixinFields0
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageMixinFields0[1].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	// messageDescUpdatedAt is the schema descriptor for updated_at field.
	messageDescUpdatedAt := messageMixinFields0[2].Descriptor()
	// message.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	message.DefaultUpdatedAt = messageDescUpdatedAt.Default.(func() time.Time)
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageMixinFields0[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() uuid.UUID)
	responseMixin := schema.Response{}.Mixin()
	responseMixinFields0 := responseMixin[0].Fields()
	_ = responseMixinFields0
	responseFields := schema.Response{}.Fields()
	_ = responseFields
	// responseDescCreatedAt is the schema descriptor for created_at field.
	responseDescCreatedAt := responseMixinFields0[1].Descriptor()
	// response.DefaultCreatedAt holds the default value on creation for the created_at field.
	response.DefaultCreatedAt = responseDescCreatedAt.Default.(func() time.Time)
	// responseDescUpdatedAt is the schema descriptor for updated_at field.
	responseDescUpdatedAt := responseMixinFields0[2].Descriptor()
	// response.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	response.DefaultUpdatedAt = responseDescUpdatedAt.Default.(func() time.Time)
	// responseDescID is the schema descriptor for id field.
	responseDescID := responseMixinFields0[0].Descriptor()
	// response.DefaultID holds the default value on creation for the id field.
	response.DefaultID = responseDescID.Default.(func() uuid.UUID)
	threadMixin := schema.Thread{}.Mixin()
	threadMixinFields0 := threadMixin[0].Fields()
	_ = threadMixinFields0
	threadFields := schema.Thread{}.Fields()
	_ = threadFields
	// threadDescCreatedAt is the schema descriptor for created_at field.
	threadDescCreatedAt := threadMixinFields0[1].Descriptor()
	// thread.DefaultCreatedAt holds the default value on creation for the created_at field.
	thread.DefaultCreatedAt = threadDescCreatedAt.Default.(func() time.Time)
	// threadDescUpdatedAt is the schema descriptor for updated_at field.
	threadDescUpdatedAt := threadMixinFields0[2].Descriptor()
	// thread.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	thread.DefaultUpdatedAt = threadDescUpdatedAt.Default.(func() time.Time)
	// threadDescID is the schema descriptor for id field.
	threadDescID := threadMixinFields0[0].Descriptor()
	// thread.DefaultID holds the default value on creation for the id field.
	thread.DefaultID = threadDescID.Default.(func() uuid.UUID)
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
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
