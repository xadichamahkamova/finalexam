-- name: CreateBudget :one
INSERT INTO budgets (user_id, category_id, amount, currency)
VALUES ($1, $2, $3, $4)
RETURNING id AS budget_id;

-- name: GetListBudget :many
SELECT
    b.id AS budget_id,
    c.name AS category,
    b.amount,
    (COALESCE(SUM(t.amount), 0))::BIGINT AS spent,
    b.currency
FROM budgets AS b
INNER JOIN categories AS c 
    ON b.category_id = c.id
LEFT JOIN transactions AS t
    ON t.category_id = b.category_id
    AND t.type = 'expense'
    AND t.user_id = $1
WHERE b.user_id = $1
GROUP BY b.id, c.name, b.amount, b.currency;


-- name: UpdateBudget :one
UPDATE budgets
SET amount = $2
WHERE id = $1
AND user_id = $3
RETURNING 'Budget updated successfully' AS message;
