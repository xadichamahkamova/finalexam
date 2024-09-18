-- name: GetTotalReports :one
SELECT
    (COALESCE(SUM(CASE WHEN type = 'income' THEN amount END), 0))::BIGINT AS total_income,
    (COALESCE(SUM(CASE WHEN type = 'expense' THEN amount END), 0))::BIGINT AS total_expenses,
    (COALESCE(SUM(CASE WHEN type = 'income' THEN amount END), 0) -
    COALESCE(SUM(CASE WHEN type = 'expense' THEN amount END), 0))::BIGINT AS net_savings
FROM
    transactions
WHERE
    user_id = $1;;


-- name: GetReportsSpendingByCategory :many
SELECT
    c.name AS category,
    (COALESCE(SUM(t.amount), 0))::BIGINT AS totalSpent
FROM
    transactions t
JOIN
    categories c ON t.category_id = c.id
WHERE
    t.user_id = $1
GROUP BY
    c.name;
