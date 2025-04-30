package quote

import (
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

	receivedQuote, err := svc.GetQuote(t.Context(), testUserID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), receivedQuote)
}

func TestGetQuote_Random(t *testing.T) {
	cfg := config.Config{QuotesConfig: config.QuotesConfig{RandomQuoteChance: 100}}

	svc, db, api := newTestService(t, &cfg)
	db.EXPECT().GetQuotes(mock.Anything, testUserID).Return(nil, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)
	api.EXPECT().GetRandomQuote(mock.Anything).Return(testQuote, nil)

	receivedQuote, err := svc.GetQuote(t.Context(), testUserID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), receivedQuote)
}

func TestGetQuote_NoQuotes(t *testing.T) {
	svc, db, api := newTestService(t, &config.Config{})
	db.EXPECT().GetQuotes(mock.Anything, testUserID).Return(nil, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)
	api.EXPECT().GetRandomQuote(mock.Anything).Return(testQuote, nil)

	receivedQuote, err := svc.GetQuote(t.Context(), testUserID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), receivedQuote)
}

func TestLikeQuote_Success(t *testing.T) {
	view := database.View{Liked: false}

	svc, db, _ := newTestService(t, &config.Config{})
	db.EXPECT().GetView(mock.Anything, testUserID, testQuote.ID).Return(view, nil)
	db.EXPECT().LikeQuote(mock.Anything, testQuote.ID).Return(nil)
	db.EXPECT().MarkAsLiked(mock.Anything, testUserID, testQuote.ID).Return(nil)

	err := svc.LikeQuote(t.Context(), testUserID, testQuote.ID)
	require.NoError(t, err)
}

func TestLikeQuote_AlreadyLiked(t *testing.T) {
	view := database.View{Liked: true}

	svc, db, _ := newTestService(t, &config.Config{})
	db.EXPECT().GetView(mock.Anything, testUserID, testQuote.ID).Return(view, nil)

	err := svc.LikeQuote(t.Context(), testUserID, testQuote.ID)
	require.NoError(t, err)
}

func TestGetSameQuote_Success(t *testing.T) {
	svc, db, _ := newTestService(t, &config.Config{})
	db.EXPECT().GetQuote(mock.Anything, testQuote.ID).Return(testQuote, nil)
	db.EXPECT().GetSameQuote(mock.Anything, testUserID, testQuote).Return(testQuote, nil)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)

	sameQuote, err := svc.GetSameQuote(t.Context(), testUserID, testQuote.ID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), sameQuote)
}

func TestGetSameQuote_Random(t *testing.T) {
	svc, db, api := newTestService(t, &config.Config{})
	db.EXPECT().GetQuote(mock.Anything, testQuote.ID).Return(testQuote, nil)
	db.EXPECT().GetSameQuote(mock.Anything, testUserID, testQuote).Return(database.Quote{}, database.ErrRecordNotFound)
	db.EXPECT().MarkAsViewed(mock.Anything, testUserID, testQuote.ID).Return(nil)
	api.EXPECT().GetRandomQuote(mock.Anything).Return(testQuote, nil)

	sameQuote, err := svc.GetSameQuote(t.Context(), testUserID, testQuote.ID)
	require.NoError(t, err)
	require.Equal(t, fromDatabaseQuoteToQuote(testQuote), sameQuote)
}

func TestChooseQuote_Success(t *testing.T) {
	quote1 := database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
		Likes:  10,
	}

	quote2 := database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
		Likes:  70,
	}

	quote3 := database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
		Likes:  20,
	}

	svc, _, _ := newTestService(t, &config.Config{})

	q, err := svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2, quote3}, 10)
	require.NoError(t, err)
	require.Equal(t, quote1, q)

	q, err = svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2, quote3}, 80)
	require.NoError(t, err)
	require.Equal(t, quote2, q)

	q, err = svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2, quote3}, 99)
	require.NoError(t, err)
	require.Equal(t, quote3, q)
}

func TestChooseQuote_ZeroLikes(t *testing.T) {
	quote1 := database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
	}

	quote2 := database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
	}

	svc, _, _ := newTestService(t, &config.Config{})

	q1, err := svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2}, 0)
	require.NoError(t, err)
	require.Equal(t, quote1, q1)

	q2, err := svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2}, 51)
	require.NoError(t, err)
	require.Equal(t, quote2, q2)
}

func TestChooseQuote_Random(t *testing.T) {
	quote1 := database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
		Likes:  1,
	}

	quote2 := database.Quote{
		ID:     gofakeit.UUID(),
		Quote:  gofakeit.Quote(),
		Author: gofakeit.Name(),
		Tags:   gofakeit.NiceColors(),
		Likes:  1,
	}

	svc, _, api := newTestService(t, &config.Config{QuotesConfig: config.QuotesConfig{RandomQuoteChance: 1}})
	api.EXPECT().GetRandomQuote(mock.Anything).Return(testQuote, nil)

	q, err := svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2}, 99)
	require.NoError(t, err)
	require.Equal(t, testQuote, q)

	q, err = svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2}, 98)
	require.NoError(t, err)
	require.Equal(t, quote2, q)

	q, err = svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2}, 21)
	require.NoError(t, err)
	require.Equal(t, quote1, q)

	svc, _, api = newTestService(t, &config.Config{QuotesConfig: config.QuotesConfig{RandomQuoteChance: 98}})
	api.EXPECT().GetRandomQuote(mock.Anything).Return(testQuote, nil)

	q, err = svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2}, 60)
	require.NoError(t, err)
	require.Equal(t, testQuote, q)

	q, err = svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2}, 1.9)
	require.NoError(t, err)
	require.Equal(t, quote2, q)

	q, err = svc.chooseQuote(t.Context(), []database.Quote{quote1, quote2}, 0)
	require.NoError(t, err)
	require.Equal(t, quote1, q)
}

func TestChooseQuote_EmptyList(t *testing.T) {
	svc, _, api := newTestService(t, &config.Config{})
	api.EXPECT().GetRandomQuote(mock.Anything).Return(testQuote, nil)

	q, err := svc.chooseQuote(t.Context(), []database.Quote{}, 19)
	require.NoError(t, err)
	require.Equal(t, testQuote, q)

	q, err = svc.chooseQuote(t.Context(), []database.Quote{}, 0)
	require.NoError(t, err)
	require.Equal(t, testQuote, q)
}

func newTestService(t *testing.T, cfg *config.Config) (*Service, *mocks.QuoteDatabase, *mocks.QuoteAPI) {
	t.Helper()

	db := mocks.NewQuoteDatabase(t)
	api := mocks.NewQuoteAPI(t)
	return NewService(cfg, db, api), db, api
}
