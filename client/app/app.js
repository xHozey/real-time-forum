import { login, register } from "./templates/auth.js";
import { registerSendData, loginSendData } from "./api/auth.js";
import { main } from "./templates/main.js";
import { getUsers, loadMessages, targetId } from "./api/users.js";
import { connectToServer } from "./websockets/connection.js";
const app = document.getElementById("app");
const link = document.getElementById("css");
let timer;
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
      link.setAttribute("href", "./app/styles/auth.css");
      app.innerHTML = register;
      registerSendData();
      break;
    case "/login":
      link.setAttribute("href", "./app/styles/auth.css");
      app.innerHTML = login;
      loginSendData();
      break;
    case "/":
      link.setAttribute("href", "./app/styles/main.css");
      app.innerHTML = main;
      getUsers();
      connectToServer();
      document
        .querySelector(".messages-container")
        .addEventListener("scroll", (event) => {
          const container = document.querySelector(".messages-container");
          clearTimeout(timer);
          if (container.scrollTop == 0) {
            timer = setTimeout(() => {
              loadMessages(container, targetId);
            }, 1000);
          }
        });
      break;
    default:
      app.innerHTML = "<h1>in progress</h1>";
      break;
  }
};

window.onpopstate = handleLocation;
window.route = route;

handleLocation();
