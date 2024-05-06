package quote

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
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
	svc, db, _ := newTestService(t, &config.Config{})
	db.EXPECT().GetQuotes(mock.Anything, testUserID).Return([]database.Quote{testQuote}, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)

	receivedQuote, err := svc.GetQuote(context.Background(), testUserID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), receivedQuote)
}

func TestGetQuote_SuccessRandom(t *testing.T) {
	cfg := config.Config{QuotesConfig: config.QuotesConfig{RandomQuoteChance: 100}}

	svc, db, api := newTestService(t, &cfg)
	db.EXPECT().GetQuotes(mock.Anything, testUserID).Return(nil, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)
	api.EXPECT().GetRandomQuote(mock.Anything).Return(testQuote, nil)

	receivedQuote, err := svc.GetQuote(context.Background(), testUserID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), receivedQuote)
}

func TestLikeQuote_Success(t *testing.T) {
	view := database.View{Liked: false}

	svc, db, _ := newTestService(t, &config.Config{})
	db.EXPECT().GetView(mock.Anything, testUserID, testQuote.ID).Return(view, nil)
	db.EXPECT().LikeQuote(mock.Anything, testQuote.ID).Return(nil)
	db.EXPECT().MarkAsLiked(mock.Anything, testUserID, testQuote.ID).Return(nil)

	err := svc.LikeQuote(context.Background(), testUserID, testQuote.ID)
	require.NoError(t, err)
}

func TestLikeQuote_AlreadyLiked(t *testing.T) {
	view := database.View{Liked: true}

	svc, db, _ := newTestService(t, &config.Config{})
	db.EXPECT().GetView(mock.Anything, testUserID, testQuote.ID).Return(view, nil)

	err := svc.LikeQuote(context.Background(), testUserID, testQuote.ID)
	require.NoError(t, err)
}

func TestGetSameQuote_Success(t *testing.T) {
	svc, db, _ := newTestService(t, &config.Config{})
	db.EXPECT().GetQuote(mock.Anything, testQuote.ID).Return(testQuote, nil)
	db.EXPECT().GetSameQuote(mock.Anything, testUserID, testQuote).Return(testQuote, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)

	sameQuote, err := svc.GetSameQuote(context.Background(), testUserID, testQuote.ID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), sameQuote)
}

func TestGetSameQuote_Random(t *testing.T) {
	svc, db, api := newTestService(t, &config.Config{})
	db.EXPECT().GetQuote(mock.Anything, testQuote.ID).Return(testQuote, nil)
	db.EXPECT().GetSameQuote(mock.Anything, testUserID, testQuote).Return(database.Quote{}, database.ErrRecordNotFound)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)
	api.EXPECT().GetRandomQuote(mock.Anything).Return(testQuote, nil)

	sameQuote, err := svc.GetSameQuote(context.Background(), testUserID, testQuote.ID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), sameQuote)
}

func newTestService(t *testing.T, cfg *config.Config) (*Service, *mocks.QuoteDatabase, *mocks.QuoteAPI) {
	t.Helper()

	db := mocks.NewQuoteDatabase(t)
	api := mocks.NewQuoteAPI(t)
	return NewService(cfg, db, api), db, api
}
