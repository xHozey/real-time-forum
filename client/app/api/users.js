import { sendMessage } from "../utils/helpers.js";
export let myNickname;
export let myId;
export let targetId;

export const getUsers = async () => {
  const usersList = document.querySelector(".users-list");
  try {
    let res = await fetch("/api/info");
    let data = await res.json();
    console.log(data);

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

export const target = async (id) => {
  let display = document.getElementById("messages-display");
  if (targetId == id) {
    targetId = 0;
    display.classList.add("hidden");
    return;
  }
  let header = document.getElementById("user-nickname");
  let messagesContainer = document.querySelector(".messages-container");
  header.innerText = document.getElementById(id).dataset.nickname;
  messagesContainer.innerHTML = "";
  targetId = id;

  try {
    let res = await fetch(`/api/client/${id}`);
    let data = await res.json();
    data.forEach((msg) => {
      sendMessage(msg.content, msg.sender);
    });
    display.classList.remove("hidden");
  } catch (err) {
    console.error(err);
  }
};

window.target = target;
