package api_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Lykalon/urlshortener/internal/api"
	"github.com/Lykalon/urlshortener/internal/config"
	"github.com/Lykalon/urlshortener/internal/database"
	"github.com/stretchr/testify/require"
)
func TestCreateShortLink(t *testing.T) {
	os.Setenv("STORAGE", "local")
	config.Init()
	storage, _ := database.InitStorage(config.GetConfig().Storage)
	storage.Init()
	storage.Save(918875086608743065, "someURL") //int decode - yPsqzw_F_Y
	config.SetStorage(storage)

	cases := []struct {
		inMethod	string
		inBody		[]byte
		inTarget	string
		outBody		[]byte
		outStatus	int
	}{
		{
			inMethod:	http.MethodPost,
			inBody:		[]byte(`{"url":"someURL"}`),
			inTarget:	"/api/create",
			outBody:	[]byte(`{"data":"yPsqzw_F_Y"}`),
			outStatus:	http.StatusOK,
		},
		{
			inMethod:	http.MethodPost,
			inBody:		[]byte(`{"url":"someURLnew"}`),
			outBody:	nil,
			inTarget:	"/api/create",
			outStatus:	http.StatusOK,
		},
		{
			inMethod:	http.MethodGet,
			inBody:		[]byte(`{"url":"someURL"}`),
			outBody:	[]byte(`Method not allowed. Use POST.`),
			inTarget:	"/api/create",
			outStatus:	http.StatusMethodNotAllowed,
		},
		{
			inMethod:	http.MethodPost,
			inBody:		[]byte(`{"url":"someURLnew"`),
			outBody:	[]byte(`unexpected EOF ¯\_(ツ)_/¯`),
			inTarget:	"/api/create",
			outStatus:	http.StatusBadRequest,
		},
		{
			inMethod:	http.MethodPost,
			inBody:		[]byte(`{"url":""}`),
			outBody:	[]byte(`Required field "url" missed. ¯\_(ツ)_/¯`),
			inTarget:	"/api/create",
			outStatus:	http.StatusBadRequest,
		},
		{
			inMethod:	http.MethodPost,
			inBody:		[]byte(`{"url1":"someURL"}`),
			outBody:	[]byte(`Required field "url" missed. ¯\_(ツ)_/¯`),
			inTarget:	"/api/create",
			outStatus:	http.StatusBadRequest,
		},
	}

	for i := range cases {
		testCase := cases[i]

		respRec := httptest.NewRecorder()
		req := httptest.NewRequest(testCase.inMethod, testCase.inTarget, bytes.NewBuffer(testCase.inBody))

		api.CreateShortLink(respRec, req)

		resp := respRec.Result()
		require.Equal(t, testCase.outStatus, resp.StatusCode)
		if testCase.outBody != nil {
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, testCase.outBody, body)
			resp.Body.Close()
		}
	}
}

func TestGetFullLink(t *testing.T) {
	os.Setenv("STORAGE", "local")
	config.Init()
	storage, _ := database.InitStorage(config.GetConfig().Storage)
	storage.Init()
	storage.Save(918875086608743065, "someURL") //int decode - yPsqzw_F_Y
	config.SetStorage(storage)

	cases := []struct {
		inMethod	string
		inBody		[]byte
		inTarget	string
		outBody		[]byte
		outStatus	int
	}{
		{
			inMethod:	http.MethodGet,
			inBody:		[]byte(`{"data":"yPsqzw_F_Y"}`),
			inTarget:	"/api/get",
			outBody:	[]byte(`{"url":"someURL"}`),
			outStatus:	http.StatusOK,
		},
		{
			inMethod:	http.MethodGet,
			inBody:		[]byte(`{"data":"yPsqzw_F_y"}`),
			outBody:	[]byte(`Full link for short link not found`),
			inTarget:	"/api/get",
			outStatus:	http.StatusNotFound,
		},
		{
			inMethod:	http.MethodPost,
			inBody:		[]byte(`{"url":"someURL"}`),
			outBody:	[]byte(`Method not allowed. Use GET.`),
			inTarget:	"/api/get",
			outStatus:	http.StatusMethodNotAllowed,
		},
		{
			inMethod:	http.MethodGet,
			inBody:		[]byte(`{"data":"yPsqzw_F_y}`),
			outBody:	[]byte(`unexpected EOF ¯\_(ツ)_/¯`),
			inTarget:	"/api/get",
			outStatus:	http.StatusBadRequest,
		},
		{
			inMethod:	http.MethodGet,
			inBody:		[]byte(`{"data":""}`),
			outBody:	[]byte(`Required field "data" missed. ¯\_(ツ)_/¯`),
			inTarget:	"/api/get",
			outStatus:	http.StatusBadRequest,
		},
		{
			inMethod:	http.MethodGet,
			inBody:		[]byte(`{"data1":"yPsqzw_F_y"}`),
			outBody:	[]byte(`Required field "data" missed. ¯\_(ツ)_/¯`),
			inTarget:	"/api/get",
			outStatus:	http.StatusBadRequest,
		},
		{
			inMethod:	http.MethodGet,
			inBody:		[]byte(`{"data":"yPsqzw_F"}`),
			outBody:	[]byte(`Wrong length for field "data". ¯\_(ツ)_/¯`),
			inTarget:	"/api/get",
			outStatus:	http.StatusBadRequest,
		},
	}

	for i := range cases {
		testCase := cases[i]

		respRec := httptest.NewRecorder()
		req := httptest.NewRequest(testCase.inMethod, testCase.inTarget, bytes.NewBuffer(testCase.inBody))

		api.GetFullLink(respRec, req)

		resp := respRec.Result()
		require.Equal(t, testCase.outStatus, resp.StatusCode)
		if testCase.outBody != nil {
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, testCase.outBody, body)
			resp.Body.Close()
		}
	}
}