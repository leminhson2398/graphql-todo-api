// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CountNumOfLikesOfCommentByID(ctx context.Context, id uuid.UUID) (int64, error)
	CountNumberOfCommentsByTodoID(ctx context.Context, todoID uuid.UUID) (int64, error)
	CreateChildCOmment(ctx context.Context, arg CreateChildCOmmentParams) (Comment, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error)
	CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error)
	CreateTodoLike(ctx context.Context, arg CreateTodoLikeParams) error
	CreateUserAccount(ctx context.Context, arg CreateUserAccountParams) (UserAccount, error)
	DeleteExpiredTokens(ctx context.Context) error
	DeleteRefreshTokenByID(ctx context.Context, id uuid.UUID) error
	DeleteRefreshTokenByUserID(ctx context.Context, userID uuid.UUID) error
	DeleteTodoByID(ctx context.Context, arg DeleteTodoByIDParams) error
	DeleteTodoLike(ctx context.Context, arg DeleteTodoLikeParams) error
	GetAllTodos(ctx context.Context) ([]Todo, error)
	GetAllUserAccounts(ctx context.Context) ([]UserAccount, error)
	GetRefreshTokenByID(ctx context.Context, id uuid.UUID) (RefreshToken, error)
	GetTodoByID(ctx context.Context, id uuid.UUID) (Todo, error)
	GetUserAccountByEmail(ctx context.Context, email string) (UserAccount, error)
	GetUserAccountByID(ctx context.Context, userID uuid.UUID) (UserAccount, error)
	GetUserByEmailOrUsername(ctx context.Context, arg GetUserByEmailOrUsernameParams) (UserAccount, error)
	LikeComment(ctx context.Context, arg LikeCommentParams) error
	SelectAllTodosOfUserByUserID(ctx context.Context, ownerID uuid.UUID) ([]Todo, error)
	SelectAllUsersWhoLikeTodoByTodoID(ctx context.Context, todoID uuid.UUID) ([]UserAccount, error)
	SelectLikeByOwnerIDAndCommentID(ctx context.Context, arg SelectLikeByOwnerIDAndCommentIDParams) (CommentLike, error)
	SelectLikeByOwnerIDAndTodoID(ctx context.Context, arg SelectLikeByOwnerIDAndTodoIDParams) (TodoLike, error)
	SelectMainCommentsByTodoId(ctx context.Context, todoID uuid.UUID) ([]Comment, error)
	SelectNumberOfLikesByTodoID(ctx context.Context, todoID uuid.UUID) (int64, error)
	SelectRoleByCode(ctx context.Context, code string) (Role, error)
	SelectSubcommentsByParentCommentId(ctx context.Context, parentCommentID uuid.UUID) ([]Comment, error)
	SetUserPassword(ctx context.Context, arg SetUserPasswordParams) (UserAccount, error)
	UnlikeComment(ctx context.Context, arg UnlikeCommentParams) error
	UpdateTodoByID(ctx context.Context, arg UpdateTodoByIDParams) (Todo, error)
	UpdateUserAccountInfo(ctx context.Context, arg UpdateUserAccountInfoParams) (UserAccount, error)
}

var _ Querier = (*Queries)(nil)
