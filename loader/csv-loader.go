package loader

import (
	"log"
	"encoding/csv"
	"os"
	"strconv"
)

// BooksLiteral is a slice literal of Bookdata struct pointers, containing a subset of the real book data
var BooksLiteral []*BookData
// LoadCsvData from csv file 
func LoadCsvData(fileName string) {
	records := readCsvFile(fileName)
	for i,book := range records {	
		if i==0{
			// You can skip headers if your csv file contains headers
		}
		avgRating, e1 := strconv.ParseFloat(book[3], 64)	
		numPages,e2 := strconv.ParseInt(book[7], 10, 32)
		ratings,e3 := strconv.ParseInt(book[8], 10, 64)
		reviews,e4 := strconv.ParseInt(book[9], 10, 64)
		if e1 != nil || e2 !=nil || e3 !=nil || e4 !=nil{
			continue
		}
		b := BookData{book[0],book[1],book[2],avgRating,book[4],book[5],book[6],int(numPages),int(ratings),int(reviews)}
		// fmt.Println(b)
		BooksLiteral = append(BooksLiteral, &b)

	}
	
}



func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}