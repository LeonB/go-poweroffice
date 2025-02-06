package poweroffice

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-poweroffice/omitempty"
	"github.com/omniboost/go-poweroffice/utils"
)

func (c *Client) NewManualJournalVoucherReversePost() ManualJournalVoucherReversePost {
	r := ManualJournalVoucherReversePost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ManualJournalVoucherReversePost struct {
	client      *Client
	queryParams *ManualJournalVoucherReversePostQueryParams
	pathParams  *ManualJournalVoucherReversePostPathParams
	method      string
	headers     http.Header
	requestBody ManualJournalVoucherReversePostBody
}

func (r ManualJournalVoucherReversePost) NewQueryParams() *ManualJournalVoucherReversePostQueryParams {
	return &ManualJournalVoucherReversePostQueryParams{}
}

type ManualJournalVoucherReversePostQueryParams struct{}

func (p ManualJournalVoucherReversePostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ManualJournalVoucherReversePost) QueryParams() *ManualJournalVoucherReversePostQueryParams {
	return r.queryParams
}

func (r ManualJournalVoucherReversePost) NewPathParams() *ManualJournalVoucherReversePostPathParams {
	return &ManualJournalVoucherReversePostPathParams{}
}

type ManualJournalVoucherReversePostPathParams struct {
	VoucherID string `schema:"voucherId"`
}

func (p *ManualJournalVoucherReversePostPathParams) Params() map[string]string {
	return map[string]string{
		"voucherId": p.VoucherID,
	}
}

func (r *ManualJournalVoucherReversePost) PathParams() *ManualJournalVoucherReversePostPathParams {
	return r.pathParams
}

func (r *ManualJournalVoucherReversePost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ManualJournalVoucherReversePost) SetMethod(method string) {
	r.method = method
}

func (r *ManualJournalVoucherReversePost) Method() string {
	return r.method
}

func (r ManualJournalVoucherReversePost) NewRequestBody() ManualJournalVoucherReversePostBody {
	return ManualJournalVoucherReversePostBody{}
}

type ManualJournalVoucherReversePostBody struct{}

func (v ManualJournalVoucherReversePostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(v)
}

func (r *ManualJournalVoucherReversePost) RequestBody() *ManualJournalVoucherReversePostBody {
	return nil
}

func (r *ManualJournalVoucherReversePost) RequestBodyInterface() interface{} {
	return nil
}

func (r *ManualJournalVoucherReversePost) SetRequestBody(body ManualJournalVoucherReversePostBody) {
	r.requestBody = body
}

func (r *ManualJournalVoucherReversePost) NewResponseBody() *ManualJournalVoucherReversePostResponseBody {
	return &ManualJournalVoucherReversePostResponseBody{}
}

type ManualJournalVoucherReversePostResponseBody struct {
	Data struct {
		Data       any  `json:"data"`
		Success    bool `json:"success"`
		Validation struct {
			ErrorLogReference string `json:"errorLogReference"`
			Summary           string `json:"summary"`
			Exception         string `json:"exception"`
			ExceptionDetails  string `json:"exceptionDetails"`
			Fields            any    `json:"fields"`
			BatchItemErrors   []any  `json:"batchItemErrors"`
		} `json:"validation"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (r *ManualJournalVoucherReversePost) URL() *url.URL {
	u := r.client.GetEndpointURL("/Voucher/ManualJournalVoucher/Reverse/{{.voucherId}}/", r.PathParams())
	return &u
}

func (r *ManualJournalVoucherReversePost) Do() (ManualJournalVoucherReversePostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
