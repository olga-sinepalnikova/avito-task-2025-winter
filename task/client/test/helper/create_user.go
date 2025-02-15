package helper

import (
	"context"
	openapiclient "github.com/olgasinepalnikova/avito-task-2025-winter/client/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
)

func CreateUser(apiClient *openapiclient.APIClient, t *testing.T) string {
	login := "created user " + strconv.FormatInt(int64(rand.Int()), 10)
	pass := "created password " + strconv.FormatInt(int64(rand.Int()), 10)

	req := openapiclient.AuthRequest{
		Username: login,
		Password: pass,
	}

	resp, httpRes, err := apiClient.DefaultAPI.ApiAuthPost(context.Background()).AuthRequest(req).Execute()
	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)

	return LoginUser(apiClient, t, login, pass)
}

func LoginUser(apiClient *openapiclient.APIClient, t *testing.T, login, pass string) string {
	req := openapiclient.AuthRequest{
		Username: login,
		Password: pass,
	}

	resp, httpRes, err := apiClient.DefaultAPI.ApiAuthPost(context.Background()).AuthRequest(req).Execute()
	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)

	return resp.GetToken()
}
