package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type Name struct {
	ID              string `json:"id"`
	Arabic          string `json:"arabic"`
	Transliteration string `json:"transliteration"`
	MeaningShaykh   string `json:"meaning_shaykh"`
	Explanation     string `json:"explanation"`
}

func main() {
	ctx := context.Background()

	db, err := mustInitDB(ctx)
	if err != nil {
		log.Fatalf("error connecting to db %s", err)
	}

	csvPath := os.Getenv("CSV_PATH")
	if csvPath == "" {
		csvPath = "./names.csv"
	}
	names, err := readCSVData(csvPath)
	if err != nil {
		log.Fatalf("error reading CSV %s", err)
	}

	if err := writeToFirestore(names, db, ctx); err != nil {
		log.Fatalf("error writing to DB %s", err)
	}
	log.Print("Upload Complete")
}

func mustInitDB(ctx context.Context) (*firestore.Client, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}

	db, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func readCSVData(path string) (*[]Name, error) {
	log.Printf("Reading CSV from %s", path)
	csv_file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer csv_file.Close()

	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	names := make([]Name, len(records))
	for i, rec := range records {
		names[i] = Name{
			ID:              rec[0],
			Arabic:          rec[1],
			Transliteration: rec[2],
			MeaningShaykh:   rec[3],
			Explanation:     rec[4],
		}
	}

	return &names, nil
}

func writeToFirestore(names *[]Name, db *firestore.Client, ctx context.Context) error {
	for _, n := range *names {
		if _, _, err := db.Collection("names").Add(ctx, n); err != nil {
			return err
		}
	}
	return nil
}
