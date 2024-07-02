function changerDePage() {
  window.location.href = "Dashboard.html";
}


//fonction pour actualiser la date 

function updateDateTime() {
  const dateTimeElement = document.getElementById('date-time');
  const now = new Date();
  const options = { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric', hour: '2-digit', minute: '2-digit' };
  const formattedDateTime = now.toLocaleDateString('fr-FR', options);
  dateTimeElement.textContent = formattedDateTime;
}

updateDateTime();



//intervalle
const cardContainer = document.getElementById('cardContainer');
const chevronLeft = document.getElementById('chevronLeft');
const chevronRight = document.getElementById('chevronRight');
let cardIndex = 0;
const cards = ["Carte 1", "Carte 2", "Carte 3", "Carte 4"];
let autoChangeInterval;

function updateCard() {
  cardContainer.textContent = cards[cardIndex];
}

function changeCard(direction) {
  if (direction === "left") {
    cardIndex = (cardIndex - 1 + cards.length) % cards.length;
  } else {
    cardIndex = (cardIndex + 1) % cards.length;
  }
  updateCard();
}

function startAutoChange() {
  autoChangeInterval = setInterval(() => changeCard("right"), 25000);
}

chevronLeft.addEventListener('click', () => {
  changeCard("left");
  clearInterval(autoChangeInterval);
  startAutoChange();
});
chevronRight.addEventListener('click', () => {
  changeCard("right");
  clearInterval(autoChangeInterval);
  startAutoChange();
});

updateCard();
startAutoChange();

// Reste du code pour le profil utilisateur (non modifié)





//profile utilisateur 
const profileImage = document.getElementById('profileImage');
const employeeNameElement = document.getElementById('employeeName');

// Remplacez ces valeurs par les données de l'employé connecté depuis votre base de données
const employeeData = {
  imageSrc: 'chemin/vers/image.jpg',
  name: 'Nom de l\'employé'
};

profileImage.src = employeeData.imageSrc;
employeeNameElement.textContent = employeeData.name;


