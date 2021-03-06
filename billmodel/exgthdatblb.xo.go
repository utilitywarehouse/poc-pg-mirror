// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// ExgthdatBlb represents a row from 'equinox.exgthdat_blb'.
type ExgthdatBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllExgthdatBlb(db XODB, callback func(x ExgthdatBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.exgthdat_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		eb := ExgthdatBlb{}

		// scan
		err = q.Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
		if err != nil {
			return err
		}
		if !callback(eb) {
			return nil
		}
	}

	return nil
}

// ExgthdatBlbByBlbLrn retrieves a row from 'equinox.exgthdat_blb' as a ExgthdatBlb.
//
// Generated from index 'exgthdat_blb_pkey'.
func ExgthdatBlbByBlbLrn(db XODB, blbLrn int64) (*ExgthdatBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.exgthdat_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	eb := ExgthdatBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
	if err != nil {
		return nil, err
	}

	return &eb, nil
}
