package server

import (
	"crud/database"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type usuario struct {
	ID uint64 		`json:"id"`
	Nome string 	`json:"nome"`
	Email string 	`json:"email"`
}

// cria um registro de usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Falha ao criar usuário!"))
		return
	}

	var usuario usuario
	if error = json.Unmarshal(body, &usuario); error != nil {
		w.Write([]byte("Falha ao converter dados!"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	stmt, error := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if error != nil {
		w.Write([]byte("Falha ao criar statement" + error.Error()))
		return
	}
	defer stmt.Close()

	insert, error := stmt.Exec(usuario.Nome, usuario.Email)
	if error != nil {
		w.Write([]byte("Falha na operação"))
		return
	}

	id, error := insert.LastInsertId()
	if error != nil {
		w.Write([]byte("Falha na recuperação do ID"))
		return
	}

	usuario.ID = uint64(id)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Location", "/usuarios/" + string(id))
	w.Header().Add("Content-Type", "application/json")
	if error = json.NewEncoder(w).Encode(usuario); error != nil {
		w.Write([]byte("Falha ao converter dados do usuário"))
		return
	}
}

// altera um usuário no banco de dados
func AlterarUsuario(w http.ResponseWriter, r *http.Request){

	parametros := mux.Vars(r)

	userId, error := strconv.ParseUint(parametros["id"], 10, 64)
	if error != nil {
		w.Write([]byte("Falha ao recuperar o parâmetro ID"))
		return
	}

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Falha ao criar usuário!"))
		return
	}

	var usuario usuario
	if error = json.Unmarshal(body, &usuario); error != nil {
		w.Write([]byte("Falha ao converter dados!"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	stmt, error := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if error != nil {
		w.Write([]byte("Falha ao criar rotina de atualização" + error.Error()))
		return
	}
	defer stmt.Close()

	_, error = stmt.Exec(usuario.Nome, usuario.Email, userId)
	if error != nil {
		w.Write([]byte("Falha na atualização do usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// recupera todos os usuários no banco de dados
func RecuperarUsuarios(w http.ResponseWriter, r *http.Request){

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	resultSet, error := db.Query("select id, nome, email from usuarios")
	if error != nil {
		w.Write([]byte("Falha ao criar consulta" + error.Error()))
		return
	}
	defer resultSet.Close()

	var usuarios []usuario

	for resultSet.Next() {
		var u usuario
		if error := resultSet.Scan(&u.ID, &u.Nome, &u.Email); error != nil {
			w.Write([]byte("Erro ao escanear usuário" + error.Error()))
			return
		}
		usuarios = append(usuarios, u)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	if error = json.NewEncoder(w).Encode(usuarios); error != nil {
		w.Write([]byte("Falha ao converter dados dos usuários"))
		return
	}
}

// recupera um usuário específico no banco de dados
func RecuperarUsuario(w http.ResponseWriter, r *http.Request){

	parametros := mux.Vars(r)

	userId, error := strconv.ParseUint(parametros["id"], 10, 64)
	if error != nil {
		w.Write([]byte("Falha ao recuperar o parâmetro ID"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	resultSet, error := db.Query("select id, nome, email from usuarios where id = ?", userId)
	if error != nil {
		w.Write([]byte("Falha ao criar consulta" + error.Error()))
		return
	}
	defer resultSet.Close()

	var usuario usuario

	if resultSet.Next() {
		if error := resultSet.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); error != nil {
			w.Write([]byte("Erro ao escanear usuário" + error.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		if error = json.NewEncoder(w).Encode(usuario); error != nil {
			w.Write([]byte("Falha ao converter dados do usuário"))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}


// exclui um usuário do banco de dados
func ExcluirUsuario(w http.ResponseWriter, r *http.Request){

	parametros := mux.Vars(r)

	userId, error := strconv.ParseUint(parametros["id"], 10, 64)
	if error != nil {
		w.Write([]byte("Falha ao recuperar o parâmetro ID"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	stmt, error := db.Prepare("delete from usuarios where id = ?")
	if error != nil {
		w.Write([]byte("Falha ao criar rotina de exclusão" + error.Error()))
		return
	}
	defer stmt.Close()

	_, error = stmt.Exec(userId)
	if error != nil {
		w.Write([]byte("Falha na exclusão do usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
