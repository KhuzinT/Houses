package rest

import (
	"Houses/internal/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) getFlats(w http.ResponseWriter, r *http.Request) {
	houseId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// ошибка при получении id дома
		SendResp(w, FailStatus, nil, http.StatusBadRequest, MsgInvalidHouseId+err.Error())
		return
	}

	flats := &[]model.Flat{}
	/*
		claims, _ := h.authorizeUser(r)
		if claims != nil && claims.UserType == model.Moderator {
			flats, err = h.s.Flat.GetFlats(uint(houseId))
		} else {
			flats, err = h.s.Flat.GetApprovedFlats(uint(houseId))
		}
	*/
	flats, err = h.s.Flat.GetFlats(uint(houseId))
	if err != nil {
		// ошибка при получении квартир по id дома
		SendResp(w, ErrorStatus, nil, http.StatusInternalServerError, MsgFailedToGetFlats+err.Error())
		return
	}

	SendResp(w, SuccessStatus, flats, http.StatusOK, MsgFlatsReceived)
}

func (h *Handler) createHouse(w http.ResponseWriter, r *http.Request) {
	/* ToDo: auth error
	claims, err := h.authorizeUser(r)
	if err != nil {
		// ошибка во время авторизации
		SendResp(w, FailStatus, nil, http.StatusBadRequest, MsgFailedToAuth+err.Error())
		return
	} else if claims != nil && claims.UserType == model.Moderator {
		// доступ запрещен
		SendResp(w, FailStatus, nil, http.StatusForbidden, MsgAccessDenied)
		return
	}*/

	var info model.HouseInfo
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		// ошибка при получении информации о доме
		SendResp(w, FailStatus, nil, http.StatusBadRequest, MsgInvalidHouseRequestPayload+err.Error())
		return
	}

	house, err := h.s.House.CreateHouse(info)
	if err != nil {
		// ошибка при создании дома
		SendResp(w, ErrorStatus, nil, http.StatusInternalServerError, MsgFailedToCreateHouse+err.Error())
		return
	}

	SendResp(w, SuccessStatus, house, http.StatusOK, MsgHouseCreated)
}
