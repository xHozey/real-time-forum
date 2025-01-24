import { myNickname } from "../api/users.js";
import { targetId } from "../services/users.js";
import { sendMessage, createUser } from "../utils/helpers.js";
import { nav } from "../utils/navigation.js";

export const connectToServer = () => {
  const conn = new WebSocket("ws://localhost:8080/ws");
  conn.onmessage = (event) => {
    let res = JSON.parse(event.data);
    switch (res.type) {
      case "chat":
        const targetDiv = document.getElementById(res.sender);
        console.log(targetDiv);
        
        if (targetId == res.sender) {
          sendMessage(res.content, targetDiv.dataset.nickname);
        } else {
          targetDiv.classList.add("new-message");
        }
        document.querySelector(".users-list").prepend(targetDiv);
        break;

      case "status":
        let div = document.getElementById(`status-${res.id}`);
        if (div) {
          div.classList.remove("offline", "online");
          if (res.status) {
            div.classList.add("online");
          } else {
            div.classList.add("offline");
          }
        } else {
          createUser(res);
        }
        break;
    }
  };

  conn.onerror = (error) => {
    console.error("WebSocket error:", error);
  };

  conn.onclose = (event) => {
    console.log("WebSocket connection closed:", event.code, event.reason);
  };

  document.getElementById("btn-message").addEventListener("click", () => {
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
      document
        .getElementById(`status-` + targetId)
        .classList.contains("offline")
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
  });
  document
    .getElementById("logout-button")
    .addEventListener("click", async () => {
      conn.close(1000, "logout");
      await fetch("/logout");
      nav("/login");
    });
};
