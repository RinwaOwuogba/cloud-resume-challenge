const menu = document.querySelector('.menu');
const navItems = document.querySelector('#nav-items');

menu.addEventListener('click', function () {
	navItems.classList.toggle('hidden');
});

const API_URL = 'http://127.0.0.1:5000/api/visits';

function makeAPIRequest(method, errorMsg) {
	return fetch(API_URL, {
		method,
	}).then((resp) => {
		if (!resp.ok) throw new Error(errorMsg);
		return resp.text();
	});
}

function getVisitCount() {
	return makeAPIRequest('PUT', 'Unable to get current visit count');
}

function updateVisitCount() {
	return makeAPIRequest('PUT', 'Unable to update visit count');
}

function setVisitCount(count) {
	const visitCountNode = document.querySelector('#visit-count');
	visitCountNode.textContent = count;
}

function main() {
	console.log('running');
	updateVisitCount()
		.then(() => getVisitCount().then((count) => setVisitCount(count)))
		.catch(console.error);
}

main();
