package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewAccessToken(t *testing.T) {
	at := NewAccessToken()
	require.Equal(t, "", at.AccessToken, "new access token should not have defined access token id")
	require.False(t, at.IsExpired(), "new access token should not be expired")
	require.Equal(t, int64(0), at.UserID, "new access token should not hace an associated user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	require.True(t, at.IsExpired())

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	require.False(t, at.IsExpired(), "access token expiring three hour from now should not be expired")
}
