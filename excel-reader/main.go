package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"math"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	fileName := "data/laporan-keuangan-pembangunan-masjid.xlsx"
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		panic(err)
	}

	data := []map[string]interface{}{}
	periode := map[string]string{}

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == "Penerimaan" || sheet.Name == "Pengeluaran" {
			for r, row := range sheet.Rows {
				if r > 0 {
					d := make(map[string]interface{})
					d["Keterangan"] = strings.Trim(row.Cells[1].String(), "")
					d["Status"] = sheet.Name

					jumlah, _ := row.Cells[2].Float()

					if sheet.Name == "Penerimaan" {
						d["Masuk"] = Round(jumlah)
						d["Keluar"] = 0
					} else {
						d["Masuk"] = 0
						d["Keluar"] = Round(jumlah)
					}

					tanggal, err := time.Parse(`2006\-01\-02`, row.Cells[0].String())
					if err != nil {
						panic(err)
					}

					d["Tanggal"] = tanggal.Format("2006-01-02")
					d["Periode"] = strings.ToUpper(tanggal.Format("January 2006"))

					data = append(data, d)

					if _, exist := periode[tanggal.Format("2006-01")+"-01"]; !exist {
						periode[tanggal.Format("2006-01")+"-01"] = d["Periode"].(string)
					}
				}
			}
		}
	}

	// laporan

	jsonData, _ := json.Marshal(data)
	jsonPrettyData, _ := prettyPrintJSON(jsonData)

	jsonFile, err := os.Create("data/laporan.json")
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	jsonFile.Write(jsonPrettyData)

	fmt.Println("JSON data written to:", jsonFile.Name())

	// periode

	periodeData, _ := json.Marshal(periode)
	periodePrettyData, _ := prettyPrintJSON(periodeData)

	periodeFile, err := os.Create("data/periode.json")
	if err != nil {
		panic(err)
	}

	defer periodeFile.Close()

	periodeFile.Write(periodePrettyData)

	fmt.Println("JSON data written to:", periodeFile.Name())
}

func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

func prettyPrintJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
