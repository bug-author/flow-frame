package server

import (
	"fmt"
	"net/url"
)

func GetAuthURL(clientId string, clientSecret string) string {
	params := url.Values{}
	params.Set("client_id", clientId)
	params.Set("client_secret", clientSecret)
	params.Set("response_type", "code")
	params.Set("scope", "tasks:write tasks:read")

	return fmt.Sprintf("https://ticktick.com/oauth/authorize?%s", params.Encode())
}
