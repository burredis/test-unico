package feiralivre

import (
	"database/sql"
	"fmt"
)

type FeiraLivreRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) FeiraLivreRepository {
	return FeiraLivreRepository{
		db: db,
	}
}

func (r FeiraLivreRepository) Insert(f FeiraLivre) error {
	tx, _ := r.db.Begin()
	stmt, err := r.db.Prepare(`INSERT INTO feiraslivres(
			lng, 
			lat, 
			setcens,
			areap,
			coddist,
			distrito,
			codsubpref,
			subprefe,
			regiao5,
			regiao8,
			nome,
			registro,
			logradouro,
			numero,
			bairro,
			referencia
		) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
	`)

	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = stmt.Exec(
		f.Lng, 
		f.Lat,
		f.Setcens,
		f.Areap, 
		f.Coddist, 
		f.Distrito,
		f.Codsubpref, 
		f.Subprefe, 
		f.Regiao5, 
		f.Regiao8, 
		f.Nome,
		f.Registro,
		f.Logradouro,
		f.Numero,
		f.Bairro,
		f.Referencia,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	tx.Commit()
	return nil
}

func (r FeiraLivreRepository) List() []FeiraLivre {
	results := make([]FeiraLivre, 0)

	rows, err := r.db.Query("SELECT * FROM feiraslivres ORDER BY id")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var f FeiraLivre

		err = rows.Scan(&f.ID, &f.Nome, &f.Lng, &f.Lat)
		if err != nil {
			fmt.Println("error to scan column value")
		}

		results = append(results, f)
	}

	return results
}
