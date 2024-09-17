-- name: RegisterUser :one
INSERT INTO users(user_name, password, email)
VALUES($1, $2, $3)
RETURNING id, user_name, email;

-- name: GetUserById :one 
SELECT id, user_name, email
FROM users 
WHERE id=$1;

-- name: GetUserByUserName :one
SELECT id, email, password
FROM users
WHERE user_name = $1;

-- name: GetUsersList :many
SELECT id, user_name, email
FROM users;