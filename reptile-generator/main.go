package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/redis/go-redis/v9"
)

var (
	hostFlag  string
	redisFlag string
	rDB       *redis.Client
)

func init() {
	flag.StringVar(&hostFlag, "h", ":3000", "listen on host")
	flag.StringVar(&redisFlag, "r", "", "Redis database")
	flag.Parse()
	rDB = redis.NewClient(&redis.Options{
		Addr:     redisFlag,
		Password: "",
		DB:       0,
	})
}

func main() {
	log.Fatal(http.ListenAndServe(hostFlag, buildRouter()))
}

func buildRouter() *httprouter.Router {
	var router = httprouter.New()
	router.GET("/api/reptiles/pick", replyJSON(ReptilePick))
	router.GET("/api/reptile/:slug", replyJSON(ReptileShow))
	// router.POST("/reptile/:slug", reptile.Create)
	// router.GET("/reptiles", replyJSON(ReptileList))
	// router.DELETE("/reptile/:slug", reptile.Delete)
	router.ServeFiles("/app/*filepath", http.Dir("frontend"))
	return router
}

type ControllerFunc func(r *http.Request, ps httprouter.Params) (any, error)

func replyJSON(controllerFunc ControllerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		object, err := controllerFunc(r, ps)
		if err != nil {
			handleErrJSON(w, err)
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(object); err != nil {
			handleErrJSON(w, err)
		}
	}
}

func handleErrJSON(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	resp := make(map[string]string)
	resp["message"] = err.Error()
	jsonResp, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	w.Write(jsonResp)
	log.Println(err)
	return
}

type Reptile struct {
	Slug      string `json:"slug"`
	Name      string `redis:"name" json:"name"`
	LatinName string `redis:"latin_name" json:"latin_name"`
	Photo     string `redis:"photo" json:"photo,omitempty"`
	Video     string `redis:"video" json:"video,omitempty"`
}

func ReptileShow(r *http.Request, ps httprouter.Params) (any, error) {
	var reptile Reptile
	slug := ps.ByName("slug")
	var ctx = context.Background()
	result := rDB.HGetAll(ctx, "reptiles:"+slug)
	if err := result.Err(); err != nil {
		return reptile, err
	}
	reptile.Slug = slug
	if err := result.Scan(&reptile); err != nil {
		return reptile, err
	}
	return reptile, nil
}

func ReptilePick(r *http.Request, _ httprouter.Params) (any, error) {
	var reptile Reptile
	var ctx = context.Background()
	rand := rDB.SRandMember(ctx, "reptiles")
	if err := rand.Err(); err != nil {
		return reptile, err
	}
	reptile.Slug = rand.Val()
	result := rDB.HGetAll(ctx, "reptiles:"+reptile.Slug)
	if err := result.Err(); err != nil {
		return reptile, err
	}
	if err := result.Scan(&reptile); err != nil {
		return reptile, err
	}
	return reptile, nil
}

// func ReptileList(r *http.Request, _ httprouter.Params) (any, error) {
// 	var reptiles []Reptile
// 	var ctx = context.Background()
// 	members := rDB.SMembers(ctx, "reptiles")
// 	if err := members.Err(); err != nil {
// 		return reptiles, err
// 	}
// 	slugs := members.Val()
// 	pipeline := rDB.Pipeline()
// 	for slug := range slugs {
// 		result := pipeline.HGetAll(ctx, "reptiles"+slug)
// 	}
// 	if err := result.Err(); err != nil {
// 		return reptiles, err
// 	}
// 	if err := result.Scan(&reptiles); err != nil {
// 		return reptiles, err
// 	}
// 	return reptiles, nil
// }
