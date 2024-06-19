/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/numtide/nixos-facter/cmd"
	"github.com/ungerik/go-sysfs"
	"log"
)

func printInfo(obj sysfs.Object) {
	log.Printf("Object: %s", obj.Name())
	for _, attr := range obj.Attributes() {
		value, err := attr.Read()
		if err != nil {
			log.Printf("%s: error(%s)", attr.Path, err.Error())
		} else {
			log.Printf("%s: %s", attr.Path, value)
		}
	}

	for _, sub := range obj.SubObjects() {
		printInfo(sub)
	}
}

func main() {

	//for _, obj := range sysfs.Block.Objects() {
	//	printInfo(obj)
	//}

	cmd.Execute()
}
