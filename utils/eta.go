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
)

// client := &http.Client{
//   CheckRedirect: redirectPolicyFunc,
// }

var base_url = "https://api.invoicing.eta.gov.eg/api/v1"

func EtaLogin() (string, error) {
	var response model.EtaLoginResponse
	// apiUrl := "https://id.preprod.eta.gov.eg"
	resource := "/connect/token"
	data := url.Values{}
	data.Set("client_id", "92fe559b-c17e-4275-a12e-132d34189ef1")
	data.Set("client_secret", "1e0c3a98-b4df-489b-b366-25e3aa5e28c6")
	data.Set("grant_type", "client_credentials")
	u, _ := url.ParseRequestURI(base_url)
	u.Path = resource
	urlStr := u.String()
	fmt.Println(urlStr)
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, "https://id.eta.gov.eg/connect/token", strings.NewReader(data.Encode())) // URL-encoded payload
	// r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(r)
	d, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(d, &response)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
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

func SubmitInvoice(authToken *string, document *model.InvoiceSubmitRequest) (*model.EtaSubmitInvoiceResponse, error) {
	client := &http.Client{}
	var response model.EtaSubmitInvoiceResponse

	jsonBody, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}
	// apiUrl := "https://api.preprod.invoicing.eta.gov.eg/api/v1"
	resource := "/documentsubmissions"
	// data.Set("client_id", "c70450b9-5b89-48dd-be15-9cf7629f7dd1")
	// data.Set("client_secret", "7825b824-841c-4f1a-81cd-d8eb60745ee6")
	// data.Set("grant_type", "client_credentials")
	u, _ := url.ParseRequestURI(base_url)
	u.Path = resource
	urlStr := u.String()
	fmt.Println(urlStr)
	r, err := http.NewRequest(http.MethodPost, "https://api.invoicing.eta.gov.eg/api/v1/documentsubmissions", bytes.NewBuffer(jsonBody)) // URL-encoded payload
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", "Bearer "+*authToken)
	r.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(r)
	d, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &response)
	if err != nil {
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
					panic("cant convert slices")
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
					itemValue := itemReflector.Field(k).Interface()
					itemValueKey := strings.ToUpper(itemReflector.Type().Field(k).Name)
					serializedString += fmt.Sprintf("%s%s%s", "\"", itemValueKey, "\"")
					serializedString += SerializeInvoice(itemValue)
				}
			}

		}

	}

	return serializedString
}

// if documentStructure is simple value type
// return """ + documentStructure.value + """
// end if

// var serializedString = ""

// foreach element in the structure:

// if element is not array type
// 	serializeString.Append (""" + element.name.uppercase + """)
// 	serializeString.Append ( Serialize(element.value) )
// end if

// if element is of array type
// 	serializeString.Append (""" + element.name.uppercase + """)
// 	foreach array element in element:
// 		// use below line for JSON because subelements of array in JSON do not have own names
// 		serializeString.Append (""" + element.name.uppercase + """)
// 		serializeString.Append ( Serialize(arrayelement.value) )
// 	end foreach
// end if

// end foreach

// return serializedString
