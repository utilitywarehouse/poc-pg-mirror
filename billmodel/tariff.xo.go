// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Tariff represents a row from 'equinox.tariff'.
type Tariff struct {
	Tariffuniquesys  sql.NullFloat64 `json:"tariffuniquesys"`  // tariffuniquesys
	Tariffid         sql.NullString  `json:"tariffid"`         // tariffid
	Tariffname       sql.NullString  `json:"tariffname"`       // tariffname
	Tariffcomments   sql.NullInt64   `json:"tariffcomments"`   // tariffcomments
	Tariffstartdate  pq.NullTime     `json:"tariffstartdate"`  // tariffstartdate
	Tariffenddate    pq.NullTime     `json:"tariffenddate"`    // tariffenddate
	Tarifffreemins   sql.NullFloat64 `json:"tarifffreemins"`   // tarifffreemins
	Tariffreducedmin sql.NullFloat64 `json:"tariffreducedmin"` // tariffreducedmin
	Tarifffreetexts  sql.NullFloat64 `json:"tarifffreetexts"`  // tarifffreetexts
	Tarifffreemintyp sql.NullString  `json:"tarifffreemintyp"` // tarifffreemintyp
	Tarifffreemincnt sql.NullString  `json:"tarifffreemincnt"` // tarifffreemincnt
	Tariffusage      sql.NullString  `json:"tariffusage"`      // tariffusage
	Tariffnewold     sql.NullString  `json:"tariffnewold"`     // tariffnewold
	Tariffcompack    sql.NullString  `json:"tariffcompack"`    // tariffcompack
	Tariffsubtariff  sql.NullString  `json:"tariffsubtariff"`  // tariffsubtariff
	Tariffsparen1    sql.NullFloat64 `json:"tariffsparen1"`    // tariffsparen1
	Tariffspared1    pq.NullTime     `json:"tariffspared1"`    // tariffspared1
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// TariffByEquinoxLrn retrieves a row from 'equinox.tariff' as a Tariff.
//
// Generated from index 'tariff_pkey'.
func TariffByEquinoxLrn(db XODB, equinoxLrn int64) (*Tariff, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tariffuniquesys, tariffid, tariffname, tariffcomments, tariffstartdate, tariffenddate, tarifffreemins, tariffreducedmin, tarifffreetexts, tarifffreemintyp, tarifffreemincnt, tariffusage, tariffnewold, tariffcompack, tariffsubtariff, tariffsparen1, tariffspared1, equinox_lrn, equinox_sec ` +
		`FROM equinox.tariff ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	t := Tariff{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&t.Tariffuniquesys, &t.Tariffid, &t.Tariffname, &t.Tariffcomments, &t.Tariffstartdate, &t.Tariffenddate, &t.Tarifffreemins, &t.Tariffreducedmin, &t.Tarifffreetexts, &t.Tarifffreemintyp, &t.Tarifffreemincnt, &t.Tariffusage, &t.Tariffnewold, &t.Tariffcompack, &t.Tariffsubtariff, &t.Tariffsparen1, &t.Tariffspared1, &t.EquinoxLrn, &t.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
