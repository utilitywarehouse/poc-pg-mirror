// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// PesratesBlb represents a row from 'equinox.pesrates_blb'.
type PesratesBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// PesratesBlbByBlbLrn retrieves a row from 'equinox.pesrates_blb' as a PesratesBlb.
//
// Generated from index 'pesrates_blb_pkey'.
func PesratesBlbByBlbLrn(db XODB, blbLrn int64) (*PesratesBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.pesrates_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	pb := PesratesBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&pb.BlbLrn, &pb.BlbData, &pb.BlbText)
	if err != nil {
		return nil, err
	}

	return &pb, nil
}
