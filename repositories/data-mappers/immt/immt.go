package immt

import (
	"encoding/gob"
	"fmt"
	"os"
)

const DATAFILE = "./data/immt/dataFile.gob"

var DATA = make(map[string]interface{})

func init() {
	load()
}

func Save() error {
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}
	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("Cannot create", DATAFILE)
		return err
	}
	defer func() {
		err = saveTo.Close()
	}()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("Cannot save to", DATAFILE)
		return err
	}
	return nil
}

func load() error {
	fmt.Println("Loading", DATAFILE)
	loadFrom, err := os.Open(DATAFILE)
	defer func() {
		err = loadFrom.Close()
	}()
	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}

	decoder := gob.NewDecoder(loadFrom)
	err = decoder.Decode(&DATA)
	return nil
}

func Add(k string, n interface{}) bool {
	if k == "" {
		return false
	}

	if Lookup(k) == nil {
		DATA[k] = n
		return true
	}
	return false
}

func Delete(k string) bool {
	if Lookup(k) != nil {
		delete(DATA, k)
		return true
	}
	return false
}

func Lookup(k string) *interface{} {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	} else {
		return nil
	}
}

func Change(k string, n interface{}) bool {
	DATA[k] = n
	return true
}

func Print() {
	for k, d := range DATA {
		fmt.Println("key: %s value: %v\n", k, d)
	}
}
