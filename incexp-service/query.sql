-- name: RegisterIncome :one 
INSERT INTO transactions(user_id, type, amount, currency, category_id)
VALUES($1, $2, $3, $4, $5)
RETURNING id;

-- name: RegisterExpense :one 
INSERT INTO transactions(user_id, type, amount, currency, category_id)
VALUES($1, $2, $3, $4, $5)
RETURNING id;


-- name: GetListIncomeVSExpense :many
SELECT 
    t.id, 
    t.type, 
    t.amount, 
    t.currency, 
    c.name AS category, 
    t.date
FROM transactions AS t
INNER JOIN categories AS c 
ON t.category_id = c.id
WHERE t.user_id = $1
ORDER BY t.date DESC;

-- name: CheckBudget :one
SELECT 
    b.id AS budget_id,
    b.user_id,
    b.category_id,
    b.amount AS budget_limit,
    COALESCE(SUM(t.amount), 0) AS total_spent,  
    b.currency,
    CASE 
        WHEN COALESCE(SUM(t.amount), 0) > b.amount THEN 'Over Budget'
        ELSE 'Within Budget'
    END AS status
FROM budgets AS b
LEFT JOIN transactions AS t
ON b.category_id = t.category_id 
AND b.user_id = t.user_id
AND t.type = 'expense'
WHERE b.user_id = $1 AND b.category_id = $2
GROUP BY b.id, b.user_id, b.category_id, b.amount, b.currency;
