// Package hem provides Hashed Email (HEM) generation following UID2/LiveRamp
// industry standards.
//
// The normalization steps applied before hashing are:
//
//   - Remove all whitespace.
//   - Lowercase.
//   - For gmail.com / googlemail.com only: strip any "+" suffix from the
//     local-part and remove all "." characters from the local-part.
//   - Canonicalize googlemail.com to gmail.com.
//
// The result is the lowercase hex-encoded SHA-256 of the normalized,
// UTF-8 encoded address.
package hem

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"unicode"
)

var gmailDomains = map[string]struct{}{
	"gmail.com":      {},
	"googlemail.com": {},
}

// normalizeEmail returns the canonical form of raw used as the SHA-256 input.
// It returns an error when the address is empty or not a single-"@" address.
func normalizeEmail(raw string) (string, error) {
	if raw == "" {
		return "", errors.New("email address must not be empty")
	}

	email := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, raw)

	if email == "" {
		return "", errors.New("email address must not be blank")
	}

	atCount := strings.Count(email, "@")
	if atCount == 0 {
		return "", errors.New("invalid email format: missing '@'")
	}
	if atCount > 1 {
		return "", errors.New("invalid email format: multiple '@' characters")
	}

	atIndex := strings.IndexByte(email, '@')
	local := email[:atIndex]
	domain := email[atIndex+1:]

	if local == "" {
		return "", errors.New("invalid email format: empty local-part")
	}
	if domain == "" {
		return "", errors.New("invalid email format: empty domain")
	}

	if _, ok := gmailDomains[domain]; ok {
		if i := strings.IndexByte(local, '+'); i >= 0 {
			local = local[:i]
		}
		local = strings.ReplaceAll(local, ".", "")
		domain = "gmail.com"
		if local == "" {
			return "", errors.New("invalid email format: empty local-part after Gmail normalization")
		}
	}

	return local + "@" + domain, nil
}

// GenerateHEM returns the lowercase hex SHA-256 digest of the normalized,
// UTF-8 encoded email address.
func GenerateHEM(raw string) (string, error) {
	normalized, err := normalizeEmail(raw)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256([]byte(normalized))
	return hex.EncodeToString(sum[:]), nil
}
