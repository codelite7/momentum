// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AgentsColumns holds the columns for the "agents" table.
	AgentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "model", Type: field.TypeString},
	}
	// AgentsTable holds the schema information for the "agents" table.
	AgentsTable = &schema.Table{
		Name:       "agents",
		Columns:    AgentsColumns,
		PrimaryKey: []*schema.Column{AgentsColumns[0]},
	}
	// BookmarksColumns holds the columns for the "bookmarks" table.
	BookmarksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "message_bookmarks", Type: field.TypeUUID, Nullable: true},
		{Name: "thread_bookmarks", Type: field.TypeUUID, Nullable: true},
		{Name: "user_bookmarks", Type: field.TypeUUID},
	}
	// BookmarksTable holds the schema information for the "bookmarks" table.
	BookmarksTable = &schema.Table{
		Name:       "bookmarks",
		Columns:    BookmarksColumns,
		PrimaryKey: []*schema.Column{BookmarksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bookmarks_messages_bookmarks",
				Columns:    []*schema.Column{BookmarksColumns[3]},
				RefColumns: []*schema.Column{MessagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bookmarks_threads_bookmarks",
				Columns:    []*schema.Column{BookmarksColumns[4]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bookmarks_users_bookmarks",
				Columns:    []*schema.Column{BookmarksColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString},
		{Name: "agent_messages", Type: field.TypeUUID, Nullable: true},
		{Name: "thread_messages", Type: field.TypeUUID},
		{Name: "user_messages", Type: field.TypeUUID, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_agents_messages",
				Columns:    []*schema.Column{MessagesColumns[4]},
				RefColumns: []*schema.Column{AgentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "messages_threads_messages",
				Columns:    []*schema.Column{MessagesColumns[5]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "messages_users_messages",
				Columns:    []*schema.Column{MessagesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ThreadsColumns holds the columns for the "threads" table.
	ThreadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "thread_children", Type: field.TypeUUID, Nullable: true},
		{Name: "user_threads", Type: field.TypeUUID},
	}
	// ThreadsTable holds the schema information for the "threads" table.
	ThreadsTable = &schema.Table{
		Name:       "threads",
		Columns:    ThreadsColumns,
		PrimaryKey: []*schema.Column{ThreadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "threads_threads_children",
				Columns:    []*schema.Column{ThreadsColumns[4]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "threads_users_threads",
				Columns:    []*schema.Column{ThreadsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgentsTable,
		BookmarksTable,
		MessagesTable,
		ThreadsTable,
		UsersTable,
	}
)

func init() {
	BookmarksTable.ForeignKeys[0].RefTable = MessagesTable
	BookmarksTable.ForeignKeys[1].RefTable = ThreadsTable
	BookmarksTable.ForeignKeys[2].RefTable = UsersTable
	MessagesTable.ForeignKeys[0].RefTable = AgentsTable
	MessagesTable.ForeignKeys[1].RefTable = ThreadsTable
	MessagesTable.ForeignKeys[2].RefTable = UsersTable
	ThreadsTable.ForeignKeys[0].RefTable = ThreadsTable
	ThreadsTable.ForeignKeys[1].RefTable = UsersTable
}
