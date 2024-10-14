package email

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-sample/internal/status"
	"io"
	"net/http"
	"net/url"
)

func GetAccessToken(ctx *fiber.Ctx) error {
	req := new(AccessTokenRequest)
	if err := ctx.QueryParser(req); err != nil {
		return err
	}
	accessToken, err := getAccessToken(req.ClientId, req.RefreshToken)
	if err != nil {
		return err
	}
	return status.Ok(ctx, accessToken)
}

const OutlookOauth2Domain = "https://login.microsoftonline.com/consumers/oauth2/v2.0/token"

// 获取 Outlook 邮箱 Oauth2 令牌信息
func getAccessToken(clientId, refreshToken string) (*AuthResponse, error) {
	fromDate := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {clientId},
		"refresh_token": {refreshToken},
	}
	response, err := http.PostForm(OutlookOauth2Domain, fromDate)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	authResp := new(AuthResponse)
	if err = json.Unmarshal(bytes, authResp); err != nil {
		return nil, err
	}
	return authResp, nil
}
