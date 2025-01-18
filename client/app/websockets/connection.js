import { myNickname, targetId, myId } from "../api/users.js";
import { escapeHtml, sendMessage } from "../utils/helpers.js";
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
        if (!div) console.log("hello world");
        
        div.classList.remove("offline", "online");
        if (res.status) {
          div.classList.add("online");
        } else {
          div.classList.add("offline");
        }
        break;
    }
  };

document.getElementById("btn-message").addEventListener("click", () => {
    let message = document.getElementById("message");
    sendMessage(message.value, myNickname);
    conn.send(targetId + " " + message.value);
    message.value = "";
  });
};
