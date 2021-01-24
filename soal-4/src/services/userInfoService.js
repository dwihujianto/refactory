const axios = require("axios");
const config = require("../config");

function getUserInfo(token, cb) {
  axios({
    method: "get",
    url: `${config.apiUrl}user`,
    headers: {
      Authorization: "Bearer " + token,
    },
  }).then((response) => {
    cb(null,response.data);
  }).catch((err) => cb(err, []));
}

module.exports = {
  getUserInfo : getUserInfo
}
