# movie-times

A Go web app that scrapes current theatre showtimes and displays them in a clean, readable table. Built to solve the problem of checking multiple theatre websites just to find a showtime.

## Tech Stack

- Go (standard library + `net/http`)
- HTML template rendering (`html/template`)
- Web scraper (`web-scraper.go`)
- GitHub Actions — CI and build validation

## Running Locally

```bash
go mod download
go run .
# → http://localhost:8080
```

## Project Structure

```
movie-times/
├── main.go                   HTTP server and routing
├── web-scraper.go            theatre website scraper
├── render/
│   └── render-template.go   HTML template renderer
├── templates/
│   ├── movie-table.html      showtime table template
│   └── static/              CSS and JS assets
└── .github/workflows/go.yml  CI pipeline
```

## Status

The scraper and rendering pipeline are in place. Theatre websites update their markup frequently — the selectors may need updating against the current site structure. Last dependency bump: September 2025.

## Next Steps

1. **Verify the scraper** — run it locally and check if it still works against the current theatre site HTML. If broken, inspect the live page and update the CSS selectors in `web-scraper.go`.
2. **Add more theatres** — the scraper is structured to extend. Add whichever theatres you actually use.
3. **Deploy as a cron** — run the scraper nightly via GitHub Actions, write the output HTML to a `gh-pages` branch, and get a free always-current showtime page with zero infrastructure.
4. **Add booking links** — deep-link each showtime row into the theatre's booking URL pattern (most use predictable query strings) so one click goes straight to checkout.
