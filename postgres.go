package main

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/flpgst/go-reportgen-producer/postgres"
)

func runPostgres(ctx context.Context, db *sql.DB) error {
	// cria pessoa
	p := models.Pessoa{
		Nome: "Joáo da Silva",
	}
	// salva pessoa no banco
	if err := p.Save(ctx, db); err != nil {
		return err
	}
	// cria transação
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// salva primeira matricula
	m0 := models.Matricula{
		Codigo:   "001",
		PessoaID: sql.NullInt64{Int64: int64(p.ID), Valid: true},
	}
	if err := m0.Save(ctx, tx); err != nil {
		return err
	}
	// salva segunda matricula
	m1 := models.Matricula{
		Codigo:   "2",
		PessoaID: sql.NullInt64{Int64: int64(p.ID), Valid: true},
	}
	if err := m1.Save(ctx, tx); err != nil {
		return err
	}
	// atualiza codigo da segunda matricula
	m1.Codigo = "002"
	if err := m1.Update(ctx, tx); err != nil {
		return err
	}
	// salva terceira matricula
	m2 := models.Matricula{
		Codigo:   "003",
		PessoaID: sql.NullInt64{Int64: int64(p.ID), Valid: true},
	}
	if err := m2.Save(ctx, tx); err != nil {
		return err
	}

	// tx commit
	if err := tx.Commit(); err != nil {
		return err
	}

	// busca matricula por codigo
	matricula0, err := models.MatriculaByCodigo(ctx, db, "001")
	if err != nil {
		return err
	}
	for _, matricula := range matricula0 {
		fmt.Printf("Matricula %d: codigo: %q\n", matricula.ID, matricula.Codigo)
		pessoa, err := matricula.Pessoa(ctx, db)
		if err != nil {
			return err
		}
		fmt.Printf("Matricula %q pessoa: %q\n", matricula.Codigo, pessoa.Nome)
	}

	// busca primeira matricula e deleta
	m3, err := models.MatriculaByID(ctx, db, 1)
	if err != nil {
		return err
	}

	if err := m3.Delete(ctx, db); err != nil {
		return err
	}
	return nil
}
