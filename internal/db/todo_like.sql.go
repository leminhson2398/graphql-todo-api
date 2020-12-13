// Code generated by sqlc. DO NOT EDIT.
// source: todo_like.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTodoLike = `-- name: CreateTodoLike :exec
INSERT INTO "todo_like" ("owner_id", "todo_id")
VALUES ($1, $2)
ON CONFLICT ("owner_id", "todo_id") DO NOTHING
`

type CreateTodoLikeParams struct {
	OwnerID uuid.UUID `json:"owner_id"`
	TodoID  uuid.UUID `json:"todo_id"`
}

func (q *Queries) CreateTodoLike(ctx context.Context, arg CreateTodoLikeParams) error {
	_, err := q.db.ExecContext(ctx, createTodoLike, arg.OwnerID, arg.TodoID)
	return err
}

const deleteTodoLike = `-- name: DeleteTodoLike :exec
DELETE FROM "todo_like"
WHERE "owner_id" = $1
    AND "todo_id" = $2
`

type DeleteTodoLikeParams struct {
	OwnerID uuid.UUID `json:"owner_id"`
	TodoID  uuid.UUID `json:"todo_id"`
}

func (q *Queries) DeleteTodoLike(ctx context.Context, arg DeleteTodoLikeParams) error {
	_, err := q.db.ExecContext(ctx, deleteTodoLike, arg.OwnerID, arg.TodoID)
	return err
}

const selectAllUsersWhoLikeTodoByTodoID = `-- name: SelectAllUsersWhoLikeTodoByTodoID :many
SELECT user_id, first_name, last_name, username, email, is_online, last_login, password_hash, created_at, role_code
FROM "user_account"
WHERE "user_id" IN (
        SELECT "owner_id"
        FROM "todo_like"
        WHERE "todo_id" = $1
    )
`

func (q *Queries) SelectAllUsersWhoLikeTodoByTodoID(ctx context.Context, todoID uuid.UUID) ([]UserAccount, error) {
	rows, err := q.db.QueryContext(ctx, selectAllUsersWhoLikeTodoByTodoID, todoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAccount
	for rows.Next() {
		var i UserAccount
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.Email,
			&i.IsOnline,
			&i.LastLogin,
			&i.PasswordHash,
			&i.CreatedAt,
			&i.RoleCode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectLikeByOwnerIDAndTodoID = `-- name: SelectLikeByOwnerIDAndTodoID :one
SELECT id, owner_id, todo_id
FROM "todo_like"
WHERE "owner_id" = $1
    AND "todo_id" = $2
`

type SelectLikeByOwnerIDAndTodoIDParams struct {
	OwnerID uuid.UUID `json:"owner_id"`
	TodoID  uuid.UUID `json:"todo_id"`
}

func (q *Queries) SelectLikeByOwnerIDAndTodoID(ctx context.Context, arg SelectLikeByOwnerIDAndTodoIDParams) (TodoLike, error) {
	row := q.db.QueryRowContext(ctx, selectLikeByOwnerIDAndTodoID, arg.OwnerID, arg.TodoID)
	var i TodoLike
	err := row.Scan(&i.ID, &i.OwnerID, &i.TodoID)
	return i, err
}

const selectNumberOfLikesByTodoID = `-- name: SelectNumberOfLikesByTodoID :one
SELECT COUNT(*)
FROM "todo_like"
WHERE "todo_id" = $1
`

func (q *Queries) SelectNumberOfLikesByTodoID(ctx context.Context, todoID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectNumberOfLikesByTodoID, todoID)
	var count int64
	err := row.Scan(&count)
	return count, err
}
