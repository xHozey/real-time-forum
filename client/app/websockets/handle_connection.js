import { targetId } from "../services/users.js";
import { sendMessage, createUser } from "../utils/helpers.js";
import { myNickname } from "../api/users.js";

export const handleIncommingMessage = (msg) => {
  switch (msg.type) {
    case "chat":
      const targetDiv = document.getElementById(msg.sender);
      if (targetId == msg.sender) {
        sendMessage(msg.content, targetDiv.dataset.nickname);
      } else {
        targetDiv.classList.add("new-message");
      }
      document.querySelector(".users-list").prepend(targetDiv);
      break;

    case "status":
      let div = document.getElementById(`status-${msg.id}`);
      if (div) {
        div.classList.remove("offline", "online");
        if (msg.status) {
          div.classList.add("online");
        } else {
          div.classList.add("offline");
        }
      } else {
        createUser(msg);
      }
      break;
  }
};

export const sendPrivateMessage = (conn) => {
  let message = document.getElementById("message");
  if (conn.readyState != WebSocket.OPEN) {
    console.error("WebSocket is closed");
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
//   const input = document.getElementById("message");
//   input.addEventListener("input", () => {
//     const targetStatus = document.getElementById(`status-${targetId}`);
//     if (targetStatus.classList.contains("offline")) return;
//     conn.send(JSON.stringify({target: targetId, type: "typing"}))
    
//   });
};
