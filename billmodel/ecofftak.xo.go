// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Ecofftak represents a row from 'equinox.ecofftak'.
type Ecofftak struct {
	Ecodate    pq.NullTime     `json:"ecodate"`     // ecodate
	Ecoea10    sql.NullFloat64 `json:"ecoea10"`     // ecoea10
	Ecoem11    sql.NullFloat64 `json:"ecoem11"`     // ecoem11
	Econe12    sql.NullFloat64 `json:"econe12"`     // econe12
	Econo13    sql.NullFloat64 `json:"econo13"`     // econo13
	Econt14    sql.NullFloat64 `json:"econt14"`     // econt14
	Econw15    sql.NullFloat64 `json:"econw15"`     // econw15
	Ecosc16    sql.NullFloat64 `json:"ecosc16"`     // ecosc16
	Ecose17    sql.NullFloat64 `json:"ecose17"`     // ecose17
	Ecoso18    sql.NullFloat64 `json:"ecoso18"`     // ecoso18
	Ecosw19    sql.NullFloat64 `json:"ecosw19"`     // ecosw19
	Ecowm20    sql.NullFloat64 `json:"ecowm20"`     // ecowm20
	Ecown21    sql.NullFloat64 `json:"ecown21"`     // ecown21
	Ecows22    sql.NullFloat64 `json:"ecows22"`     // ecows22
	Ecoun23    sql.NullFloat64 `json:"ecoun23"`     // ecoun23
	EquinoxPrn sql.NullInt64   `json:"equinox_prn"` // equinox_prn
	EquinoxLrn int64           `json:"equinox_lrn"` // equinox_lrn
	EquinoxSec sql.NullInt64   `json:"equinox_sec"` // equinox_sec
}

func AllEcofftak(db XODB, callback func(x Ecofftak) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`ecodate, ecoea10, ecoem11, econe12, econo13, econt14, econw15, ecosc16, ecose17, ecoso18, ecosw19, ecowm20, ecown21, ecows22, ecoun23, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ecofftak `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		e := Ecofftak{}

		// scan
		err = q.Scan(&e.Ecodate, &e.Ecoea10, &e.Ecoem11, &e.Econe12, &e.Econo13, &e.Econt14, &e.Econw15, &e.Ecosc16, &e.Ecose17, &e.Ecoso18, &e.Ecosw19, &e.Ecowm20, &e.Ecown21, &e.Ecows22, &e.Ecoun23, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(e) {
			return nil
		}
	}

	return nil
}

// EcofftakByEquinoxLrn retrieves a row from 'equinox.ecofftak' as a Ecofftak.
//
// Generated from index 'ecofftak_pkey'.
func EcofftakByEquinoxLrn(db XODB, equinoxLrn int64) (*Ecofftak, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ecodate, ecoea10, ecoem11, econe12, econo13, econt14, econw15, ecosc16, ecose17, ecoso18, ecosw19, ecowm20, ecown21, ecows22, ecoun23, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ecofftak ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Ecofftak{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Ecodate, &e.Ecoea10, &e.Ecoem11, &e.Econe12, &e.Econo13, &e.Econt14, &e.Econw15, &e.Ecosc16, &e.Ecose17, &e.Ecoso18, &e.Ecosw19, &e.Ecowm20, &e.Ecown21, &e.Ecows22, &e.Ecoun23, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
