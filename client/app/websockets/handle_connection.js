import { targetId } from "../services/users.js";
import { sendMessage, createUser } from "../utils/helpers.js";
import { myNickname } from "../api/users.js";

export const handleIncommingMessage = (msg) => {
  const typingIndicator = document.getElementById("typing-indicator");
  const targetDiv = document.getElementById(msg.sender);
  const user = document.getElementById(`typing-indicator-${msg.sender}`);
  switch (msg.type) {
    case "chat":
      if (targetId == msg.sender) {
        typingIndicator.classList.add("hidden");
        sendMessage(msg.content, targetDiv.dataset.nickname);
      } else {
        targetDiv.classList.add("new-message");
      }
      document.querySelector(".users-list").prepend(targetDiv);
      break;

    case "status":
      let userStatus = document.getElementById(`status-${msg.id}`);
      if (userStatus) {
        userStatus.classList.remove("offline", "online");
        if (msg.status) {
          userStatus.classList.add("online");
        } else {
          userStatus.classList.add("offline");
        }
      } else {
        createUser(msg);
      }
      break;

    case "typing":
      if (targetId == msg.sender) {
        document.getElementById(
          "typing-user"
        ).innerText = `${targetDiv.dataset.nickname} is typing`;
        typingIndicator.classList.remove("hidden");
      } else {
        user.classList.remove("hidden");
      }
      break;
    case "stop-typing":
      typingIndicator.classList.add("hidden");
      user.classList.add("hidden");

  }
};

export const sendPrivateMessage = (conn) => {
  let message = document.getElementById("message");
  if (conn.readyState != WebSocket.OPEN) {
    console.error("WebS9.3ocket is closed");
    return;
  }
  if (message.value.length > 500 || !message.value.trim()) {
    message.classList.add("error");
    setTimeout(() => {
      message.classList.remove("error");
    }, 1000);
    return;
  }
  if (
    document.getElementById(`status-` + targetId).classList.contains("offline")
  ) {
    message.classList.add("error");
    setTimeout(() => {
      message.classList.remove("error");
    }, 1000);
  } else {
    message.classList.remove("error");
    sendMessage(message.value, myNickname);
    document
      .querySelector(".users-list")
      .prepend(document.getElementById(targetId));
    conn.send(JSON.stringify({ target: targetId, content: message.value }));
    message.value = "";
  }
};

export const typingStatus = (conn) => {
  const input = document.getElementById("message");
  let debounce = false;
  let timer;
  input.addEventListener("input", () => {
    const targetStatus = document.getElementById(`status-${targetId}`);
    if (targetStatus.classList.contains("offline")) return;
    clearTimeout(timer);
    if (!debounce) {
      conn.send(
        JSON.stringify({ target: targetId, type: "typing", content: "typing" })
      );
      debounce = true;
    }
    timer = setTimeout(() => {
      conn.send(
        JSON.stringify({ target: targetId, type: "typing", content: "stop-typing" })
      );
      debounce = false;
    }, 500);
  });
};
