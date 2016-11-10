// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Gastran represents a row from 'equinox.gastrans'.
type Gastran struct {
	Gtrandatecreated pq.NullTime     `json:"gtrandatecreated"` // gtrandatecreated
	Gtranfiletype    sql.NullString  `json:"gtranfiletype"`    // gtranfiletype
	Gtransource      sql.NullString  `json:"gtransource"`      // gtransource
	Gtranfileproduce sql.NullString  `json:"gtranfileproduce"` // gtranfileproduce
	Gtrantype        sql.NullString  `json:"gtrantype"`        // gtrantype
	Gtranstatus      sql.NullString  `json:"gtranstatus"`      // gtranstatus
	Gtranref         sql.NullString  `json:"gtranref"`         // gtranref
	Gtransc          sql.NullString  `json:"gtransc"`          // gtransc
	Gtranregref      sql.NullString  `json:"gtranregref"`      // gtranregref
	Gtrancf          sql.NullFloat64 `json:"gtrancf"`          // gtrancf
	Gtranlastinspd   pq.NullTime     `json:"gtranlastinspd"`   // gtranlastinspd
	Gtranappointd    pq.NullTime     `json:"gtranappointd"`    // gtranappointd
	Gtranrolecodeins sql.NullString  `json:"gtranrolecodeins"` // gtranrolecodeins
	Gtranmamid       sql.NullString  `json:"gtranmamid"`       // gtranmamid
	Gtranourstatus   sql.NullString  `json:"gtranourstatus"`   // gtranourstatus
	Gtransenttoxoser pq.NullTime     `json:"gtransenttoxoser"` // gtransenttoxoser
	Gtranxoserveresp sql.NullString  `json:"gtranxoserveresp"` // gtranxoserveresp
	Gtranxoserverdat pq.NullTime     `json:"gtranxoserverdat"` // gtranxoserverdat
	Gtranrecordcount sql.NullInt64   `json:"gtranrecordcount"` // gtranrecordcount
	Gtranremmsn      sql.NullString  `json:"gtranremmsn"`      // gtranremmsn
	Gtranremlc       sql.NullString  `json:"gtranremlc"`       // gtranremlc
	Gtranremlcword   sql.NullString  `json:"gtranremlcword"`   // gtranremlcword
	Gtranremlcnote   sql.NullInt64   `json:"gtranremlcnote"`   // gtranremlcnote
	Gtranremastatusc sql.NullString  `json:"gtranremastatusc"` // gtranremastatusc
	Gtranremmtypec   sql.NullString  `json:"gtranremmtypec"`   // gtranremmtypec
	Gtranremmtypew   sql.NullString  `json:"gtranremmtypew"`   // gtranremmtypew
	Gtranremmechc    sql.NullString  `json:"gtranremmechc"`    // gtranremmechc
	Gtranremmechw    sql.NullString  `json:"gtranremmechw"`    // gtranremmechw
	Gtranremoamidate pq.NullTime     `json:"gtranremoamidate"` // gtranremoamidate
	Gtranremrolecode sql.NullString  `json:"gtranremrolecode"` // gtranremrolecode
	Gtranremmeascap  sql.NullFloat64 `json:"gtranremmeascap"`  // gtranremmeascap
	Gtranremmdials   sql.NullInt64   `json:"gtranremmdials"`   // gtranremmdials
	Gtranremiorm     sql.NullString  `json:"gtranremiorm"`     // gtranremiorm
	Gtranremrf       sql.NullFloat64 `json:"gtranremrf"`       // gtranremrf
	Gtranremreadd    pq.NullTime     `json:"gtranremreadd"`    // gtranremreadd
	Gtranremreadttz  sql.NullInt64   `json:"gtranremreadttz"`  // gtranremreadttz
	Gtranremreading  sql.NullString  `json:"gtranremreading"`  // gtranremreading
	Gtranremtobill   sql.NullInt64   `json:"gtranremtobill"`   // gtranremtobill
	Gtranrembilltox  sql.NullInt64   `json:"gtranrembilltox"`  // gtranrembilltox
	Gtranrembilld    pq.NullTime     `json:"gtranrembilld"`    // gtranrembilld
	Gtranrembillttz  sql.NullInt64   `json:"gtranrembillttz"`  // gtranrembillttz
	Gtranrembillread sql.NullString  `json:"gtranrembillread"` // gtranrembillread
	Gtraninsmake     sql.NullString  `json:"gtraninsmake"`     // gtraninsmake
	Gtraninsmodel    sql.NullString  `json:"gtraninsmodel"`    // gtraninsmodel
	Gtraninsyman     sql.NullString  `json:"gtraninsyman"`     // gtraninsyman
	Gtraninsmsn      sql.NullString  `json:"gtraninsmsn"`      // gtraninsmsn
	Gtraninslc       sql.NullString  `json:"gtraninslc"`       // gtraninslc
	Gtraninslcword   sql.NullString  `json:"gtraninslcword"`   // gtraninslcword
	Gtraninslcnote   sql.NullInt64   `json:"gtraninslcnote"`   // gtraninslcnote
	Gtraninsstatusc  sql.NullString  `json:"gtraninsstatusc"`  // gtraninsstatusc
	Gtraninsmettypec sql.NullString  `json:"gtraninsmettypec"` // gtraninsmettypec
	Gtraninsmetw     sql.NullString  `json:"gtraninsmetw"`     // gtraninsmetw
	Gtraninsmetmechc sql.NullString  `json:"gtraninsmetmechc"` // gtraninsmetmechc
	Gtraninsmetmechw sql.NullString  `json:"gtraninsmetmechw"` // gtraninsmetmechw
	Gtraninsoamidate pq.NullTime     `json:"gtraninsoamidate"` // gtraninsoamidate
	Gtraninsrolecode sql.NullString  `json:"gtraninsrolecode"` // gtraninsrolecode
	Gtraninsmeascap  sql.NullFloat64 `json:"gtraninsmeascap"`  // gtraninsmeascap
	Gtraninsmdials   sql.NullInt64   `json:"gtraninsmdials"`   // gtraninsmdials
	Gtraninsiorm     sql.NullString  `json:"gtraninsiorm"`     // gtraninsiorm
	Gtraninsrf       sql.NullFloat64 `json:"gtraninsrf"`       // gtraninsrf
	Gtraninsreadd    pq.NullTime     `json:"gtraninsreadd"`    // gtraninsreadd
	Gtraninsreadttz  sql.NullInt64   `json:"gtraninsreadttz"`  // gtraninsreadttz
	Gtraninsreading  sql.NullString  `json:"gtraninsreading"`  // gtraninsreading
	Gtraninstobill   sql.NullInt64   `json:"gtraninstobill"`   // gtraninstobill
	Gtraninsbilld    pq.NullTime     `json:"gtraninsbilld"`    // gtraninsbilld
	Gtraninsbillttz  sql.NullInt64   `json:"gtraninsbillttz"`  // gtraninsbillttz
	Gtraninsbillread sql.NullString  `json:"gtraninsbillread"` // gtraninsbillread
	Gtranrolecoder   sql.NullString  `json:"gtranrolecoder"`   // gtranrolecoder
	Gtraninsbilltox  sql.NullInt64   `json:"gtraninsbilltox"`  // gtraninsbilltox
	Gjrsfile         sql.NullString  `json:"gjrsfile"`         // gjrsfile
	Gtransc3         sql.NullString  `json:"gtransc3"`         // gtransc3
	Gtransn1         sql.NullInt64   `json:"gtransn1"`         // gtransn1
	Gtransn2         sql.NullInt64   `json:"gtransn2"`         // gtransn2
	Gtransn3         sql.NullInt64   `json:"gtransn3"`         // gtransn3
	Gtransd1         pq.NullTime     `json:"gtransd1"`         // gtransd1
	Gtransd2         pq.NullTime     `json:"gtransd2"`         // gtransd2
	Gtransd3         pq.NullTime     `json:"gtransd3"`         // gtransd3
	Gtransc4         sql.NullString  `json:"gtransc4"`         // gtransc4
	Gtransc5         sql.NullString  `json:"gtransc5"`         // gtransc5
	Gtransoutput     sql.NullString  `json:"gtransoutput"`     // gtransoutput
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// GastranByEquinoxLrn retrieves a row from 'equinox.gastrans' as a Gastran.
//
// Generated from index 'gastrans_pkey'.
func GastranByEquinoxLrn(db XODB, equinoxLrn int64) (*Gastran, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gtrandatecreated, gtranfiletype, gtransource, gtranfileproduce, gtrantype, gtranstatus, gtranref, gtransc, gtranregref, gtrancf, gtranlastinspd, gtranappointd, gtranrolecodeins, gtranmamid, gtranourstatus, gtransenttoxoser, gtranxoserveresp, gtranxoserverdat, gtranrecordcount, gtranremmsn, gtranremlc, gtranremlcword, gtranremlcnote, gtranremastatusc, gtranremmtypec, gtranremmtypew, gtranremmechc, gtranremmechw, gtranremoamidate, gtranremrolecode, gtranremmeascap, gtranremmdials, gtranremiorm, gtranremrf, gtranremreadd, gtranremreadttz, gtranremreading, gtranremtobill, gtranrembilltox, gtranrembilld, gtranrembillttz, gtranrembillread, gtraninsmake, gtraninsmodel, gtraninsyman, gtraninsmsn, gtraninslc, gtraninslcword, gtraninslcnote, gtraninsstatusc, gtraninsmettypec, gtraninsmetw, gtraninsmetmechc, gtraninsmetmechw, gtraninsoamidate, gtraninsrolecode, gtraninsmeascap, gtraninsmdials, gtraninsiorm, gtraninsrf, gtraninsreadd, gtraninsreadttz, gtraninsreading, gtraninstobill, gtraninsbilld, gtraninsbillttz, gtraninsbillread, gtranrolecoder, gtraninsbilltox, gjrsfile, gtransc3, gtransn1, gtransn2, gtransn3, gtransd1, gtransd2, gtransd3, gtransc4, gtransc5, gtransoutput, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.gastrans ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	g := Gastran{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&g.Gtrandatecreated, &g.Gtranfiletype, &g.Gtransource, &g.Gtranfileproduce, &g.Gtrantype, &g.Gtranstatus, &g.Gtranref, &g.Gtransc, &g.Gtranregref, &g.Gtrancf, &g.Gtranlastinspd, &g.Gtranappointd, &g.Gtranrolecodeins, &g.Gtranmamid, &g.Gtranourstatus, &g.Gtransenttoxoser, &g.Gtranxoserveresp, &g.Gtranxoserverdat, &g.Gtranrecordcount, &g.Gtranremmsn, &g.Gtranremlc, &g.Gtranremlcword, &g.Gtranremlcnote, &g.Gtranremastatusc, &g.Gtranremmtypec, &g.Gtranremmtypew, &g.Gtranremmechc, &g.Gtranremmechw, &g.Gtranremoamidate, &g.Gtranremrolecode, &g.Gtranremmeascap, &g.Gtranremmdials, &g.Gtranremiorm, &g.Gtranremrf, &g.Gtranremreadd, &g.Gtranremreadttz, &g.Gtranremreading, &g.Gtranremtobill, &g.Gtranrembilltox, &g.Gtranrembilld, &g.Gtranrembillttz, &g.Gtranrembillread, &g.Gtraninsmake, &g.Gtraninsmodel, &g.Gtraninsyman, &g.Gtraninsmsn, &g.Gtraninslc, &g.Gtraninslcword, &g.Gtraninslcnote, &g.Gtraninsstatusc, &g.Gtraninsmettypec, &g.Gtraninsmetw, &g.Gtraninsmetmechc, &g.Gtraninsmetmechw, &g.Gtraninsoamidate, &g.Gtraninsrolecode, &g.Gtraninsmeascap, &g.Gtraninsmdials, &g.Gtraninsiorm, &g.Gtraninsrf, &g.Gtraninsreadd, &g.Gtraninsreadttz, &g.Gtraninsreading, &g.Gtraninstobill, &g.Gtraninsbilld, &g.Gtraninsbillttz, &g.Gtraninsbillread, &g.Gtranrolecoder, &g.Gtraninsbilltox, &g.Gjrsfile, &g.Gtransc3, &g.Gtransn1, &g.Gtransn2, &g.Gtransn3, &g.Gtransd1, &g.Gtransd2, &g.Gtransd3, &g.Gtransc4, &g.Gtransc5, &g.Gtransoutput, &g.EquinoxPrn, &g.EquinoxLrn, &g.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
