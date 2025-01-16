import { handleLocation } from "../app.js";
export const nav = (path) => {
  window.history.pushState({}, "", path);
  handleLocation();
};
