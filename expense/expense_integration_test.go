package expense

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
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

	res := request(http.MethodPost, uri(), body)
	err := res.Decode(&eb)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, eb.Id)
	assert.Equal(t, "anuchito smoothie", eb.Title)
	assert.Equal(t, int64(99), eb.Amount)
	assert.Equal(t, 2, len(eb.Tags))
}

func TestGetExpenseByID(t *testing.T) {
	var eb ExpenseBody

	res := request(http.MethodGet, uri("1"), nil)
	err := res.Decode(&eb)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, eb.Id)
	assert.Equal(t, "apple smoothie", eb.Title)
	assert.Equal(t, "no discount", eb.Note)
	assert.Equal(t, int64(89), eb.Amount)
	assert.Equal(t, 1, len(eb.Tags))
}

func TestGetAllExpense(t *testing.T) {
	var eb []ExpenseBody

	res := request(http.MethodGet, uri(), nil)
	err := res.Decode(&eb)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, len(eb))
}

func TestUpdateExpenseByID(t *testing.T) {
	body := bytes.NewBufferString(`{
		"id": 1,
		"title": "anuchito smoothie",
		"amount": 99,
		"note": "kbtg",
		"tags": ["InW"]
	}`)
	var eb ExpenseBody

	res := request(http.MethodPost, uri(), body)
	err := res.Decode(&eb)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, eb.Id)
	assert.Equal(t, "anuchito smoothie", eb.Title)
	assert.Equal(t, int64(99), eb.Amount)
	assert.Equal(t, 1, len(eb.Tags))
	assert.Equal(t, "InW", eb.Tags[0])
}

func uri(paths ...string) string {
	host := "http://localhost:2565/expenses"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
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
