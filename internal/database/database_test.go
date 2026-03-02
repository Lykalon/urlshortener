package database_test

import (
	"testing"

	"github.com/Lykalon/urlshortener/internal/database"
	"github.com/stretchr/testify/require"
)

func TestInitStogae(t *testing.T) {
	cases := []struct {
		typeIn	string
		isErr	bool
	}{
		{
			typeIn: "local",
			isErr: false,
		},
		{
			typeIn: "postgres",
			isErr: false,
		},
		{
			typeIn: "something",
			isErr: true,
		},
	}

	for i := range cases {
		testCase := cases[i]

		storage, err := database.InitStorage(testCase.typeIn)

		if testCase.isErr {
			require.Equal(t, nil, storage)
			require.ErrorIs(t, err, database.ErrorInvalidArg)
		} else {
			var storageType string
			switch storage.(type) {
			case *database.LocalStorage:
				storageType = "local"
			case *database.PgStorage:
				storageType = "postgres"
			}

			require.Equal(t, storageType, testCase.typeIn)
			require.NoError(t, err)
		}
	}
}

func TestLocalStorage(t *testing.T) {
	storage, err := database.InitStorage("local")
	storage.Init()
	require.NoError(t, err)
	storage.Save(947554746059674745, "4q0gRu_yLZ")

	cases := []struct {
		shortLink		int64
		fullLink		string
		shortLinkFound	int64
		fullLinkFound	string
		found			bool
	}{
		{
			shortLink: 947554746059674745,
			fullLink: "4q0gRu_yLZ",
			shortLinkFound: 947554746059674745,
			fullLinkFound: "4q0gRu_yLZ",
			found: true,
		},
		{
			shortLink: 947554746059674740,
			fullLink: "4q0gRu_yLz",
			shortLinkFound: 0,
			fullLinkFound: "",
			found: false,
		},
	}
	for i := range cases {
		testCase := cases[i]

		shortLink, found := storage.FindShort(testCase.fullLink)
		require.Equal(t, testCase.shortLinkFound, shortLink)
		require.Equal(t, found, testCase.found)

		fullLink, found := storage.FindFull(testCase.shortLink)
		require.Equal(t, testCase.fullLinkFound, fullLink)
		require.Equal(t, found, testCase.found)
	}
	storage.Close()
}