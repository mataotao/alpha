package user

import (
	"alpha/config"
	"alpha/repositories/data-mappers/model"
	"go.uber.org/zap"
	"time"
)

type ListResponse struct {
	Count uint64              `json:"count"`
	List  []*ListInfoResponse `json:"list"`
}
type ListInfoResponse struct {
	Id       uint64    `json:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Mobile   uint64    `json:"mobile"`
	Avatar   string    `json:"avatar"`
	LastTime time.Time `json:"last_time"`
	LastIp   string    `json:"last_ip"`
	Status   byte      `json:"status"`
}

func List(user *model.UserModel, page, limit uint64) (*ListResponse, error) {
	listResponse := &ListResponse{
		List: make([]*ListInfoResponse, 0, limit),
	}

	list, count, err := user.List("*", page, limit)
	if err != nil {
		config.Logger.Error("user list",
			zap.Error(err),
		)
		return listResponse, err
	}
	for i := range list[:] {
		info := &ListInfoResponse{
			Id:       list[i].Id,
			Username: list[i].Username,
			Name:     list[i].Name,
			Mobile:   list[i].Mobile,
			Avatar:   list[i].Avatar,
			LastTime: list[i].LastTime,
			LastIp:   list[i].LastIp,
			Status:   list[i].Status,
		}
		listResponse.List = append(listResponse.List, info)
	}
	listResponse.Count = count
	return listResponse, nil
}
