package utils

import (
	"bytes"
	"encoding/json"
	"eta/config"
	"eta/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// client := &http.Client{
//   CheckRedirect: redirectPolicyFunc,
// }

var base_url = GenerateConfigBasedOnEnv("INVOICING_URL")
var auth_url = GenerateConfigBasedOnEnv("AUTH_URL")
var client_id = GenerateConfigBasedOnEnv("CLIENT_ID")
var client_secret = GenerateConfigBasedOnEnv("CLIENT_SECRET")
var pos_secret = GenerateConfigBasedOnEnv("POS_SECRET")
var pos_version = config.Config("POS_VERSION")
var pos_serial = config.Config("POS_SERIAL")
var client = &http.Client{}
var location, _ = time.LoadLocation("Africa/Cairo")
var lastLoginExpiresAt time.Time
var token string

func GenerateConfigBasedOnEnv(key string) string {
	fullKey := fmt.Sprintf("%s_%s", key, config.Config("ENVIRONMENT"))
	return config.Config(fullKey)
}

func SetContentType(req *http.Request) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

func SetAuthorization(req *http.Request) {
	isTokenExpired := time.Now().Before(lastLoginExpiresAt)
	if isTokenExpired || token == "" {
		EtaLoginPos()
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+token)
}

func GenerateRequestBody() url.Values {
	data := url.Values{}
	data.Set("client_id", client_id)
	data.Set("client_secret", client_secret)
	data.Set("grant_type", "client_credentials")
	return data
}
func EtaLogin() (string, error) {
	var response model.EtaLoginResponse
	resource := "/connect/token"
	data := GenerateRequestBody()
	url := auth_url + resource

	r, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	SetContentType(r)
	resp, _ := client.Do(r)
	d, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(d, &response)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	lastLoginExpiresAt = time.Now().In(location).Add(1 * time.Hour)
	return response.AccessToken, nil
}

func EtaLoginPos() (string, error) {
	var response model.EtaLoginResponse
	resource := "connect/token"
	data := GenerateRequestBody()
	url := auth_url + resource

	r, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	SetContentType(r)
	r.Header.Add("posserial", pos_serial)
	r.Header.Add("pososversion", pos_version)
	r.Header.Add("presharedkey", pos_secret)
	resp, _ := client.Do(r)
	d, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(d, &response)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	lastLoginExpiresAt = time.Now().In(location).Add(1 * time.Hour)
	token = response.AccessToken
	return response.AccessToken, nil
}

func SignInvoices(invoices *[]model.Invoice) (*model.InvoiceSubmitRequest, error) {
	var doc model.InvoiceSubmitRequest
	jsonValue, _ := json.Marshal(invoices)

	resp, err := http.Post(config.Config("SIGNER_URL")+"sign", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &doc)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &doc, nil
}

func SubmitReceipt(document *model.ReceiptSubmitRequest) (*model.EtaSubmitInvoiceResponse, error) {
	client := &http.Client{}
	var response model.EtaSubmitInvoiceResponse

	jsonBody, err := json.Marshal(document)
	if err != nil {
		fmt.Println("error while parsing doc")
		return nil, err
	}
	resource := "receiptsubmissions"

	url := base_url + resource

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("error forming request")
		return nil, err
	}

	SetAuthorization(r)
	resp, _ := client.Do(r)
	d, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("error forming request")
		return nil, err
	}
	err = json.Unmarshal(d, &response)
	if err != nil {
		fmt.Println("error parsing response")
		fmt.Println(d)
		if resp.StatusCode == 401 {
			EtaLoginPos()
		}
		return nil, err
	}

	return &response, nil
}

func SerializeInvoice(invoice interface{}) string {
	invoiceReflector := reflect.ValueOf(invoice)
	if invoiceReflector.Kind() != reflect.Struct && invoiceReflector.Kind() != reflect.Slice {
		return fmt.Sprintf("%s%v%s", "\"", invoice, "\"")
	}
	var serializedString string
	for i := 0; i < invoiceReflector.NumField(); i++ {
		currentValue := invoiceReflector.Field(i).Interface()
		currentValueReflictor := reflect.ValueOf(currentValue)
		currentValueType := currentValueReflictor.Kind()
		currentValueKey := strings.ToUpper(invoiceReflector.Type().Field(i).Name)
		if currentValueType != reflect.Struct && currentValueType != reflect.Slice {
			serializedString += fmt.Sprintf("%s%s%s", "\"", currentValueKey, "\"")
			serializedString += SerializeInvoice(currentValue)
		}
		if currentValueType == reflect.Struct {
			serializedString += fmt.Sprintf("%s%s%s", "\"", currentValueKey, "\"")
			for j := 0; j < currentValueReflictor.NumField(); j++ {
				currentValue2 := currentValueReflictor.Field(j).Interface()
				currentValueKey2 := strings.ToUpper(currentValueReflictor.Type().Field(j).Name)
				serializedString += fmt.Sprintf("%s%s%s", "\"", currentValueKey2, "\"")
				serializedString += SerializeInvoice(currentValue2)
			}
		}
		if currentValueType == reflect.Slice {
			serializedString += fmt.Sprintf("%s%s%s", "\"", currentValueKey, "\"")
			slice, ok := currentValueReflictor.Interface().([]model.ItemData)
			if !ok {
				slice2, ok2 := currentValueReflictor.Interface().([]float64)
				if !ok2 {
					slice3, ok3 := currentValueReflictor.Interface().([]model.TaxTotals)
					for k := 0; k < len(slice3); k++ {
						serializedString += fmt.Sprintf("%s%s%s", "\"", currentValueKey, "\"")
						serializedString += SerializeInvoice(slice3[k])
						// serializedString += fmt.Sprintf("%v", currentValueReflictor.Interface())
					}
					if !ok3 {
						panic("cant convert slices")
					}
				} else {
					for k := 0; k < len(slice2); k++ {
						serializedString += fmt.Sprintf("%s%s%s", "\"", currentValueKey, "\"")
						serializedString += SerializeInvoice(slice2[k])
						// serializedString += fmt.Sprintf("%v", currentValueReflictor.Interface())
					}
				}

			} else {

				for k := 0; k < len(slice); k++ {
					itemReflector := reflect.ValueOf(slice[k])
					serializedString += fmt.Sprintf("%s%s%s", "\"", currentValueKey, "\"")
					// itemValue := itemReflector.Field(k).Interface()
					// itemValueKey := strings.ToUpper(itemReflector.Type().Field(k).Name)
					for j := 0; j < itemReflector.NumField(); j++ {

						itemValue2 := itemReflector.Field(j).Interface()
						itemValueKey2 := strings.ToUpper(itemReflector.Type().Field(j).Name)
						serializedString += fmt.Sprintf("%s%s%s", "\"", itemValueKey2, "\"")
						serializedString += SerializeInvoice(itemValue2)
					}

				}
			}

		}

	}
	return serializedString
}
