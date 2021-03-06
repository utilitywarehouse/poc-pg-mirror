// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Prospero represents a row from 'equinox.prospero'.
type Prospero struct {
	Prosaccnumber    sql.NullString  `json:"prosaccnumber"`    // prosaccnumber
	Proscustaccount  sql.NullString  `json:"proscustaccount"`  // proscustaccount
	Prosaccactive    sql.NullInt64   `json:"prosaccactive"`    // prosaccactive
	Prosnameoncard   sql.NullString  `json:"prosnameoncard"`   // prosnameoncard
	Prosdob          pq.NullTime     `json:"prosdob"`          // prosdob
	Prosissued       pq.NullTime     `json:"prosissued"`       // prosissued
	Prosnameoncard2  sql.NullString  `json:"prosnameoncard2"`  // prosnameoncard2
	Prosdob2         pq.NullTime     `json:"prosdob2"`         // prosdob2
	Prosissued2      pq.NullTime     `json:"prosissued2"`      // prosissued2
	Prosnectarcard   sql.NullString  `json:"prosnectarcard"`   // prosnectarcard
	Prosmainsupermkt sql.NullString  `json:"prosmainsupermkt"` // prosmainsupermkt
	Prosminbalance   sql.NullInt64   `json:"prosminbalance"`   // prosminbalance
	Prostopupvalue   sql.NullInt64   `json:"prostopupvalue"`   // prostopupvalue
	Prosfixedmonthly sql.NullInt64   `json:"prosfixedmonthly"` // prosfixedmonthly
	Prosdebitno      sql.NullString  `json:"prosdebitno"`      // prosdebitno
	Prosdebitstart   sql.NullString  `json:"prosdebitstart"`   // prosdebitstart
	Prosdebitexpiry  sql.NullString  `json:"prosdebitexpiry"`  // prosdebitexpiry
	Prosdebitissue   sql.NullString  `json:"prosdebitissue"`   // prosdebitissue
	Prosdebitcv2     sql.NullString  `json:"prosdebitcv2"`     // prosdebitcv2
	Prosdebitcardtyp sql.NullString  `json:"prosdebitcardtyp"` // prosdebitcardtyp
	Prosinitialload  pq.NullTime     `json:"prosinitialload"`  // prosinitialload
	Proslatestbalanc sql.NullFloat64 `json:"proslatestbalanc"` // proslatestbalanc
	Prossavingmtd    sql.NullFloat64 `json:"prossavingmtd"`    // prossavingmtd
	Prossavingytd    sql.NullFloat64 `json:"prossavingytd"`    // prossavingytd
	Prosmobilenumber sql.NullString  `json:"prosmobilenumber"` // prosmobilenumber
	Prosopeningbal   sql.NullFloat64 `json:"prosopeningbal"`   // prosopeningbal
	Prosclosingbal   sql.NullFloat64 `json:"prosclosingbal"`   // prosclosingbal
	Proscardtype     sql.NullString  `json:"proscardtype"`     // proscardtype
	Proscardstatus   sql.NullString  `json:"proscardstatus"`   // proscardstatus
	Proscardstatus2  sql.NullString  `json:"proscardstatus2"`  // proscardstatus2
	Proslastsains    pq.NullTime     `json:"proslastsains"`    // proslastsains
	Proslimit        sql.NullInt64   `json:"proslimit"`        // proslimit
	Proslocation     sql.NullString  `json:"proslocation"`     // proslocation
	Prosexecid       sql.NullString  `json:"prosexecid"`       // prosexecid
	Prosappformno    sql.NullString  `json:"prosappformno"`    // prosappformno
	Prosdontpay      sql.NullString  `json:"prosdontpay"`      // prosdontpay
	Prosbpaid        pq.NullTime     `json:"prosbpaid"`        // prosbpaid
	Prosbpaid2       pq.NullTime     `json:"prosbpaid2"`       // prosbpaid2
	Prospromo        sql.NullInt64   `json:"prospromo"`        // prospromo
	Prosexissued     pq.NullTime     `json:"prosexissued"`     // prosexissued
	Prosexissued2    pq.NullTime     `json:"prosexissued2"`    // prosexissued2
	Prosshopperref   sql.NullString  `json:"prosshopperref"`   // prosshopperref
	Proscard1statmis sql.NullInt64   `json:"proscard1statmis"` // proscard1statmis
	Proscard2statmis sql.NullInt64   `json:"proscard2statmis"` // proscard2statmis
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllProspero(db XODB, callback func(x Prospero) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`prosaccnumber, proscustaccount, prosaccactive, prosnameoncard, prosdob, prosissued, prosnameoncard2, prosdob2, prosissued2, prosnectarcard, prosmainsupermkt, prosminbalance, prostopupvalue, prosfixedmonthly, prosdebitno, prosdebitstart, prosdebitexpiry, prosdebitissue, prosdebitcv2, prosdebitcardtyp, prosinitialload, proslatestbalanc, prossavingmtd, prossavingytd, prosmobilenumber, prosopeningbal, prosclosingbal, proscardtype, proscardstatus, proscardstatus2, proslastsains, proslimit, proslocation, prosexecid, prosappformno, prosdontpay, prosbpaid, prosbpaid2, prospromo, prosexissued, prosexissued2, prosshopperref, proscard1statmis, proscard2statmis, equinox_lrn, equinox_sec ` +
		`FROM equinox.prospero `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := Prospero{}

		// scan
		err = q.Scan(&p.Prosaccnumber, &p.Proscustaccount, &p.Prosaccactive, &p.Prosnameoncard, &p.Prosdob, &p.Prosissued, &p.Prosnameoncard2, &p.Prosdob2, &p.Prosissued2, &p.Prosnectarcard, &p.Prosmainsupermkt, &p.Prosminbalance, &p.Prostopupvalue, &p.Prosfixedmonthly, &p.Prosdebitno, &p.Prosdebitstart, &p.Prosdebitexpiry, &p.Prosdebitissue, &p.Prosdebitcv2, &p.Prosdebitcardtyp, &p.Prosinitialload, &p.Proslatestbalanc, &p.Prossavingmtd, &p.Prossavingytd, &p.Prosmobilenumber, &p.Prosopeningbal, &p.Prosclosingbal, &p.Proscardtype, &p.Proscardstatus, &p.Proscardstatus2, &p.Proslastsains, &p.Proslimit, &p.Proslocation, &p.Prosexecid, &p.Prosappformno, &p.Prosdontpay, &p.Prosbpaid, &p.Prosbpaid2, &p.Prospromo, &p.Prosexissued, &p.Prosexissued2, &p.Prosshopperref, &p.Proscard1statmis, &p.Proscard2statmis, &p.EquinoxLrn, &p.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(p) {
			return nil
		}
	}

	return nil
}

// ProsperoByEquinoxLrn retrieves a row from 'equinox.prospero' as a Prospero.
//
// Generated from index 'prospero_pkey'.
func ProsperoByEquinoxLrn(db XODB, equinoxLrn int64) (*Prospero, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`prosaccnumber, proscustaccount, prosaccactive, prosnameoncard, prosdob, prosissued, prosnameoncard2, prosdob2, prosissued2, prosnectarcard, prosmainsupermkt, prosminbalance, prostopupvalue, prosfixedmonthly, prosdebitno, prosdebitstart, prosdebitexpiry, prosdebitissue, prosdebitcv2, prosdebitcardtyp, prosinitialload, proslatestbalanc, prossavingmtd, prossavingytd, prosmobilenumber, prosopeningbal, prosclosingbal, proscardtype, proscardstatus, proscardstatus2, proslastsains, proslimit, proslocation, prosexecid, prosappformno, prosdontpay, prosbpaid, prosbpaid2, prospromo, prosexissued, prosexissued2, prosshopperref, proscard1statmis, proscard2statmis, equinox_lrn, equinox_sec ` +
		`FROM equinox.prospero ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Prospero{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Prosaccnumber, &p.Proscustaccount, &p.Prosaccactive, &p.Prosnameoncard, &p.Prosdob, &p.Prosissued, &p.Prosnameoncard2, &p.Prosdob2, &p.Prosissued2, &p.Prosnectarcard, &p.Prosmainsupermkt, &p.Prosminbalance, &p.Prostopupvalue, &p.Prosfixedmonthly, &p.Prosdebitno, &p.Prosdebitstart, &p.Prosdebitexpiry, &p.Prosdebitissue, &p.Prosdebitcv2, &p.Prosdebitcardtyp, &p.Prosinitialload, &p.Proslatestbalanc, &p.Prossavingmtd, &p.Prossavingytd, &p.Prosmobilenumber, &p.Prosopeningbal, &p.Prosclosingbal, &p.Proscardtype, &p.Proscardstatus, &p.Proscardstatus2, &p.Proslastsains, &p.Proslimit, &p.Proslocation, &p.Prosexecid, &p.Prosappformno, &p.Prosdontpay, &p.Prosbpaid, &p.Prosbpaid2, &p.Prospromo, &p.Prosexissued, &p.Prosexissued2, &p.Prosshopperref, &p.Proscard1statmis, &p.Proscard2statmis, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
