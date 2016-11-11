// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// GinvBlb represents a row from 'equinox.ginv_blb'.
type GinvBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllGinvBlb(db XODB, callback func(x GinvBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.ginv_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		gb := GinvBlb{}

		// scan
		err = q.Scan(&gb.BlbLrn, &gb.BlbData, &gb.BlbText)
		if err != nil {
			return err
		}
		if !callback(gb) {
			return nil
		}
	}

	return nil
}

// GinvBlbByBlbLrn retrieves a row from 'equinox.ginv_blb' as a GinvBlb.
//
// Generated from index 'ginv_blb_pkey'.
func GinvBlbByBlbLrn(db XODB, blbLrn int64) (*GinvBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.ginv_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	gb := GinvBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&gb.BlbLrn, &gb.BlbData, &gb.BlbText)
	if err != nil {
		return nil, err
	}

	return &gb, nil
}
