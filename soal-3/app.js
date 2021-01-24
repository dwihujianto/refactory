const fs = require('fs');

const rawdata = fs.readFileSync('data.json');
const items = JSON.parse(rawdata);

//Find items in Meeting Room.
itemInMeetingRoom = items.find(item => {
	return item.placement.name == "Meeting Room";
});

console.log("Find items in Meeting Room.", itemInMeetingRoom);

//Find all electronic devices.
electronics = items.filter(item => {
	return item.type == "electronic";
});

console.log("Find all electronic devices.", electronics);

//Find all furnitures.
furnitures = items.filter(item => {
	return item.type == "furniture";
});

console.log("Find all furnitures.", furnitures);

//Find all items was purchased at 16 Januari 2020.
purchases = items.filter(item => {
	date = new Date(item.purchased_at);
	return date.getFullYear() == 2020 && date.getDate() == 16 && date.getMonth() == 0;
});

console.log("Find all furnitures.", purchases);

//Find all items with brown color.
brownColors = items.filter(item => {
	return item.name.toLowerCase().includes('brown');
});

console.log("Find all items with brown color.", brownColors)