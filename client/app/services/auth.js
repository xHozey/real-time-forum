import { registerSendData, loginSendData } from "../api/auth.js";

export const registerService = () => {
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
      if (!nickname.value.trim()) {
        showError(nickname, "Nickname is required.");
        return;
      }
      if (!age.value || isNaN(age.value)) {
        showError(age, "Age is required.");
        return;
      }
      if (!gender.value.trim()) {
        showError(gender, "Gender is required.");
        return;
      }
      if (!firstName.value.trim()) {
        showError(firstName, "First name is required.");
        return;
      }
      if (!lastName.value.trim()) {
        showError(lastName, "Last name is required.");
        return;
      }
      if (!email.value.trim()) {
        showError(email, "Email is required.");
        return;
      }
      if (!password.value.trim()) {
        showError(password, "Password is required.");
        return;
      }
      clearError();
      const user = {
        nickname: nickname.value,
        age: parseInt(age.value, 10),
        gender: gender.value,
        firstName: firstName.value,
        lastName: lastName.value,
        email: email.value,
        password: password.value,
      };
      await registerSendData(user);
    });
};

export const loginSerive = () => {
  document.getElementById("login-btn").addEventListener("click", async () => {
    const username = document.getElementById("username");
    const password = document.getElementById("password");
    if (!username.value.trim()) {
      showError(username, "username is required.");
      return;
    }
    if (!password.value.trim()) {
      showError(password, "Password is required.");
      return;
    }

    clearError();
    const user = {
      nickname: username.value,
      password: password.value,
    };
    loginSendData(user);
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
