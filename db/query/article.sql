-- name: CreateArticle :one
INSERT INTO articles (
  content,
  title,
  poster_id
) VALUES (
  $1,
  $2,
  $3
) RETURNING *;

-- name: GetArticle :one
SELECT * FROM articles
WHERE id = $1 LIMIT 1;

-- name: ListArticles :many
SELECT * FROM articles
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateArticle :exec
UPDATE articles
SET content = $2,
title = $3
WHERE id = $1
RETURNING *;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1;