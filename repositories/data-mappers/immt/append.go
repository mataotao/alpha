package immt

import (
	"alpha/config"
	"encoding/gob"
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"

	"time"
)

const (
	maxPendingChannel int = 10
	maxPendingData    int = 10
	maxTime               = 10 * time.Minute
	maxDataTime           = 1 * time.Minute
)

func appending() {
	pendData := make([]*command, 0, maxPendingData)
	pendingChannel := make(chan *command, maxPendingChannel)
	execChan := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	wg.Add(4)
	go func() {
		defer wg.Done()
		for v := range MT.append {
			pendingChannel <- v
		}
	}()
	go func() {
		defer wg.Done()
		for {
			d := <-pendingChannel
			execChan <- struct{}{}
			pendData = append(pendData, d)
			if len(pendData) == maxPendingData {
				loadSaveDisk(pendData)
				pendData = make([]*command, 0, maxPendingData)
			}
			<-execChan
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(maxDataTime):
				execChan <- struct{}{}
				loadSaveDisk(pendData)
				<-execChan
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(maxTime):
				execChan <- struct{}{}
				if err := MT.Save(); err != nil {
					config.Logger.Error("immt error",
						zap.Error(err),
					)
				}
				<-execChan
			}
		}
	}()
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:

	}
}
func loadSaveDisk(commands []*command) {
	saveDisk <- struct{}{}
	defer func() {
		<-saveDisk
	}()
	saveData := make(map[string]interface{})

	if _, err := os.Stat(MT.dataFile); !os.IsNotExist(err) {
		loadFrom, err := os.Open(MT.dataFile)
		if err != nil {
			config.Logger.Error("immt error",
				zap.Error(err),
			)
			return
		}
		decoder := gob.NewDecoder(loadFrom)
		err = decoder.Decode(&saveData)
		if err != nil {
			config.Logger.Error("immt error",
				zap.Error(err),
			)
			return
		}
		err = loadFrom.Close()
		if err != nil {
			config.Logger.Error("immt error",
				zap.Error(err),
			)
			return
		}
	}

	for i := range commands[:] {
		switch commands[i].action {
		case "SET":
			saveData[commands[i].key] = commands[i].value
		case "DELETE":
			delete(saveData, commands[i].key)
		}
	}

	if err := os.Remove(MT.dataFile); err != nil {
		fmt.Println(err)
	}
	saveTo, err := os.Create(MT.dataFile)
	if err != nil {
		config.Logger.Error("immt error",
			zap.Error(err),
		)
		return
	}
	defer func() {
		err = saveTo.Close()
	}()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(saveData)
	if err != nil {
		config.Logger.Error("immt error",
			zap.Error(err),
		)
		return
	}
}
