package quoteapi

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"quote/test/mocks"
)

var testQuote = RandomQuote{
	ID:      gofakeit.UUID(),
	Content: gofakeit.Quote(),
	Author:  gofakeit.Name(),
	Tags:    gofakeit.NiceColors(),
}

func TestGetRandomQuote_Success(t *testing.T) {
	svc, db := newService(t)

	responder, err := httpmock.NewJsonResponder(200, testQuote)
	require.NoError(t, err)
	httpmock.RegisterResponder(http.MethodGet, randomQuoteURL, responder)
	db.EXPECT().SaveQuote(mock.Anything, testQuote.toDatabase()).Return(nil)

	quote, err := svc.GetRandomQuote(context.Background())
	require.NoError(t, err)
	require.Equal(t, testQuote.toDatabase(), quote)
}

func newService(t *testing.T) (*Service, *mocks.QuoteapiDatabase) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	t.Cleanup(httpmock.Deactivate)

	db := mocks.NewQuoteapiDatabase(t)
	return NewService(db, client), db
}
