package store

import (
	"fmt"
	"testing"
)

func TestNewStore(t *testing.T) {
	db := NewStore("./db_cache")
	err := db.Save([]byte("abc"), []byte("alice"))

	fmt.Println("save", err)
	v := db.Get([]byte("abc"))
	fmt.Println(string(v))
	db.Close()



	db = NewStore("./db_cache")
	err = db.Save([]byte("abcd"), []byte("alice2"))

	fmt.Println("save", err)
	v = db.Get([]byte("abcd"))
	fmt.Println(string(v))
	db.Close()


}
