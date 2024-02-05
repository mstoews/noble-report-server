-- name: ListGeneralLedger :many
SELECT * from gl_accounts order by 1,2
LIMIT 10000;

-- name: CreateGLAccount :one
insert into gl_accounts (account, child, parent_account, type, sub_type, description, balance, comments, create_date,
                         create_user, update_date, update_user)
values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12    
) RETURNING *;

-- name: UpdateGLAccount :one
UPDATE gl_accounts set 
account = $1, 
child = $2, 
parent_account = $3, 
type = $4, 
sub_type = $5, 
description = $6, 
balance = $7, 
comments = $8, 
create_date = $9,
create_user = $10, 
update_date = $11,
update_user = $12 where account = $1 and child = $2
RETURNING *;

-- name: GetGLAccount :one
SELECT *
FROM gl_accounts
WHERE account = $1 and child = $2
ORDER BY 1,2
LIMIT 10000;



