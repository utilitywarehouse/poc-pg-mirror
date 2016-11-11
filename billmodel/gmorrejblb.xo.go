// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// GmorrejBlb represents a row from 'equinox.gmorrej_blb'.
type GmorrejBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllGmorrejBlb(db XODB, callback func(x GmorrejBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gmorrej_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		gb := GmorrejBlb{}

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

// GmorrejBlbByBlbLrn retrieves a row from 'equinox.gmorrej_blb' as a GmorrejBlb.
//
// Generated from index 'gmorrej_blb_pkey'.
func GmorrejBlbByBlbLrn(db XODB, blbLrn int64) (*GmorrejBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gmorrej_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	gb := GmorrejBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&gb.BlbLrn, &gb.BlbData, &gb.BlbText)
	if err != nil {
		return nil, err
	}

	return &gb, nil
}
