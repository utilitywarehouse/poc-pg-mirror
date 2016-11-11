// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// CrsarBlb represents a row from 'equinox.crsar_blb'.
type CrsarBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllCrsarBlb(db XODB, callback func(x CrsarBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.crsar_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		cb := CrsarBlb{}

		// scan
		err = q.Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
		if err != nil {
			return err
		}
		if !callback(cb) {
			return nil
		}
	}

	return nil
}

// CrsarBlbByBlbLrn retrieves a row from 'equinox.crsar_blb' as a CrsarBlb.
//
// Generated from index 'crsar_blb_pkey'.
func CrsarBlbByBlbLrn(db XODB, blbLrn int64) (*CrsarBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.crsar_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	cb := CrsarBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
	if err != nil {
		return nil, err
	}

	return &cb, nil
}
