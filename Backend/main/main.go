package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"site_teste/Backend/conn" // Importa o nosso pacote de conexão
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.Conn
func main() {
	var err error
	//Aqui estamos utilizando a funçao de conectar ao banco de dados que esta disponivel no pacote conn que criamos
	// e guardando nossa resposta nas variáveis database (que será a partir dela que manipularemos nosso banco) e err (caso exista um erro)
	db, err = conn.ConnectDatabase()

	//Se algum erro for retornado da funçao, a execuçao da funcao será interrompida
	if err != nil {
		panic(err)
	}

	// 2. Definir as Rotas do Servidor
	// Rota para servir a página inicial (o formulário)
	http.HandleFunc("/", serveIndex)

	// Rota para receber os dados do formulário
	http.HandleFunc("/registro", registrarUsuario)

	// 3. Iniciar o Servidor
	fmt.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// serveIndex envia o arquivo index.html para o navegador.
func serveIndex(w http.ResponseWriter, r *http.Request) {
	// O caminho para o index.html é relativo à raiz do projeto
	path := filepath.Join("..", "..", "index.html")
	http.ServeFile(w, r, path)
}
// registrarUsuario processa os dados recebidos do formulário.
func registrarUsuario(w http.ResponseWriter, r *http.Request) {
	// Garante que o método da requisição seja POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Parseia os dados do formulário
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao parsear o formulário", http.StatusBadRequest)
		return
	}

	// Coleta os dados do formulário
	nome := r.FormValue("nome")
	email := r.FormValue("email")
	senha := r.FormValue("senha")
	genero := r.FormValue("genero")
	posicaoSex := r.FormValue("posicao_sex")
	cidade := r.FormValue("cidade")
	dataNascimento := r.FormValue("data_nascimento")

	// **SEGURANÇA**: Criptografa a senha antes de salvar
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao proteger a senha", http.StatusInternalServerError)
		return
	}

// Prepara a instrução SQL para inserir os dados
sqlStatement := `
		INSERT INTO usuarios (nome, email, senha, genero, posicao_sex, cidade, data_nascimento)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

// Executa a instrução SQL
	_, err = db.Exec(context.Background(), sqlStatement, nome, email, string(senhaHash), genero, posicaoSex, cidade, dataNascimento)
	if err != nil {
		// Loga o erro no servidor para debugging
		log.Printf("Erro ao inserir usuário no banco de dados: %v", err)
		http.Error(w, "Erro ao registrar usuário. O e-mail já pode existir.", http.StatusInternalServerError)
		return
	}

// Responde ao usuário com sucesso
	fmt.Fprintln(w, "<h1>Registro realizado com sucesso!</h1>")
	log.Printf("Novo usuário registrado: %s (%s)", nome, email)
}
