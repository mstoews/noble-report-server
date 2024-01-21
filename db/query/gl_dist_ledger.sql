-- name: ListDistLedger :many
SELECT *
FROM gl_distribution_ledger
ORDER BY 3,1,2
LIMIT 10000;


-- name: getDistLedgerByRef :one
SELECT *
FROM gl_distribution_ledger WHERE account = $1 and child = $2 and period = $3 
ORDER BY 3,1,2;
