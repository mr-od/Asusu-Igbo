-- name: SearchProduct :many
SELECT *
FROM products
WHERE textsearchable_index_col @@ to_tsquery(sqlc.arg(search_query))
ORDER BY created_at DESC
LIMIT 100;