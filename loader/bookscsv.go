package loader

import (
	"os"
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

// BooksLiteral 
var BooksLiteral []*BookData

// LoadCSV function for loading data from CSV
func LoadCSV() {
	csvfile, err := os.Open("./assets/books.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		avgRating, e1 := strconv.ParseFloat(record[3], 64)
		numPages, e2 := strconv.ParseInt(record[7], 10, 32)
		ratings, e3 := strconv.ParseInt(record[8], 10, 64)
		reviews, e4 := strconv.ParseInt(record[9], 10, 64)

		if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
			continue
		}

		//fmt.Printf("Question: %s Answer %s & %s\n", record[0], record[1], record[2])
		//fmt.Println(bookList)
		b := BookData{record[0], record[1], record[2], avgRating, record[4], record[5], record[6], int(numPages), int(ratings), int(reviews)}
		BooksLiteral = append(BooksLiteral, &b)
	}

}