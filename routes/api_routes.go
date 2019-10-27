package routes

import (
	CO "Iris-Mini-Social-Network/config"
	"os"
	"time"

	"github.com/badoux/checkmail"
	"github.com/kataras/iris/v12"
)

// CreateNewPost route
func CreateNewPost(ctx iris.Context) {

	title := ctx.PostValueTrim("title")
	content := ctx.PostValueTrim("content")
	id, _ := CO.AllSessions(ctx)

	db := CO.DB()

	stmt, _ := db.Prepare("INSERT INTO posts(title, content, createdBy, createdAt) VALUES (?, ?, ?, ?)")
	rs, iErr := stmt.Exec(title, content, id, time.Now())
	CO.Err(iErr)

	insertID, _ := rs.LastInsertId()

	resp := map[string]interface{}{
		"postID": insertID,
		"mssg":   "Post Created!!",
	}
	json(ctx, resp)
}

// DeletePost route
func DeletePost(ctx iris.Context) {
	post := ctx.FormValue("post")
	db := CO.DB()

	_, dErr := db.Exec("DELETE FROM posts WHERE postID=?", post)
	CO.Err(dErr)

	json(ctx, map[string]interface{}{
		"mssg": "Post Deleted!!",
	})
}

// UpdatePost route
func UpdatePost(ctx iris.Context) {
	postID := ctx.PostValue("postID")
	title := ctx.PostValue("title")
	content := ctx.PostValue("content")

	db := CO.DB()
	db.Exec("UPDATE posts SET title=?, content=? WHERE postID=?", title, content, postID)

	json(ctx, map[string]interface{}{
		"mssg": "Post Updated!!",
	})
}

// UpdateProfile route
func UpdateProfile(ctx iris.Context) {
	resp := make(map[string]interface{})

	id, _ := CO.AllSessions(ctx)
	username := ctx.PostValueTrim("username")
	email := ctx.PostValueTrim("email")
	bio := ctx.PostValueTrim("bio")

	mailErr := checkmail.ValidateFormat(email)
	db := CO.DB()

	if username == "" || email == "" {
		resp["mssg"] = "Some values are missing!!"
	} else if mailErr != nil {
		resp["mssg"] = "Invalid email format!!"
	} else {
		_, iErr := db.Exec("UPDATE users SET username=?, email=?, bio=? WHERE id=?", username, email, bio, id)
		CO.Err(iErr)

		session := CO.GetSession(ctx)
		session.Set("username", username)

		resp["mssg"] = "Profile updated!!"
		resp["success"] = true
	}

	json(ctx, resp)
}

// ChangeAvatar route
func ChangeAvatar(ctx iris.Context) {
	resp := make(map[string]interface{})
	id, _ := CO.AllSessions(ctx)

	dir, _ := os.Getwd()
	dest := dir + "/public/users/" + id + "/avatar.png"

	dErr := os.Remove(dest)
	CO.Err(dErr)

	// avatar key of post form file, but let's grab all of them,
	// the `ctx.FormFile` can be used to manually upload files per post key to the server.
	_, upErr := ctx.UploadFormFiles(dest)

	if upErr != nil {
		resp["mssg"] = "An error occured!!"
	} else {
		resp["mssg"] = "Avatar changed!!"
		resp["success"] = true
	}

	json(ctx, resp)
}

// Follow route
func Follow(ctx iris.Context) {
	id, _ := CO.AllSessions(ctx)
	user := ctx.PostValue("user")
	username := CO.Get(user, "username")

	db := CO.DB()
	stmt, _ := db.Prepare("INSERT INTO follow(followBy, followTo, followTime) VALUES(?, ?, ?)")
	_, exErr := stmt.Exec(id, user, time.Now())
	CO.Err(exErr)

	json(ctx, iris.Map{
		"mssg": "Followed " + username + "!!",
	})
}

// Unfollow route
func Unfollow(ctx iris.Context) {
	id, _ := CO.AllSessions(ctx)
	user := ctx.PostValue("user")
	username := CO.Get(user, "username")

	db := CO.DB()
	stmt, _ := db.Prepare("DELETE FROM follow WHERE followBy=? AND followTo=?")
	_, dErr := stmt.Exec(id, user)
	CO.Err(dErr)

	json(ctx, iris.Map{
		"mssg": "Unfollowed " + username + "!!",
	})
}

// Like post route
func Like(ctx iris.Context) {
	post := ctx.PostValue("post")
	db := CO.DB()
	id, _ := CO.AllSessions(ctx)

	stmt, _ := db.Prepare("INSERT INTO likes(postID, likeBy, likeTime) VALUES (?, ?, ?)")
	_, err := stmt.Exec(post, id, time.Now())
	CO.Err(err)

	json(ctx, iris.Map{
		"mssg": "Post Liked!!",
	})
}

// Unlike post route
func Unlike(ctx iris.Context) {
	post := ctx.PostValue("post")
	id, _ := CO.AllSessions(ctx)
	db := CO.DB()

	stmt, _ := db.Prepare("DELETE FROM likes WHERE postID=? AND likeBy=?")
	_, err := stmt.Exec(post, id)
	CO.Err(err)

	json(ctx, iris.Map{
		"mssg": "Post Unliked!!",
	})
}

// DeactivateAcc route post method
func DeactivateAcc(ctx iris.Context) {
	session := CO.GetSession(ctx)
	id, _ := CO.AllSessions(ctx)
	db := CO.DB()
	var postID int

	db.Exec("DELETE FROM profile_views WHERE viewBy=?", id)
	db.Exec("DELETE FROM profile_views WHERE viewTo=?", id)
	db.Exec("DELETE FROM follow WHERE followBy=?", id)
	db.Exec("DELETE FROM follow WHERE followTo=?", id)
	db.Exec("DELETE FROM likes WHERE likeBy=?", id)

	rows, _ := db.Query("SELECT postID FROM posts WHERE createdBy=?", id)
	for rows.Next() {
		rows.Scan(&postID)
		db.Exec("DELETE FROM likes WHERE postID=?", postID)
	}

	db.Exec("DELETE FROM posts WHERE createdBy=?", id)
	db.Exec("DELETE FROM users WHERE id=?", id)

	dir, _ := os.Getwd()
	userPath := dir + "/public/users/" + id

	rmErr := os.RemoveAll(userPath)
	CO.Err(rmErr)

	// session.Delete("id")
	// session.Delete("username")
	// or
	session.Destroy()

	json(ctx, iris.Map{
		"mssg": "Deactivated your account!!",
	})
}
