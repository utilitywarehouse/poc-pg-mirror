// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// AffiliatBlb represents a row from 'equinox.affiliat_blb'.
type AffiliatBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// AffiliatBlbByBlbLrn retrieves a row from 'equinox.affiliat_blb' as a AffiliatBlb.
//
// Generated from index 'affiliat_blb_pkey'.
func AffiliatBlbByBlbLrn(db XODB, blbLrn int64) (*AffiliatBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.affiliat_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	ab := AffiliatBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&ab.BlbLrn, &ab.BlbData, &ab.BlbText)
	if err != nil {
		return nil, err
	}

	return &ab, nil
}
