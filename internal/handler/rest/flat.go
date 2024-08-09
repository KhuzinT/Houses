package rest

import (
	"Houses/internal/model"
	"Houses/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) createFlat(w http.ResponseWriter, r *http.Request) {
	var info model.FlatInfo
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		// ошибка при получении информации о квартире
		SendResp(w, FailStatus, nil, http.StatusBadRequest, MsgInvalidFlatRequestPayload+err.Error())
		return
	}

	houseId, err := strconv.Atoi(r.URL.Query().Get("houseId"))
	if err != nil {
		// ошибка при получении id дома
		SendResp(w, FailStatus, nil, http.StatusBadRequest, MsgInvalidHouseId+err.Error())
		return
	}

	flat, err := h.s.Flat.CreateFlat(uint(houseId), info)
	if err != nil {
		// ошибка при создании квартиры
		SendResp(w, ErrorStatus, nil, http.StatusInternalServerError, MsgFailedToCreateFlat)
		return
	}

	SendResp(w, SuccessStatus, flat, http.StatusOK, MsgFlatCreated)
}

func (h *Handler) updateFlat(w http.ResponseWriter, r *http.Request) {
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

	flatId, err := strconv.Atoi(r.URL.Query().Get("flatId"))
	if err != nil {
		// ошибка при получении id квартиры
		SendResp(w, FailStatus, nil, http.StatusBadRequest, MsgInvalidFlatId+err.Error())
		return
	}

	status, err := convertStrToFlatStatus(r.URL.Query().Get("status"))
	if err != nil {
		// ошибка при получении нового статуса квартиры
		SendResp(w, FailStatus, nil, http.StatusBadRequest, MsgInvalidFlatStatus+err.Error())
		return
	}

	// по полученному id квартиры проверяем ее статус в базе
	// если квартира находится на модерации и модератор не мы,
	// то изменить статус мы не сможем

	_ /*flat*/, err = h.s.Flat.GetById(uint(flatId))
	if err != nil {
		// ошибка при получении информации о квартире
		SendResp(w, ErrorStatus, nil, http.StatusInternalServerError, MsgFailedToGetFlat+err.Error())
		return
	}

	/*
		if flat.FlatStatus == model.OnModeration && flat.ModeratorID != claims.UserID {
			// ошибка модерации
			SendResp(w, ErrorStatus, nil, http.StatusForbidden, MsgAlreadyOnModeration)
			return
		}
	*/

	if err := h.s.Flat.UpdateFlatStatus(uint(flatId), 0 /*claims.UserID*/, status); err != nil {
		// ошибка при обновлении статуса
		SendResp(w, ErrorStatus, nil, http.StatusInternalServerError, MsgFailedToUpdateFlat+err.Error())
		return
	}

	SendResp(w, SuccessStatus, nil, http.StatusOK, MsgFlatUpdated)
}

func convertStrToFlatStatus(statusStr string) (model.FlatStatus, error) {
	switch statusStr {
	case "created":
		return model.Created, nil
	case "approved":
		return model.Approved, nil
	case "declined":
		return model.Declined, nil
	case "on_moderation":
		return model.OnModeration, nil
	default:
		return model.Created, utils.ErrUnknownFlatStatus
	}
}
