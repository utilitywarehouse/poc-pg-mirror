// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Imei represents a row from 'equinox.imei'.
type Imei struct {
	ImeiIemi        sql.NullString  `json:"imei_iemi"`        // imei_iemi
	ImeiMake        sql.NullString  `json:"imei_make"`        // imei_make
	ImeiModel       sql.NullString  `json:"imei_model"`       // imei_model
	ImeiReplacecost sql.NullFloat64 `json:"imei_replacecost"` // imei_replacecost
	ImeiRPC         sql.NullInt64   `json:"imei_rpc"`         // imei_rpc
	ImeiMpNone      sql.NullFloat64 `json:"imei_mp_none"`     // imei_mp_none
	ImeiMp1         sql.NullFloat64 `json:"imei_mp_1"`        // imei_mp_1
	ImeiMp2         sql.NullFloat64 `json:"imei_mp_2"`        // imei_mp_2
	ImeiMp3         sql.NullFloat64 `json:"imei_mp_3"`        // imei_mp_3
	ImeiRPCCode     sql.NullString  `json:"imei_rpc_code"`    // imei_rpc_code
	ImeiSubLevel    sql.NullString  `json:"imei_sub_level"`   // imei_sub_level
	ImeiSparen1     sql.NullFloat64 `json:"imei_sparen1"`     // imei_sparen1
	ImeiDateentered pq.NullTime     `json:"imei_dateentered"` // imei_dateentered
	ImeiSupplier    sql.NullString  `json:"imei_supplier"`    // imei_supplier
	ImeiUpc         sql.NullString  `json:"imei_upc"`         // imei_upc
	EquinoxLrn      int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// ImeiByEquinoxLrn retrieves a row from 'equinox.imei' as a Imei.
//
// Generated from index 'imei_pkey'.
func ImeiByEquinoxLrn(db XODB, equinoxLrn int64) (*Imei, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`imei_iemi, imei_make, imei_model, imei_replacecost, imei_rpc, imei_mp_none, imei_mp_1, imei_mp_2, imei_mp_3, imei_rpc_code, imei_sub_level, imei_sparen1, imei_dateentered, imei_supplier, imei_upc, equinox_lrn, equinox_sec ` +
		`FROM equinox.imei ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	i := Imei{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&i.ImeiIemi, &i.ImeiMake, &i.ImeiModel, &i.ImeiReplacecost, &i.ImeiRPC, &i.ImeiMpNone, &i.ImeiMp1, &i.ImeiMp2, &i.ImeiMp3, &i.ImeiRPCCode, &i.ImeiSubLevel, &i.ImeiSparen1, &i.ImeiDateentered, &i.ImeiSupplier, &i.ImeiUpc, &i.EquinoxLrn, &i.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &i, nil
}
