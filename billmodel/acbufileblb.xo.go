// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// AcbufileBlb represents a row from 'equinox.acbufile_blb'.
type AcbufileBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllAcbufileBlb(db XODB, callback func(x AcbufileBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.acbufile_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		ab := AcbufileBlb{}

		// scan
		err = q.Scan(&ab.BlbLrn, &ab.BlbData, &ab.BlbText)
		if err != nil {
			return err
		}
		if !callback(ab) {
			return nil
		}
	}

	return nil
}

// AcbufileBlbByBlbLrn retrieves a row from 'equinox.acbufile_blb' as a AcbufileBlb.
//
// Generated from index 'acbufile_blb_pkey'.
func AcbufileBlbByBlbLrn(db XODB, blbLrn int64) (*AcbufileBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.acbufile_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	ab := AcbufileBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&ab.BlbLrn, &ab.BlbData, &ab.BlbText)
	if err != nil {
		return nil, err
	}

	return &ab, nil
}
