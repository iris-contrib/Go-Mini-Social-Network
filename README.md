# Iris-Mini-Social-Network
A mini Social-Network created with the awesome Iris Golang💖💖!!

This is a fork of: https://github.com/yTakkar/Go-Mini-Social-Network.

Created by [@kataras](https://twitter.com/MakisMaropoulos) as an example on converting a web app from gin to iris. The structure is the same as the original repository, nothing changed in order to be easier to watch the changes between them.

# Quick Links
1. [Screenshots](#screenshots)
2. [Requirements](#requirements)
3. [Usage](#usage)

# Screenshots
![alt text](https://raw.githubusercontent.com/iris-contrib/Iris-Mini-Social-Network/master/screenshots/Snap%202017-09-26%20at%2001.11.55.png)
![alt text](https://raw.githubusercontent.com/iris-contrib/Iris-Mini-Social-Network/master/screenshots/Snap%202017-09-26%20at%2001.12.18.png)
![alt text](https://raw.githubusercontent.com/iris-contrib/Iris-Mini-Social-Network/master/screenshots/Snap%202017-09-26%20at%2013.11.39.png)
![alt text](https://raw.githubusercontent.com/iris-contrib/Iris-Mini-Social-Network/master/screenshots/Snap%202017-09-26%20at%2001.13.22.png)
![alt text](https://raw.githubusercontent.com/iris-contrib/Iris-Mini-Social-Network/master/screenshots/Snap%202017-09-26%20at%2001.12.03.png)
![alt text](https://raw.githubusercontent.com/iris-contrib/Iris-Mini-Social-Network/master/screenshots/Snap%202017-09-26%20at%2001.13.07.png)
![alt text](https://raw.githubusercontent.com/iris-contrib/Iris-Mini-Social-Network/master/screenshots/Snap%202017-09-26%20at%2001.13.29.png)

# Requirements
1. Make sure you keep this project inside `src/` of your Golang project folder ($GOPATH).
2. Following packages should be installed.
    1. [iris](https://github.com/kataras/iris)
    2. [checkmail](https://github.com/badoux/checkmail)
    3. [MySQL driver](https://github.com/go-sql-driver/mysql)
    4. [bcrypt](https://golang.org/x/crypto/bcrypt)
    5. [sessions](https://github.com/gorilla/sessions)
    6. [godotenv](https://github.com/joho/godotenv)

# Usage

1. Open PHPMyAdmin or any other db tool, create a db & import `db.sql`.

2. Install all the dependencies.
```bash
# with npm
npm install

# or with yarn
yarn
```

3. Create a `.env` file & insert the following code. Replace values with yours!!
```javascript
PORT=YOUR PORT [STRING]
SESSION_HASH_SECRET=SESSION_HASH_KEY [STRING]
SESSION_BLOCK_SECRET=SESSION_BLOCK_KEY [STRING]
DB_USER=DB_USER [STRING]
DB_PASSWORD=DB PASSWORD [STRING]
DB=DB YOU JUST CREATE [STRING]
```

> See the default `.env` file to see an example of it

4. My root folder name is `Iris-Mini-Social-Network`, if yours is different then go ahead & change it as it used for imports. It can be done easily by searching the whole project.

5. Now run the app.
```javascript
npm start [OR] yarn start
```

6. Run the app in browser.
```javascript
localhost:[PORT]
```

7. Enjoy 💖💖!!
