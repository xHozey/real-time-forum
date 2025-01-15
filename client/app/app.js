import { login, register } from "./templates/auth.js";
import { registerSendData, loginSendData } from "./api/auth.js";
import { main } from "./templates/main.js";
const app = document.getElementById("app");

const route = (event) => {
  event = event || window.event;
  event.preventDefault();
  window.history.pushState({}, "", event.target.href);
  handleLocation();
};

export const handleLocation = async () => {
  const path = window.location.pathname;
  switch (path) {
    case "/register":
      app.innerHTML = register;
      registerSendData();
      break;
    case "/login":
      app.innerHTML = login;
      loginSendData();
      break;
    case "/":
      app.innerHTML = main;
      break;
    default:
      app.innerHTML = "<h1>in progress</h1>";
      break;
  }
};

window.onpopstate = handleLocation;
window.route = route;

handleLocation();
