package main

import (
	// "fmt"
	// "net/http"
	// "os"
	// "runtime"
	"time"
)

type Users struct {
	ID       int	`json:"id"`
	Username string `json:"username"`
	Password string	`json:"password"`
	createAt time.Time	`json:"createAt"`
	updateAt time.Time	`json:"updateAt"`
}

type follows struct {
	ID       int `json:"id"`
	follower_id int `json:"follower_id"`
	followed_id int `json:"followed_id"`
	createAt time.Time `json:"createAt"`
	updateAt time.Time `json:"updateAt"`
}

type tweets struct {
	ID       int	`json:"id"`
	user_id int	`json:"user_id"`
	content string	`json:"content"`
	createAt time.Time	`json:"createAt"`
	updateAt time.Time	`json:"updateAt"`
}

type likes struct {
	ID       int	`json:"id"`
	user_id int	`json:"user_id"`
	tweet_id int	`json:"tweet_id"`
	createAt time.Time	`json:"createAt"`
	updateAt time.Time	`json:"updateAt"`
}

type retweets struct {
	ID       int	`json:"id"`
	user_id int	`json:"user_id"`
	tweet_id int	`json:"tweet_id"`
	createAt time.Time	`json:"createAt"`
	updateAt time.Time	`json:"updateAt"`
}



const (
	dbFileName = "db.mysql"

	createUsersTable = `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`

	createFollowsTable = `
		CREATE TABLE IF NOT EXISTS follows (
			id INT AUTO_INCREMENT PRIMARY KEY,
			follower_id INT NOT NULL,
			followed_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`

	createTweetsTable = `
		CREATE TABLE IF NOT EXISTS tweets (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`

	createLikesTable = `
		CREATE TABLE IF NOT EXISTS likes (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			tweet_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`

	createRetweetsTable = `
		CREATE TABLE IF NOT EXISTS retweets (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			tweet_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`

	insertUser = `
		INSERT INTO users (username, password) VALUES (?, ?);
	`

	insertFollow = `
		INSERT INTO follows (follower_id, followed_id) VALUES (?, ?);
	`

	insertTweet = `
		INSERT INTO tweets (user_id, content) VALUES (?, ?);
	`

	insertLike = `
		INSERT INTO likes (user_id, tweet_id) VALUES (?, ?);
	`

	insertRetweet = `
		INSERT INTO retweets (user_id, tweet_id) VALUES (?, ?);
	`

	getUser = `
		SELECT id, username, password, created_at, updated_at FROM users WHERE id = ?;	
	`

	getFollow = `
		SELECT id, follower_id, followed_id, created_at, updated_at FROM follows WHERE id = ?;
	`



)


func main() {

}


func decodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}

func responseJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		panic(err)
	}
}