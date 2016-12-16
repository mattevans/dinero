package dinero

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	packageVersion = "0.1.0"
	backendURL     = "http://openexchangerates.org"
	userAgent      = "dinero/" + packageVersion
)

// Client holds a connection to the OXR API.
type Client struct {
	client     *http.Client
	AppID      string
	UserAgent  string
	BackendURL *url.URL

	// Services used for communicating with the API.
	Rates  *RatesService
	Update *UpdateService
	Cache  *CacheService
}

type service struct {
	client *Client
}

// Response is a OXR response. This wraps the standard http.Response
// returned from the OXR API.
type Response struct {
	*http.Response
	ErrorCode int64
	Message   string
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	*http.Response
	ErrorCode   int64  `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%d %v", r.Response.StatusCode, r.Description)
}

// NewClient creates a new Client with the appropriate connection details and
// services used for communicating with the API.
func NewClient(oxrAppID string) *Client {
	// Init new http.Client.
	httpClient := http.DefaultClient

	// Parse BE URL.
	baseURL, _ := url.Parse(backendURL)

	c := &Client{
		client:     httpClient,
		BackendURL: baseURL,
		UserAgent:  userAgent,
		AppID:      oxrAppID,
	}

	c.Update = &UpdateService{client: c}
	c.Rates = &RatesService{client: c}
	c.Cache = &CacheService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlPath,
// which will be resolved to the BackendURL of the Client.
func (c *Client) NewRequest(method, urlPath string, body interface{}) (*http.Request, error) {
	// Append out OXR App ID to URL, :-(
	urlPath = fmt.Sprintf("%s&app_id=%s", urlPath, c.AppID)

	// Parse our URL.
	rel, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	// Resolve to absolute URI.
	u := c.BackendURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Create the request.
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Add our libraries UA.
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in 'v', or returned as an error if an API (if found).
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	// Wrap our response.
	response := &Response{Response: resp}

	// Check for any errors that may have occurred.
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(&v)
			if err != nil {
				return nil, err
			}
		}

	}

	return response, err
}

// CheckResponse checks the API response for errors. A response is considered an
// error if it has a status code outside the 200 range. API error responses map
// to ErrorResponse.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}

	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}
	return errorResponse
}
