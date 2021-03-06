// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// MinicritBlb represents a row from 'equinox.minicrit_blb'.
type MinicritBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllMinicritBlb(db XODB, callback func(x MinicritBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.minicrit_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		mb := MinicritBlb{}

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

// MinicritBlbByBlbLrn retrieves a row from 'equinox.minicrit_blb' as a MinicritBlb.
//
// Generated from index 'minicrit_blb_pkey'.
func MinicritBlbByBlbLrn(db XODB, blbLrn int64) (*MinicritBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.minicrit_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	mb := MinicritBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&mb.BlbLrn, &mb.BlbData, &mb.BlbText)
	if err != nil {
		return nil, err
	}

	return &mb, nil
}
