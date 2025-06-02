package modelos

import (
	"app/src/config"
	"app/src/requisicoes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	FotoPerfil  string       `json:"fotoPerfil" db:"foto_perfil"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)
	for range 4 {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("erro ao buscar o usuario")
			}
			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar os seguidores")
			}
			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar os seguindo")
			}
			seguindo = seguindoCarregados

		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("erro ao buscar as publicacoes")
			}
			publicacoes = publicacoesCarregadas

		}

	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

// BuscarDadosDoUsuario invokes the API to fetch the user data
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

// BuscarSeguidores invokes the API to fetch the followers
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

// BuscarSeguindo invokes the API to fetch the followed
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}
	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguindo
}

// BuscarPublicacoes invokes the API to fetch the posts
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao

	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}
	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
