// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// FsecureBlb represents a row from 'equinox.fsecure_blb'.
type FsecureBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllFsecureBlb(db XODB, callback func(x FsecureBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.fsecure_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		fb := FsecureBlb{}

		// scan
		err = q.Scan(&fb.BlbLrn, &fb.BlbData, &fb.BlbText)
		if err != nil {
			return err
		}
		if !callback(fb) {
			return nil
		}
	}

	return nil
}

// FsecureBlbByBlbLrn retrieves a row from 'equinox.fsecure_blb' as a FsecureBlb.
//
// Generated from index 'fsecure_blb_pkey'.
func FsecureBlbByBlbLrn(db XODB, blbLrn int64) (*FsecureBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.fsecure_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	fb := FsecureBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&fb.BlbLrn, &fb.BlbData, &fb.BlbText)
	if err != nil {
		return nil, err
	}

	return &fb, nil
}
