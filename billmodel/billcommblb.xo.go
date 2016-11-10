// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// BillcommBlb represents a row from 'equinox.billcomm_blb'.
type BillcommBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// BillcommBlbByBlbLrn retrieves a row from 'equinox.billcomm_blb' as a BillcommBlb.
//
// Generated from index 'billcomm_blb_pkey'.
func BillcommBlbByBlbLrn(db XODB, blbLrn int64) (*BillcommBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.billcomm_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	bb := BillcommBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&bb.BlbLrn, &bb.BlbData, &bb.BlbText)
	if err != nil {
		return nil, err
	}

	return &bb, nil
}
