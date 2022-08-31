-- name: CreateTrade :one
INSERT INTO trade (trd_recordno,trd_glosstraderef,trd_versiono,trd_origin,trd_tradetype,trd_settlementstatus,trd_tradestatus,trd_originversion) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING *;

-- INSERT INTO trade (trd_recordno,trd_glosstraderef,trd_versiono,trd_origin,trd_tradetype,trd_settlementstatus,trd_tradestatus,trd_originversion) VALUES (148931,00000000000000138893,1,'TE','ECAG','CANC','C',2);