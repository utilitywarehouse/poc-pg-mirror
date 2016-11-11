// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// GcommpanBlb represents a row from 'equinox.gcommpan_blb'.
type GcommpanBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllGcommpanBlb(db XODB, callback func(x GcommpanBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gcommpan_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		gb := GcommpanBlb{}

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

// GcommpanBlbByBlbLrn retrieves a row from 'equinox.gcommpan_blb' as a GcommpanBlb.
//
// Generated from index 'gcommpan_blb_pkey'.
func GcommpanBlbByBlbLrn(db XODB, blbLrn int64) (*GcommpanBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gcommpan_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	gb := GcommpanBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&gb.BlbLrn, &gb.BlbData, &gb.BlbText)
	if err != nil {
		return nil, err
	}

	return &gb, nil
}
