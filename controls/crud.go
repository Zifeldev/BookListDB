package controls

import (
	"context"
	"example/db"

	// "github.com/jackc/pgx/v5"
	"fmt"
)

type book struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

func InsertBook(title, author string) error{
	_, err := db.Conn.Exec(context.Background(),"INSERT INTO  books(title,author) VALUES($1,$2)",title,author)
	if err != nil {
		fmt.Println("❌ Ошибка вставки:", err)
	}
	return err
}

func DeleteBook(id int) error{
	_,err := db.Conn.Exec(context.Background(),"DELETE FROM books where id=$1",id)

	if err != nil{
		fmt.Println("Error",err)
	}
	return err
}

func GetAllBooks() ([]book,error){
	rows, err := db.Conn.Query(context.Background(),"SELECT id,title,author FROM books")
	if err != nil{
		fmt.Println("Querry error",err)
		return nil, err
	}

	defer rows.Close()

	var books []book
	for rows.Next(){
		var b book
		if err := rows.Scan(&b.ID,&b.Title,&b.Author); err != nil{
			fmt.Println("Error with reading",err)
			return nil,err
		}
		books = append(books, b)
	}
	return books, nil
}

func UpdateBook(id int, newTitle, newAuthor string) error {
	_,err := db.Conn.Exec(context.Background(),"UPDATE books SET title=$1, author=$2 where id=$3",newTitle,newAuthor,id)
	if err != nil{
		fmt.Println("Error with update", err)
	}
	return err
}