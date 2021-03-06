// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Band represents a row from 'equinox.band'.
type Band struct {
	Banduniquesys    sql.NullFloat64 `json:"banduniquesys"`    // banduniquesys
	Bandid           sql.NullString  `json:"bandid"`           // bandid
	Bandchrgopt      sql.NullString  `json:"bandchrgopt"`      // bandchrgopt
	Bandfixed        sql.NullString  `json:"bandfixed"`        // bandfixed
	Bandrebate       sql.NullInt64   `json:"bandrebate"`       // bandrebate
	Bandfreemins     sql.NullInt64   `json:"bandfreemins"`     // bandfreemins
	Bandvatrate      sql.NullFloat64 `json:"bandvatrate"`      // bandvatrate
	Bandvol1type1    sql.NullFloat64 `json:"bandvol1type1"`    // bandvol1type1
	Bandvol1type2    sql.NullFloat64 `json:"bandvol1type2"`    // bandvol1type2
	Bandvol1type3    sql.NullFloat64 `json:"bandvol1type3"`    // bandvol1type3
	Bandvol1type4    sql.NullFloat64 `json:"bandvol1type4"`    // bandvol1type4
	Bandvol1type5    sql.NullFloat64 `json:"bandvol1type5"`    // bandvol1type5
	Bandvol2type1    sql.NullFloat64 `json:"bandvol2type1"`    // bandvol2type1
	Bandvol2type2    sql.NullFloat64 `json:"bandvol2type2"`    // bandvol2type2
	Bandvol2type3    sql.NullFloat64 `json:"bandvol2type3"`    // bandvol2type3
	Bandvol2type4    sql.NullFloat64 `json:"bandvol2type4"`    // bandvol2type4
	Bandvol2type5    sql.NullFloat64 `json:"bandvol2type5"`    // bandvol2type5
	Bandvol3type1    sql.NullFloat64 `json:"bandvol3type1"`    // bandvol3type1
	Bandvol3type2    sql.NullFloat64 `json:"bandvol3type2"`    // bandvol3type2
	Bandvol3type3    sql.NullFloat64 `json:"bandvol3type3"`    // bandvol3type3
	Bandvol3type4    sql.NullFloat64 `json:"bandvol3type4"`    // bandvol3type4
	Bandvol3type5    sql.NullFloat64 `json:"bandvol3type5"`    // bandvol3type5
	Bandvol4type1    sql.NullFloat64 `json:"bandvol4type1"`    // bandvol4type1
	Bandvol4type2    sql.NullFloat64 `json:"bandvol4type2"`    // bandvol4type2
	Bandvol4type3    sql.NullFloat64 `json:"bandvol4type3"`    // bandvol4type3
	Bandvol4type4    sql.NullFloat64 `json:"bandvol4type4"`    // bandvol4type4
	Bandvol4type5    sql.NullFloat64 `json:"bandvol4type5"`    // bandvol4type5
	Bandvol5type1    sql.NullFloat64 `json:"bandvol5type1"`    // bandvol5type1
	Bandvol5type2    sql.NullFloat64 `json:"bandvol5type2"`    // bandvol5type2
	Bandvol5type3    sql.NullFloat64 `json:"bandvol5type3"`    // bandvol5type3
	Bandvol5type4    sql.NullFloat64 `json:"bandvol5type4"`    // bandvol5type4
	Bandvol5type5    sql.NullFloat64 `json:"bandvol5type5"`    // bandvol5type5
	Bandvol6type1    sql.NullFloat64 `json:"bandvol6type1"`    // bandvol6type1
	Bandvol6type2    sql.NullFloat64 `json:"bandvol6type2"`    // bandvol6type2
	Bandvol6type3    sql.NullFloat64 `json:"bandvol6type3"`    // bandvol6type3
	Bandvol6type4    sql.NullFloat64 `json:"bandvol6type4"`    // bandvol6type4
	Bandvol6type5    sql.NullFloat64 `json:"bandvol6type5"`    // bandvol6type5
	Bandvol7type1    sql.NullFloat64 `json:"bandvol7type1"`    // bandvol7type1
	Bandvol7type2    sql.NullFloat64 `json:"bandvol7type2"`    // bandvol7type2
	Bandvol7type3    sql.NullFloat64 `json:"bandvol7type3"`    // bandvol7type3
	Bandvol7type4    sql.NullFloat64 `json:"bandvol7type4"`    // bandvol7type4
	Bandvol7type5    sql.NullFloat64 `json:"bandvol7type5"`    // bandvol7type5
	Bandvol8type1    sql.NullFloat64 `json:"bandvol8type1"`    // bandvol8type1
	Bandvol8type2    sql.NullFloat64 `json:"bandvol8type2"`    // bandvol8type2
	Bandvol8type3    sql.NullFloat64 `json:"bandvol8type3"`    // bandvol8type3
	Bandvol8type4    sql.NullFloat64 `json:"bandvol8type4"`    // bandvol8type4
	Bandvol8type5    sql.NullFloat64 `json:"bandvol8type5"`    // bandvol8type5
	Bandvol9type1    sql.NullFloat64 `json:"bandvol9type1"`    // bandvol9type1
	Bandvol9type2    sql.NullFloat64 `json:"bandvol9type2"`    // bandvol9type2
	Bandvol9type3    sql.NullFloat64 `json:"bandvol9type3"`    // bandvol9type3
	Bandvol9type4    sql.NullFloat64 `json:"bandvol9type4"`    // bandvol9type4
	Bandvol9type5    sql.NullFloat64 `json:"bandvol9type5"`    // bandvol9type5
	Bandvol10type1   sql.NullFloat64 `json:"bandvol10type1"`   // bandvol10type1
	Bandvol10type2   sql.NullFloat64 `json:"bandvol10type2"`   // bandvol10type2
	Bandvol10type3   sql.NullFloat64 `json:"bandvol10type3"`   // bandvol10type3
	Bandvol10type4   sql.NullFloat64 `json:"bandvol10type4"`   // bandvol10type4
	Bandvol10type5   sql.NullFloat64 `json:"bandvol10type5"`   // bandvol10type5
	Bandinclfreemins sql.NullInt64   `json:"bandinclfreemins"` // bandinclfreemins
	Banincindiscount sql.NullInt64   `json:"banincindiscount"` // banincindiscount
	Bandbasetype1    sql.NullFloat64 `json:"bandbasetype1"`    // bandbasetype1
	Bandbasetype2    sql.NullFloat64 `json:"bandbasetype2"`    // bandbasetype2
	Bandbasetype3    sql.NullFloat64 `json:"bandbasetype3"`    // bandbasetype3
	Bandbasetype4    sql.NullFloat64 `json:"bandbasetype4"`    // bandbasetype4
	Bandbasetype5    sql.NullFloat64 `json:"bandbasetype5"`    // bandbasetype5
	Bandmarkup       sql.NullString  `json:"bandmarkup"`       // bandmarkup
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllBand(db XODB, callback func(x Band) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`banduniquesys, bandid, bandchrgopt, bandfixed, bandrebate, bandfreemins, bandvatrate, bandvol1type1, bandvol1type2, bandvol1type3, bandvol1type4, bandvol1type5, bandvol2type1, bandvol2type2, bandvol2type3, bandvol2type4, bandvol2type5, bandvol3type1, bandvol3type2, bandvol3type3, bandvol3type4, bandvol3type5, bandvol4type1, bandvol4type2, bandvol4type3, bandvol4type4, bandvol4type5, bandvol5type1, bandvol5type2, bandvol5type3, bandvol5type4, bandvol5type5, bandvol6type1, bandvol6type2, bandvol6type3, bandvol6type4, bandvol6type5, bandvol7type1, bandvol7type2, bandvol7type3, bandvol7type4, bandvol7type5, bandvol8type1, bandvol8type2, bandvol8type3, bandvol8type4, bandvol8type5, bandvol9type1, bandvol9type2, bandvol9type3, bandvol9type4, bandvol9type5, bandvol10type1, bandvol10type2, bandvol10type3, bandvol10type4, bandvol10type5, bandinclfreemins, banincindiscount, bandbasetype1, bandbasetype2, bandbasetype3, bandbasetype4, bandbasetype5, bandmarkup, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.band `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		b := Band{}

		// scan
		err = q.Scan(&b.Banduniquesys, &b.Bandid, &b.Bandchrgopt, &b.Bandfixed, &b.Bandrebate, &b.Bandfreemins, &b.Bandvatrate, &b.Bandvol1type1, &b.Bandvol1type2, &b.Bandvol1type3, &b.Bandvol1type4, &b.Bandvol1type5, &b.Bandvol2type1, &b.Bandvol2type2, &b.Bandvol2type3, &b.Bandvol2type4, &b.Bandvol2type5, &b.Bandvol3type1, &b.Bandvol3type2, &b.Bandvol3type3, &b.Bandvol3type4, &b.Bandvol3type5, &b.Bandvol4type1, &b.Bandvol4type2, &b.Bandvol4type3, &b.Bandvol4type4, &b.Bandvol4type5, &b.Bandvol5type1, &b.Bandvol5type2, &b.Bandvol5type3, &b.Bandvol5type4, &b.Bandvol5type5, &b.Bandvol6type1, &b.Bandvol6type2, &b.Bandvol6type3, &b.Bandvol6type4, &b.Bandvol6type5, &b.Bandvol7type1, &b.Bandvol7type2, &b.Bandvol7type3, &b.Bandvol7type4, &b.Bandvol7type5, &b.Bandvol8type1, &b.Bandvol8type2, &b.Bandvol8type3, &b.Bandvol8type4, &b.Bandvol8type5, &b.Bandvol9type1, &b.Bandvol9type2, &b.Bandvol9type3, &b.Bandvol9type4, &b.Bandvol9type5, &b.Bandvol10type1, &b.Bandvol10type2, &b.Bandvol10type3, &b.Bandvol10type4, &b.Bandvol10type5, &b.Bandinclfreemins, &b.Banincindiscount, &b.Bandbasetype1, &b.Bandbasetype2, &b.Bandbasetype3, &b.Bandbasetype4, &b.Bandbasetype5, &b.Bandmarkup, &b.EquinoxPrn, &b.EquinoxLrn, &b.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(b) {
			return nil
		}
	}

	return nil
}

// BandByEquinoxLrn retrieves a row from 'equinox.band' as a Band.
//
// Generated from index 'band_pkey'.
func BandByEquinoxLrn(db XODB, equinoxLrn int64) (*Band, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`banduniquesys, bandid, bandchrgopt, bandfixed, bandrebate, bandfreemins, bandvatrate, bandvol1type1, bandvol1type2, bandvol1type3, bandvol1type4, bandvol1type5, bandvol2type1, bandvol2type2, bandvol2type3, bandvol2type4, bandvol2type5, bandvol3type1, bandvol3type2, bandvol3type3, bandvol3type4, bandvol3type5, bandvol4type1, bandvol4type2, bandvol4type3, bandvol4type4, bandvol4type5, bandvol5type1, bandvol5type2, bandvol5type3, bandvol5type4, bandvol5type5, bandvol6type1, bandvol6type2, bandvol6type3, bandvol6type4, bandvol6type5, bandvol7type1, bandvol7type2, bandvol7type3, bandvol7type4, bandvol7type5, bandvol8type1, bandvol8type2, bandvol8type3, bandvol8type4, bandvol8type5, bandvol9type1, bandvol9type2, bandvol9type3, bandvol9type4, bandvol9type5, bandvol10type1, bandvol10type2, bandvol10type3, bandvol10type4, bandvol10type5, bandinclfreemins, banincindiscount, bandbasetype1, bandbasetype2, bandbasetype3, bandbasetype4, bandbasetype5, bandmarkup, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.band ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	b := Band{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&b.Banduniquesys, &b.Bandid, &b.Bandchrgopt, &b.Bandfixed, &b.Bandrebate, &b.Bandfreemins, &b.Bandvatrate, &b.Bandvol1type1, &b.Bandvol1type2, &b.Bandvol1type3, &b.Bandvol1type4, &b.Bandvol1type5, &b.Bandvol2type1, &b.Bandvol2type2, &b.Bandvol2type3, &b.Bandvol2type4, &b.Bandvol2type5, &b.Bandvol3type1, &b.Bandvol3type2, &b.Bandvol3type3, &b.Bandvol3type4, &b.Bandvol3type5, &b.Bandvol4type1, &b.Bandvol4type2, &b.Bandvol4type3, &b.Bandvol4type4, &b.Bandvol4type5, &b.Bandvol5type1, &b.Bandvol5type2, &b.Bandvol5type3, &b.Bandvol5type4, &b.Bandvol5type5, &b.Bandvol6type1, &b.Bandvol6type2, &b.Bandvol6type3, &b.Bandvol6type4, &b.Bandvol6type5, &b.Bandvol7type1, &b.Bandvol7type2, &b.Bandvol7type3, &b.Bandvol7type4, &b.Bandvol7type5, &b.Bandvol8type1, &b.Bandvol8type2, &b.Bandvol8type3, &b.Bandvol8type4, &b.Bandvol8type5, &b.Bandvol9type1, &b.Bandvol9type2, &b.Bandvol9type3, &b.Bandvol9type4, &b.Bandvol9type5, &b.Bandvol10type1, &b.Bandvol10type2, &b.Bandvol10type3, &b.Bandvol10type4, &b.Bandvol10type5, &b.Bandinclfreemins, &b.Banincindiscount, &b.Bandbasetype1, &b.Bandbasetype2, &b.Bandbasetype3, &b.Bandbasetype4, &b.Bandbasetype5, &b.Bandmarkup, &b.EquinoxPrn, &b.EquinoxLrn, &b.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
