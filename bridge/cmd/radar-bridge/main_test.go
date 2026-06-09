package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNormalizeFavoriteURL(t *testing.T) {
	got, err := normalizeFavoriteURL("HTTPS://Example.COM/article/?utm_source=radar&b=2&a=1#section")
	if err != nil {
		t.Fatal(err)
	}
	if got != "https://example.com/article/?a=1&b=2" {
		t.Fatalf("unexpected URL: %s", got)
	}
}

func TestFavoriteIDStableForTrackingParameters(t *testing.T) {
	base := FavoriteCandidate{
		Kind:       "link",
		ReportDate: "2026-06-09",
		ReportType: "ai-hn",
	}
	first := base
	first.URL = "https://example.com/post?utm_source=one"
	second := base
	second.URL = "https://example.com/post?utm_source=two"

	firstID, err := favoriteID(first)
	if err != nil {
		t.Fatal(err)
	}
	secondID, err := favoriteID(second)
	if err != nil {
		t.Fatal(err)
	}
	if firstID != secondID {
		t.Fatal("tracking parameters should not create duplicate favorites")
	}
}

func TestSafeFilename(t *testing.T) {
	got := safeFilename(`A/B: "Agent" | Notes?`)
	if strings.ContainsAny(got, `/\:*?"<>|`) {
		t.Fatalf("unsafe filename: %s", got)
	}
}

func TestSafeJoinRejectsTraversal(t *testing.T) {
	root := t.TempDir()
	if _, err := safeJoin(root, "../../secret"); err == nil {
		t.Fatal("expected traversal to be rejected")
	}
	got, err := safeJoin(root, "2026-06-09/ai-hn.md")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(got, filepath.Clean(root)+string(filepath.Separator)) {
		t.Fatalf("path escaped root: %s", got)
	}
}

func TestExcerptFavoriteIDNormalizesWhitespace(t *testing.T) {
	first := FavoriteCandidate{
		Kind:       "excerpt",
		Excerpt:    "Agent   memory\narchitecture",
		ReportDate: "2026-06-09",
		ReportType: "ai-agents",
	}
	second := first
	second.Excerpt = "agent memory architecture"
	firstID, err := favoriteID(first)
	if err != nil {
		t.Fatal(err)
	}
	secondID, err := favoriteID(second)
	if err != nil {
		t.Fatal(err)
	}
	if firstID != secondID {
		t.Fatal("normalized excerpts should have the same ID")
	}
}
