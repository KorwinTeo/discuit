package server

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/discuitnet/discuit/core"
)

// SitemapURL represents a single URL in the sitemap
type SitemapURL struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	LastMod    string   `xml:"lastmod,omitempty"`
	ChangeFreq string   `xml:"changefreq,omitempty"`
	Priority   string   `xml:"priority,omitempty"`
}

// SitemapURLSet represents the root element of a sitemap
type SitemapURLSet struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []SitemapURL `xml:"url"`
}

// handleSitemap generates and serves the sitemap.xml
func (s *Server) handleSitemap(w *responseWriter, r *request) error {
	baseURL := fmt.Sprintf("https://%s", r.req.Host)

	urlset := SitemapURLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  []SitemapURL{},
	}

	// Add homepage
	urlset.URLs = append(urlset.URLs, SitemapURL{
		Loc:        baseURL + "/",
		LastMod:    time.Now().Format("2006-01-02"),
		ChangeFreq: "daily",
		Priority:   "1.0",
	})

	// Add communities
	communities, err := core.GetCommunities(r.ctx, s.db, core.CommunitiesSortDefault, "all", 1000, nil)
	if err == nil {
		for _, community := range communities {
			urlset.URLs = append(urlset.URLs, SitemapURL{
				Loc:        baseURL + "/" + community.Name,
				LastMod:    community.CreatedAt.Format("2006-01-02"),
				ChangeFreq: "weekly",
				Priority:   "0.8",
			})
		}
	}

	// Add recent posts using GetFeed
	feedOpts := &core.FeedOptions{
		Sort:      core.FeedSortHot,
		Viewer:    nil,
		Community: nil,
		Homefeed:  false,
		Limit:     500,
		Next:      "",
	}

	feedResult, err := core.GetFeed(r.ctx, s.db, feedOpts)
	if err == nil && feedResult != nil {
		for _, post := range feedResult.Posts {
			urlset.URLs = append(urlset.URLs, SitemapURL{
				Loc:        baseURL + "/" + post.CommunityName + "/post/" + post.PublicID,
				LastMod:    post.CreatedAt.Format("2006-01-02"),
				ChangeFreq: "weekly",
				Priority:   "0.6",
			})
		}
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)

	// Write XML declaration
	w.Write([]byte(xml.Header))

	// Encode and write the sitemap
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")
	return encoder.Encode(urlset)
}

// handleRobotsTxt serves the robots.txt file
func (s *Server) handleRobotsTxt(w *responseWriter, r *request) error {
	robotsTxt := `User-agent: *
Disallow: /search?q=*
Disallow: /*/modtools/*
Disallow: /admin/*
Disallow: /api/*

Sitemap: https://` + r.req.Host + `/sitemap.xml`

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(robotsTxt))
	return nil
}
