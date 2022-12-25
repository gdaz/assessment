package expense

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewExpense(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "anuchito smoothie",
		"amount": 99,
		"note": "kbtg",
		"tags": ["food", "beverage"]
	}`)
	var eb ExpenseBody

	res := request(http.MethodPost, "http://localhost:2565/expenses", body)
	err := res.Decode(&eb)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, eb.Id)
	assert.Equal(t, "anuchito smoothie", eb.Title)
	assert.Equal(t, int64(99), eb.Amount)
	assert.Equal(t, 2, len(eb.Tags))
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	// req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
