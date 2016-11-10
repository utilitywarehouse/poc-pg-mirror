// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// AclusersBlb represents a row from 'equinox.aclusers_blb'.
type AclusersBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// AclusersBlbByBlbLrn retrieves a row from 'equinox.aclusers_blb' as a AclusersBlb.
//
// Generated from index 'aclusers_blb_pkey'.
func AclusersBlbByBlbLrn(db XODB, blbLrn int64) (*AclusersBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.aclusers_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	ab := AclusersBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&ab.BlbLrn, &ab.BlbData, &ab.BlbText)
	if err != nil {
		return nil, err
	}

	return &ab, nil
}
