package coding

import (
	"context"
	"github.com/guoyk93/gitdump"
	"github.com/guoyk93/grace"
	"github.com/stretchr/testify/require"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestHosting_List(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	h := Hosting{}
	opts := gitdump.HostingOptions{
		URL:      os.Getenv("TEST_CODING_URL"),
		Username: os.Getenv("TEST_CODING_USERNAME"),
		Password: os.Getenv("TEST_CODING_PASSWORD"),
	}
	u := grace.Must(url.Parse(opts.URL))
	tenant := strings.Split(u.Hostname(), ".")[0]
	repos, err := h.List(context.Background(), opts)
	require.NoError(t, err)
	require.NotEmpty(t, repos)
	repo := repos[rand.Intn(len(repos))]
	require.True(t, strings.HasPrefix(repo.URL, "https://e.coding.net/"+tenant))
	require.True(t, strings.HasPrefix(repo.SubDir, u.Hostname()))
	require.True(t, repo.Username == opts.Username)
	require.True(t, repo.Password == opts.Password)
}
