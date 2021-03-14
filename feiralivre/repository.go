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

func (f *FeiraLivre) scanColumns(rows *sql.Rows) error {
	err := rows.Scan(
		&f.ID,
		&f.Lng,
		&f.Lat,
		&f.Setcens,
		&f.Areap,
		&f.Coddist,
		&f.Distrito,
		&f.Codsubpref,
		&f.Subprefe,
		&f.Regiao5,
		&f.Regiao8,
		&f.Nome,
		&f.Registro,
		&f.Logradouro,
		&f.Numero,
		&f.Bairro,
		&f.Referencia,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r FeiraLivreRepository) Insert(f FeiraLivre) error {
	tx, _ := r.db.Begin()
	stmt, err := r.db.Prepare(`
		INSERT INTO 
			feiraslivres(lng, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome, registro, logradouro, numero, bairro, referencia) 
			VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
	`)
	if err != nil {
		fmt.Println("Insert:Prepare", err)
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
		tx.Rollback()
		fmt.Println("Insert:exec", err)
		return err
	}
	tx.Commit()
	return nil
}

func qlike(s string) string {
	return "%" + s + "%"
}

func (r FeiraLivreRepository) Search(q string) []FeiraLivre {
	results := make([]FeiraLivre, 0)
	stmt, err := r.db.Prepare(`
		SELECT * 
			FROM feiraslivres 
			WHERE nome LIKE ? OR distrito LIKE ? OR regiao5 LIKE ? OR bairro LIKE ?`)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := stmt.Query(
		qlike(q), qlike(q), qlike(q), qlike(q),
	)
	for rows.Next() {
		f := FeiraLivre{}
		err := f.scanColumns(rows)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, f)
	}
	rows.Close()
	return results
}

func (r FeiraLivreRepository) FindById(id int) FeiraLivre {
	rows, err := r.db.Query(`
		SELECT * 
			FROM feiraslivres 
			WHERE id=? 
			LIMIT 1`, id)
	if err != nil {
		fmt.Println(err)
	}
	f := FeiraLivre{}
	for rows.Next() {
		err := f.scanColumns(rows)
		if err != nil {
			fmt.Println(err)
		}
	}
	rows.Close()
	return f
}

func (r FeiraLivreRepository) Update(id int, f FeiraLivre) error {
	tx, _ := r.db.Begin()
	stmt, err := r.db.Prepare(`
		UPDATE feiraslivres 
			SET lng=?, lat=?, setcens=?, areap=?, coddist=?, distrito=?, codsubpref=?, subprefe=?, regiao5=?, regiao8=?, nome=?, registro=?, logradouro=?, numero=?, bairro=?, referencia=? 
			WHERE id=?
	`)
	if err != nil {
		fmt.Println("Update:Prepare", err)
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
		id,
	)
	if err != nil {
		tx.Rollback()
		fmt.Println("Update:exec", err)
		return err
	}
	tx.Commit()
	return nil
}

func (r FeiraLivreRepository) Remove(id int) error {
	tx, _ := r.db.Begin()
	stmt, err := r.db.Prepare(`
		DELETE 
			FROM feiraslivres 
			WHERE id=?
		`)
	if err != nil {
		fmt.Println("Remove:Prepare", err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		fmt.Println("Remove:exec", err)
		return err
	}
	tx.Commit()
	return nil
}
