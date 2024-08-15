package lysand

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"io"
	"net/http"
)

func (c *FederationClient) ValidateSignatureHeader(req *http.Request) (bool, error) {
	fedHeaders, err := ExtractFederationHeaders(req.Header)
	if err != nil {
		return false, err
	}

	// TODO: Fetch user from database instead of using the URI
	user, err := c.GetUser(req.Context(), fedHeaders.SignedBy)
	if err != nil {
		return false, err
	}

	body, err := copyBody(req)
	if err != nil {
		return false, err
	}

	v := Verifier{ed25519.PublicKey(user.PublicKey.PublicKey)}
	valid := v.Verify(req.Method, req.URL, body, fedHeaders)

	return valid, nil
}

func hashSHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func must[In any, Out any](fn func(In) (Out, error), v In) Out {
	out, err := fn(v)
	if err != nil {
		panic(err)
	}

	return out
}

func copyBody(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	if err := req.Body.Close(); err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}
