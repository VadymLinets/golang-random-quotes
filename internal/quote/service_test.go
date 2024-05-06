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
	testUserID = gofakeit.UUID()
	testQuote  = database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
	}
)

func TestGetQuote_Success(t *testing.T) {
	svc, db := newService(t, &config.Config{}, false)
	db.EXPECT().GetQuotes(mock.Anything, testUserID).Return([]database.Quote{testQuote}, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)

	receivedQuote, err := svc.GetQuote(context.Background(), testUserID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), receivedQuote)
}

func TestGetQuote_SuccessRandom(t *testing.T) {
	cfg := config.Config{QuotesConfig: config.QuotesConfig{RandomQuoteChance: 100}}

	svc, db := newService(t, &cfg, true)
	db.EXPECT().GetQuotes(mock.Anything, testUserID).Return(nil, nil)
	db.EXPECT().SaveQuote(mock.Anything, testQuote).Return(nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)

	responder, err := httpmock.NewJsonResponder(200, RandomQuote{
		ID:      testQuote.ID,
		Content: testQuote.Quote,
		Author:  testQuote.Author,
		Tags:    testQuote.Tags,
	})
	require.NoError(t, err)
	httpmock.RegisterResponder(http.MethodGet, randomQuoteURL, responder)

	receivedQuote, err := svc.GetQuote(context.Background(), testUserID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), receivedQuote)
}

func TestLikeQuote_Success(t *testing.T) {
	view := database.View{Liked: false}

	svc, db := newService(t, &config.Config{}, false)
	db.EXPECT().GetView(mock.Anything, testUserID, testQuote.ID).Return(view, nil)
	db.EXPECT().LikeQuote(mock.Anything, testQuote.ID).Return(nil)
	db.EXPECT().MarkAsLiked(mock.Anything, testUserID, testQuote.ID).Return(nil)

	err := svc.LikeQuote(context.Background(), testUserID, testQuote.ID)
	require.NoError(t, err)
}

func TestLikeQuote_AlreadyLiked(t *testing.T) {
	view := database.View{Liked: true}

	svc, db := newService(t, &config.Config{}, false)
	db.EXPECT().GetView(mock.Anything, testUserID, testQuote.ID).Return(view, nil)

	err := svc.LikeQuote(context.Background(), testUserID, testQuote.ID)
	require.NoError(t, err)
}

func TestGetSameQuote_Success(t *testing.T) {
	svc, db := newService(t, &config.Config{}, false)
	db.EXPECT().GetQuote(mock.Anything, testQuote.ID).Return(testQuote, nil)
	db.EXPECT().GetSameQuote(mock.Anything, testUserID, testQuote).Return(testQuote, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)

	sameQuote, err := svc.GetSameQuote(context.Background(), testUserID, testQuote.ID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), sameQuote)
}

func TestGetSameQuote_Random(t *testing.T) {
	svc, db := newService(t, &config.Config{}, true)
	db.EXPECT().GetQuote(mock.Anything, testQuote.ID).Return(testQuote, nil)
	db.EXPECT().GetSameQuote(mock.Anything, testUserID, testQuote).Return(database.Quote{}, database.ErrRecordNotFound)
	db.EXPECT().SaveQuote(mock.Anything, testQuote).Return(nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)

	responder, err := httpmock.NewJsonResponder(200, RandomQuote{
		ID:      testQuote.ID,
		Content: testQuote.Quote,
		Author:  testQuote.Author,
		Tags:    testQuote.Tags,
	})
	require.NoError(t, err)
	httpmock.RegisterResponder(http.MethodGet, randomQuoteURL, responder)

	sameQuote, err := svc.GetSameQuote(context.Background(), testUserID, testQuote.ID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), sameQuote)
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
