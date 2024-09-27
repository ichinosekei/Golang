package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

type First_generator struct {
	active bool
	which_generator
}
type Second_generator struct {
	active bool
	which_generator
}

type which_generator interface {
	Generate_id(title string) string
	Change_generator(l storageSlice)
	//change_generator(generator which_generator)

}

type Book struct {
	id     string
	Title  string
	Author string
	//wh_gen int
}

type storageMap struct {
	books map[int]Book
	Library
}

type storageSlice struct {
	books []Book
	//Seacher
	Library
}

type Library interface {
	book_s() []Book
	Change_generator(generator which_generator, archieve Library)
	//ChangeStorage()
	//Add(book Book, generator which_generator)
}

func (l *storageSlice) book_s() []Book {
	return l.books
}

func (mapp *storageMap) book_s() []Book {

	array := make([]Book, 0, len(mapp.books))
	for i, _ := range mapp.books {
		array = append(array, mapp.books[i])
	}
	return array
}

func (First_generator) Generate_id(title string) string {
	hash := sha1.New()
	hash.Write([]byte(title))
	return fmt.Sprintf("%x", hash.Sum(nil))
	//fmt.Println(book.id)
}

func (Second_generator) Generate_id(title string) string {
	hash := md5.New()
	hash.Write([]byte(title))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (l *storageSlice) Add(book Book, generator which_generator) {
	book.id = generator.Generate_id(book.Title)
	l.books = append(l.books, book)
}

func (mapp *storageMap) Add(book Book, generator which_generator) {
	if mapp.books == nil {
		mapp.books = make(map[int]Book)
	}
	book.id = generator.Generate_id(book.Title)
	mapp.books[len(mapp.books)] = book

}

func ChangeGenerator(generator which_generator, archieve Library) {
	for i, book := range archieve.book_s() {
		archieve.book_s()[i].id = generator.Generate_id(book.Title)
	}
}

func (v *storageSlice) ChangeStorage() storageMap {
	books := map[int]Book{}
	for i, book := range v.books {
		books[i] = book
	}
	return storageMap{books: books}

}

func (v *storageMap) ChangeStorage() storageSlice {
	books := make([]Book, 0, len(v.books))
	for _, book := range v.books {
		books = append(books, book)
	}
	return storageSlice{books: books}
}

//func ChangeStorage(ArchieveOld *Library) Library {
//    unpoint := *ArchieveOld
//	switch v := unpoint.(type) {
//	case *storageSlice:
//		//storageMap{books: map[int]Book{},nil}
//		ArchieveOld = (v.ChangetoMap())
//		ArchieveOld = v.ChangetoMap()
//		println(v.books)
//		return ArchieveOld
//
//	case *storageMap:
//		v.ChangeToMap()
//		return v.ChangetoSlice()
//	}
//
//	//println("not change")
//	return ArchieveOld
//}

func find(id string, archieve Library) (Book, bool) {
	for _, book := range archieve.book_s() {
		if book.id == id {
			return book, true
		}
	}
	return Book{}, false
}

func Search(name string, generator which_generator, archieve Library) (Book, bool) {
	var book Book
	var ok bool
	book, ok = find(generator.Generate_id(name), archieve)
	if ok == true {
		return book, true
	}
	return Book{}, false

}

func main() {
	hash := sha1.New()
	hash.Write([]byte("Book1"))
	s1 := fmt.Sprintf("%x", hash.Sum(nil))

	hash = sha1.New()
	hash.Write([]byte("Book2"))
	s2 := fmt.Sprintf("%x", hash.Sum(nil))

	hash = sha1.New()
	hash.Write([]byte("Book3"))
	s3 := fmt.Sprintf("%x", hash.Sum(nil))

	hash = sha1.New()
	hash.Write([]byte("Book4"))
	s4 := fmt.Sprintf("%x", hash.Sum(nil))

	hash = sha1.New()
	hash.Write([]byte("Book5"))
	s5 := fmt.Sprintf("%x", hash.Sum(nil))

	//fmt.Sprintf("%x", sha1.New().Write([]byte("Book1")).Sum(nil))
	Book1 := Book{s1, "Book1", "Author1"}
	Book2 := Book{s2, "Book2", "Author2"}
	Book3 := Book{s3, "Book3", "Author3"}
	Book4 := Book{s4, "Book4", "Author4"}
	//Book5 := Book{fmt.Sprintf("%x", sha1.New().Write([]byte("Book5")).Sum(nil)), "Book5", "Author5"}
	storage1 := storageSlice{[]Book{Book1, Book2, Book3, Book4}, nil}
	book, ok := (Search("Book5", First_generator{}, &storage1))

	storage1.Add(Book{s5, "Book5", "Author5"}, First_generator{})

	println(storage1.books[4].id)
	ChangeGenerator(Second_generator{}, &storage1)
	println(storage1.books[4].id)

	book, ok = (Search("Book3", Second_generator{}, &storage1))
	if ok == true {
		println(book.id)
	} else {
		println("not found")
	}
	///////// map
	println("map")
	mapp := map[int]Book{
		0: Book{s1, "Book1", "Author1"},
		1: Book{s2, "Book2", "Author2"},
		2: Book{s3, "Book3", "Author3"},
		3: Book{s4, "Book4", "Author4"},
		4: Book{s5, "Book5", "Author5"},
	}
	storage2 := storageMap{mapp, nil}
	storage2.Add(Book{s5, "Book5", "Author5"}, First_generator{})
	println(storage2.books[5].id)
	ChangeGenerator(Second_generator{}, &storage2)

	////// change storageSlice
	println("change storageSlice")
	storage3 := storageSlice{[]Book{Book1, Book2, Book3, Book4}, nil}
	//r := storageMap{map[int]Book{}, nil}
	def := storage3.ChangeStorage()
	book, ok = (Search("Book3", First_generator{}, &def))
	if ok == true {
		println(book.id)
	} else {
		println("not found")
	}

}
