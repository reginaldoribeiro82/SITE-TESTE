package main

import import (
	// 1. Pacotes da biblioteca padrão do Go
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	// 2. Pacotes de terceiros (que baixamos)
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"

	// 3. Pacotes internos do nosso projeto
	"site_teste/Backend/conn"
)

var db *pgx.Conn
func main() {
	var err error

	// CORREÇÃO 2: O nome da função é 'ConectaDB', e não 'ConnectDatabase'.
	db, err = conn.ConectaDB()
	if err != nil {
		log.Fatalf("Erro fatal ao conectar no DB: %v", err)
	}
	defer db.Close(context.Background())

	fmt.Println("Conexão com o PostgreSQL estabelecida com sucesso!")

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/registro", registrarUsuario)

	fmt.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "index.html")
	http.ServeFile(w, r, path)
}

func registrarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao parsear o formulário", http.StatusBadRequest)
		return
	}

	nome := r.FormValue("nome")
	email := r.FormValue("email")
	senha := r.FormValue("senha")
	genero := r.FormValue("genero")
	posicaoSex := r.FormValue("posicao_sex")
	cidade := r.FormValue("cidade")
	dataNascimento := r.FormValue("data_nascimento")

	senhaHash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao proteger a senha", http.StatusInternalServerError)
		return
	}

	sqlStatement := `
		INSERT INTO usuarios (nome, email, senha, genero, posicao_sex, cidade, data_nascimento)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = db.Exec(context.Background(), sqlStatement, nome, email, string(senhaHash), genero, posicaoSex, cidade, dataNascimento)
	if err != nil {
		log.Printf("Erro ao inserir usuário no banco de dados: %v", err)
		http.Error(w, "Erro ao registrar usuário. O e-mail já pode existir.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "<h1>Registro realizado com sucesso!</h1>")
	log.Printf("Novo usuário registrado: %s (%s)", nome, email)
}