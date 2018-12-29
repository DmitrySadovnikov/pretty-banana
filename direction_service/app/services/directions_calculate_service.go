package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	p "direction_service/app/proto"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

type DirectionsCalculateService struct{}

type path struct {
	Distance float64 `json:"distance"`
	Time     int64   `json:"time"`
}

type graphhopperResponse struct {
	Paths []path `json:"paths"`
}

func (_ DirectionsCalculateService) Call(request *p.Calculate_Request) *p.Calculate_Response {
	redisCodec := redisCodec()
	cacheKey := fmt.Sprintf(
		"%f,%f,%f,%f",
		request.StartPoint.Lat,
		request.StartPoint.Lng,
		request.EndPoint.Lat,
		request.EndPoint.Lng,
	)

	if result, ok := fetchCachedResult(redisCodec, cacheKey); ok {
		return result
	}
	result := fetchResult(request)
	go writeResultToCache(redisCodec, cacheKey, result)
	return result
}

func fetchResult(request *p.Calculate_Request) *p.Calculate_Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("GRAPHHOPPER_URL"), nil)

	if err != nil {
		log.Print(err)
		panic(err)
	}

	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("key", os.Getenv("GRAPHHOPPER_KEY"))
	q.Add("point", fmt.Sprintf("%f,%f", request.StartPoint.Lat, request.StartPoint.Lng))
	q.Add("point", fmt.Sprintf("%f,%f", request.EndPoint.Lat, request.EndPoint.Lng))
	q.Add("vehicle", "car")
	q.Add("calc_points", "false")
	q.Add("points_encoded", "false")
	q.Add("instructions", "false")
	q.Add("optimize", "true")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		panic(err)
	}

	defer resp.Body.Close()
	rawBody, _ := ioutil.ReadAll(resp.Body)

	var graphhopperResponse graphhopperResponse
	err = json.Unmarshal(rawBody, &graphhopperResponse)
	if err != nil {
		log.Print(err)
		panic(err)
	}

	return &p.Calculate_Response{
		Distance: graphhopperResponse.Paths[0].Distance,
		Time:     graphhopperResponse.Paths[0].Time,
	}
}

func redisCodec() *cache.Codec {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server": os.Getenv("REDIS_URL"),
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	return codec
}

func fetchCachedResult(codec *cache.Codec, key string) (*p.Calculate_Response, bool) {
	var result *p.Calculate_Response
	err := codec.Get(key, &result)

	if err == nil {
		log.Printf("Successfully read from cache by key %s", key)
		return result, true
	}
	log.Printf("Unsuccessfully read from cache by key %s %v", key, err)
	return result, false
}

func writeResultToCache(codec *cache.Codec, key string, result *p.Calculate_Response) {
	codec.Set(&cache.Item{
		Key:        key,
		Object:     result,
		Expiration: time.Minute,
	})
	log.Printf("Successfully write to cache key %s", key)
}
