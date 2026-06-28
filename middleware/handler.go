package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Printf("Succesfully connected db")

	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stocks

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("error at decoding response body %v", err)
	}

	insertId := insertStock(stock)
	res := response{
		ID:      insertId,
		Message: "stock inserted successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("some problem in converting the id %v", err)
	}

	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("issue in fetching the stock %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getStocks()

	if err != nil {
		log.Fatalf("issue in fetching all stock %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("some problem in converting the id %v", err)
	}

	var stock models.Stocks

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("error at decoding response body %v", err)
	}

	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("stock updated successfully , total rows affected are %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("some problem in converting the id %v", err)
	}

	deletedRows := deleteStock(int64(id))

	msg := fmt.Sprintf("stock deleted successfully , total rows affected are %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stocks) int64 {
	db := createConnection()
	defer db.Close()
	sqlstat := `INSERT INTO stocks (name,price,company) VALUES ($1,$2,$3) RETURNING stockid`
	var id int64

	err := db.QueryRow(sqlstat, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v ", err)
	}

	fmt.Printf("Inserted a single Row %v", id)
	return id
}

func getStock(id int64) (models.Stocks, error) {
	db := createConnection()
	defer db.Close()

	sqlstat := `SELECT * FROM stocks WHERE id=$1`

	row := db.QueryRow(sqlstat, id)

	var stock models.Stocks

	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Printf("No rows returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan row %v", err)
	}

	return stock, err

}

func getStocks() ([]models.Stocks, error) {
	db := createConnection()
	defer db.Close()

	var stocks []models.Stocks

	sqlstat := `SELECT * FROM stocks`

	rows, err := db.Query(sqlstat)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	for rows.Next() {
		var stock models.Stocks
		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}

		stocks = append(stocks, stock)
	}
	return stocks, err
}

func updateStock(id int64, stock models.Stocks) int64 {
	db := createConnection()
	defer db.Close()

	sqlstat := `UPDATE stocks SET name=$1,price=$2,company=$3 WHERE id=$1`

	res, err := db.Exec(sqlstat, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error at checking rows affected %v", err)
	}

	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()

	sqlstat := `DELETE FROM stocks WHERE id=$1`

	res, err := db.Exec(sqlstat, id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error at checking rows affected %v", err)
	}

	return rowsAffected
}
