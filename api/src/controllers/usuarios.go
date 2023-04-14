package controllers

import "net/http"

func CriarUsuario (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}
func BuscaTodosUsuarios (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}
func BuscandoUsuarioPorID (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário por ID"))
}
func AtualizaUsuario (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}
func DeletaUsuario (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}