-- name: CreateTrade :one
INSERT INTO trade (trd_recordno,trd_glosstraderef,trd_versiono,trd_origin,
trd_tradetype,trd_settlementstatus,trd_tradestatus,trd_originversion,trd_date) 
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING *;

-- name: ListTrades :many
SELECT * FROM trade
WHERE 
    trd_date = $1 
ORDER BY trd_recordno
LIMIT $2
OFFSET $3;

-- INSERT INTO trade (trd_recordno,trd_glosstraderef,trd_versiono,trd_origin,trd_tradetype,trd_settlementstatus,trd_tradestatus,trd_originversion) VALUES (148931,00000000000000138893,1,'TE','ECAG','CANC','C',2, '2022-02-01);

-- name: GetTrade :many
SELECT * FROM trade
WHERE 
    trd_recordno = $1
ORDER BY trd_recordno
LIMIT 1;