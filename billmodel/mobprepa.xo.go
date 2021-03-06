// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Mobprepa represents a row from 'equinox.mobprepa'.
type Mobprepa struct {
	Mppcliuniquesys sql.NullFloat64 `json:"mppcliuniquesys"` // mppcliuniquesys
	Mppcli          sql.NullString  `json:"mppcli"`          // mppcli
	Mppsim          sql.NullString  `json:"mppsim"`          // mppsim
	Mppaccountno    sql.NullString  `json:"mppaccountno"`    // mppaccountno
	Mppminbalance   sql.NullInt64   `json:"mppminbalance"`   // mppminbalance
	Mppmaxbalance   sql.NullInt64   `json:"mppmaxbalance"`   // mppmaxbalance
	Mppdebitno      sql.NullString  `json:"mppdebitno"`      // mppdebitno
	Mppdebitstart   sql.NullString  `json:"mppdebitstart"`   // mppdebitstart
	Mppdebitexpiry  sql.NullString  `json:"mppdebitexpiry"`  // mppdebitexpiry
	Mppdebitissue   sql.NullString  `json:"mppdebitissue"`   // mppdebitissue
	Mppdebitcv2     sql.NullString  `json:"mppdebitcv2"`     // mppdebitcv2
	Mpptopuppin     sql.NullString  `json:"mpptopuppin"`     // mpptopuppin
	Mpplivedate     pq.NullTime     `json:"mpplivedate"`     // mpplivedate
	Mppenddate      pq.NullTime     `json:"mppenddate"`      // mppenddate
	Mpptexttopup    sql.NullInt64   `json:"mpptexttopup"`    // mpptexttopup
	Mppccunique     sql.NullFloat64 `json:"mppccunique"`     // mppccunique
	Mppshopperref   sql.NullString  `json:"mppshopperref"`   // mppshopperref
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec
}

func AllMobprepa(db XODB, callback func(x Mobprepa) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`mppcliuniquesys, mppcli, mppsim, mppaccountno, mppminbalance, mppmaxbalance, mppdebitno, mppdebitstart, mppdebitexpiry, mppdebitissue, mppdebitcv2, mpptopuppin, mpplivedate, mppenddate, mpptexttopup, mppccunique, mppshopperref, equinox_lrn, equinox_sec ` +
		`FROM equinox.mobprepa `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		m := Mobprepa{}

		// scan
		err = q.Scan(&m.Mppcliuniquesys, &m.Mppcli, &m.Mppsim, &m.Mppaccountno, &m.Mppminbalance, &m.Mppmaxbalance, &m.Mppdebitno, &m.Mppdebitstart, &m.Mppdebitexpiry, &m.Mppdebitissue, &m.Mppdebitcv2, &m.Mpptopuppin, &m.Mpplivedate, &m.Mppenddate, &m.Mpptexttopup, &m.Mppccunique, &m.Mppshopperref, &m.EquinoxLrn, &m.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(m) {
			return nil
		}
	}

	return nil
}

// MobprepaByEquinoxLrn retrieves a row from 'equinox.mobprepa' as a Mobprepa.
//
// Generated from index 'mobprepa_pkey'.
func MobprepaByEquinoxLrn(db XODB, equinoxLrn int64) (*Mobprepa, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`mppcliuniquesys, mppcli, mppsim, mppaccountno, mppminbalance, mppmaxbalance, mppdebitno, mppdebitstart, mppdebitexpiry, mppdebitissue, mppdebitcv2, mpptopuppin, mpplivedate, mppenddate, mpptexttopup, mppccunique, mppshopperref, equinox_lrn, equinox_sec ` +
		`FROM equinox.mobprepa ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	m := Mobprepa{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&m.Mppcliuniquesys, &m.Mppcli, &m.Mppsim, &m.Mppaccountno, &m.Mppminbalance, &m.Mppmaxbalance, &m.Mppdebitno, &m.Mppdebitstart, &m.Mppdebitexpiry, &m.Mppdebitissue, &m.Mppdebitcv2, &m.Mpptopuppin, &m.Mpplivedate, &m.Mppenddate, &m.Mpptexttopup, &m.Mppccunique, &m.Mppshopperref, &m.EquinoxLrn, &m.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
