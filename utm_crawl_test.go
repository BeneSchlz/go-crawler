package main

import (
	"net/url"
	"testing"
)

func TestHasUTMParams(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected bool
	}{
		{
			name:     "has utm_source",
			inputURL: "https://example.com/path?utm_source=seo",
			expected: true,
		},
		{
			name:     "has utm_term",
			inputURL: "https://example.com/path?utm_term=spring",
			expected: true,
		},
		{
			name:     "has utm_content",
			inputURL: "https://example.com/path?utm_content=cta",
			expected: true,
		},
		{
			name:     "has all",
			inputURL: "https://example.com/path?utm_source=google&utm_term=abc&utm_content=xyz",
			expected: true,
		},
		{
			name:     "has none",
			inputURL: "https://example.com/path",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := url.Parse(tt.inputURL)
			if err != nil {
				t.Fatalf("error parsing URL: %v", err)
			}
			result := hasUTMParams(u)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
