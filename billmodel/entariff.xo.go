// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Entariff represents a row from 'equinox.entariff'.
type Entariff struct {
	Enversion       sql.NullInt64   `json:"enversion"`       // enversion
	Enrevenuegas    sql.NullFloat64 `json:"enrevenuegas"`    // enrevenuegas
	Enrevenueelec   sql.NullFloat64 `json:"enrevenueelec"`   // enrevenueelec
	Enrevenuee7     sql.NullFloat64 `json:"enrevenuee7"`     // enrevenuee7
	Enscgasddse     sql.NullFloat64 `json:"enscgasddse"`     // enscgasddse
	Enscgasddho     sql.NullFloat64 `json:"enscgasddho"`     // enscgasddho
	Enscgasddgo     sql.NullFloat64 `json:"enscgasddgo"`     // enscgasddgo
	Enscgasddgoh    sql.NullFloat64 `json:"enscgasddgoh"`    // enscgasddgoh
	Enscgascase     sql.NullFloat64 `json:"enscgascase"`     // enscgascase
	Enscgascaho     sql.NullFloat64 `json:"enscgascaho"`     // enscgascaho
	Enscgascago     sql.NullFloat64 `json:"enscgascago"`     // enscgascago
	Enscgascagoh    sql.NullFloat64 `json:"enscgascagoh"`    // enscgascagoh
	Enscgasppse     sql.NullFloat64 `json:"enscgasppse"`     // enscgasppse
	Enscgasppho     sql.NullFloat64 `json:"enscgasppho"`     // enscgasppho
	Enscgasppgo     sql.NullFloat64 `json:"enscgasppgo"`     // enscgasppgo
	Enscgasppgoh    sql.NullFloat64 `json:"enscgasppgoh"`    // enscgasppgoh
	Enunitgasddse   sql.NullFloat64 `json:"enunitgasddse"`   // enunitgasddse
	Enunitgasddho   sql.NullFloat64 `json:"enunitgasddho"`   // enunitgasddho
	Enunitgasddgo   sql.NullFloat64 `json:"enunitgasddgo"`   // enunitgasddgo
	Enunitgasddgoh  sql.NullFloat64 `json:"enunitgasddgoh"`  // enunitgasddgoh
	Enunitgascase   sql.NullFloat64 `json:"enunitgascase"`   // enunitgascase
	Enunitgascaho   sql.NullFloat64 `json:"enunitgascaho"`   // enunitgascaho
	Enunitgascago   sql.NullFloat64 `json:"enunitgascago"`   // enunitgascago
	Enunitgascagoh  sql.NullFloat64 `json:"enunitgascagoh"`  // enunitgascagoh
	Enunitgasppse   sql.NullFloat64 `json:"enunitgasppse"`   // enunitgasppse
	Enunitgasppho   sql.NullFloat64 `json:"enunitgasppho"`   // enunitgasppho
	Enunitgasppgo   sql.NullFloat64 `json:"enunitgasppgo"`   // enunitgasppgo
	Enunitgasppgoh  sql.NullFloat64 `json:"enunitgasppgoh"`  // enunitgasppgoh
	Enscelecddse    sql.NullFloat64 `json:"enscelecddse"`    // enscelecddse
	Enscelecddho    sql.NullFloat64 `json:"enscelecddho"`    // enscelecddho
	Enscelecddgo    sql.NullFloat64 `json:"enscelecddgo"`    // enscelecddgo
	Enscelecddgoh   sql.NullFloat64 `json:"enscelecddgoh"`   // enscelecddgoh
	Ensceleccase    sql.NullFloat64 `json:"ensceleccase"`    // ensceleccase
	Ensceleccaho    sql.NullFloat64 `json:"ensceleccaho"`    // ensceleccaho
	Ensceleccago    sql.NullFloat64 `json:"ensceleccago"`    // ensceleccago
	Ensceleccagoh   sql.NullFloat64 `json:"ensceleccagoh"`   // ensceleccagoh
	Enscelecppse    sql.NullFloat64 `json:"enscelecppse"`    // enscelecppse
	Enscelecppho    sql.NullFloat64 `json:"enscelecppho"`    // enscelecppho
	Enscelecppgo    sql.NullFloat64 `json:"enscelecppgo"`    // enscelecppgo
	Enscelecppgoh   sql.NullFloat64 `json:"enscelecppgoh"`   // enscelecppgoh
	Enunitelecddse  sql.NullFloat64 `json:"enunitelecddse"`  // enunitelecddse
	Enunitelecddho  sql.NullFloat64 `json:"enunitelecddho"`  // enunitelecddho
	Enunitelecddgo  sql.NullFloat64 `json:"enunitelecddgo"`  // enunitelecddgo
	Enunitelecddgoh sql.NullFloat64 `json:"enunitelecddgoh"` // enunitelecddgoh
	Enuniteleccase  sql.NullFloat64 `json:"enuniteleccase"`  // enuniteleccase
	Enuniteleccaho  sql.NullFloat64 `json:"enuniteleccaho"`  // enuniteleccaho
	Enuniteleccago  sql.NullFloat64 `json:"enuniteleccago"`  // enuniteleccago
	Enuniteleccagoh sql.NullFloat64 `json:"enuniteleccagoh"` // enuniteleccagoh
	Enunitelecppse  sql.NullFloat64 `json:"enunitelecppse"`  // enunitelecppse
	Enunitelecppho  sql.NullFloat64 `json:"enunitelecppho"`  // enunitelecppho
	Enunitelecppgo  sql.NullFloat64 `json:"enunitelecppgo"`  // enunitelecppgo
	Enunitelecppgoh sql.NullFloat64 `json:"enunitelecppgoh"` // enunitelecppgoh
	Ensce7ddse      sql.NullFloat64 `json:"ensce7ddse"`      // ensce7ddse
	Ensce7ddho      sql.NullFloat64 `json:"ensce7ddho"`      // ensce7ddho
	Ensce7ddgo      sql.NullFloat64 `json:"ensce7ddgo"`      // ensce7ddgo
	Ensce7ddgoh     sql.NullFloat64 `json:"ensce7ddgoh"`     // ensce7ddgoh
	Ensce7case      sql.NullFloat64 `json:"ensce7case"`      // ensce7case
	Ensce7caho      sql.NullFloat64 `json:"ensce7caho"`      // ensce7caho
	Ensce7cago      sql.NullFloat64 `json:"ensce7cago"`      // ensce7cago
	Ensce7cagoh     sql.NullFloat64 `json:"ensce7cagoh"`     // ensce7cagoh
	Ensce7ppse      sql.NullFloat64 `json:"ensce7ppse"`      // ensce7ppse
	Ensce7ppho      sql.NullFloat64 `json:"ensce7ppho"`      // ensce7ppho
	Ensce7ppgo      sql.NullFloat64 `json:"ensce7ppgo"`      // ensce7ppgo
	Ensce7ppgoh     sql.NullFloat64 `json:"ensce7ppgoh"`     // ensce7ppgoh
	Enunite7ddse    sql.NullFloat64 `json:"enunite7ddse"`    // enunite7ddse
	Enunite7ddho    sql.NullFloat64 `json:"enunite7ddho"`    // enunite7ddho
	Enunite7ddgo    sql.NullFloat64 `json:"enunite7ddgo"`    // enunite7ddgo
	Enunite7ddgoh   sql.NullFloat64 `json:"enunite7ddgoh"`   // enunite7ddgoh
	Enunite7case    sql.NullFloat64 `json:"enunite7case"`    // enunite7case
	Enunite7caho    sql.NullFloat64 `json:"enunite7caho"`    // enunite7caho
	Enunite7cago    sql.NullFloat64 `json:"enunite7cago"`    // enunite7cago
	Enunite7cagoh   sql.NullFloat64 `json:"enunite7cagoh"`   // enunite7cagoh
	Enunite7ppse    sql.NullFloat64 `json:"enunite7ppse"`    // enunite7ppse
	Enunite7ppho    sql.NullFloat64 `json:"enunite7ppho"`    // enunite7ppho
	Enunite7ppgo    sql.NullFloat64 `json:"enunite7ppgo"`    // enunite7ppgo
	Enunite7ppgoh   sql.NullFloat64 `json:"enunite7ppgoh"`   // enunite7ppgoh
	Ennunite7ddse   sql.NullFloat64 `json:"ennunite7ddse"`   // ennunite7ddse
	Ennunite7ddho   sql.NullFloat64 `json:"ennunite7ddho"`   // ennunite7ddho
	Ennunite7ddgo   sql.NullFloat64 `json:"ennunite7ddgo"`   // ennunite7ddgo
	Ennunite7ddgoh  sql.NullFloat64 `json:"ennunite7ddgoh"`  // ennunite7ddgoh
	Ennunite7case   sql.NullFloat64 `json:"ennunite7case"`   // ennunite7case
	Ennunite7caho   sql.NullFloat64 `json:"ennunite7caho"`   // ennunite7caho
	Ennunite7cago   sql.NullFloat64 `json:"ennunite7cago"`   // ennunite7cago
	Ennunite7cagoh  sql.NullFloat64 `json:"ennunite7cagoh"`  // ennunite7cagoh
	Ennunite7ppse   sql.NullFloat64 `json:"ennunite7ppse"`   // ennunite7ppse
	Ennunite7ppho   sql.NullFloat64 `json:"ennunite7ppho"`   // ennunite7ppho
	Ennunite7ppgo   sql.NullFloat64 `json:"ennunite7ppgo"`   // ennunite7ppgo
	Ennunite7ppgoh  sql.NullFloat64 `json:"ennunite7ppgoh"`  // ennunite7ppgoh
	Ensebill        sql.NullFloat64 `json:"ensebill"`        // ensebill
	Enhobill        sql.NullFloat64 `json:"enhobill"`        // enhobill
	Engobill        sql.NullFloat64 `json:"engobill"`        // engobill
	Enesebill       sql.NullFloat64 `json:"enesebill"`       // enesebill
	Enehobill       sql.NullFloat64 `json:"enehobill"`       // enehobill
	Enegobill       sql.NullFloat64 `json:"enegobill"`       // enegobill
	Ene7sebill      sql.NullFloat64 `json:"ene7sebill"`      // ene7sebill
	Ene7hobill      sql.NullFloat64 `json:"ene7hobill"`      // ene7hobill
	Ene7gobill      sql.NullFloat64 `json:"ene7gobill"`      // ene7gobill
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec
}

func AllEntariff(db XODB, callback func(x Entariff) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`enversion, enrevenuegas, enrevenueelec, enrevenuee7, enscgasddse, enscgasddho, enscgasddgo, enscgasddgoh, enscgascase, enscgascaho, enscgascago, enscgascagoh, enscgasppse, enscgasppho, enscgasppgo, enscgasppgoh, enunitgasddse, enunitgasddho, enunitgasddgo, enunitgasddgoh, enunitgascase, enunitgascaho, enunitgascago, enunitgascagoh, enunitgasppse, enunitgasppho, enunitgasppgo, enunitgasppgoh, enscelecddse, enscelecddho, enscelecddgo, enscelecddgoh, ensceleccase, ensceleccaho, ensceleccago, ensceleccagoh, enscelecppse, enscelecppho, enscelecppgo, enscelecppgoh, enunitelecddse, enunitelecddho, enunitelecddgo, enunitelecddgoh, enuniteleccase, enuniteleccaho, enuniteleccago, enuniteleccagoh, enunitelecppse, enunitelecppho, enunitelecppgo, enunitelecppgoh, ensce7ddse, ensce7ddho, ensce7ddgo, ensce7ddgoh, ensce7case, ensce7caho, ensce7cago, ensce7cagoh, ensce7ppse, ensce7ppho, ensce7ppgo, ensce7ppgoh, enunite7ddse, enunite7ddho, enunite7ddgo, enunite7ddgoh, enunite7case, enunite7caho, enunite7cago, enunite7cagoh, enunite7ppse, enunite7ppho, enunite7ppgo, enunite7ppgoh, ennunite7ddse, ennunite7ddho, ennunite7ddgo, ennunite7ddgoh, ennunite7case, ennunite7caho, ennunite7cago, ennunite7cagoh, ennunite7ppse, ennunite7ppho, ennunite7ppgo, ennunite7ppgoh, ensebill, enhobill, engobill, enesebill, enehobill, enegobill, ene7sebill, ene7hobill, ene7gobill, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.entariff `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		e := Entariff{}

		// scan
		err = q.Scan(&e.Enversion, &e.Enrevenuegas, &e.Enrevenueelec, &e.Enrevenuee7, &e.Enscgasddse, &e.Enscgasddho, &e.Enscgasddgo, &e.Enscgasddgoh, &e.Enscgascase, &e.Enscgascaho, &e.Enscgascago, &e.Enscgascagoh, &e.Enscgasppse, &e.Enscgasppho, &e.Enscgasppgo, &e.Enscgasppgoh, &e.Enunitgasddse, &e.Enunitgasddho, &e.Enunitgasddgo, &e.Enunitgasddgoh, &e.Enunitgascase, &e.Enunitgascaho, &e.Enunitgascago, &e.Enunitgascagoh, &e.Enunitgasppse, &e.Enunitgasppho, &e.Enunitgasppgo, &e.Enunitgasppgoh, &e.Enscelecddse, &e.Enscelecddho, &e.Enscelecddgo, &e.Enscelecddgoh, &e.Ensceleccase, &e.Ensceleccaho, &e.Ensceleccago, &e.Ensceleccagoh, &e.Enscelecppse, &e.Enscelecppho, &e.Enscelecppgo, &e.Enscelecppgoh, &e.Enunitelecddse, &e.Enunitelecddho, &e.Enunitelecddgo, &e.Enunitelecddgoh, &e.Enuniteleccase, &e.Enuniteleccaho, &e.Enuniteleccago, &e.Enuniteleccagoh, &e.Enunitelecppse, &e.Enunitelecppho, &e.Enunitelecppgo, &e.Enunitelecppgoh, &e.Ensce7ddse, &e.Ensce7ddho, &e.Ensce7ddgo, &e.Ensce7ddgoh, &e.Ensce7case, &e.Ensce7caho, &e.Ensce7cago, &e.Ensce7cagoh, &e.Ensce7ppse, &e.Ensce7ppho, &e.Ensce7ppgo, &e.Ensce7ppgoh, &e.Enunite7ddse, &e.Enunite7ddho, &e.Enunite7ddgo, &e.Enunite7ddgoh, &e.Enunite7case, &e.Enunite7caho, &e.Enunite7cago, &e.Enunite7cagoh, &e.Enunite7ppse, &e.Enunite7ppho, &e.Enunite7ppgo, &e.Enunite7ppgoh, &e.Ennunite7ddse, &e.Ennunite7ddho, &e.Ennunite7ddgo, &e.Ennunite7ddgoh, &e.Ennunite7case, &e.Ennunite7caho, &e.Ennunite7cago, &e.Ennunite7cagoh, &e.Ennunite7ppse, &e.Ennunite7ppho, &e.Ennunite7ppgo, &e.Ennunite7ppgoh, &e.Ensebill, &e.Enhobill, &e.Engobill, &e.Enesebill, &e.Enehobill, &e.Enegobill, &e.Ene7sebill, &e.Ene7hobill, &e.Ene7gobill, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(e) {
			return nil
		}
	}

	return nil
}

// EntariffByEquinoxLrn retrieves a row from 'equinox.entariff' as a Entariff.
//
// Generated from index 'entariff_pkey'.
func EntariffByEquinoxLrn(db XODB, equinoxLrn int64) (*Entariff, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`enversion, enrevenuegas, enrevenueelec, enrevenuee7, enscgasddse, enscgasddho, enscgasddgo, enscgasddgoh, enscgascase, enscgascaho, enscgascago, enscgascagoh, enscgasppse, enscgasppho, enscgasppgo, enscgasppgoh, enunitgasddse, enunitgasddho, enunitgasddgo, enunitgasddgoh, enunitgascase, enunitgascaho, enunitgascago, enunitgascagoh, enunitgasppse, enunitgasppho, enunitgasppgo, enunitgasppgoh, enscelecddse, enscelecddho, enscelecddgo, enscelecddgoh, ensceleccase, ensceleccaho, ensceleccago, ensceleccagoh, enscelecppse, enscelecppho, enscelecppgo, enscelecppgoh, enunitelecddse, enunitelecddho, enunitelecddgo, enunitelecddgoh, enuniteleccase, enuniteleccaho, enuniteleccago, enuniteleccagoh, enunitelecppse, enunitelecppho, enunitelecppgo, enunitelecppgoh, ensce7ddse, ensce7ddho, ensce7ddgo, ensce7ddgoh, ensce7case, ensce7caho, ensce7cago, ensce7cagoh, ensce7ppse, ensce7ppho, ensce7ppgo, ensce7ppgoh, enunite7ddse, enunite7ddho, enunite7ddgo, enunite7ddgoh, enunite7case, enunite7caho, enunite7cago, enunite7cagoh, enunite7ppse, enunite7ppho, enunite7ppgo, enunite7ppgoh, ennunite7ddse, ennunite7ddho, ennunite7ddgo, ennunite7ddgoh, ennunite7case, ennunite7caho, ennunite7cago, ennunite7cagoh, ennunite7ppse, ennunite7ppho, ennunite7ppgo, ennunite7ppgoh, ensebill, enhobill, engobill, enesebill, enehobill, enegobill, ene7sebill, ene7hobill, ene7gobill, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.entariff ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Entariff{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Enversion, &e.Enrevenuegas, &e.Enrevenueelec, &e.Enrevenuee7, &e.Enscgasddse, &e.Enscgasddho, &e.Enscgasddgo, &e.Enscgasddgoh, &e.Enscgascase, &e.Enscgascaho, &e.Enscgascago, &e.Enscgascagoh, &e.Enscgasppse, &e.Enscgasppho, &e.Enscgasppgo, &e.Enscgasppgoh, &e.Enunitgasddse, &e.Enunitgasddho, &e.Enunitgasddgo, &e.Enunitgasddgoh, &e.Enunitgascase, &e.Enunitgascaho, &e.Enunitgascago, &e.Enunitgascagoh, &e.Enunitgasppse, &e.Enunitgasppho, &e.Enunitgasppgo, &e.Enunitgasppgoh, &e.Enscelecddse, &e.Enscelecddho, &e.Enscelecddgo, &e.Enscelecddgoh, &e.Ensceleccase, &e.Ensceleccaho, &e.Ensceleccago, &e.Ensceleccagoh, &e.Enscelecppse, &e.Enscelecppho, &e.Enscelecppgo, &e.Enscelecppgoh, &e.Enunitelecddse, &e.Enunitelecddho, &e.Enunitelecddgo, &e.Enunitelecddgoh, &e.Enuniteleccase, &e.Enuniteleccaho, &e.Enuniteleccago, &e.Enuniteleccagoh, &e.Enunitelecppse, &e.Enunitelecppho, &e.Enunitelecppgo, &e.Enunitelecppgoh, &e.Ensce7ddse, &e.Ensce7ddho, &e.Ensce7ddgo, &e.Ensce7ddgoh, &e.Ensce7case, &e.Ensce7caho, &e.Ensce7cago, &e.Ensce7cagoh, &e.Ensce7ppse, &e.Ensce7ppho, &e.Ensce7ppgo, &e.Ensce7ppgoh, &e.Enunite7ddse, &e.Enunite7ddho, &e.Enunite7ddgo, &e.Enunite7ddgoh, &e.Enunite7case, &e.Enunite7caho, &e.Enunite7cago, &e.Enunite7cagoh, &e.Enunite7ppse, &e.Enunite7ppho, &e.Enunite7ppgo, &e.Enunite7ppgoh, &e.Ennunite7ddse, &e.Ennunite7ddho, &e.Ennunite7ddgo, &e.Ennunite7ddgoh, &e.Ennunite7case, &e.Ennunite7caho, &e.Ennunite7cago, &e.Ennunite7cagoh, &e.Ennunite7ppse, &e.Ennunite7ppho, &e.Ennunite7ppgo, &e.Ennunite7ppgoh, &e.Ensebill, &e.Enhobill, &e.Engobill, &e.Enesebill, &e.Enehobill, &e.Enegobill, &e.Ene7sebill, &e.Ene7hobill, &e.Ene7gobill, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
