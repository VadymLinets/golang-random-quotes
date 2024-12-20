package test

import (
	"context"
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

	client := resty.New()
	cfg, db := runEssentials(t)

	err := db.SaveQuote(context.Background(), testQuote)
	require.NoError(t, err)

	getQuote(t, cfg, client, db)
	likeQuote(t, cfg, client, db)
	getSameQuote(t, cfg, client, db)
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

	dbq, err := db.GetQuote(context.Background(), testQuote.ID)
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

	q, err := db.GetQuote(context.Background(), testQuote.ID)
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

	err := db.SaveQuote(context.Background(), sameQuote)
	require.NoError(t, err)

	err = db.SaveQuote(context.Background(), randomQuote)
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
