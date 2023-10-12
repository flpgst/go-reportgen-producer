package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/flpgst/go-reportgen-producer/models"
)

func runPostgres(ctx context.Context, db *sql.DB) error {
	// cria pessoa
	p, err := models.PessoaByID(ctx, db, 1)
	if err != nil {
		return err
	}

	// cria transação
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// tx commit
	if err := tx.Commit(); err != nil {
		return err
	}

	matriculasPessoa, err := models.MatriculaByPessoaID(ctx, db, p.ID)
	if err != nil {
		return err
	}

	for _, matricula := range matriculasPessoa {
		fmt.Printf("Matricula %d: codigo: %q\n", matricula.ID, matricula.Codigo.String)
		pessoa, err := matricula.Pessoa(ctx, db)
		if err != nil {
			return err
		}
		fmt.Printf("Matricula %q pessoa: %q\n", matricula.Codigo.String, pessoa.Nome.String)
	}

	return nil
}
