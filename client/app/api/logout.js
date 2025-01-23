import { getRequest } from "../utils/helpers.js";
import { nav } from "../utils/navigation.js";

export const logout = () => {
  document
    .getElementById("logout-button")
    .addEventListener("click", async () => {
      await fetch("logout");
      nav("/login");
    });
};
