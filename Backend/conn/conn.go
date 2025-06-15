package conn

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jackc/pgx/v5"
	"os"
)

//Variaveis com informacoes necessarias para conexao com o banco de dados
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "Site_Teste"
)

// ConectaDB estabelece a conexão com o banco de dados PostgreSQL.
func ConectaDB() (*pgx.Conn, error) {
	// Formato da string de conexão: "postgres://username:password@localhost:5432/database_name"
	// Substitua com seu usuário, senha e nome do banco de dados.
	connString := "postgres://seu_usuario:sua_senha@localhost:5432/site_teste"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Não foi possível conectar ao banco de dados: %v\n", err)
		return nil, err
	}

	// O defer conn.Close() será chamado no main.go para garantir que a conexão seja fechada.
	return conn, nil
}