package utils

import (
	"bytes"
	"io"
	"net/http"
)

// ReadAndLogBody reads the request body, logs it, and returns the body as a byte slice.
// It also resets the body so it can be read again.
func ReadBody(r *http.Request) ([]byte, error) {
	// Read the body into a buffer
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Reset the body reader so it can be read again
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
