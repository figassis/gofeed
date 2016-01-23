package feed_test

import (
	"io/ioutil"
	"testing"

	"github.com/mmcdole/go-feed"
	"github.com/stretchr/testify/assert"
)

func TestParseRSSFeed_DetectVersion_RSS090(t *testing.T) {
	f, _ := ioutil.ReadFile("testdata/simple_rss090.xml")

	rss, err := feed.ParseRSSFeed(string(f))

	expected := "0.9"
	assert.Nil(t, err, "Failed to parse feed")
	assert.Equal(t, expected, rss.Version, "Expected Version %s, got %s", expected, rss.Version)
}

func TestParseRSSFeed_DetectVersion_RSS091(t *testing.T) {
	f, _ := ioutil.ReadFile("testdata/simple_rss091.xml")

	rss, err := feed.ParseRSSFeed(string(f))

	expected := "0.91"
	assert.Nil(t, err, "Failed to parse feed")
	assert.Equal(t, expected, rss.Version, "Expected Version %s, got %s", expected, rss.Version)
}

func TestParseRSSFeed_DetectVersion_RSS092(t *testing.T) {
	f, _ := ioutil.ReadFile("testdata/simple_rss092.xml")

	rss, err := feed.ParseRSSFeed(string(f))

	expected := "0.92"
	assert.Nil(t, err, "Failed to parse feed")
	assert.Equal(t, expected, rss.Version, "Expected Version %s, got %s", expected, rss.Version)
}

func TestParseRSSFeed_DetectVersion_RSS10(t *testing.T) {
	f, _ := ioutil.ReadFile("testdata/simple_rss10.xml")

	rss, err := feed.ParseRSSFeed(string(f))

	expected := "1.0"
	assert.Nil(t, err, "Failed to parse feed")
	assert.Equal(t, expected, rss.Version, "Expected Version %s, got %s", expected, rss.Version)
}

func TestParseRSSFeed_DetectVersion_RSS20(t *testing.T) {
	f, _ := ioutil.ReadFile("testdata/simple_rss20.xml")

	rss, err := feed.ParseRSSFeed(string(f))

	expected := "2.0"
	assert.Nil(t, err, "Failed to parse feed")
	assert.Equal(t, expected, rss.Version, "Expected Version %s, got %s", expected, rss.Version)
}
