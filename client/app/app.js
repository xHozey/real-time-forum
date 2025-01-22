import { login, register } from "./templates/auth.js";
import { registerSendData, loginSendData } from "./api/auth.js";
import { main } from "./templates/main.js";
import { getUsers, loadMessages, targetId } from "./api/users.js";
import { connectToServer } from "./websockets/connection.js";
import { throttle } from "./utils/helpers.js";
import { showPostPanel } from "./api/add_post.js";
import { getPosts } from "./api/get_posts.js";
const app = document.getElementById("app");
const link = document.getElementById("css");
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
      link.innerHTML = `<link rel="stylesheet" href="./app/styles/auth.css">`
      app.innerHTML = register;
      registerSendData();
      break;
    case "/login":
      link.innerHTML = `<link rel="stylesheet" href="./app/styles/auth.css">`
      app.innerHTML = login;
      loginSendData();
      break;
    case "/":
      link.innerHTML = `
      <link rel="stylesheet" href="./app/styles/main.css">
      <link rel="stylesheet" href="./app/styles/add_post.css">
      <link rel="stylesheet" href="./app/styles/posts.css">

      `
      app.innerHTML = main;
      getUsers();
      connectToServer();
      throttle();
      showPostPanel()
      getPosts()
      break;
    default:
      app.innerHTML = "<h1>in progress</h1>";
      break;
  }
};

window.onpopstate = handleLocation;
window.route = route;

handleLocation();
