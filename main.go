package main

import (
	R "Iris-Mini-Social-Network/routes"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())

	app.RegisterView(iris.HTML("./views", ".html"))
	app.HandleDir("/", "./public")

	user := app.Party("/user")
	{
		user.Post("/signup", R.UserSignup)
		user.Post("/login", R.UserLogin)
	}

	app.Get("/", R.Index)
	app.Get("/welcome", R.Welcome)
	app.Get("/explore", R.Explore)
	app.Get("/404", R.NotFound)
	app.Get("/signup", R.Signup)
	app.Get("/login", R.Login)
	app.Get("/logout", R.Logout)
	app.Get("/deactivate", R.Deactivate)
	app.Get("/edit_profile", R.EditProfile)
	app.Get("/create_post", R.CreatePost)

	app.Get("/profile/:id", R.Profile)
	app.Get("/profile", R.NotFound)

	app.Get("/view_post/:id", R.ViewPost)
	app.Get("/view_post", R.NotFound)

	app.Get("/edit_post/:id", R.EditPost)
	app.Get("/edit_post", R.NotFound)

	app.Get("/followers/:id", R.Followers)
	app.Get("/followers", R.NotFound)

	app.Get("/followings/:id", R.Followings)
	app.Get("/followings", R.NotFound)

	app.Get("/likes/:id", R.Likes)
	app.Get("/likes", R.NotFound)

	api := app.Party("/api")
	{
		api.Post("/create_new_post", R.CreateNewPost)
		api.Post("/delete_post", R.DeletePost)
		api.Post("/update_post", R.UpdatePost)
		api.Post("/update_profile", R.UpdateProfile)
		api.Post("/change_avatar", R.ChangeAvatar)
		api.Post("/follow", R.Follow)
		api.Post("/unfollow", R.Unfollow)
		api.Post("/like", R.Like)
		api.Post("/unlike", R.Unlike)
		api.Post("/deactivate-account", R.DeactivateAcc)
	}

	app.Run(iris.Addr(os.Getenv("PORT")))
}
