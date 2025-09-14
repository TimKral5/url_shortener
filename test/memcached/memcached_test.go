package memcached_test

import (
	"os"
	"testing"

	"github.com/timkral5/url_shortener/internal/cache"
	"github.com/timkral5/url_shortener/test/testdata"
)

var conn *cache.MemcachedConnection

func TestNewMemcachedConnection(t *testing.T) {
	var err error

	connStr := os.Getenv("SHORTENER_MEMCACHED_URL")

	conn, err = cache.NewMemcachedConnection(connStr)
	if err != nil {
		t.Error("Connection to Memcached failed:")
		t.Error(err)

		return
	}
}

func TestAddURL(t *testing.T) {
	data := testdata.ReadStaticTestData()

	for hash, url := range data.TestURLs {
		hash = hash[:10]

		t.Log("Adding URL", url, "with hash", hash, "...")

		err := conn.AddURL(hash, url)
		if err != nil {
			t.Error("Operation failed:", err)

			return
		}

		t.Log("Operation successful.")
	}
}

func TestGetURL(t *testing.T) {
	data := testdata.ReadStaticTestData()

	for hash, url := range data.TestURLs {
		hash = hash[:10]

		t.Log("Fetching URL for hash", hash, "...")

		result, err := conn.GetURL(hash)
		if err != nil {
			t.Error("Operation failed:", err)

			return
		}

		if result != url {
			t.Error("URL does not match expected value:")
			t.Error("Expected", url, "but got", result)

			return
		}

		t.Log("Operation successful.")
	}
}

func TestDisconnect(t *testing.T) {
	err := conn.Disconnect()
	if err != nil {
		t.Error("Disconnect from Memcached failed:", err)

		return
	}
}


func BenchmarkURLOperations(b *testing.B) {
	var err error

	connStr := os.Getenv("SHORTENER_MEMCACHED_URL")

	conn, err = cache.NewMemcachedConnection(connStr)
	if err != nil {
		b.Fail()

		return
	}

	runBenchmark(b, conn)

	err = conn.Disconnect()
	if err != nil {
		b.Error("Disconnect from Memcached failed:", err)

		return
	}
}

func runBenchmark(b *testing.B, conn *cache.MemcachedConnection) {
	b.Helper()
	b.StopTimer()
	b.ResetTimer()

	for i := range b.N {
		hash, url := testdata.GenerateTestValues(i)

		b.StartTimer()

		err := conn.AddURL(hash, url)
		if err != nil {
			b.Error(err)

			return
		}

		_, err = conn.GetURL(hash)
		if err != nil {
			b.Error(err)

			return
		}

		b.StopTimer()
	}
}
