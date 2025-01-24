import { nav } from "../utils/navigation.js";
export const registerSendData = async (user) => {
  await sendAuthRequest("/api/register", user, "/login");
};

export const loginSendData = async (user) => {
  await sendAuthRequest("/api/login", user, "/");
};

const sendAuthRequest = async (endpoint, data, redirect) => {
  try {
    const res = await fetch(endpoint, {
      method: "POST",
      body: JSON.stringify(data),
    });
    if (!res.ok) {
      document.querySelector(".error-message").innerText = await res.json();
    } else {
      nav(redirect);
    }
  } catch (err) {
    console.error(err);
  }
};
