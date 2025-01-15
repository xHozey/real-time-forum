import { nav } from "../utils/navigation.js";

export const registerSendData = () => {
  document
    .getElementById("register-btn")
    .addEventListener("click", async () => {
      const nickname = document.getElementById("nickname");
      const age = document.getElementById("age");
      const gender = document.getElementById("gender");
      const firstName = document.getElementById("firstName");
      const lastName = document.getElementById("lastName");
      const email = document.getElementById("email");
      const password = document.getElementById("password");
      if (!nickname.value) {
        showError(nickname, "Nickname is required.");
        return;
      }
      if (!age.value || isNaN(age.value)) {
        showError(age, "Age is required.");
        return;
      }
      if (!gender.value) {
        showError(gender, "Gender is required.");
        return;
      }
      if (!firstName.value) {
        showError(firstName, "First name is required.");
        return;
      }
      if (!lastName.value) {
        showError(lastName, "Last name is required.");
        return;
      }
      if (!email.value) {
        showError(email, "Email is required.");
        return;
      }
      if (!password.value) {
        showError(password, "Password is required.");
        return;
      }

      try {
        clearError();
        const user = {
          nickname: document.getElementById("nickname").value,
          age: parseInt(document.getElementById("age").value, 10),
          gender: document.getElementById("gender").value,
          firstName: document.getElementById("firstName").value,
          lastName: document.getElementById("lastName").value,
          email: document.getElementById("email").value,
          password: document.getElementById("password").value,
        };
        const res = await fetch("/api/register", {
          method: "POST",
          body: JSON.stringify(user),
        });
        if (!res.ok) {
          document.querySelector(".error-message").innerText = await res.json();
        } else {
          nav("/login");
        }
      } catch (err) {
        console.error(err);
      }
    });
};

export const loginSendData = () => {
  document.getElementById("login-btn").addEventListener("click", async () => {
    const username = document.getElementById("username");
    const password = document.getElementById("password");
    if (!username.value) {
      showError(username, "username is required.");
      return;
    }
    if (!password.value) {
      showError(password, "Password is required.");
      return;
    }

    try {
      clearError();
      const user = {
        nickname: document.getElementById("username").value,
        password: document.getElementById("password").value,
      };
      const res = await fetch("/api/login", {
        method: "POST",
        body: JSON.stringify(user),
      });
      if (!res.ok) {
        document.querySelector(".error-message").innerText = await res.json();
      } else {
        nav("/");
      }
    } catch (err) {
      console.error(err);
    }
  });
};

const showError = (field, error) => {
  clearError();
  document.querySelector(".error-message").innerText = error;
  field.classList.add("error");
};

const clearError = () => {
  const fields = document.querySelectorAll(".error");
  document.querySelector(".error-message").innerText = "";
  if (fields) {
    fields.forEach((field) => {
      field.classList.remove("error");
    });
  }
};
