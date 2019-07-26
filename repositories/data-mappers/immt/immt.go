package immt

import (
	"alpha/config"

	"go.uber.org/zap"

	"encoding/gob"
	"fmt"
	"os"
)

const (
	maxSendData int = 10
)

var MT *Mt
var saveDisk = make(chan struct{}, 1)

type Mt struct {
	data     map[string]interface{}
	looking  chan struct{}
	dataFile string
	append   chan *command
}
type command struct {
	action string
	key    string
	value  interface{}
}

func init() {
	MT = &Mt{
		data:     make(map[string]interface{}),
		looking:  make(chan struct{}, 1),
		dataFile: "./data/immt/dataFile.gob",
		append:   make(chan *command, maxSendData),
	}
	if err := MT.Init(); err != nil {
		config.Logger.Error("immt error",
			zap.Error(err),
		)
	}
	go func() {
		appending()
	}()
}

func (mt *Mt) Save() error {
	saveDisk <- struct{}{}
	defer func() {
		<-saveDisk
	}()
	fmt.Println("Saving", mt.dataFile)
	if err := os.Remove(mt.dataFile); err != nil {
		fmt.Println(err)
	}
	saveTo, err := os.Create(mt.dataFile)
	if err != nil {
		fmt.Println("Cannot create", mt.dataFile)
		return err
	}
	defer func() {
		err = saveTo.Close()
	}()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(mt.data)
	if err != nil {
		fmt.Println("Cannot save to", mt.dataFile)
		return err
	}
	return nil
}

func (mt *Mt) Init() error {
	fmt.Println("Loading", mt.dataFile)
	loadFrom, err := os.Open(mt.dataFile)
	defer func() {
		err = loadFrom.Close()
	}()
	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}

	decoder := gob.NewDecoder(loadFrom)
	err = decoder.Decode(&mt.data)
	return nil
}

func (mt *Mt) Set(k string, v interface{}) bool {
	if k == "" {
		return false
	}
	if mt.Get(k) == nil {
		mt.data[k] = v
		go func() {
			command := &command{
				action: "SET",
				key:    k,
				value:  v,
			}
			mt.append <- command
		}()
		return true
	}

	return false
}

func (mt *Mt) Delete(k string) bool {
	if mt.Get(k) == nil {
		delete(mt.data, k)
		go func() {
			command := &command{
				action: "DELETE",
				key:    k,
			}
			mt.append <- command
		}()
		return true
	}
	return false
}

func (mt *Mt) Get(k string) *interface{} {
	mt.looking <- struct{}{}
	defer func() {
		<-mt.looking
	}()
	_, ok := mt.data[k]
	if ok {
		n := mt.data[k]
		return &n
	} else {
		return nil
	}
}

func (mt *Mt) Print() {
	for k, d := range mt.data {
		fmt.Println("key: %s value: %v\n", k, d)
	}
}
