// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// GdchargeBlb represents a row from 'equinox.gdcharge_blb'.
type GdchargeBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllGdchargeBlb(db XODB, callback func(x GdchargeBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gdcharge_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		gb := GdchargeBlb{}

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

// GdchargeBlbByBlbLrn retrieves a row from 'equinox.gdcharge_blb' as a GdchargeBlb.
//
// Generated from index 'gdcharge_blb_pkey'.
func GdchargeBlbByBlbLrn(db XODB, blbLrn int64) (*GdchargeBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gdcharge_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	gb := GdchargeBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&gb.BlbLrn, &gb.BlbData, &gb.BlbText)
	if err != nil {
		return nil, err
	}

	return &gb, nil
}
