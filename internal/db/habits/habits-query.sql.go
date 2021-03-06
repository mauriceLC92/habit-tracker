// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: habits-query.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createHabit = `-- name: CreateHabit :one
INSERT INTO habits (
  id, user_id, title
) VALUES (
  $1, $2, $3
)
RETURNING id, user_id, title, completed, created_at, updated_at
`

type CreateHabitParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Title  string
}

func (q *Queries) CreateHabit(ctx context.Context, arg CreateHabitParams) (Habit, error) {
	row := q.db.QueryRowContext(ctx, createHabit, arg.ID, arg.UserID, arg.Title)
	var i Habit
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getHabitById = `-- name: GetHabitById :one
SELECT id, user_id, title, completed, created_at, updated_at FROM habits
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetHabitById(ctx context.Context, id uuid.UUID) (Habit, error) {
	row := q.db.QueryRowContext(ctx, getHabitById, id)
	var i Habit
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getHabitsForUserId = `-- name: GetHabitsForUserId :one
SELECT id, user_id, title, completed, created_at, updated_at FROM habits
WHERE user_id = $1
`

func (q *Queries) GetHabitsForUserId(ctx context.Context, userID uuid.UUID) (Habit, error) {
	row := q.db.QueryRowContext(ctx, getHabitsForUserId, userID)
	var i Habit
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateHabitById = `-- name: UpdateHabitById :one
UPDATE habits
set title = $2,
completed = $3
WHERE id = $1
RETURNING id, user_id, title, completed, created_at, updated_at
`

type UpdateHabitByIdParams struct {
	ID        uuid.UUID
	Title     string
	Completed bool
}

func (q *Queries) UpdateHabitById(ctx context.Context, arg UpdateHabitByIdParams) (Habit, error) {
	row := q.db.QueryRowContext(ctx, updateHabitById, arg.ID, arg.Title, arg.Completed)
	var i Habit
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
