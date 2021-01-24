# CodeDebugging

Penyelesaian error code

### problem 1

Env yang tidak terload disebabkan karena tidak dilakunan inisialisasi terlebih dahulu ketika package di import

file `./config/index.js` :
code sebelumnya
```js
const dotenv = require("dotenv");
dan
const envFound = dotenv.config();
```

Solusi
```js
const dotenv = require("dotenv").config();
const envFound = dotenv;
```

### problem 2
Config yang tidak bisa terload di app.js karena di export dengan object config baru, sehingga harus chain 2 kali untuk memanggil config

file `./config/index.js` :
Code sebelumnya
```js
module.exports = { config };
```

solusi
```js
module.exports = config;
```

### problem 3
aithService tidak terload karena export module kurang huruf `s`

file `./src/services/authService.js` :
Code sebelumnya
```js
module.export = {
     redirectUri: redirectUri
}
```

solusi
```js
module.exports = {
     redirectUri: redirectUri
}
```

### problem 4
Code request access token perlu di refactory, parameter cukup menggunakan query string, tidak perlu menggunakn body, dan resolve promise then tidak perlu dipanggil 2 kali

file `./src/services/authCallbackService.js` :
Code sebelumnya
```js
function callback(req, res) {
	const body = {
	  client_id: config.clientId,
	  client_secret: config.clientSecret,
	  code: req.query.code,
	};
	const options = { headers: { accept: "application/json" } };
	axios
	  .post(`${config.oauthUrl}/access_token`, body, options)
	  .then((res) => resp.data["accessToken"])
	  .then((accessToken) => {
	//..........lanjutan
```

Solusi
```js
function callback(req, res) {
	axios({
	  method: 'post',
	  url: `${config.oauthUrl}/access_token?client_id=${config.clientId}&client_secret=${config.clientSecret}&code=${req.query.code}`,
	  headers: {
	    accept: 'application/json'
	  }
	}).then((response) => {
	//..........lanjutan
```

### problem 5
code request user info terdapat kesalahan export fungsi, sehingga code akan dibaca sebagai object bukan fungsi. Code request info user harusnya dibuat callback karena code akan dipanggil kembali setelah mendapatkan access token (chaining).

file `./src/services/userInfoService.js` :
Code sebelumnya
```js
function getUserInfo(token) {
  axios({
    method: "get",
    url: `${config.apiUrl}/users`,
    headers: {
      Authorization: "token " + token,
    },
  }).then((response) => {
    return response.data;
  });
}
module.export = getUserInfo
```

Solusi
```js
function getUserInfo(token, cb) {
  axios({
    method: "get",
    url: `${config.apiUrl}user`,
    headers: {
      Authorization: "Bearer "  token,
    },
  }).then((response) => {
    cb(null,response.data);
  }).catch((err) => cb(err, []));
}

module.exports = {
  getUserInfo : getUserInfo
}
```

file `./src/services/authCallbackService.js` :

full sulusi code untuk callback
```js
function callback(req, res) {
  axios({
    method: 'post',
    url: `${config.oauthUrl}/access_token?client_id=${config.clientId}&client_secret=${config.clientSecret}&code=${req.query.code}`,
    headers: {
      accept: 'application/json'
    }
  }).then((response) => {
    const accessToken = response.data.access_token;
    UserServices.getUserInfo(accessToken, (err, user) => {
      if (err) {
        res.status(500).json({message : err.message})
      } else {
        res.json({
          data: {
            login: user.login,
            githubId: user.id,
            avatar: user.avatar_url,
            email: user.email,
            name: user.name,
            location: user.location,
          },
        });
      }
    });
  }).catch((err) => res.status(500).json({ message: err.message }));
}
```