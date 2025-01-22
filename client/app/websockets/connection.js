import { myNickname, targetId, myId } from "../api/users.js";
import { sendMessage, createUser } from "../utils/helpers.js";
export const connectToServer = () => {
  const conn = new WebSocket("ws://localhost:8080/ws");
  conn.onmessage = (event) => {
    let res = JSON.parse(event.data);
    switch (res.type) {
      case "chat":
        const targetDiv = document.getElementById(res.sender);
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

  conn.onclose = (event) => {
    console.log("WebSocket connection closed:", event.reason);
  };
  conn.onerror = (event) => {
    console.error("WebSocket error:", error);
  };

  document.getElementById("btn-message").addEventListener("click", () => {
    let message = document.getElementById("message");
    if (conn.readyState != WebSocket.OPEN) {
      console.error("WebSocket is closed");
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
      conn.send(targetId + " " + message.value);
      message.value = "";
    }
  });
};
