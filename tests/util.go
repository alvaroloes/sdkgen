package tests

import "net/url"

type FormattedDiff []string

func (diffs FormattedDiff) String() string {
	var s string
	for _, diff := range diffs {
		s += diff + "\n"
	}
	return s
}

func MustParseURL(urlString string) *url.URL {
	url, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	return url
}
