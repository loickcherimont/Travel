/* JS is mixed with CSS */
	const cards = document.querySelectorAll('.card');
	const popupPages = document.querySelectorAll('.popup-container');
	const quitPopups = document.querySelectorAll('.fa-times');

	cards.forEach((card, index) => {
		card.onclick = () => {
			popupPages[index].style.transform = "initial";
		};
	});

	quitPopups.forEach((close, index)=> {
		close.onclick = () => {
			popupPages[index].style.transform = "scale(0)";
		};
	});