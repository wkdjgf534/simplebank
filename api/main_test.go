package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	db "simplebank/db/sqlc"
	"simplebank/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

// Add "github.com/lib/pq" before execituing this test
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) // reduce log's amounts during tests
	os.Exit(m.Run())
}
