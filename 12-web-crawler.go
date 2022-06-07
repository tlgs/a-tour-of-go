package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Cache struct {
	mux sync.Mutex
	m   map[string]bool
}

func (c *Cache) Add(url string) {
	c.mux.Lock()
	c.m[url] = true
	c.mux.Unlock()
}

func (c *Cache) Exists(url string) bool {
	c.mux.Lock()
	_, ok := c.m[url]
	c.mux.Unlock()
	return ok
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, outer chan string) {
	defer close(outer)

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	cache.Add(url)
	if err != nil {
		outer <- err.Error()
		return
	}
	outer <- fmt.Sprintf("found: %s %q", url, body)

	filtered := []string{}
	for _, u := range urls {
		if !cache.Exists(u) {
			filtered = append(filtered, u)
		}
	}

	inner := make([]chan string, len(filtered))
	for i, u := range filtered {
		inner[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, inner[i])
	}

	// drain children goroutines
	for i := range inner {
		for v := range inner[i] {
			outer <- v
		}
	}
}

var cache = Cache{m: make(map[string]bool)}

func main() {
	ch := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	for v := range ch {
		fmt.Println(v)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
