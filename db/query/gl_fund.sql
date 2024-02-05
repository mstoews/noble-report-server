-- name: ListFunds :many
SELECT * from gl_funds order by 1,2
LIMIT 10000;

-- name: CreateFund :one
insert into gl_funds (fund, description, create_date, create_user)
values ( $1, $2, $3, $4 ) RETURNING *;

-- name: UpdateFund :one
UPDATE gl_funds set 
fund = $1, 
description = $2, 
create_date = $3,
create_user = $4 
where fund = $1 
RETURNING *;

-- name: GetFund :one
SELECT * FROM gl_funds
WHERE fund = $1
ORDER BY 1
LIMIT 10000;


