package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/internal/dto"
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/internal/entity"
	"github.com/Guilherme-G-Cadilhe/golang-projects/GolangExpertFullCycle/module-2/4-apis-advanced/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserInterface
}

type Error struct {
	Message string `json:"message"`
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

// GetJWT godoc
// # Este endpoint cria um Token
// @Summary         Get a User JWT
// @Description     Get a User JWT
// @Accept          json
// @Produce         json
// @Tags            users
// @Param           request    body    dto.GetJWTInput   true   "user credentials"
// @Success         200     {object} dto.GetJWTOutput
// @Failure         404     {object} Error
// @Failure         500     {object} Error
// @Router          /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var userJwt dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&userJwt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	user, err := h.UserDB.FindByEmail(userJwt.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !user.ValidatePassword(userJwt.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		// sub = subject
		"sub": user.ID.String(),
		// exp = expiresIn (Tem que ser exp se não não é lido pelo middleware)
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	// Forma simples de serializar o json de saida
	acessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)

}

// CreateUser - Cria um novo usuario
// # Este endpoint cria um novo usuario e retorna status 201 Created
// @Summary         Cria um novo usuario
// @Description     Cria um novo usuario com os dados enviados no corpo da requisição
// @ID              create-user
// @Accept          json
// @Produce         json
// @Tags            users
// @Param           request    body    dto.CreateUserInput   true   "user request"
// @Success         201
// @Failure         400     {object} Error
// @Failure         500     {object} Error
// @Router          /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	newUser, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
