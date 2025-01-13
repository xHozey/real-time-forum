const socket = new WebSocket("ws://localhost:8080/ws");
const container = document.getElementById("container");
const users = [];
let display = 10;
socket.onopen = async () => {
  try {
    let res = await fetch("/api/info");
    let data = await res.json();
    data.sort((a, b) => {
      return a.nickname.toLowerCase().localeCompare(b.nickname.toLowerCase());
    });
    data.forEach((user) => {
      let div = document.createElement("div");
      div.id = user.user_id;
      div.setAttribute("nickname", user.nickname);
      div.innerHTML = `<span>User: ${user.nickname}
       Status: ${div.status ? "online" : "offline"}</span>
       <input type="text" id="${user.user_id}-text">
       <button onclick="send(${user.user_id})">Submit</button>
       `;
      container.append(div);
    });
  } catch (err) {
    console.log(err);
  }
};

const send = (id) => {
  let div = document.getElementById(id);
  let message = document.getElementById(`${id}-text`).value;
  div.innerHTML += `<p>You: ${escapeHtml(message)}</p>`;
  socket.send(id + " " + message);
};

socket.onmessage = (event) => {
  let message = event.data.split(/ (.+)/);
  let div = document.getElementById(message[0]);
  if (div) {
    div.innerHTML += `<p>${div.getAttribute("nickname")}: ${escapeHtml(message[1])}</p>`;
  } else {
    console.error(`unkown user`);
  }
};

function escapeHtml(unsafe)
{
    return unsafe
         .replace(/&/g, "&amp;")
         .replace(/</g, "&lt;")
         .replace(/>/g, "&gt;")
         .replace(/"/g, "&quot;")
         .replace(/'/g, "&#039;");
 }