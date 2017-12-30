package shortener

import (
	"reflect"
	"testing"
)

func TestGetUrl(t *testing.T) {

	url := "google.com"

	result, errorFunc := getUrl(url)
	if errorFunc != nil {
		if errorFunc.Error() != "Dados invalidos\n" {
			t.Log(errorFunc)
			t.Errorf("Era esperado a mensagem de erro dados invalidos")
		}
	} else {
		resultType := reflect.TypeOf(result)
		if resultType.Kind() != reflect.String {
			t.Errorf("O valor esperado era uma string, porém foi recebido %v", resultType.Kind())
		}
	}

}

func TestGetUrlForNilValue(t *testing.T) {

	url := ""

	result, errorFunc := getUrl(url)
	if errorFunc != nil {
		if errorFunc.Error() != "Dados invalidos\n" {
			t.Log(errorFunc)
			t.Errorf("Era esperado a mensagem de erro dados invalidos")
		}
	} else {
		resultType := reflect.TypeOf(result)
		if resultType.Kind() != reflect.String {
			t.Errorf("O valor esperado era uma string, porém foi recebido %v", resultType.Kind())
		}
	}

}

func TestShort(t *testing.T) {
	tests := []struct {
		url    string
		expect string
	}{
		{"google.com", "https://goo.gl/mR2d"},
		{"", "Dados invalidos\n"},
	}

	for _, test := range tests {
		c, err := Short(test.url)
		errorFunc := <-err
		if errorFunc.Error() != "false" {
			if errorFunc.Error() != test.expect {
				t.Errorf("Era esperado a mensagem %v mais vou recebido %v", test.expect, errorFunc.Error())
			}

		} else {
			result := <-c
			if result != test.expect {
				t.Errorf("O valor esperado era %v e o recebido foi %v", test.expect, result)
			}

		}
	}

}
