package main

import (
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"github.com/marcelocorreia/r2d2/converter"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/signal"
	"syscall"
)

var (
	generate     = kingpin.Command("generate", "Generates Stuff")
	convert      = kingpin.Command("convert", "Generates Stuff")
	generateType = generate.Flag("type", "type to be generate. E.g.: tfvars").Short('t').String()
	convertType  = convert.Flag("type", "type of convertion. E.g.: json2hcl, hcl2json").Required().Short('t').String()
)

func main() {
	gracefulKill()

	switch kingpin.Parse() {
	case "generate":
		generator(*generateType)
	case "convert":
		convertIt(*convertType)
	default:
		println("blah")
	}
}

func convertIt(convertType string) {
	switch convertType {
	case "json2hcl":
		println("converting")
		tf := converter.TFVarsManager{}
		tf.ToHCL()

	case "hcl2json":
		println("converting")
		tf := converter.TFVarsManager{}
		tf.ToJSON()

	}
}

func generator(generateType string) {
	switch generateType {
	case "tfvars":
		println("tfvars")
	default:
		println("not found")
	}
}

//
func gracefulKill() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Shutting down gracefully...")
		ct.ResetColor()
		defer fmt.Println("Done.")
		os.Exit(0)
	}()
}
