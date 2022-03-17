package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx/v3"
	"gopkg.in/ini.v1"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	initConf()

	// Or optionally use the hook with an API key:
	log.AddHook(logruseq.NewSeqHook("http://localhost:5341",
		logruseq.OptionAPIKey("H7vifFTVyZ3RpSYg3NvI")))

	log.
		// WithFields(log.Fields{
		// 	"animal": "walrus",
		// }).
		Info("A walrus appears")
	xlsxTT()
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
		if err != nil {
			panic(err)
		}
		adcodeValue, err := adcode.FormattedValue()
		if err != nil {
			panic(err)
		}
		getAddressFactorySingleInstance().addressCode[adcodeValue] = nameValue

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
	log.Info("Save chinaCode to redis success!")
}
