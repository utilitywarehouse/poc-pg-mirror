// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Binlist represents a row from 'equinox.binlist'.
type Binlist struct {
	Binno           sql.NullInt64  `json:"binno"`           // binno
	Binissuer       sql.NullString `json:"binissuer"`       // binissuer
	Bincardscheme   sql.NullString `json:"bincardscheme"`   // bincardscheme
	Binprodtype     sql.NullString `json:"binprodtype"`     // binprodtype
	Binusage        sql.NullString `json:"binusage"`        // binusage
	Binswissueno    sql.NullString `json:"binswissueno"`    // binswissueno
	Binswissuelen   sql.NullString `json:"binswissuelen"`   // binswissuelen
	Binswcdstandard sql.NullString `json:"binswcdstandard"` // binswcdstandard
	Binpanlength    sql.NullString `json:"binpanlength"`    // binpanlength
	Binbintype      sql.NullString `json:"binbintype"`      // binbintype
	Binlastupdate   pq.NullTime    `json:"binlastupdate"`   // binlastupdate
	Binlastupdateby sql.NullString `json:"binlastupdateby"` // binlastupdateby
	Binentered      pq.NullTime    `json:"binentered"`      // binentered
	Binspare1       sql.NullString `json:"binspare1"`       // binspare1
	Binspare2       sql.NullString `json:"binspare2"`       // binspare2
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec
}

func AllBinlist(db XODB, callback func(x Binlist) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`binno, binissuer, bincardscheme, binprodtype, binusage, binswissueno, binswissuelen, binswcdstandard, binpanlength, binbintype, binlastupdate, binlastupdateby, binentered, binspare1, binspare2, equinox_lrn, equinox_sec ` +
		`FROM equinox.binlist `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		b := Binlist{}

		// scan
		err = q.Scan(&b.Binno, &b.Binissuer, &b.Bincardscheme, &b.Binprodtype, &b.Binusage, &b.Binswissueno, &b.Binswissuelen, &b.Binswcdstandard, &b.Binpanlength, &b.Binbintype, &b.Binlastupdate, &b.Binlastupdateby, &b.Binentered, &b.Binspare1, &b.Binspare2, &b.EquinoxLrn, &b.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(b) {
			return nil
		}
	}

	return nil
}

// BinlistByEquinoxLrn retrieves a row from 'equinox.binlist' as a Binlist.
//
// Generated from index 'binlist_pkey'.
func BinlistByEquinoxLrn(db XODB, equinoxLrn int64) (*Binlist, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`binno, binissuer, bincardscheme, binprodtype, binusage, binswissueno, binswissuelen, binswcdstandard, binpanlength, binbintype, binlastupdate, binlastupdateby, binentered, binspare1, binspare2, equinox_lrn, equinox_sec ` +
		`FROM equinox.binlist ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	b := Binlist{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&b.Binno, &b.Binissuer, &b.Bincardscheme, &b.Binprodtype, &b.Binusage, &b.Binswissueno, &b.Binswissuelen, &b.Binswcdstandard, &b.Binpanlength, &b.Binbintype, &b.Binlastupdate, &b.Binlastupdateby, &b.Binentered, &b.Binspare1, &b.Binspare2, &b.EquinoxLrn, &b.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
