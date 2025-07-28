package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"

	"quote/config"
	"quote/internal/quote"
	"quote/pkg/database"
	"quote/test/common"
	"quote/test/docker"
)

var (
	testUserID = gofakeit.UUID()
	testQuote  = database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
	}
)

func TestIntegration(t *testing.T) {
	if os.Getenv("MUTATION_TESTING") != "" {
		t.Skip("Skipping integration testing for mutation tests")
	}

	postgresConfig := docker.StartPostgres(t)
	cfg := common.GetConfig(t, postgresConfig)
	db := common.PostgresClient(t, cfg)
	common.RunApp(t, cfg)

	httpClient := resty.New()

	err := db.SaveQuote(t.Context(), testQuote)
	require.NoError(t, err)

	getQuote(t, cfg, httpClient, db)
	likeQuote(t, cfg, httpClient, db)
	getSameQuote(t, cfg, httpClient, db)
}

func getQuote(t *testing.T, cfg *config.Config, client *resty.Client, db *database.Postgres) {
	t.Helper()

	params := map[string]string{
		"user_id": testUserID,
	}

	resp, err := client.R().SetQueryParams(params).Get(fmt.Sprintf("http://%s/", cfg.Addr))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var q quote.Quote

	err = json.Unmarshal(resp.Body(), &q)
	require.NoError(t, err)

	require.Equal(t, quote.Quote{
		ID:     testQuote.ID,
		Quote:  testQuote.Quote,
		Author: testQuote.Author,
		Tags:   testQuote.Tags,
	}, q)

	dbq, err := db.GetQuote(t.Context(), testQuote.ID)
	require.NoError(t, err)
	require.Equal(t, testQuote, dbq)
}

func likeQuote(t *testing.T, cfg *config.Config, client *resty.Client, db *database.Postgres) {
	t.Helper()

	params := map[string]string{
		"user_id":  testUserID,
		"quote_id": testQuote.ID,
	}

	resp, err := client.R().SetQueryParams(params).Patch(fmt.Sprintf("http://%s/like", cfg.Addr))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	q, err := db.GetQuote(t.Context(), testQuote.ID)
	require.NoError(t, err)
	require.Equal(t, int64(1), q.Likes)
}

func getSameQuote(t *testing.T, cfg *config.Config, client *resty.Client, db *database.Postgres) {
	t.Helper()

	sameQuote := database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: testQuote.Author,
		Tags:   testQuote.Tags,
	}

	randomQuote := database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
	}

	err := db.SaveQuote(t.Context(), sameQuote)
	require.NoError(t, err)

	err = db.SaveQuote(t.Context(), randomQuote)
	require.NoError(t, err)

	params := map[string]string{
		"user_id":  testUserID,
		"quote_id": testQuote.ID,
	}

	resp, err := client.R().SetQueryParams(params).Get(fmt.Sprintf("http://%s/same", cfg.Addr))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	var q quote.Quote

	err = json.Unmarshal(resp.Body(), &q)
	require.NoError(t, err)

	require.Equal(t, quote.Quote{
		ID:     sameQuote.ID,
		Quote:  sameQuote.Quote,
		Author: sameQuote.Author,
		Tags:   sameQuote.Tags,
	}, q)
}
