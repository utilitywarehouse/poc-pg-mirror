// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// MpmmBlb represents a row from 'equinox.mpmm_blb'.
type MpmmBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllMpmmBlb(db XODB, callback func(x MpmmBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.mpmm_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		mb := MpmmBlb{}

		// scan
		err = q.Scan(&mb.BlbLrn, &mb.BlbData, &mb.BlbText)
		if err != nil {
			return err
		}
		if !callback(mb) {
			return nil
		}
	}

	return nil
}

// MpmmBlbByBlbLrn retrieves a row from 'equinox.mpmm_blb' as a MpmmBlb.
//
// Generated from index 'mpmm_blb_pkey'.
func MpmmBlbByBlbLrn(db XODB, blbLrn int64) (*MpmmBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.mpmm_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	mb := MpmmBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&mb.BlbLrn, &mb.BlbData, &mb.BlbText)
	if err != nil {
		return nil, err
	}

	return &mb, nil
}
