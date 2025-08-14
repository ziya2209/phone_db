package handler

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"unicode"

	"github.com/ziya2209/goproject/phonedb/internal/repo"
)

type dbResetter struct {
	d repo.Dao
}

func (s dbResetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reseting database")
	if err := s.d.Reset(); err != nil {
		fmt.Fprintln(w, err)
	}

}

func ResetDatabase(d repo.Dao) http.Handler {

	return dbResetter{d: d}
}
func AddPhoneNumber(d repo.Dao) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("adding phone numbers")
		file, err := os.Open("number.txt")
		if err != nil {
			panic(err)
		}
		reader := bufio.NewScanner(file)
		for reader.Scan() {
			num := reader.Text()
			n := normalize(num)
			err = d.AddPhoneNumber(n)
			if err != nil {
				fmt.Println("failed to store number into database")
				panic(err)
			}

		}

	}
}
func GetAllPhoneNumbers(d repo.Dao) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getting phone numbers")
		nums, err := d.GetAllPhoneNumbers()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		for _, num := range nums {
			fmt.Fprintln(w, num)

		}
	}
}

func normalize(n string) string {
	char := ""
	for _, v := range n {
		if unicode.IsNumber(v) {
			char += string(v)
		}
	}

	return char
}
