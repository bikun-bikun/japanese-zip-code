package japanese_zip_code

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const defaultUrl = "https://zip-cloud.appspot.com/api/search"

type Client struct {
	Url string
	*http.Client
}

func NewClient() *Client {
	return &Client{defaultUrl, http.DefaultClient}
}

type resultZipCode struct {
	Message string    `json:"message,omitempty"`
	Results []ZipCode `json:"results,omitempty"`
	Status  int       `json:"status,omitempty"`
}

type ZipCode struct {
	Zipcode  string `json:"zipcode,omitempty"`
	Prefcode string `json:"prefcode,omitempty"`
	Address1 string `json:"address1,omitempty"`
	Address2 string `json:"address2,omitempty"`
	Address3 string `json:"address3,omitempty"`
	Kana1    string `json:"kana1,omitempty"`
	Kana2    string `json:"kana2,omitempty"`
	Kana3    string `json:"kana3,omitempty"`
}

func (c *Client) GetZipCode(zipCode string) ([]ZipCode, error) {
	z, err := excludeHyphen(zipCode)
	if err != nil {
		return nil, err
	}

	resp, err := c.Get(fmt.Sprintf("%v?zipcode=%v", c.Url, z))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("ZipCloud API is not abailable.")
	}

	var r resultZipCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	if r.Status != http.StatusOK {
		return nil, errors.New(r.Message)
	}

	return r.Results, nil
}

func excludeHyphen(zipCode string) (string, error) {
	ss := strings.Split(zipCode, "-")
	if len(ss) > 2 {
		return "", errors.New("The zip code format is invalid.")
	}

	z := strings.Join(ss, "")
	if len(z) != 7 {
		return "", errors.New("The number of digits in the zip code is invalid.")
	}

	if _, err := strconv.Atoi(z); err != nil {
		return "", errors.New("Bat zip code.")
	}

	return z, nil
}
