-- name: GetHabitById :one
SELECT * FROM habits
WHERE id = $1 LIMIT 1;

-- name: GetHabitsForUserId :one
SELECT * FROM habits
WHERE user_id = $1;

-- name: CreateHabit :one
INSERT INTO habits (
  id, user_id, title
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateHabitById :one
UPDATE habits
set title = $2,
completed = $3
WHERE id = $1
RETURNING *;