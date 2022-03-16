package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/tealeg/xlsx/v3"
	"gopkg.in/ini.v1"
)

func main() {
	xlsxTT()
	initConf()
	// ExampleClient(newRedisOptions())

	saveAddressToReids(newRedisOptions())
}

type redisConf struct {
	addr     string
	username string
	password string
	db       int
}

var redisServer = &redisConf{}

func initConf() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		log.Fatalf("cannot load ini :%v", err)
	}
	err = cfg.Section("redis").MapTo(redisServer)
	if err != nil {
		log.Fatalf("cannot map config from ini :%v", err)
	}
}

func newRedisOptions() *redis.Options {

	return &redis.Options{
		Addr:     redisServer.addr,
		Password: redisServer.password, // no password set default ""
		Username: redisServer.username, //no password set default ""
		DB:       redisServer.db,       // use default DB 0
	}
}

func xlsxTT() {
	// open an existing file
	wb, err := xlsx.OpenFile("citycode.xlsx")
	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	fmt.Println("Sheets in this file:")
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name)
		sh, ok := wb.Sheet[sh.Name]
		if !ok {
			fmt.Println("Sheet does not exist")
			return
		}
		fmt.Println("Max row in sheet:", sh.MaxRow)
		err = sh.ForEachRow(rowVisitor)
		fmt.Println("Err=", err)
	}

	fmt.Println("----")
	fmt.Printf(" map :%v", getAddressFactorySingleInstance().addressCode)
}
func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cnum, _ := c.GetCoordinates()
		fmt.Printf(" [%d] value: %s \n", cnum, value)

		// getAddressFactorySingleInstance().addressCode[value] =cnum
	}
	return err
}

func rowVisitor(r *xlsx.Row) error {
	cnum := r.GetCoordinate()

	if cnum >= 1 {
		name := r.GetCell(0)
		adcode := r.GetCell(1)
		nameValue, err := name.FormattedValue()
		adcodeValue, err := adcode.FormattedValue()
		if err != nil {
			fmt.Println(err.Error())
			return nil
		} else {
			getAddressFactorySingleInstance().addressCode[adcodeValue] = nameValue
		}

	}
	return nil
}

type addressCode struct {
	addressCode map[string]string
}

var (
	dressFactorySingleInstance = &addressCode{
		addressCode: make(map[string]string),
	}
)

func getAddressFactorySingleInstance() *addressCode {
	return dressFactorySingleInstance
}

var ctx = context.Background()

func saveAddressToReids(opt *redis.Options) {

	rdb := redis.NewClient(opt)
	var code = "chinaCode"
	err := rdb.HSet(ctx, code, getAddressFactorySingleInstance().addressCode).Err()
	if err != nil {
		panic(err)
	}

}
