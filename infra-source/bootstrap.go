package main

import (
	"os"
	"github.com/allegro/bigcache"
	"time"
	"encoding/json"
	"github.com/fatih/color"
	"github.com/boltdb/bolt"
)


// Function that read the config.json file and populates
//  the Config singleton to use further in the app during
//  runtime.
//
func loadAppConfig() {
	//sync.Once{}.Do()
	//current_dir, _ := os.Getwd()
	configFile, _ := os.Open("./config/app.json")
	loadStaticPages, _ := os.Open("./config/static.json")

	defer func() {
		configFile.Close()
		loadStaticPages.Close()
	}()

	configFileParser := json.NewDecoder(configFile)
	configFileParser.Decode(&Config)

	staticsFileParser := json.NewDecoder(loadStaticPages)
	staticsFileParser.Decode(&StaticPages)

	color.Green(" * Configurations Loaded SuccessFully ")
}

var Db *bolt.DB

func initdb(){
	if Config.Database.Driver == "boltDB"{
		Db, _ = bolt.Open(Config.Database.DatabaseName, 0600, nil)
		Db.Update(func(tx *bolt.Tx) error{
			_,err := tx.CreateBucketIfNotExists([]byte("users"))
			if err != nil {
				color.Red(" * Error During Creating Users Bucket :",err.Error())
			}
			return err
		})
	}
}

func initCache() (*bigcache.BigCache, error) {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,
		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute,
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// prints information about additional memory allocation
		Verbose: true,
		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,
		// callback fired when the oldest entry is removed because of its
		// expiration time or no space left for the new entry. Default value is nil which
		// means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,
	}
	return bigcache.NewBigCache(config)
}

func init(){
	UserSession.Options.HttpOnly = true
}