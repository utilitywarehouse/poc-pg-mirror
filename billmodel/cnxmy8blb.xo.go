// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Cnxmy8Blb represents a row from 'equinox.cnxmy8_blb'.
type Cnxmy8Blb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// Cnxmy8BlbByBlbLrn retrieves a row from 'equinox.cnxmy8_blb' as a Cnxmy8Blb.
//
// Generated from index 'cnxmy8_blb_pkey'.
func Cnxmy8BlbByBlbLrn(db XODB, blbLrn int64) (*Cnxmy8Blb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.cnxmy8_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	cb := Cnxmy8Blb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
	if err != nil {
		return nil, err
	}

	return &cb, nil
}
