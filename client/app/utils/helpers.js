export const escapeHtml = (unsafe) => {
  return unsafe
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#039;");
};

export const sendMessage = (message, name) => {
  let messagesContainer = document.querySelector(".messages-container");
  let div = document.createElement("div");
  div.classList.add("message");
  let author = document.createElement("div");
  author.classList.add("message-author");
  author.innerText = name;
  div.append(author);
  let content = document.createElement("div");
  content.classList.add("message-content");
  content.innerText = message;
  div.append(content);
  messagesContainer.append(div);
  messagesContainer.scroll(0, messagesContainer.scrollHeight);
};

export const prepandMessage = (message, name) => {
  let messagesContainer = document.querySelector(".messages-container");
  let div = document.createElement("div");
  div.classList.add("message");
  let author = document.createElement("div");
  author.classList.add("message-author");
  author.innerText = name;
  div.append(author);
  let content = document.createElement("div");
  content.classList.add("message-content");
  content.innerText = message;
  div.append(content);
  messagesContainer.prepend(div);
};

export const createUser = (user) => {
  const usersList = document.querySelector(".users-list");
  const div = document.createElement("div");
  div.classList.add("user-item");
  div.id = user.id;
  div.setAttribute("onclick", `target(${user.id})`);
  div.setAttribute("data-nickname", `${user.nickname}`);
  div.innerHTML = `${user.nickname} <span id="status-${
    user.id
  }" class="${user.status ? "online" : "offline"}"></span>`;
  usersList.append(div);
}