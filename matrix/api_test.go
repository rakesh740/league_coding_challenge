package matrix

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	req := httptest.NewRequest("POST", "/echo", nil)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.csv")
	part.Write([]byte("1,2,3\n4,5,6\n7,8,9\n"))

	writer.Close()
	req.Body = ioutil.NopCloser(body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	Echo(w, req)

	resp := w.Result()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if bodyString != expected {
		t.Errorf("Expected %s, got %s", expected, bodyString)
	}
}

func TestFlatten(t *testing.T) {
	req := httptest.NewRequest("POST", "/flatten", nil)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.csv")
	part.Write([]byte("4,5,6\n1,2,3\n7,8,9\n"))

	writer.Close()
	req.Body = ioutil.NopCloser(body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	Flatten(w, req)

	resp := w.Result()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	// Check the response
	expected := "4,5,6,1,2,3,7,8,9"
	if strings.TrimSpace(bodyString) != expected {
		t.Errorf("Expected %s, got %s", expected, bodyString)
	}
}

func TestInvert(t *testing.T) {
	req := httptest.NewRequest("POST", "/invert", nil)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.csv")
	part.Write([]byte("4,5,6\n1,2,3\n7,8,9\n"))

	writer.Close()
	req.Body = ioutil.NopCloser(body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	Invert(w, req)

	resp := w.Result()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	// Check the response
	expected := "4,1,7\n5,2,8\n6,3,9\n"
	if bodyString != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, bodyString)
	}
}

// Helper function to create a mock CSV file
func createMockCSVFile(content string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "test.csv")
	if err != nil {
		return nil, err
	}
	_, err = part.Write([]byte(content))
	if err != nil {
		return nil, err
	}
	writer.Close()

	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

func TestSum(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"Non-Square Matrix", "1,2,3\n4,5,6\n", "error row count and col count not equal"},
		{"Invalid Integer", "1,2,3\n4,abc,6\n7,8,9\n", "error strconv.ParseInt: parsing \"abc\": invalid syntax: Invalid Input"},
		{"Success", "1,2,3\n4,5,6\n7,8,9\n", "45"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := createMockCSVFile(tt.input)
			if err != nil {
				t.Fatalf("Failed to create mock CSV file: %v", err)
			}

			w := httptest.NewRecorder()
			Sum(w, req)

			resp := w.Result()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := strings.TrimSpace(string(bodyBytes))

			if bodyString != tt.expectedOutput {
				t.Errorf("Expected '%s', got '%s'", tt.expectedOutput, bodyString)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"Non-Square Matrix", "1,2,3\n4,5,6\n", "error row count and col count not equal"},
		{"Invalid Integer", "1,2,3\n4,abc,6\n7,8,9\n", "error strconv.ParseInt: parsing \"abc\": invalid syntax: Invalid Input"},
		{"Success", "1,2,3\n4,5,6\n7,8,9\n", "362880"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := createMockCSVFile(tt.input)
			if err != nil {
				t.Fatalf("Failed to create mock CSV file: %v", err)
			}

			w := httptest.NewRecorder()
			Multiply(w, req)

			resp := w.Result()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := strings.TrimSpace(string(bodyBytes))

			if bodyString != tt.expectedOutput {
				t.Errorf("Expected '%s', got '%s'", tt.expectedOutput, bodyString)
			}
		})
	}
}
