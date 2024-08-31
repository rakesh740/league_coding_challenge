package matrix

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	rowCount := len(records)
	colCount := len(records[0])
	if rowCount != colCount {
		w.Write([]byte(fmt.Sprintf("error row count and col count not equal")))
		return
	}

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

func Flatten(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	rowCount := len(records)
	colCount := len(records[0])
	if rowCount != colCount {
		w.Write([]byte(fmt.Sprintf("error row count and col count not equal")))
		return
	}

	var response string
	for i, row := range records {
		for j, s := range row {
			if i == len(records)-1 && j == len(row)-1 {
				response = fmt.Sprintf("%s%s", response, s)
				break
			}
			response = fmt.Sprintf("%s%s,", response, s)
		}
	}
	fmt.Fprint(w, response)
}

func Invert(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	rowCount := len(records)
	colCount := len(records[0])
	if rowCount != colCount {
		w.Write([]byte(fmt.Sprintf("error row count and col count not equal")))
		return
	}

	var response string
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {

			if j == colCount-1 {
				response = fmt.Sprintf("%s%s", response, records[j][i])
				break
			}
			response = fmt.Sprintf("%s%s,", response, records[j][i])
		}
		response = fmt.Sprintf("%s\n", response)
	}

	fmt.Fprint(w, response)
}

func Sum(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	rowCount := len(records)
	colCount := len(records[0])
	if rowCount != colCount {
		w.Write([]byte(fmt.Sprintf("error row count and col count not equal")))
		return
	}

	var response int64
	for _, row := range records {
		for _, s := range row {
			d, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("error %s: Invalid Input", err.Error())))
				return
			}
			response = response + d
		}
	}

	fmt.Fprint(w, response)
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	rowCount := len(records)
	colCount := len(records[0])
	if rowCount != colCount {
		w.Write([]byte(fmt.Sprintf("error row count and col count not equal")))
		return
	}

	var response int64 = 1
	for _, row := range records {
		for _, s := range row {
			d, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("error %s: Invalid Input", err.Error())))
				return
			}
			response = response * d
		}
	}

	fmt.Fprint(w, response)
}
