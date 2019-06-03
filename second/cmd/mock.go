package main

var (
	defaultResponse        = []byte("Default response")
	testProperlyformedYAML = []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)

	testMalformedYAML = []byte(`
	- path: /urlshort
		url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
		url: https://github.com/gophercises/urlshort/tree/solution
`)

	testProperlyformedJSON = []byte(`
	[{
		"path": "/some-path",
		"url": "https://the-url-to-redirect-to.com/well-done"
	},
	{
		"path": "/some-path",
		"url": "https://the-url-to-redirect-to.com/well-done"
	}]
`)

	testMalformedJSON = []byte(`
	- path: /urlshort
		url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
		url: https://github.com/gophercises/urlshort/tree/solution
`)
)
