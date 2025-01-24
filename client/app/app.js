import { login, register } from "./templates/auth.js";
import { registerService, loginSerive } from "./services/auth.js";
import { main } from "./templates/main.js";
import { getUsers } from "./api/users.js";
import { connectToServer } from "./websockets/connection.js";
import { throttleMessages, throttlePosts } from "./utils/helpers.js";
import { showPostPanel } from "./services/threads.js";
import { getPosts } from "./api/get_threads.js";
import { addComments } from "./services/threads.js";

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
      link.innerHTML = `<link rel="stylesheet" href="./app/styles/auth.css">`;
      app.innerHTML = register;
      registerService();
      break;
    case "/login":
      link.innerHTML = `<link rel="stylesheet" href="./app/styles/auth.css">`;
      app.innerHTML = login;
      loginSerive();
      break;
    case "/":
      link.innerHTML = `
      <link rel="stylesheet" href="./app/styles/main.css">
      <link rel="stylesheet" href="./app/styles/add_post.css">
      <link rel="stylesheet" href="./app/styles/posts.css">
      <link rel="stylesheet" href="./app/styles/comments.css">
      <link rel="stylesheet" href="./app/styles/users.css">
      <link rel="stylesheet" href="./app/styles/navbar.css">
      `;
      app.innerHTML = main;
      getUsers();
      connectToServer();
      throttleMessages();
      showPostPanel();
      getPosts();
      addComments();
      throttlePosts();
      break;
    default:
      app.innerHTML = "<h1>in progress</h1>";
      break;
  }
};

window.onpopstate = handleLocation;
window.route = route;

handleLocation();
