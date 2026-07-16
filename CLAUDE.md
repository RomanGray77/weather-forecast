# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project status

This repository currently contains only planning documents (`README.md`, `concept.md`) and a bare `go.mod`. No Go source code has been written yet. When implementing, follow the plan in `concept.md` rather than inventing a different architecture.

## Module

- Module path: `github.com/RomanGray77/weather-forecast`
- Go version: 1.26.1 (see `go.mod`)

## Commands

Since there is no source code yet, standard Go tooling applies once files exist:

- `go run .` — run the server (once `main.go` exists)
- `go build ./...` — build all packages
- `go test ./...` — run all tests
- `go test ./... -run TestName` — run a single test
- `go vet ./...` — vet the code
- `go mod tidy` — sync `go.mod`/`go.sum` after adding dependencies (e.g. `chi`)

## Intended architecture (from concept.md)

This is a small learning project — a Go web app that fetches weather data and displays it:

1. **HTTP server**: a local [chi](https://github.com/go-chi/chi) router serving on port 8080.
2. **`/weather` page**: renders a simple HTML form where the user enters a city name and submits it.
3. **Handler logic**: on submit, the handler calls the [wttr.in](https://github.com/chubin/wttr.in/) API for the given city, equivalent to:
   ```
   curl 'wttr.in/<City>?format=j2' | jq '.weather[] | {date, avgtempC, sunrise: .astronomy[0].sunrise, sunset: .astronomy[0].sunset}'
   ```
   i.e. it should extract, per day: `date`, `avgtempC`, `astronomy[0].sunrise`, `astronomy[0].sunset`.
4. **Persistence**: the parsed result is saved to a local `weather_forecast.json` file.
5. **Response**: the same data is rendered back on the `/weather` page below the input form.

When implementing, keep the request → fetch-from-wttr.in → persist-to-JSON → render flow intact, since that round trip is the core of the app.
