package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type state struct {
	id               int
	name             string
	abbreviation     string
	censusRegionName string
}

func parseState(columns map[string]int, record []string) (*state, error) {
	// because struct cannot be nil in order to check for error, all struct can be
	// is zero or empty string, therefore, a pointer is used to, because pointer
	// can be nil.
	id, err := strconv.Atoi(record[columns["id"]])
	name := record[columns["name"]]
	abbreviation := record[columns["abbreviation"]]
	censusRegionName := record[columns["census_region_name"]]
	if err != nil {
		return nil, err
	}

	return &state{
		id:               id,
		name:             name,
		abbreviation:     abbreviation,
		censusRegionName: censusRegionName,
	}, nil
}

func main() {
	// open file
	f, err := os.Open("state_table.csv")
	if err != nil {
		log.Fatalln("error opning a file", err)
	}
	defer f.Close()

	stateLookup := map[string]state{}
	// parse a csv file
	csvReader := csv.NewReader(f)
	columns := make(map[string]int)
	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		if rowCount == 0 {
			for idx, column := range record {
				columns[column] = idx
			}
		} else {
			state, err := parseState(columns, record)
			if err != nil {
				log.Fatalln(err)
			}
			stateLookup[state.abbreviation] = *state

			//log.Println(state)
		}

	}

	if len(os.Args) < 2 {
		log.Fatalln("expected state abbreviation")

	}

	abbreviation := os.Args[1]
	state, ok := stateLookup[abbreviation]

	if !ok {
		log.Fatalln("invalid abbreviation")
	}

	fmt.Println(state)
}
