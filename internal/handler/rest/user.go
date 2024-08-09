package rest

import (
	"Houses/internal/model"
	"Houses/internal/utils"
	"encoding/json"
	"net/http"
)

const authString = "Authorization"
const bearerString = "Bearer "

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var login model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		// ошибка при получении информации о пользователе
		SendResp(w, FailStatus, nil, http.StatusBadRequest, MsgInvalidLoginRequestPayload+err.Error())
		return
	}

	if err := h.s.User.Login(login); err != nil {
		// неверный логин или пароль
		SendResp(w, FailStatus, nil, http.StatusUnauthorized, MsgInvalidEmailOrPassword+err.Error())
		return
	}

	user, err := h.s.User.GetUserByEmail(login.Email)
	if err != nil {
		// ошибка при получении информации о пользователе по email
		SendResp(w, ErrorStatus, nil, http.StatusInternalServerError, MsgFailedToRetrieveUser+err.Error())
		return
	}

	token, err := h.a.GenerateToken(user.ID, user.UserType)
	if err != nil {
		// ошибка при генерации токена
		SendResp(w, ErrorStatus, nil, http.StatusInternalServerError, MsgFailedToGenToken+err.Error())
		return
	}

	w.Header().Add(authString, bearerString+token)

	SendResp(w, SuccessStatus, map[string]string{"token": token}, http.StatusOK, MsgUserLoggedIn)
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var login model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		// ошибка при получении информации о пользователе
		SendResp(w, FailStatus, nil, http.StatusBadRequest, MsgInvalidLoginRequestPayload+err.Error())
		return
	}

	utype := model.Client
	if r.URL.Query().Get("type") == "moderator" {
		utype = model.Moderator
	}

	if err := h.s.User.Register(login, utype); err != nil {
		// ошибка при регистрации
		SendResp(w, FailStatus, nil, http.StatusNotAcceptable, MsgFailedToRegister+err.Error())
		return
	}

	SendResp(w, SuccessStatus, nil, http.StatusOK, MsgUserRegistered)
}

func (h *Handler) authorizeUser(r *http.Request) (*utils.Claims, error) {
	header := r.Header.Get(authString)
	if header == "" {
		return nil, utils.ErrMissingAuthHeader
	}

	token := header[len(bearerString):]
	claims, err := h.a.ParseToken(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
