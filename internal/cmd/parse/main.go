package main

import (
	"fmt"
	"os"

	pgq "github.com/lfittl/pg_query_go"
)

func main() {
	query := `INSERT INTO users (id,level) VALUES (1,0) ON CONFLICT (id) DO UPDATE SET level = users.level + 1`
	// query := `INSERT INTO customers (name, email) VALUES ('Microsoft', 'hotline@microsoft.com') ON CONFLICT (name) DO NOTHING`
	// query := `SELECT a.id,first_name,last_name FROM customer AS a,laptops AS l INNER JOIN payment ON payment.id = a.id LEFT JOIN people ON people.id = a.id JOIN cars ON cars.id = a.id`
	// query := `SELECT * FROM weather`
	// query := `SELECT city, (temp_hi+temp_lo)/2 AS temp_avg, date FROM weather`
	// query := `SELECT 4 OPERATOR(pg_catalog.*) 4`
	// query := `SELECT * FROM weather WHERE city = 'San Francisco' AND prcp > 0.0`
	// query := `SELECT * FROM weather WHERE city = 'San Francisco'`
	// query := `SELECT * FROM weather ORDER BY city, temp_lo`
	// query := `SELECT DISTINCT ON (bcolor) bcolor,fcolor FROM t1 ORDER BY bcolor,fcolor`
	// query := `SELECT DISTINCT city FROM weather`
	// query := `SELECT blue = 11`
	// query := `SELECT max(11, 20) = 20`
	// query := `SELECT '100'::integer != 100`
	// query := `SELECT a.customer_id[0].blue AS c, NULL, 1 = 2,email,amount FROM table_name`
	// query := `INSERT INTO contacts (contact_id,first_name) VALUES (NULL,'John')`
	// query := `INSERT INTO users (firstname, lastname) VALUES ('Joe', 'Cool') RETURNING id, firstname`
	// query := `DELETE FROM link USING link_tmp WHERE link.id = link_tmp.id`
	// query := `UPDATE products SET blue.price[1] = price * 1.1 WHERE price <= 99.99 RETURNING name, price AS new_price`
	// query := `UPDATE stock SET retail = stock_backup.retail FROM stock_backup WHERE stock.isbn = stock_backup.isbn`
	// query := `UPDATE stock SET retail = (cost * ((retail / cost) + 0.1::numeric))`
	// query := `SELECT isbn, retail, cost FROM stock ORDER BY isbn ASC, cost DESC, retail USING > LIMIT 3`
	// query := `SELECT * FROM stock OFFSET 33 + 1 LIMIT 3`

	tree, err := pgq.ParseToJSON(query)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "%s\n", query)
	fmt.Printf("%s\n", tree)
}
