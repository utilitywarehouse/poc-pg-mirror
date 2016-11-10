// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// MpcauditBlb represents a row from 'equinox.mpcaudit_blb'.
type MpcauditBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// MpcauditBlbByBlbLrn retrieves a row from 'equinox.mpcaudit_blb' as a MpcauditBlb.
//
// Generated from index 'mpcaudit_blb_pkey'.
func MpcauditBlbByBlbLrn(db XODB, blbLrn int64) (*MpcauditBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.mpcaudit_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	mb := MpcauditBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&mb.BlbLrn, &mb.BlbData, &mb.BlbText)
	if err != nil {
		return nil, err
	}

	return &mb, nil
}
