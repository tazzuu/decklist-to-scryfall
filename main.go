package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const (
	baseURL       = "https://scryfall.com/search"
	defaultFilter = "(game:paper) unique:prints"
	defaultRemove = "-t:basic"
)

// Regular expression to match card count prefix (e.g., "2 " or "2x ")
var cardCountRegex = regexp.MustCompile(`^\d+(?:x\s|\s)`)

func main() {
	var filter string
	var remove string
	flag.StringVar(&filter, "filter", defaultFilter, "Scryfall filter to append (e.g. '(game:paper) unique:prints')")
	flag.StringVar(&remove, "remove", defaultRemove, "Scryfall removal filter to append (e.g. '-type:basic')")
	flag.Parse()

	// If no file argument is given, read from stdin
	var input io.Reader = os.Stdin
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	// Read and process the input
	var cards []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Remove any leading card count (e.g., "2 " or "2x ")
		line = cardCountRegex.ReplaceAllString(line, "")

		// Replace any double quotes with single quotes to avoid breaking the query
		line = strings.ReplaceAll(line, "\"", "'")
		cards = append(cards, fmt.Sprintf("%q", line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	if len(cards) == 0 {
		fmt.Fprintln(os.Stderr, "No card names found in input")
		os.Exit(0)
	}

	// Process cards in chunks of 20
	const chunkSize = 20
	for i := 0; i < len(cards); i += chunkSize {
		end := i + chunkSize
		if end > len(cards) {
			end = len(cards)
		}
		chunk := cards[i:end]

		// Add ! prefix to each quoted card name
		prefixedChunk := make([]string, len(chunk))
		for j, card := range chunk {
			prefixedChunk[j] = "!" + card
		}
		// fmt.Printf("chunk: %q\n", prefixedChunk)

		// Build the query string for this chunk
		query := "(" + strings.Join(prefixedChunk, " or ") + ")" // Wrap OR conditions in parentheses
		if remove != "" {
			query = query + " " + remove // Apply removal filter first
		}
		if filter != "" {
			query = query + " " + filter // Apply other filters last
		}
		// fmt.Printf("query: %q\n", query)

		// Construct and encode the URL
		params := url.Values{}
		params.Set("as", "grid")
		params.Set("order", "name")
		params.Set("q", query)

		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
		fmt.Printf("Cards %d-%d:\n%s\n\n", i+1, end, finalURL)
	}
}
