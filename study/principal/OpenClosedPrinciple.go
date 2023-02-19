package principal

import (
    "fmt"
    "io"
    "os"
)

/*
开闭原则

Go 是使用组合来实现的继承
在 Go 中没有 extends 关键字，Go 使用 interface 实现的功能叫组合
Go 使用组合来代替的继承
 */

type IBook interface {
    GetName() string
    GetPrice() int
}

type Book struct {
    Name string
    Price int
}

func (n *Book) GetName() string {
    return n.Name
}

func (n *Book) GetPrice() int {
    return n.Price
}

type Novel struct {
    *Book
}

// OffNovel 打五折
type OffNovel struct {
    *Book
}

func (n *OffNovel) GetPrice() int {
    return n.Price / 2
}

func PrintBook(n IBook) {
    fmt.Println(n.GetName())
    fmt.Println(n.GetPrice())
}

// main
func mainBck() {
    book := &Book{
        Name:  "hello",
        Price: 10,
    }

    novel := &Novel{
      book,
    }
    PrintBook(novel)

    n := &OffNovel{
        book,
    }
    PrintBook(n)

    // // Save writes the contents of doc to the file f.
    // func Save(f *os.File, doc *Document) error
    //
    // // Save writes the contents of doc to the supplied ReadWriter.
    // func Save(rw io.ReadWriter, doc *Document) error
    //
    // // Save writes the contents of doc to the supplied Writer.
    // func Save(w io.Writer, doc *Document) error
}
