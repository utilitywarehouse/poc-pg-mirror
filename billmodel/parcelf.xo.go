// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Parcelf represents a row from 'equinox.parcelf'.
type Parcelf struct {
	PfUniqueid     sql.NullFloat64 `json:"pf_uniqueid"`     // pf_uniqueid
	PfAddress1     sql.NullString  `json:"pf_address1"`     // pf_address1
	PfAddress2     sql.NullString  `json:"pf_address2"`     // pf_address2
	PfAddress3     sql.NullString  `json:"pf_address3"`     // pf_address3
	PfAddress4     sql.NullString  `json:"pf_address4"`     // pf_address4
	PfPostcode     sql.NullString  `json:"pf_postcode"`     // pf_postcode
	PfBuildingname sql.NullString  `json:"pf_buildingname"` // pf_buildingname
	PfBuildingno   sql.NullString  `json:"pf_buildingno"`   // pf_buildingno
	PfStreet       sql.NullString  `json:"pf_street"`       // pf_street
	PfLocale       sql.NullString  `json:"pf_locale"`       // pf_locale
	PfPosttown     sql.NullString  `json:"pf_posttown"`     // pf_posttown
	PfCounty       sql.NullString  `json:"pf_county"`       // pf_county
	PfPacks        sql.NullInt64   `json:"pf_packs"`        // pf_packs
	PfName         sql.NullString  `json:"pf_name"`         // pf_name
	PfAccountname  sql.NullString  `json:"pf_accountname"`  // pf_accountname
	PfAccountno    sql.NullString  `json:"pf_accountno"`    // pf_accountno
	PfPrinted      sql.NullInt64   `json:"pf_printed"`      // pf_printed
	PfEntereddate  pq.NullTime     `json:"pf_entereddate"`  // pf_entereddate
	PfSource       sql.NullString  `json:"pf_source"`       // pf_source
	PfCreatefile   sql.NullInt64   `json:"pf_createfile"`   // pf_createfile
	PfSparec1      sql.NullString  `json:"pf_sparec1"`      // pf_sparec1
	PfSparec2      sql.NullString  `json:"pf_sparec2"`      // pf_sparec2
	PfSpared1      pq.NullTime     `json:"pf_spared1"`      // pf_spared1
	PfSpared2      pq.NullTime     `json:"pf_spared2"`      // pf_spared2
	PfSparen1      sql.NullFloat64 `json:"pf_sparen1"`      // pf_sparen1
	PfSparen2      sql.NullFloat64 `json:"pf_sparen2"`      // pf_sparen2
	PfSparel1      sql.NullInt64   `json:"pf_sparel1"`      // pf_sparel1
	PfSparel2      sql.NullInt64   `json:"pf_sparel2"`      // pf_sparel2
	PfConsignid    sql.NullString  `json:"pf_consignid"`    // pf_consignid
	PfDeliverytype sql.NullString  `json:"pf_deliverytype"` // pf_deliverytype
	EquinoxLrn     int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec     sql.NullInt64   `json:"equinox_sec"`     // equinox_sec
}

// ParcelfByEquinoxLrn retrieves a row from 'equinox.parcelf' as a Parcelf.
//
// Generated from index 'parcelf_pkey'.
func ParcelfByEquinoxLrn(db XODB, equinoxLrn int64) (*Parcelf, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pf_uniqueid, pf_address1, pf_address2, pf_address3, pf_address4, pf_postcode, pf_buildingname, pf_buildingno, pf_street, pf_locale, pf_posttown, pf_county, pf_packs, pf_name, pf_accountname, pf_accountno, pf_printed, pf_entereddate, pf_source, pf_createfile, pf_sparec1, pf_sparec2, pf_spared1, pf_spared2, pf_sparen1, pf_sparen2, pf_sparel1, pf_sparel2, pf_consignid, pf_deliverytype, equinox_lrn, equinox_sec ` +
		`FROM equinox.parcelf ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Parcelf{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.PfUniqueid, &p.PfAddress1, &p.PfAddress2, &p.PfAddress3, &p.PfAddress4, &p.PfPostcode, &p.PfBuildingname, &p.PfBuildingno, &p.PfStreet, &p.PfLocale, &p.PfPosttown, &p.PfCounty, &p.PfPacks, &p.PfName, &p.PfAccountname, &p.PfAccountno, &p.PfPrinted, &p.PfEntereddate, &p.PfSource, &p.PfCreatefile, &p.PfSparec1, &p.PfSparec2, &p.PfSpared1, &p.PfSpared2, &p.PfSparen1, &p.PfSparen2, &p.PfSparel1, &p.PfSparel2, &p.PfConsignid, &p.PfDeliverytype, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
