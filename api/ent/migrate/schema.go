// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AgentsColumns holds the columns for the "agents" table.
	AgentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "provider", Type: field.TypeString},
		{Name: "model", Type: field.TypeString},
		{Name: "api_key", Type: field.TypeString},
	}
	// AgentsTable holds the schema information for the "agents" table.
	AgentsTable = &schema.Table{
		Name:       "agents",
		Columns:    AgentsColumns,
		PrimaryKey: []*schema.Column{AgentsColumns[0]},
	}
	// BookmarksColumns holds the columns for the "bookmarks" table.
	BookmarksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "tenant_id", Type: field.TypeString},
		{Name: "message_bookmarks", Type: field.TypeString, Nullable: true},
		{Name: "response_bookmarks", Type: field.TypeString, Nullable: true},
		{Name: "thread_bookmarks", Type: field.TypeString, Nullable: true},
		{Name: "user_bookmarks", Type: field.TypeString},
	}
	// BookmarksTable holds the schema information for the "bookmarks" table.
	BookmarksTable = &schema.Table{
		Name:       "bookmarks",
		Columns:    BookmarksColumns,
		PrimaryKey: []*schema.Column{BookmarksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bookmarks_tenants_tenant",
				Columns:    []*schema.Column{BookmarksColumns[3]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "bookmarks_messages_bookmarks",
				Columns:    []*schema.Column{BookmarksColumns[4]},
				RefColumns: []*schema.Column{MessagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "bookmarks_responses_bookmarks",
				Columns:    []*schema.Column{BookmarksColumns[5]},
				RefColumns: []*schema.Column{ResponsesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bookmarks_threads_bookmarks",
				Columns:    []*schema.Column{BookmarksColumns[6]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "bookmarks_users_bookmarks",
				Columns:    []*schema.Column{BookmarksColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString},
		{Name: "tenant_id", Type: field.TypeString},
		{Name: "thread_messages", Type: field.TypeString},
		{Name: "user_messages", Type: field.TypeString},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_tenants_tenant",
				Columns:    []*schema.Column{MessagesColumns[4]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "messages_threads_messages",
				Columns:    []*schema.Column{MessagesColumns[5]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "messages_users_messages",
				Columns:    []*schema.Column{MessagesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ResponsesColumns holds the columns for the "responses" table.
	ResponsesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString, Nullable: true},
		{Name: "agent_responses", Type: field.TypeString},
		{Name: "message_response", Type: field.TypeString, Unique: true},
		{Name: "tenant_id", Type: field.TypeString},
	}
	// ResponsesTable holds the schema information for the "responses" table.
	ResponsesTable = &schema.Table{
		Name:       "responses",
		Columns:    ResponsesColumns,
		PrimaryKey: []*schema.Column{ResponsesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "responses_agents_responses",
				Columns:    []*schema.Column{ResponsesColumns[4]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "responses_messages_response",
				Columns:    []*schema.Column{ResponsesColumns[5]},
				RefColumns: []*schema.Column{MessagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "responses_tenants_tenant",
				Columns:    []*schema.Column{ResponsesColumns[6]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TenantsColumns holds the columns for the "tenants" table.
	TenantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TenantsTable holds the schema information for the "tenants" table.
	TenantsTable = &schema.Table{
		Name:       "tenants",
		Columns:    TenantsColumns,
		PrimaryKey: []*schema.Column{TenantsColumns[0]},
	}
	// ThreadsColumns holds the columns for the "threads" table.
	ThreadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "tenant_id", Type: field.TypeString},
		{Name: "thread_children", Type: field.TypeString, Nullable: true},
		{Name: "user_threads", Type: field.TypeString},
	}
	// ThreadsTable holds the schema information for the "threads" table.
	ThreadsTable = &schema.Table{
		Name:       "threads",
		Columns:    ThreadsColumns,
		PrimaryKey: []*schema.Column{ThreadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "threads_tenants_tenant",
				Columns:    []*schema.Column{ThreadsColumns[4]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "threads_threads_children",
				Columns:    []*schema.Column{ThreadsColumns[5]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "threads_users_threads",
				Columns:    []*schema.Column{ThreadsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "tenant_id", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_tenants_tenant",
				Columns:    []*schema.Column{UsersColumns[4]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// WorkosEventCursorsColumns holds the columns for the "workos_event_cursors" table.
	WorkosEventCursorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_created_cursor", Type: field.TypeString, Nullable: true},
	}
	// WorkosEventCursorsTable holds the schema information for the "workos_event_cursors" table.
	WorkosEventCursorsTable = &schema.Table{
		Name:       "workos_event_cursors",
		Columns:    WorkosEventCursorsColumns,
		PrimaryKey: []*schema.Column{WorkosEventCursorsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgentsTable,
		BookmarksTable,
		MessagesTable,
		ResponsesTable,
		TenantsTable,
		ThreadsTable,
		UsersTable,
		WorkosEventCursorsTable,
	}
)

func init() {
	BookmarksTable.ForeignKeys[0].RefTable = TenantsTable
	BookmarksTable.ForeignKeys[1].RefTable = MessagesTable
	BookmarksTable.ForeignKeys[2].RefTable = ResponsesTable
	BookmarksTable.ForeignKeys[3].RefTable = ThreadsTable
	BookmarksTable.ForeignKeys[4].RefTable = UsersTable
	MessagesTable.ForeignKeys[0].RefTable = TenantsTable
	MessagesTable.ForeignKeys[1].RefTable = ThreadsTable
	MessagesTable.ForeignKeys[2].RefTable = UsersTable
	ResponsesTable.ForeignKeys[0].RefTable = AgentsTable
	ResponsesTable.ForeignKeys[1].RefTable = MessagesTable
	ResponsesTable.ForeignKeys[2].RefTable = TenantsTable
	ThreadsTable.ForeignKeys[0].RefTable = TenantsTable
	ThreadsTable.ForeignKeys[1].RefTable = ThreadsTable
	ThreadsTable.ForeignKeys[2].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = TenantsTable
}
