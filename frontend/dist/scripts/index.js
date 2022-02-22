const menu = document.querySelector('.menu');
const navItems = document.querySelector('#nav-items');

menu.addEventListener('click', function () {
	navItems.classList.toggle('hidden');
});
