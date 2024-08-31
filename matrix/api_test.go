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
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"Non-Square Matrix", "1,2,3\n4,5,6\n", "error row count and col count not equal"},
		{"Success", "1,2,3\n4,5,6\n7,8,9\n", "1,2,3\n4,5,6\n7,8,9"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := createTestRequest(tt.input)
			if err != nil {
				t.Fatalf("Failed to create mock CSV file: %v", err)
			}

			w := httptest.NewRecorder()
			Echo(w, req)

			resp := w.Result()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := strings.TrimSpace(string(bodyBytes))

			if bodyString != tt.expectedOutput {
				t.Errorf("Expected '%s', got '%s'", tt.expectedOutput, bodyString)
			}
		})
	}
}

func TestFlatten(t *testing.T) {

	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"Non-Square Matrix", "1,2,3\n4,5,6\n", "error row count and col count not equal"},
		{"Success", "4,5,6\n1,2,3\n7,8,9\n", "4,5,6,1,2,3,7,8,9"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := createTestRequest(tt.input)
			if err != nil {
				t.Fatalf("Failed to create mock CSV file: %v", err)
			}

			w := httptest.NewRecorder()
			Flatten(w, req)

			resp := w.Result()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := strings.TrimSpace(string(bodyBytes))

			if bodyString != tt.expectedOutput {
				t.Errorf("Expected '%s', got '%s'", tt.expectedOutput, bodyString)
			}
		})
	}
}

func TestInvert(t *testing.T) {

	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"Non-Square Matrix", "1,2,3\n4,5,6\n", "error row count and col count not equal"},
		{"Success", "1,2,3\n4,5,6\n7,8,9\n", "1,4,7\n2,5,8\n3,6,9"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := createTestRequest(tt.input)
			if err != nil {
				t.Fatalf("Failed to create mock CSV file: %v", err)
			}

			w := httptest.NewRecorder()
			Invert(w, req)

			resp := w.Result()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := strings.TrimSpace(string(bodyBytes))

			if bodyString != tt.expectedOutput {
				t.Errorf("Expected '%s', got '%s'", tt.expectedOutput, bodyString)
			}
		})
	}
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
			req, err := createTestRequest(tt.input)
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
			req, err := createTestRequest(tt.input)
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

// Helper function to create a request
func createTestRequest(content string) (*http.Request, error) {
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
