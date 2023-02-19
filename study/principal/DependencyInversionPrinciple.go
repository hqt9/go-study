package principal

import "fmt"

// IDriver 司机抽象接口
type IDriver interface {
    Drive(car ICar)
}

type Driver struct {}

func (d *Driver) Drive(car ICar) {
    car.Run()
}

// ICar 汽车抽象接口
type ICar interface {
    Run()
}

type Bmw struct {}

func (b *Bmw) Run() {
    fmt.Println("I drive BMW")
}

type Benz struct {}

func (z *Benz) Run() {
    fmt.Println("I drive Benz")
}

func mainBck() {
    driver := &Driver{}

    driver.Drive(&Bmw{})
    driver.Drive(&Benz{})
}
