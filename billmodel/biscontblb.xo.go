// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// BiscontBlb represents a row from 'equinox.biscont_blb'.
type BiscontBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllBiscontBlb(db XODB, callback func(x BiscontBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.biscont_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		bb := BiscontBlb{}

		// scan
		err = q.Scan(&bb.BlbLrn, &bb.BlbData, &bb.BlbText)
		if err != nil {
			return err
		}
		if !callback(bb) {
			return nil
		}
	}

	return nil
}

// BiscontBlbByBlbLrn retrieves a row from 'equinox.biscont_blb' as a BiscontBlb.
//
// Generated from index 'biscont_blb_pkey'.
func BiscontBlbByBlbLrn(db XODB, blbLrn int64) (*BiscontBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.biscont_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	bb := BiscontBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&bb.BlbLrn, &bb.BlbData, &bb.BlbText)
	if err != nil {
		return nil, err
	}

	return &bb, nil
}
