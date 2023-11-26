package main

import (
	"context"
	"database/sql"
	"flag"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var redisFlag = flag.String("r", "", "Redis database")
var postgresFlag = flag.String("p", "", "Postgres database")

type Creator struct {
	ID        string `redis:"id"`
	Username  string `redis:"username"`
	FirstName string `redis:"firstname"`
	LastName  string `redis:"lastname"`
}

func main() {
	flag.Parse()

	// Open Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     *redisFlag,
		Password: "",
		DB:       0,
	})
	var ctx = context.Background()

	// Open Postgres
	pdb, err := sql.Open("mysql", *postgresFlag)
	if err != nil {
		panic(err)
	}
	defer pdb.Close()
	pdb.SetConnMaxLifetime(time.Minute * 3)
	pdb.SetMaxOpenConns(10)
	pdb.SetMaxIdleConns(10)

	// Get creators
	creators, err := getCreatorsPostgres(pdb)
	if err != nil {
		panic(err)
	}

	// Set creators
	setCreatorsRedis(ctx, creators, rdb)

}

func getCreatorRedis(ctx context.Context, id string, rdb *redis.Client) (Creator, error) {
	var creator Creator
	result := rdb.HGetAll(ctx, "brickbros:creator:"+id)
	if err := result.Err(); err != nil {
		return creator, err
	}
	err := result.Scan(creator)
	if err != nil {
		return creator, err
	}
	return creator, nil
}

func setCreatorsRedis(ctx context.Context, creators []Creator, rdb *redis.Client) {
	for _, creator := range creators {
		rdb.HSet(ctx, "brickbros:creator:"+creator.ID, creator)
		rdb.LPush(ctx, "brickbros:creators", creator.ID)
	}
}

func getCreatorsPostgres(db *sql.DB) ([]Creator, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(id) FROM Creators WHERE `status` = 'normal'").Scan(&count)
	if err != nil {
		return nil, err
	}
	creators := make([]Creator, count)
	rows, err := db.Query("SELECT id, cname, fname, lname FROM Creators WHERE `status` = 'normal'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var index int
	for rows.Next() {
		err := rows.Scan(&creators[index].ID, &creators[index].Username, &creators[index].FirstName, &creators[index].LastName)
		if err != nil {
			return nil, err
		}
		index++
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return creators, nil
}
