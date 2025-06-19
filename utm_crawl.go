package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func hasUTMParams(u *url.URL) bool {
	q := u.Query()
	return q.Has("utm_source") || q.Has("utm_term") || q.Has("utm_content")
}

func getUTMsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var utmURLs []string
	var traverseNodes func(*html.Node)
	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key == "href" {
					href, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v' : %v\n", anchor.Val, err)
						continue
					}

					resolvedURL := baseURL.ResolveReference(href)
					if hasUTMParams(resolvedURL) {
						utmURLs = append(utmURLs, resolvedURL.String())
					}

				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}
	}
	traverseNodes(doc)

	return utmURLs, nil
}
