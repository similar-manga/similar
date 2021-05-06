/*
 * MangaDex API
 *
 * MangaDex is an ad-free manga reader offering high-quality images!  Here is some generic stuff about the API  # Authentication  You can login with `/auth/login` endpoints. It will return a JWT that remains for 15min and that have a 4h refresh token.  # Rating limits  We are using rating limits in order to avoid too much calls on our endpoints, here is how is it configured:  | Endpoint                | Calls            | Time frame                | |-------------------------|------------------|---------------------------| | `/account/create` | 1 | 60 minutes | | `/account/activate/{code}` | 30 | 60 minutes | | `/account/activate/resend`, `/account/recover`, `/account/recover/{code}` | 5 | 60 minutes | | `/auth/login` | 30 | 60 minutes | | `/auth/refresh` | 30 | 60 minutes | | `/chapter/{id}/read` | 300              | 10 minutes                | | `/upload/begin`, `/upload/{id}`, `/upload/{id}/commit` | 30               | 1 minute                  | | `PUT /chapter/{id}` | 10               | 1 minute                  | | `DELETE /chapter/{id}` | 10               | 1 minute                  | | `POST /manga` | 10               | 60 minutes                | | `PUT /manga/{id}` | 10               | 1 minute                  | | `DELETE /manga/{id}` | 10               | 10 minutes                | | `POST /group` | 10               | 60 minutes                | | `PUT /group/{id}` | 10               | 1 minute                  | | `DELETE /group/{id}` | 10               | 10 minutes                | | `POST /author` | 10               | 60 minutes                | | `PUT /author` | 10               | 1 minutes                 | | `DELETE /author/{id}` | 10               | 10 minutes                | | `POST /captcha/solve` | 10 | 10 minutes |  You can get details about how your rate limit is going by reading following headers:  | Header                  | Description                                              | |-------------------------|-----------------------------------------------------------| | X-RateLimit-Limit       | Number of max requests allowed in the current time period | | X-RateLimit-Remaining   | Number of remaining requests in the current time period   | | X-RateLimit-Retry-After | Timestamp of end of current time period as UNIX timestamp |  # Captchas  Some endpoints may require captchas to proceed, in order to slow down automated malicious traffic. Regular users might see a couple of captchas, based on the frequency of write requests or on certain endpoints like user signup.  Once an endpoint decides that a captcha needs to be solved, a 403 Forbidden response will be returned, where the error title is `captcha_required_exception`. The sitekey needed for recaptcha to function is provided in both the `X-Captcha-Sitekey` header field, as well as in the error context, specified as the `siteKey` parameter.  The captcha result of the client can either be passed into the repeated original request with the `X-Captcha-Result` header or alternatively the `POST /captcha/solve` endpoint can be called to solve this captcha. The time a solved captcha is remembered varies across different endpoints and can also be influenced by individual client behavior.  Authentication is not required for the `POST /captcha/solve` endpoint, captchas are tracked separately for client ip and user id. If you are logged in, you want to send the session token so you validate the captcha for your client ip and user id at the same time, but it is not required.  # Chapter pages processing  ## Pages processing  When you fetch a chapter response you'll have 4 fields that you need for pages processing:  | Field                       | Type     | Description        | |-----------------------------|----------|--------------------| | `data.id`                   | `string` | API identifier     | | `data.attributes.hash`      | `string` | MD@H identifier    | | `data.attributes.data`      | `array`  | High quality pages | | `data.attributes.dataSaver` | `array`  | Low quality pages  |  From this point you miss one thing: a MD@H backend server to get images from, to get that make a request to `GET /at-home/server/{data.id}`, it will return the server url to use.  Then to build the pages, you have to build your url as following: `https://{md@h server node}/data/{data.attributes.hash}/{data.attributes.data}`  Or if you want to use the low quality files: `https://{md@h server node}/data-saver/{data.attributes.hash}/{data.attributes.dataSaver}`  Here is an example of what the url could looks like:  | Field                       | Value | |-----------------------------|-------| | `data.id`                   | `000002b1-e8de-4281-9781-8e81e869f579` | | `data.attributes.hash`      | `caad0c22434276b9e3e56a78fe2e7993` | | `data.attributes.data`      | `[\"x1-a87ae6522fa5c244fd76985c7d953ccf3975bec66ce9b8e813549e642b38a47a.png\", ...]` | | `data.attributes.dataSaver` | `[\"x1-a1d3047dfccd77b3117a86ccf19a9c5403e09baec6a78893ed1d3825d2c71256.jpg\", ...]` |  As a \"fake\" MD@H node we'll use `https://s2.mangadex.org/` server.  So for high quality we'll have an URL like that: https://s2.mangadex.org/data/caad0c22434276b9e3e56a78fe2e7993/x1-a87ae6522fa5c244fd76985c7d953ccf3975bec66ce9b8e813549e642b38a47a.png  And for low quality: https://s2.mangadex.org/data-saver/caad0c22434276b9e3e56a78fe2e7993/x1-a1d3047dfccd77b3117a86ccf19a9c5403e09baec6a78893ed1d3825d2c71256.jpg  ## Report  In order to make everything works well, we keep statistics over MD@H nodes and how they perform. In order to keep theses statistics you have to post data for each page you fetch from a MD@H node.  Here is an example: ```curl POST https://api.mangadex.network/report {   \"url\": \"https://s2.mangadex.org/data/caad0c22434276b9e3e56a78fe2e7993/x1-a87ae6522fa5c244fd76985c7d953ccf3975bec66ce9b8e813549e642b38a47a.png\",   \"success\": true,   \"bytes\": 800000, // size of the loaded image   \"duration\": 213, // miliseconds to load the image   \"cached\": false, // X-Cache header of the MDAH node == 'HIT' ? } ```  # Static data  ## Manga publication demographic  | Value            | Description               | |------------------|---------------------------| | shonen           | Manga is a Shonen         | | shoujo           | Manga is a Shoujo         | | josei            | Manga is a Josei          | | seinen           | Manga is a Seinen         |  ## Manga status  | Value            | Description               | |------------------|---------------------------| | ongoing          | Manga is still going on   | | completed        | Manga is completed        | | hiatus           | Manga is paused           | | abandoned        | Manga has been abandoned  |  ## Manga reading status  | Value            | |------------------| | reading          | | on_hold          | | plan\\_to\\_read   | | dropped          | | re\\_reading      | | completed        |  ## Manga content rating  | Value            | Description               | |------------------|---------------------------| | safe             | Safe content              | | suggestive       | Suggestive content        | | erotica          | Erotica content           | | pornographic     | Pornographic content      |  ## CustomList visibility  | Value            | Description               | |------------------|---------------------------| | public           | CustomList is public      | | private          | CustomList is private     |  ## Relationship types  | Value            | Description                    | |------------------|--------------------------------| | manga            | Manga resource                 | | chapter          | Chapter resource               | | author           | Author resource                | | artist           | Author resource (drawers only) | | scanlation_group | ScanlationGroup resource       | | tag              | Tag resource                   | | user             | User resource                  | | custom_list      | CustomList resource            |
 *
 * API version: 5.0.0
 * Contact: mangadexstaff@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mangadex

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the MangaDex API API v5.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccountApi *AccountApiService

	AtHomeApi *AtHomeApiService

	AuthApi *AuthApiService

	AuthorApi *AuthorApiService

	CaptchaApi *CaptchaApiService

	ChapterApi *ChapterApiService

	CustomListApi *CustomListApiService

	FeedApi *FeedApiService

	InfrastructureApi *InfrastructureApiService

	LegacyApi *LegacyApiService

	MangaApi *MangaApiService

	ScanlationGroupApi *ScanlationGroupApiService

	SearchApi *SearchApiService

	UserApi *UserApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccountApi = (*AccountApiService)(&c.common)
	c.AtHomeApi = (*AtHomeApiService)(&c.common)
	c.AuthApi = (*AuthApiService)(&c.common)
	c.AuthorApi = (*AuthorApiService)(&c.common)
	c.CaptchaApi = (*CaptchaApiService)(&c.common)
	c.ChapterApi = (*ChapterApiService)(&c.common)
	c.CustomListApi = (*CustomListApiService)(&c.common)
	c.FeedApi = (*FeedApiService)(&c.common)
	c.InfrastructureApi = (*InfrastructureApiService)(&c.common)
	c.LegacyApi = (*LegacyApiService)(&c.common)
	c.MangaApi = (*MangaApiService)(&c.common)
	c.ScanlationGroupApi = (*ScanlationGroupApiService)(&c.common)
	c.SearchApi = (*SearchApiService)(&c.common)
	c.UserApi = (*UserApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
		if strings.Contains(contentType, "application/xml") {
			if err = xml.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		} else if strings.Contains(contentType, "application/json") {
			if err = json.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}
