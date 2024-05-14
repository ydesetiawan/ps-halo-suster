package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
)

type UserInfo struct {
	UserId string
	Role   string
}

func ExtractUserInfo(ctx echo.Context) (*UserInfo, error) {
	result := &UserInfo{}
	user := ctx.Get("user_info")
	jsonData, err := json.Marshal(user)

	if err != nil {
		return result, fmt.Errorf("failed to marshal context value. error: %v", err)
	}

	var userInfo map[string]interface{}
	err = json.Unmarshal(jsonData, &userInfo)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal json. error: %v", err)
	}

	return &UserInfo{
		UserId: fmt.Sprintf("%v", userInfo["user_id"]),
		Role:   fmt.Sprintf("%v", userInfo["role"]),
	}, nil
}
