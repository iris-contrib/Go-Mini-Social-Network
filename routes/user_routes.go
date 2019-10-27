package routes

import (
	CO "Iris-Mini-Social-Network/config"
	"os"
	"strconv"
	"time"

	"github.com/badoux/checkmail"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

// Signup route
func Signup(ctx iris.Context) {
	notLoggedIn(ctx)
	renderTemplate(ctx, "signup", iris.Map{
		"title": "Signup For Free",
	})
}

// Login route
func Login(ctx iris.Context) {
	notLoggedIn(ctx)
	renderTemplate(ctx, "login", iris.Map{
		"title": "Login To Continue",
	})
}

// Logout route
func Logout(ctx iris.Context) {
	loggedIn(ctx, "")
	session := CO.GetSession(ctx)
	session.Destroy()
	// or
	// session.Delete("id")
	// session.Delete("username")
	// if just delete the values.
	ctx.Redirect("/login")
}

// UserSignup function to register user
func UserSignup(ctx iris.Context) {
	resp := make(map[string]interface{})

	username := ctx.PostValueTrim("username")
	email := ctx.PostValueTrim("email")
	password := ctx.PostValueTrim("password")
	passwordAgain := ctx.PostValueTrim("password_again")

	mailErr := checkmail.ValidateFormat(email)

	db := CO.DB()

	var (
		userCount  int
		emailCount int
	)

	db.QueryRow("SELECT COUNT(id) AS userCount FROM users WHERE username=?", username).Scan(&userCount)
	db.QueryRow("SELECT COUNT(id) AS emailCount FROM users WHERE email=?", email).Scan(&emailCount)

	if username == "" || email == "" || password == "" || passwordAgain == "" {
		resp["mssg"] = "Some values are missing!!"
	} else if len(username) < 4 || len(username) > 32 {
		resp["mssg"] = "Username should be between 4 and 32"
	} else if mailErr != nil {
		resp["mssg"] = "Invalid email format!!"
	} else if password != passwordAgain {
		resp["mssg"] = "Passwords don't match"
	} else if userCount > 0 {
		resp["mssg"] = "Username already exists!!"
	} else if emailCount > 0 {
		resp["mssg"] = "Email already exists!!"
	} else {

		stmt, _ := db.Prepare("INSERT INTO users(username, email, password, joined) VALUES (?, ?, ?, ?)")
		rs, iErr := stmt.Exec(username, email, hash(password), time.Now())
		CO.Err(iErr)
		insertID, _ := rs.LastInsertId()
		insStr := strconv.FormatInt(insertID, 10)

		dir, _ := os.Getwd()
		userPath := dir + "/public/users/" + insStr

		dirErr := os.Mkdir(userPath, 0655)
		CO.Err(dirErr)

		linkErr := os.Link(dir+"/public/images/golang.png", userPath+"/avatar.png")
		CO.Err(linkErr)

		session := CO.GetSession(ctx)
		session.Set("id", insStr)
		session.Set("username", username)

		resp["success"] = true
		resp["mssg"] = "Hello, " + username + "!!"

	}
	json(ctx, resp)
}

// UserLogin function to log user in
func UserLogin(ctx iris.Context) {
	resp := make(map[string]interface{})

	rusername := ctx.PostValueTrim("username")
	rpassword := ctx.PostValueTrim("password")

	db := CO.DB()
	var (
		userCount int
		id        int
		username  string
		password  string
	)

	db.QueryRow("SELECT COUNT(id) AS userCount, id, username, password FROM users WHERE username=?", rusername).Scan(&userCount, &id, &username, &password)

	encErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(rpassword))

	if rusername == "" || rpassword == "" {
		resp["mssg"] = "Some values are missing!!"
	} else if userCount == 0 {
		resp["mssg"] = "Invalid username!!"
	} else if encErr != nil {
		resp["mssg"] = "Invalid password!!"
	} else {
		session := CO.GetSession(ctx)
		session.Set("id", strconv.Itoa(id))
		session.Set("username", username)

		resp["mssg"] = "Hello, " + username + "!!"
		resp["success"] = true
	}
	json(ctx, resp)
}
