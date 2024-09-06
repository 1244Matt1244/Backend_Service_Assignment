```bash
#!/bin/bash

# Fetch Magic the Gathering cards
echo "Fetching Magic the Gathering cards..."
if ! curl -s "https://api.magicthegathering.io/v1/cards" -o mtg_cards.json; then
    echo "Failed to fetch Magic the Gathering cards"
    exit 1
fi

# Import cards into the database
echo "Importing cards into the database..."
go run main.go import-mtg mtg_cards.json

# Fetch traffic camera data
echo "Fetching traffic camera data..."
if ! curl -s "https://docs.google.com/spreadsheets/d/your-spreadsheet-id/export?format=csv" -o cameras.csv; then
    echo "Failed to fetch camera data"
    exit 1
fi

# Import cameras into the database
echo "Importing cameras into the database..."
go run main.go import-cameras cameras.csv

echo "Import completed!"
