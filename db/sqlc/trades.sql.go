// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: trades.sql

package db

import (
	"context"
)

const createTrade = `-- name: CreateTrade :one
INSERT INTO trade (trd_recordno,trd_glosstraderef,trd_versiono,trd_origin,trd_tradetype,trd_settlementstatus,trd_tradestatus,trd_originversion) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING trd_recordno, trd_glosstraderef, trd_versiono, trd_origin, trd_tradetype, trd_settlementstatus, trd_tradestatus, trd_originversion
`

type CreateTradeParams struct {
	TrdRecordno         int32  `json:"trd_recordno"`
	TrdGlosstraderef    int32  `json:"trd_glosstraderef"`
	TrdVersiono         int32  `json:"trd_versiono"`
	TrdOrigin           string `json:"trd_origin"`
	TrdTradetype        string `json:"trd_tradetype"`
	TrdSettlementstatus string `json:"trd_settlementstatus"`
	TrdTradestatus      string `json:"trd_tradestatus"`
	TrdOriginversion    int32  `json:"trd_originversion"`
}

func (q *Queries) CreateTrade(ctx context.Context, arg CreateTradeParams) (Trade, error) {
	row := q.db.QueryRowContext(ctx, createTrade,
		arg.TrdRecordno,
		arg.TrdGlosstraderef,
		arg.TrdVersiono,
		arg.TrdOrigin,
		arg.TrdTradetype,
		arg.TrdSettlementstatus,
		arg.TrdTradestatus,
		arg.TrdOriginversion,
	)
	var i Trade
	err := row.Scan(
		&i.TrdRecordno,
		&i.TrdGlosstraderef,
		&i.TrdVersiono,
		&i.TrdOrigin,
		&i.TrdTradetype,
		&i.TrdSettlementstatus,
		&i.TrdTradestatus,
		&i.TrdOriginversion,
	)
	return i, err
}