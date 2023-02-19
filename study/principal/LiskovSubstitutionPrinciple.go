package principal

import (
    "encoding/binary"
    "fmt"
)

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

func Write(w ReadWriter, p []byte) (int, error) {
    _, err := w.Write(p)
    if err != nil {
        return 0, err
    }
    return w.Read(p)
}

type ReadContent struct {
    // Content string
}

func (content *ReadContent) Read(p []byte) (int, error)  {
    fmt.Println(p)
    return 1, nil
}

func (content *ReadContent) Write(p []byte) (int, error)  {
    data := binary.LittleEndian.Uint16(p)
    data2 := binary.LittleEndian.Uint32(p)
    data3 := binary.LittleEndian.Uint64(p)
    fmt.Println(data, data2, data3)
    return 1, nil
}

func mainBck() {

    content := &ReadContent{}

    _, err := Write(content, []byte("12345678"))
    if err != nil {
        return
    }
}
