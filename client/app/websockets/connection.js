import { myNickname, targetId, myId } from "../api/users.js";
import { escapeHtml, sendMessage } from "../utils/helpers.js";
export const connectToServer = () => {
  const conn = new WebSocket("ws://localhost:8080/ws");
  conn.onclose = () => {
    console.log("bye");
  };

  conn.onmessage = (event) => {
    try {
      let res = JSON.parse(event.data);
      let div = document.getElementById(`status-${res.id}`);
      div.classList.remove("offline", "online");
      if (res.status) {
        div.classList.add("online");
      } else {
        div.classList.add("offline");
      }
    } catch (err) {
      let message = event.data.split(/ (.+)/);
      const targetDiv = document.getElementById(message[0])
      if (targetId == message[0]) {
        sendMessage(
          message[1],
          targetDiv.dataset.nickname
        );
      } else {
        targetDiv.classList.add("new-message")
      }
      document.querySelector(".users-list").prepend(targetDiv)
    }
  };

  document.getElementById("btn-message").addEventListener("click", () => {
    let message = document.getElementById("message")
    sendMessage(message.value, myNickname);
    conn.send(targetId + " " + message.value);
    message.value = ""
  });
};
