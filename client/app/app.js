import { register } from "./components/features/register_component.js";
import { login } from "./components/features/login_component.js";
const appView = document.getElementById("app");

const handleRouteChanges = () => {
  let viewComponent = null;
  const path = window.location.pathname;
  switch (path) {
    case "/":
      break;
    case "/login":
      viewComponent = new login();
      break;
    case "/register":
      viewComponent = new register();
      break;
  }
  if (viewComponent) {
    appView.innerHTML = "";
    appView.append(viewComponent);
  }
};

const route = (event) => {
  event = event || window.event
  event.preventDefault();
  window.history.pushState({}, "", event.target.href);
  handleRouteChanges();
};

document.addEventListener("DOMContentLoaded", () => {
  document.querySelector(".href").addEventListener("click", route);
});

window.onpopstate = handleRouteChanges;
window.route = route;
handleRouteChanges();
