-- name: ListPeriods :many
SELECT * FROM gl_period order by 1, 2;

-- name: CreatePeriod :one
INSERT into gl_period (period_id, period_year, start_date, end_date, description, create_date, create_user, update_date, update_user)
values ( 
    $1, $2, $3, $4, $5, $6, $7, $8, $9  
) RETURNING *;

-- name: UpdatePeriod :one
UPDATE gl_period set
period_id = $1,
period_year = $2, 
start_date = $3, 
end_date = $4, 
description = $5, 
create_date = $6, 
create_user = $7, 
update_date = $8, 
update_user = $9
where period_id = $1 and period_year = $2
RETURNING *;

