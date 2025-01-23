import { getRequest, prepandMessage, createUser } from "../utils/helpers.js";
export let myNickname;
export let myId;
export let targetId;
const limit = 10;
let messages = [];
export const getUsers = async () => {
  try {
    let res = await fetch("/api/info");
    let data = await res.json();
    myNickname = data.Nickname;
    myId = data.UserId;
    data.Clients.forEach((user) => {
      createUser(user);
    });
  } catch (err) {
    console.error(err);
  }
};

const target = async (id) => {
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
  await loadMessages(messagesContainer, targetId);
  display.classList.remove("hidden");
  messagesContainer.scroll(0, messagesContainer.scrollHeight);
};

export const loadMessages = async (container, id) => {
  const url =
    `/api/client/${id}?` + new URLSearchParams({ offset: messages[id] });
  const data = await getRequest(url);
  if (data.length != 0) {
    data.forEach((msg) => {
      prepandMessage(msg.content, msg.sender, new Date(msg.creation).toLocaleTimeString());
    });
    messages[id] += limit;
    container.scroll(0, 85 * limit);
  }
};

window.target = target;
