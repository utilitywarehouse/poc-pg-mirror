// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// GmamBlb represents a row from 'equinox.gmam_blb'.
type GmamBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllGmamBlb(db XODB, callback func(x GmamBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gmam_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		gb := GmamBlb{}

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

// GmamBlbByBlbLrn retrieves a row from 'equinox.gmam_blb' as a GmamBlb.
//
// Generated from index 'gmam_blb_pkey'.
func GmamBlbByBlbLrn(db XODB, blbLrn int64) (*GmamBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.gmam_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	gb := GmamBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&gb.BlbLrn, &gb.BlbData, &gb.BlbText)
	if err != nil {
		return nil, err
	}

	return &gb, nil
}
