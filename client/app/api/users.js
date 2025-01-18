import { prepandMessage } from "../utils/helpers.js";
export let myNickname;
export let myId;
export let targetId;
const limit = 10;
let messages = [];
export const getUsers = async () => {
  const usersList = document.querySelector(".users-list");
  try {
    let res = await fetch("/api/info");
    let data = await res.json();
    myNickname = data.Nickname;
    myId = data.UserId;
    data.Clients.forEach((user) => {
      const div = document.createElement("div");
      div.classList.add("user-item");
      div.id = user.user_id;
      div.setAttribute("onclick", `target(${user.user_id})`);
      div.setAttribute("data-nickname", `${user.nickname}`);
      div.innerHTML = `${user.nickname} <span id="status-${
        user.user_id
      }" class="${user.status ? "online" : "offline"}"></span>`;
      usersList.append(div);
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
  let y = container.scrollHeight;
  const query = new URLSearchParams({ offset: messages[id] });
  try {
    let res = await fetch(`/api/client/${id}?` + query);
    let data = await res.json();
    data.forEach((msg) => {
      prepandMessage(msg.content, msg.sender);
    });
    messages[id] += limit;
  } catch (err) {
    console.error(err);
  } finally {
    container.scroll(
      0,
      document.querySelector(".message").scrollHeight * limit
    );
  }
};

window.target = target;
