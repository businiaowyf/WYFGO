package main

import (
	// "github.com/bitly/go-simplejson"
	// "database/sql"
	"github.com/jessevdk/go-flags"
	"gopkg.in/ini.v1"
	"gopkg.in/olivere/elastic.v3"
	"log"
	"os"
	// "strconv"
	// "time"
)

type Options struct {
	// Example of verbosity with level
	Verbose []bool `short:"v" long:"verbose" description:"Verbose output"`

	// Example of optional value
	ClientConf string `short:"c" long:"conf" description:"Clinet Config" optional:"no"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

var infoLog *log.Logger

func main() {
	// 处理命令行参数
	if _, err := parser.Parse(); err != nil {
		log.Fatalln("parse cmd line failed!")
	}

	if options.ClientConf == "" {
		log.Fatalln("Must input config file name")
	}

	// log.Println("config name =", options.ClientConf)
	// 读取配置文件
	cfg, err := ini.Load([]byte(""), options.ClientConf)
	if err != nil {
		log.Printf("load config file=%s failed", options.ClientConf)
		return
	}
	// 配置文件只读 设置此标识提升性能
	cfg.BlockMode = false
	// 定义一个文件
	fileName := cfg.Section("LOG").Key("path").MustString("/home/ht/clinet.log")
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
		return
	}

	// 创建一个日志对象
	infoLog = log.New(logFile, "[Info]", log.LstdFlags)
	// 配置log的Flag参数
	infoLog.SetFlags(infoLog.Flags() | log.LstdFlags)

	// 读取ip+port
	serverIp := cfg.Section("OUTER_SERVER").Key("server_ip").MustString("127.0.0.3")
	serverPort := cfg.Section("OUTER_SERVER").Key("server_port").MustString("9200")

	infoLog.Printf("server_ip=%v server_port=%v\n", serverIp, serverPort)

	// Obtain a client. You can also provide your own HTTP client here.
	// Create a client and connect to http://192.168.2.10:9201
	client, err := elastic.NewClient(elastic.SetURL("http://"+serverIp+":"+serverPort),
		elastic.SetMaxRetries(10),
		elastic.SetBasicAuth("hellotalk_admin", "hellotalk_admin123456"),
		elastic.SetErrorLog(infoLog))
	if err != nil {
		// Handle error
		panic(err)
	}
	// Trace request and response details like this
	//client.SetTracer(log.New(os.Stdout, "", 0))

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	infoLog.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	infoLog.Printf("Elasticsearch version %s", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("logstash-client-log-2016.09.12").Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// log
		infoLog.Println("unfound index logstash-client-log-2016.09.12")
	} else {
		infoLog.Println("exist index logstash-client-log-2016.09.12")
	}

	// Create an aggregation for users and a sub-aggregation for a date histogram of tweets (per year).
	timeline := elastic.NewTermsAggregation().Field("userid").Size(10).OrderByCountDesc()
	// histogram := elastic.NewDateHistogramAggregation().Field("created").Interval("year")
	// timeline = timeline.SubAggregation("history", histogram)

	// Search with a term query
	searchResult, err := client.Search().
		Index("logstash-client-log-2016.09.12"). // search in index "twitter"
		Query(elastic.NewMatchAllQuery()).       // return all results, but ...
		SearchType("count").                     // ... do not return hits, just the count
		Aggregation("timeline", timeline).       // add our aggregation to the query
		Pretty(true).                            // pretty print request and response JSON
		Do()                                     // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// Access "timeline" aggregate in search result.
	agg, found := searchResult.Aggregations.Terms("timeline")
	if !found {
		infoLog.Printf("we sould have a terms aggregation called %q\n", "timeline")
		return
	}
	for _, userBucket := range agg.Buckets {
		// Every bucket should have the user field as key.
		cmd := userBucket.Key
		infoLog.Printf("cmd=%v count=%v", cmd, userBucket.DocCount)
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
