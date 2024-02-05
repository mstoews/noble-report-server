-- name: ListDistLedger :many
SELECT *
FROM gl_distribution_ledger
ORDER BY 3,1,2
LIMIT 10000;

-- name: GetDistLedgerByRef :one
SELECT *
FROM gl_distribution_ledger WHERE account = $1 and child = $2 and period = $3 
ORDER BY 3,1,2;


-- name: ListDistLedgerByPeriod :many
SELECT *
FROM gl_distribution_ledger WHERE period = $1 
ORDER BY 3,1,2;

-- name: ListAccountTypes :many
SELECT *
FROM gl_account_type
ORDER BY 1;

-- name: ListJournalDetails :many
SELECT * FROM gl_journal_detail order by 1,2;

-- name: CreateJournalDetails :one
INSERT INTO gl_journal_detail  
 (
    journal_id, 
    journal_subid, 
    account, 
    child, 
    sub_type, 
    description, 
    debit, 
    credit,
    create_date, 
    create_user)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateJournalDetails :one
UPDATE gl_journal_detail 
SET 
    account = $3, 
    child = $4, 
    sub_type = $5, 
    description = $6, 
    debit = $7, 
    credit = $8,
    create_date = $9, 
    create_user = $10
WHERE journal_id = $1 and journal_subid = $2
RETURNING *;

-- name: ListJournalHeader :many
SELECT * FROM gl_journal_header ORDER BY 1;

-- name: CreateJournalHeader :one
INSERT INTO gl_journal_header 
(   journal_id, 
    description, 
    booked, 
    booked_date, 
    booked_user, 
    create_date, 
    create_user,
    period, 
    transaction_date, 
    status, 
    type, 
    amount)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: UpdateJournalHeader :one
UPDATE gl_journal_header 
SET 
    journal_id = $1, 
    description = $2, 
    booked = $3, 
    booked_date = $4, 
    booked_user = $5, 
    create_date = $6, 
    create_user = $7,
    period = $8, 
    transaction_date = $9, 
    status = $10, 
    type = $11, 
    amount = $12
WHERE journal_id = $1
RETURNING *;

