document.addEventListener("DOMContentLoaded", function () {
  const sidebarLinks = document.querySelectorAll(".sidebar a");

  sidebarLinks.forEach(function (link) {
    link.addEventListener("click", function (event) {
      event.preventDefault();

      const targetId = this.getAttribute("href").substring(1);
      const targetSection = document.getElementById(targetId);

      if (targetSection) {
        // Masquer toutes les sections
        document
          .querySelectorAll(".content section")
          .forEach(function (section) {
            section.classList.remove("active");
          });

        // Afficher uniquement la section ciblée
        targetSection.classList.add("active");
      }
    });
  });
});

document
  .getElementById("modifyButton")
  .addEventListener("click", function (event) {
    event.preventDefault();
    toggleFields();
  });

function toggleFields() {
  var fields = document.querySelectorAll("#myForm input");
  var modifyButton = document.getElementById("modifyButton");
  var engebtn = document.getElementById("engebtn");

  if (modifyButton.textContent === "Modifier") {
    fields.forEach(function (field) {
      field.classList.remove("disabled");
    });

    modifyButton.textContent = "Annuler";
    engebtn.style.display = "block";
  } else {
    fields.forEach(function (field) {
      field.classList.add("disabled");
    });
    modifyButton.textContent = "Modifier";
    engebtn.style.display = "none";
  }
}

const table = document.getElementById("myTable");
const formContainer = document.getElementById("form-container");
const form = document.getElementById("myForm");
const siteInput = document.querySelector(".site");
const numeroInput = document.querySelector(".numero");
const clientInput = document.querySelector(".client");
const volumesInput = document.querySelector(".volumes");
const rechargeInput = document.querySelector(".recharge");
const expirationInput = document.querySelector(".expiration");
const Btnfermer = document.querySelector(".Btnfermer");
const customerId = document.querySelector(".customer-id");
const siteId = document.querySelector(".site-id");
const numberId = document.querySelector(".number-id");
const rechargeId = document.querySelector(".recharge-id");
const delSiteId = document.querySelector(".delete-site-id");
const delNumberId = document.querySelector(".delete-number-id");
const delRechargeId = document.querySelector(".delete-recharge-id");
var fieldDATA = document.querySelectorAll("#myForm input");
var modifyButtonDATA = document.getElementById("modifyButton");
var engebtnDATA = document.getElementById("engebtn");

table.addEventListener("click", function (event) {
  const clickedRow = event.target.parentNode;
  const cells = clickedRow.querySelectorAll("td");

  siteInput.value = cells[1].textContent;
  numeroInput.value = cells[2].textContent;
  clientInput.value = cells[3].textContent;
  volumesInput.value = cells[4].textContent;
  rechargeInput.value = cells[5].textContent;
  expirationInput.value = cells[6].textContent;
  customerId.value = cells[7].textContent;
  siteId.value = cells[8].textContent;
  numberId.value = cells[9].textContent;
  rechargeId.value = cells[10].textContent;
  delSiteId.value = cells[8].textContent;
  delNumberId.value = cells[9].textContent;
  delRechargeId.value = cells[10].textContent;
  formContainer.classList.add("show");
});

Btnfermer.addEventListener("click", function () {
  formContainer.classList.remove("show");

  if (modifyButtonDATA.textContent === "Annuler") {
    fieldDATA.forEach(function (field) {
      field.classList.add("disabled");
    });
    modifyButtonDATA.textContent = "Modifier";
    engebtnDATA.style.display = "none";
  }
});

document.addEventListener("DOMContentLoaded", function () {
  var filtresBtn = document.getElementById("filtresBtn");
  var autres = document.getElementById("autres");

  filtresBtn.addEventListener("click", function () {
    autres.style.display = autres.style.display === "flex" ? "none" : "flex";
  });
});

// **********************USER FORMULAIRE

document.addEventListener("DOMContentLoaded", function () {
  var cliqueUser = document.getElementById("cliqueUser");
  var formulaireU = document.getElementById("formulaireU");
  var closeBtnU = document.getElementById("closeBtnU");
  cliqueUser.addEventListener("click", function () {
    formulaireU.style.display =
      formulaireU.style.display === "flex" ? "none" : "flex";
  });
  closeBtnU.addEventListener("click", function () {
    formulaireU.style.display = "none";
  });
});

/***************** espace dossier******************************* */
document.addEventListener("DOMContentLoaded", function () {
  var filtresBtnd = document.getElementById("filtresBtnd");
  var autres1 = document.getElementById("autres1");

  filtresBtnd.addEventListener("click", function () {
    autres1.style.display = autres1.style.display === "flex" ? "none" : "flex";
  });
});

//**********gestion du formulaire de modification des dossiers

// Fonction pour afficher le formulaire de modification
function afficherFormulaire() {
  document.getElementById("modificationForm").classList.remove("hidden");
  document.getElementById("modificationForm").classList.add("visible");
}
// Fonction pour cacher le formulaire de modification
function cacherFormulaire() {
  document.getElementById("modificationForm").classList.remove("visible");
  document.getElementById("modificationForm").classList.add("hidden");
}
// Fonction pour remplir le formulaire avec les valeurs par défaut
function remplirFormulaire(
  project,
  secteur,
  contact,
  id,
  updateId,
  description
) {
  //document.getElementById("inputNomSociete").value = nomSociete;
  document.getElementById("project-name-dash").value = project;
  //document.getElementById("inputNomCommercial").value = nomCommercial;
  document.getElementById("inputSecteur").value = secteur;
  document.getElementById("inputContact").value = contact;
  document.getElementById("project-id-hidden").value = id;
  document.getElementById("project-id-update").value = updateId;
  document.getElementById("project-description").value = description;
}

// Événement de clic sur une boîte
document.querySelectorAll(".boxx").forEach((box) => {
  box.addEventListener("click", () => {
    // Remplacer les valeurs par défaut du formulaire avec les valeurs des span correspondants
    remplirFormulaire(
      //box.querySelector("#customer-name").innerText,
      box.querySelector("#project-name-dash").innerText,
      //box.querySelector("#commercial-name").innerText,
      box.querySelector("#secteur").innerText,
      box.querySelector("#contact").innerText,
      box.querySelector("#project-id").innerText,
      box.querySelector("#project-id").innerText,
      box.querySelector("#description").innerText
    );
    // Afficher le formulaire de modification
    afficherFormulaire();
  });
});
document.getElementById("closeBtnMD").addEventListener("click", () => {
  cacherFormulaire();
  TO;
});

// Événement de clic sur le bouton Supprimer
document.getElementById("supprimer").addEventListener("click", () => {
  // Code pour supprimer le dossier
});

//**********tableau dossier

// Sélectionnez toutes les lignes de tableau
var rows = document.querySelectorAll("tr");

// Parcourez chaque ligne de tableau
rows.forEach(function (row) {
  // Ajoutez un écouteur d'événements 'click' à chaque ligne de tableau
  row.addEventListener("click", function () {
    // Supprimez la classe 'selected' de toutes les lignes de tableau
    rows.forEach(function (row) {
      row.classList.remove("selected");
    });
    // Ajoutez la classe 'selected' à la ligne de tableau cliquée
    row.classList.add("selected");
  });
});
//****** formulaire dossier

document.addEventListener("DOMContentLoaded", function () {
  var newdossier = document.getElementById("newdossier");
  var ajoutclient = document.getElementById("ajoutclient");
  var closeBtn = document.getElementById("closeBtn");

  newdossier.addEventListener("click", function () {
    ajoutclient.style.display =
      ajoutclient.style.display === "flex" ? "none" : "flex";
  });
  closeBtn.addEventListener("click", function () {
    ajoutclient.style.display = "none";
  });
});

let currentStep = 1;

function nextStep() {
  document.getElementById("step" + currentStep).style.display = "none";
  currentStep++;
  document.getElementById("step" + currentStep).style.display = "block";
}

function prevStep() {
  document.getElementById("step" + currentStep).style.display = "none";
  currentStep--;
  document.getElementById("step" + currentStep).style.display = "block";
}

function enregistrer() {
  // Code pour enregistrer les données
  alert("Données enregistrées avec succès !");
}

function previewImage(event) {
  const preview = document.getElementById("preview");
  preview.style.display = "block";
  preview.src = URL.createObjectURL(event.target.files[0]);
}

/*****************************espace gestion des incidents */
//document.addEventListener("DOMContentLoaded", function() {
//  var interventionBTN = document.getElementById("interventionBTN");
//  var intervention = document.getElementById("intervention");
//  var closeBtninc = document.getElementById("closeBtninc");
//  interventionBTN.addEventListener("click", function() {
//    intervention.style.display =
//      intervention.style.display === "flex" ? "none" : "flex";
//  });
//  closeBtninc.addEventListener("click", function() {
//    intervention.style.display = "none";
//  });
//});

//commercialUpdateBtn = document.querySelector("#commercial-update-button")
//
//commercialUpdateEvent = (e) => {
//  e.preventDefault()
//  deleteTrElement('#commercial-table tr', '.commercial-id-update')
//  commercialForm.submit()
//}
//
//commercialUpdateBtn.addEventListener('click', commercialUpdateEvent)
//commercialUpdateBtn.removeEventListener('click', commercialUpdateEvent)
//
//commercialTable = document.querySelector('#commercial-table')
//commercialForm = document.querySelector("#commercial-form")
//commercialFormContainer = document.querySelector(".commercial-form-container")
//commercialName = document.querySelector('.agent-name')
//commercialFirstName = document.querySelector('.agent-first-name')
//commercialRole = document.querySelector('.agent-role')
//commercialAddress = document.querySelector('.agent-address')
//commercialId = document.querySelector('.agent-id')
//deleteCommercialId = document.querySelector('.delete-commercial-id')
//
//let commercialCloseBtn = document.querySelector(".commercial-close-btn");
//
//commercialTable.addEventListener("click", function(event) {
//  const clickedRow = event.target.parentNode;
//  const cells = clickedRow.querySelectorAll("td");
//
//  commercialName.value = cells[0].textContent;
//  commercialFirstName.value = cells[1].textContent;
//  commercialRole.value = cells[2].textContent;
//  commercialAddress.value = cells[3].textContent;
//  commercialId.value = cells[4].textContent;
//  deleteCommercialId.value = cells[4].textContent;
//
//  commercialFormContainer.classList.add("show");
//});
//
//commercialCloseBtn.addEventListener("click", function() {
//  commercialFormContainer.classList.remove("show");
//});

function showHiddenForm(id, deleteId, container, actor, btn) {
  let rows = document.querySelectorAll(actor + "-table" + " tr");
  let inputs = document.querySelectorAll(actor + "-form" + " input");
  container = document.querySelector(container);
  btn = document.querySelector(btn);

  rows.forEach((row) => {
    row.addEventListener("click", () => {
      cols = row.querySelectorAll("td");

      for (let i = 0; i < cols.length; i++) {
        console.log(inputs[i].className + " = " + cols[i].textContent);
        inputs[i].value = cols[i].textContent;
      }

      deleteRow(id, deleteId);
      container.classList.add("show");
    });
  });

  btn.addEventListener("click", function () {
    container.classList.remove("show");
  });
}

function deleteRow(id, deleteId) {
  deleteId = document.querySelector(deleteId);
  id = document.querySelector(id);

  deleteId.value = id.value;
}

showHiddenForm(
  "#section6 .commercial-id",
  "#section6 .delete-commercial-id",
  ".commercial-form-container",
  "#commercial",
  ".commercial-close-btn"
);
showHiddenForm(
  "#section6 .customer-id",
  "#section6 .delete-customer-id",
  ".customer-form-container",
  "#customer",
  ".customer-close-btn"
);
showHiddenForm(
  "#section6 .user-id",
  "#section6 .delete-user-id",
  ".user-form-container",
  "#user",
  ".user-close-btn"
);
showHiddenForm(
  "#section6 .service-id",
  "#section6 .delete-service-id",
  ".service-form-container",
  "#service",
  ".service-close-btn"
);
showHiddenForm(
  "#section5 .incident-id",
  "#section5 .delete-incident-id",
  ".incident-form-container",
  "#incident",
  ".incident-close-btn"
);
showHiddenForm(
  "#section4 .prospect-id",
  "#section4 .delete-prospect-id",
  ".prospect-form-container",
  "#prospect",
  ".prospect-close-btn"
);

/*****************************************code modification */
document
  .getElementById("modifierx")
  .addEventListener("click", function (event) {
    event.preventDefault();
    toggleFieldsx();
  });

function toggleFieldsx() {
  var fields1 = document.querySelectorAll("#myForm1 input");
  var modifierx = document.getElementById("modifierx");
  var engebtn1 = document.getElementById("bnx");

  if (modifierx.textContent === "Modifier") {
    fields1.forEach(function (field) {
      field.classList.remove("disabled");
    });

    modifierx.textContent = "Annuler";
    engebtn1.style.display = "block";
  } else {
    fields1.forEach(function (field) {
      field.classList.add("disabled");
    });
    modifierx.textContent = "Modifier";
    engebtn1.style.display = "none";
  }
}

/****************************modification suivie prospects*/

document
  .getElementById("modifsuive")
  .addEventListener("click", function (event) {
    event.preventDefault();
    toggleFieldsz();
  });

function toggleFieldsz() {
  var fields4 = document.querySelectorAll("#prospect-form input ");
  var modifsuive = document.getElementById("modifsuive");
  var BXX = document.getElementById("BXX");

  if (modifsuive.textContent === "Modifier") {
    fields4.forEach(function (field) {
      field.classList.remove("disabled");
    });

    modifsuive.textContent = "Annuler";
    BXX.style.display = "block";
  } else {
    fields4.forEach(function (field) {
      field.classList.add("disabled");
    });
    modifsuive.textContent = "Modifier";
    BXX.style.display = "none";
  }
}

/***********************modific tion suivie prospect */
document
  .getElementById("modifierP")
  .addEventListener("click", function (event) {
    event.preventDefault();
    toggleFieldsP();
  });

function toggleFieldsP() {
  var fields0 = document.querySelectorAll("#incident-form input ");
  var modifierP = document.getElementById("modifierP");
  var STAY = document.getElementById("STAY");

  if (modifierP.textContent === "Modifier") {
    fields0.forEach(function (field) {
      field.classList.remove("disabled");
    });

    modifierP.textContent = "Annuler";
    STAY.style.display = "block";
  } else {
    fields0.forEach(function (field) {
      field.classList.add("disabled");
    });
    modifierP.textContent = "Modifier";
    STAY.style.display = "none";
  }
}

/****************************************code special acueille */

//************************************** gerer les cartes
const cardContainer = document.querySelector(".card-container");
const cards = document.querySelectorAll(".card");
const prevButton = document.getElementById("prevBtn");
const nextButton = document.getElementById("nextBtn");

let currentIndex = 0;

function showCard(index) {
  cards.forEach((card) => {
    card.style.transform = `translateX(-${index * 100}%)`;
  });
}

function showNextCard() {
  currentIndex = (currentIndex + 1) % cards.length;
  showCard(currentIndex);
}

function showPrevCard() {
  currentIndex = (currentIndex - 1 + cards.length) % cards.length;
  showCard(currentIndex);
}

nextButton.addEventListener("click", showNextCard);
prevButton.addEventListener("click", showPrevCard);

// Auto-scroll
setInterval(showNextCard, 4000); // Change card every 3 seconds

const allSideMenu = document.querySelectorAll("#sidebar .side-menu.top li a");

allSideMenu.forEach((item) => {
  const li = item.parentElement;

  item.addEventListener("click", function () {
    allSideMenu.forEach((i) => {
      i.parentElement.classList.remove("active");
    });
    li.classList.add("active");
  });
});

/************************************************************************************************ */

const ctx = document.getElementById("dataChart").getContext("2d");

// Fonction pour générer des données aléatoires pour les deux courbes
function generateRandomData() {
  const data = [];
  for (let i = 0; i < 10; i++) {
    data.push(Math.floor(Math.random() * 20));
  }
  return data;
}

let airtelData = generateRandomData();
let mtnData = generateRandomData();

const chart = new Chart(ctx, {
  type: "line",
  data: {
    labels: [
      "Mars",
      "Avril",
      "Mai",
      "Juin",
      "Juillet",
      "Août",
      "Septembre",
      "Octobre",
      "Novembre",
      "Décembre",
    ],
    datasets: [
      {
        label: "Airtel",
        data: airtelData,
        borderColor: "red",
        borderWidth: 2,
        fill: false,
      },
      {
        label: "MTN",
        data: mtnData,
        borderColor: "yellow",
        borderWidth: 2,
        fill: false,
      },
    ],
  },
  options: {
    animation: {
      onProgress: function () {
        // À chaque mise à jour, générez de nouvelles données aléatoires pour les courbes
        airtelData = generateRandomData();
        mtnData = generateRandomData();
        chart.update();
      },
    },
    scales: {
      y: {
        beginAtZero: true,
        max: 20,
        title: {
          display: true,
          text: "Consommation (Go)",
        },
      },
      x: {
        title: {
          display: true,
          text: "Mois",
        },
      },
    },
  },
});

//***********gerer le calendrier */
function mettreAJourCalendrier() {
  var dateActuelle = new Date();
  var jourActuel = dateActuelle.getDate();
  var moisActuel = dateActuelle.getMonth() + 1;
  var elementsDate = document.getElementsByClassName("date");
  var semaineActuelle = dateActuelle.getDate();

  //***************mettre a jour la semaine

  let currentDate = new Date();

  // Obtenir le numéro du jour de la semaine (0 pour Dimanche, 1 pour Lundi, etc.)
  let currentDay = currentDate.getDay();

  // Obtenir le jour du mois
  let currentMonthDay = currentDate.getDate();
  // Mettre à jour la liste des jours dans le calendrier
  let datesContainer = document.getElementById("dates");
  datesContainer.innerHTML = "";

  // Boucle pour ajouter les jours de la semaine en cours
  for (let i = 1; i < 8; i++) {
    let day = document.createElement("span");
    day.className = "date";
    // Calculer le jour en cours de la semaine
    let currentWeekDay = currentMonthDay - currentDay + i;
    // Ajouter le jour au conteneur des dates
    day.textContent = currentWeekDay;
    datesContainer.appendChild(day);
  }

  // Retirer la classe "aujourd'hui" de tous les éléments de date
  for (var i = 0; i < elementsDate.length; i++) {
    elementsDate[i].classList.remove("aujourd-hui");
  }

  // Trouver l'élément de date correspondant à la date actuelle et lui ajouter la classe "aujourd'hui"
  for (var j = 0; j < elementsDate.length; j++) {
    if (parseInt(elementsDate[j].innerText) === jourActuel) {
      elementsDate[j].classList.add("aujourd-hui");
      break;
    }
  }

  //****************Mettre à jour l'heure actuelle
  var heureElement = document.getElementById("heure");
  var heures = dateActuelle.getHours().toString().padStart(2, "0");
  var minutes = dateActuelle.getMinutes().toString().padStart(2, "0");
  heureElement.innerText = heures + ":" + minutes;

  // ***************Mettre à jour le mois si nécessaire
  var nomMoisElement = document.getElementById("nom-mois");
  var dernierJourMois = new Date(
    dateActuelle.getFullYear(),
    moisActuel,
    0
  ).getDate();
  if (jourActuel === dernierJourMois) {
    // Si c'est le dernier jour du mois, passer au mois suivant
    moisActuel = (moisActuel % 12) + 1;
    nomMoisElement.innerText = obtenirNomMois(moisActuel);
  }
}

function obtenirNomMois(numeroMois) {
  var nomsMois = [
    "Janvier",
    "Février",
    "Mars",
    "Avril",
    "Mai",
    "Juin",
    "Juillet",
    "Août",
    "Septembre",
    "Octobre",
    "Novembre",
    "Décembre",
  ];
  return nomsMois[numeroMois - 1];
}

// Appeler la fonction pour la première fois
mettreAJourCalendrier();

// Mettre à jour le calendrier toutes les secondes
setInterval(mettreAJourCalendrier, 1000);
