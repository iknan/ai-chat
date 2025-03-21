// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		Conversation:       newConversation(db, opts...),
		ConversationMember: newConversationMember(db, opts...),
		Friendship:         newFriendship(db, opts...),
		Message:            newMessage(db, opts...),
		MessageStatus:      newMessageStatus(db, opts...),
		User:               newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Conversation       conversation
	ConversationMember conversationMember
	Friendship         friendship
	Message            message
	MessageStatus      messageStatus
	User               user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		Conversation:       q.Conversation.clone(db),
		ConversationMember: q.ConversationMember.clone(db),
		Friendship:         q.Friendship.clone(db),
		Message:            q.Message.clone(db),
		MessageStatus:      q.MessageStatus.clone(db),
		User:               q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		Conversation:       q.Conversation.replaceDB(db),
		ConversationMember: q.ConversationMember.replaceDB(db),
		Friendship:         q.Friendship.replaceDB(db),
		Message:            q.Message.replaceDB(db),
		MessageStatus:      q.MessageStatus.replaceDB(db),
		User:               q.User.replaceDB(db),
	}
}

type queryCtx struct {
	Conversation       *conversationDo
	ConversationMember *conversationMemberDo
	Friendship         *friendshipDo
	Message            *messageDo
	MessageStatus      *messageStatusDo
	User               *userDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Conversation:       q.Conversation.WithContext(ctx),
		ConversationMember: q.ConversationMember.WithContext(ctx),
		Friendship:         q.Friendship.WithContext(ctx),
		Message:            q.Message.WithContext(ctx),
		MessageStatus:      q.MessageStatus.WithContext(ctx),
		User:               q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
