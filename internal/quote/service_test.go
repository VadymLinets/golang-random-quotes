package quote

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"quote/config"
	"quote/pkg/database"
	"quote/test/mocks"
)

var (
	testUserID  = gofakeit.UUID()
	testQuoteID = gofakeit.UUID()
	testQuote   = gofakeit.Quote()
)

func TestGetQuote_Success(t *testing.T) {
	quote := database.Quote{ID: testQuoteID, Quote: testQuote}

	svc, db := newService(t, &config.Config{}, false)
	db.EXPECT().GetQuotes(mock.Anything, testUserID).Return([]database.Quote{quote}, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuoteID).Return(nil)

	receivedQuote, err := svc.GetQuote(context.Background(), testUserID)
	require.NoError(t, err)
	require.Equal(t, Quote{ID: quote.ID, Quote: quote.Quote}, receivedQuote)
}

func TestGetQuote_SuccessRandom(t *testing.T) {
	quote := database.Quote{ID: testQuoteID, Quote: testQuote}
	cfg := config.Config{QuotesConfig: config.QuotesConfig{RandomQuoteChance: 100}}

	svc, db := newService(t, &cfg, true)
	db.EXPECT().GetQuotes(mock.Anything, testUserID).Return(nil, nil)
	db.EXPECT().SaveQuote(mock.Anything, quote).Return(nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuoteID).Return(nil)

	responder, err := httpmock.NewJsonResponder(200, RandomQuote{ID: testQuoteID, Content: testQuote})
	require.NoError(t, err)
	httpmock.RegisterResponder(http.MethodGet, randomQuoteURL, responder)

	receivedQuote, err := svc.GetQuote(context.Background(), testUserID)
	require.NoError(t, err)
	require.Equal(t, Quote{ID: quote.ID, Quote: quote.Quote}, receivedQuote)
}

func TestLikeQuote_Success(t *testing.T) {
	view := database.View{Liked: false}

	svc, db := newService(t, &config.Config{}, false)
	db.EXPECT().GetView(context.Background(), testUserID, testQuoteID).Return(view, nil)
	db.EXPECT().LikeQuote(context.Background(), testQuoteID).Return(nil)
	db.EXPECT().MarkAsLiked(context.Background(), testUserID, testQuoteID).Return(nil)

	err := svc.LikeQuote(context.Background(), testUserID, testQuoteID)
	require.NoError(t, err)
}

func TestLikeQuote_AlreadyLiked(t *testing.T) {
	view := database.View{Liked: true}

	svc, db := newService(t, &config.Config{}, false)
	db.EXPECT().GetView(context.Background(), testUserID, testQuoteID).Return(view, nil)

	err := svc.LikeQuote(context.Background(), testUserID, testQuoteID)
	require.NoError(t, err)
}

func TestGetSameQuote_Success(t *testing.T) {
	quote := database.Quote{ID: testQuoteID, Quote: testQuote}

	svc, db := newService(t, &config.Config{}, false)
	db.EXPECT().GetSameQuote(context.Background(), testUserID, testQuoteID).Return(quote, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuoteID).Return(nil)

	sameQuote, err := svc.GetSameQuote(context.Background(), testUserID, testQuoteID)
	require.NoError(t, err)
	require.Equal(t, Quote{ID: quote.ID, Quote: quote.Quote}, sameQuote)
}

func TestGetSameQuote_Random(t *testing.T) {
	quote := database.Quote{ID: testQuoteID, Quote: testQuote}

	svc, db := newService(t, &config.Config{}, true)
	db.EXPECT().GetSameQuote(context.Background(), testUserID, testQuoteID).Return(database.Quote{}, database.ErrRecordNotFound)
	db.EXPECT().SaveQuote(mock.Anything, quote).Return(nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuoteID).Return(nil)

	responder, err := httpmock.NewJsonResponder(200, RandomQuote{
		ID:      testQuoteID,
		Content: testQuote,
	})
	require.NoError(t, err)
	httpmock.RegisterResponder(http.MethodGet, randomQuoteURL, responder)

	sameQuote, err := svc.GetSameQuote(context.Background(), testUserID, testQuoteID)
	require.NoError(t, err)
	require.Equal(t, Quote{
		ID:    testQuoteID,
		Quote: testQuote,
	}, sameQuote)
}

func newService(t *testing.T, cfg *config.Config, mockHTTP bool) (*Service, *mocks.QuoteDatabase) {
	t.Helper()

	var client *resty.Client
	if mockHTTP {
		client = resty.New()
		httpmock.ActivateNonDefault(client.GetClient())
		t.Cleanup(httpmock.Deactivate)
	} else {
		client = resty.New()
	}

	db := mocks.NewQuoteDatabase(t)
	return NewService(cfg, db, client), db
}
