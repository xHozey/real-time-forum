import { register } from "./components/features/register_component.js";
import { login } from "./components/features/login_component.js";
const appView = document.getElementById("app");

const handleRouteChanges = () => {
  let viewComponent = null
  const path = window.location.pathname;
  switch (path) {
    case "/post":
      break;
    case "/login":
        viewComponent = new login()
      break;
    case "/register":
        viewComponent = new register()
      break;
    default:
      break;
  }
  if (viewComponent) {
    appView.innerHTML = '';  
    appView.append(viewComponent);  
  }
};

window.addEventListener("popstate", handleRouteChanges);

handleRouteChanges();
