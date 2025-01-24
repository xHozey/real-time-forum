import { loadMessages, limit } from "../api/users.js";
export let targetId;
export let messages = [];

export const target = async (id) => {
  let display = document.getElementById("messages-display");
  if (targetId == id) {
    targetId = 0;
    display.classList.add("hidden");
    return;
  }
  document.getElementById("message").classList.remove("error");
  let targetDiv = document.getElementById(id);
  targetDiv.classList.remove("new-message");
  let messagesContainer = document.querySelector(".messages-container");
  let header = document.getElementById("user-nickname");
  header.innerText = targetDiv.dataset.nickname;
  messagesContainer.innerHTML = "";
  targetId = id;
  messages[targetId] = 0;
  await loadMessages(messagesContainer, targetId, messages[targetId]);
  messages[id] += limit;
  display.classList.remove("hidden");
  messagesContainer.scroll(0, messagesContainer.scrollHeight);
};

window.target = target;
