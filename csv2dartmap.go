package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"log"
	"bytes"
	"time"
)

func generateDartMapString(records [][]string) string {
	var buf bytes.Buffer
	buf.WriteString("// generated " + time.Now().String() + "\n")
	buf.WriteString("Map<String, Map<String, String>> localized_values = {\n")

	if len(records) > 0 {
		
		header := records[:1]
		data := records[1:]
		
		for headerIdx, lang := range header[0][1:] {
			buf.WriteString("\t'")
			buf.WriteString(lang)
			buf.WriteString("': {\n")
			
			for _, row := range data {
				buf.WriteString("\t\t'")
				buf.WriteString(row[0])
				buf.WriteString("': '")
				buf.WriteString(row[headerIdx+1])
				buf.WriteString("',\n")				
			}
			buf.WriteString("\t},\n")
		} 		
	}
		
	buf.WriteString("};")
	return buf.String()
}

func csv2dartmap(srcName string, destName string)(err error) {
	file, err := os.Open(srcName)
	if err != nil {
		return
	}

	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	file.Close()

	file, err = os.Create(destName)
	if err != nil {
		return
	}
	
	dartString := generateDartMapString(records)
	_, err = fmt.Fprintf(file, dartString)	
	return
}

func main() {
	// File open
	err := csv2dartmap("./localized_values.csv", "./generated.dart")
	if err != nil {
		log.Fatal(err)
	}
}