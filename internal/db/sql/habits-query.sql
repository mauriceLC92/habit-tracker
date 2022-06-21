-- name: GetHabit :one
SELECT * FROM habits
WHERE id = $1 LIMIT 1;