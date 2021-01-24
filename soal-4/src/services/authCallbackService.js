const axios = require("axios");
const UserServices = require("./userInfoService");
const config = require("../config");

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

module.exports = {
  callback: callback,
};
