const fs = require('fs');

const rawdata = fs.readFileSync('data.json');
const users = JSON.parse(rawdata);

// Find users who doesn't have any phone numbers.
doesntHavePhone = users.filter(user => {
	return user.profile.phones.length == 0;
});
console.log("users who doesn't have any phone numbers.", doesntHavePhone);

// Find users who have articles..
haveArticles = users.filter(user => {
	return user.articles.length > 0;
});
console.log("Find users who have articles..", haveArticles)

// Find users who have "annis" on their name.
annis = users.filter(user => {
	haveAnnis = false;
	if (user.profile.length > 0) {
		haveAnnis = user.profile.full_name.toLowerCase().includes(/annis/);
	}
	return haveAnnis;
});
console.log("Find users who have 'annis' on their name.", annis);

// Find users who have articles on year 2020.
article2020 = users.filter(user => {
	const articlePublishAt2020 = user.articles.filter(article => {
		return new Date(article.title.published_at).getFullYear() == 2020;
	});

	return articlePublishAt2020.length > 0;
});
console.log("Find users who have articles on year 2020.", article2020);

//Find users who are born on 1986.
usersBornOn1986 = users.find(user => {
	return new Date(user.profile.birthday).getFullYear() == 1986;
});
console.log("Find users who are born on 1986.", usersBornOn1986);

//Find articles that contain "tips" on the title.
containsTips = users.filter(user => {
	const containsTips = user.articles.filter(article => {
		return article.title.toLowerCase().includes("tips");
	});

	return containsTips.length > 0;
});

console.log("Find articles that contain \"tips\" on the title.", containsTips);

//Find articles published before August 2019.
publishedBefore2019 = users.filter(user => {
	const articlePublishAt2019 = user.articles.filter(article => {
		return new Date(article.title.published_at).getFullYear() < 2019;
	});

	return articlePublishAt2019.length > 0;
});

console.log("Find articles published before August 2019.", publishedBefore2019);