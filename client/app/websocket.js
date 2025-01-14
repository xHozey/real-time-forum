const socket = new WebSocket("ws://localhost:8080/ws");

const container = document.getElementById("userList");
const messagesContainer = document.getElementById("messages");
const messageInput = document.getElementById("messageInput");
const sendMessageBtn = document.getElementById("sendMessageBtn");
let UserID = 0;
let display = 10;
let targetId = 0;
socket.onopen = async () => {
  try {
    let = "";
    let res = await fetch("/api/info");
    let data = await res.json();
    UserID = data.UserID;
    data.Clients.sort((a, b) => {
      return a.nickname.toLowerCase().localeCompare(b.nickname.toLowerCase());
    });
    data.Clients.forEach((user) => {
      let div = document.createElement("div");
      div.classList.add("user");
      div.id = `user-${user.user_id}`;
      div.innerHTML = `<span>${user.nickname}</span>`;
      div.setAttribute("onclick", `changeTarget(${user.user_id})`);
      if (!user.status) {
        div.classList.add("offline");
      } else {
        div.classList.add("online");
      }
      container.append(div);
    });
  } catch (err) {
    console.log(err);
  }
};

const changeTarget = async (id) => {
  if (targetId == id) return;
  messagesContainer.innerHTML = "";
  targetId = id;
  try {
    let res = await fetch(`/api/client/${id}`);
    let data = await res.json();
    data.forEach((msg) => {
      messagesContainer.innerHTML += `<p>${msg.receiver}: ${msg.content}`
    })
  } catch (err) {
    console.error(err);
  }
};

sendMessageBtn.addEventListener("click", () => {
  messagesContainer.innerHTML += `<p>You: ${escapeHtml(
    messageInput.value
  )}</p>`;
  socket.send(targetId + " " + messageInput.value);
});

socket.onmessage = (event) => {
  try {
    let res = JSON.parse(event.data);
    let div = document.getElementById(`user-${res.id}`);
    div.classList.remove("offline", "online");
    if (res.status) {
      div.classList.add("online");
    } else {
      div.classList.add("offline");
    }
  } catch (err) {
    let message = event.data.split(/ (.+)/);
    if (targetId == message[0]) {
      messagesContainer.innerHTML += `<p>${message[1]}</p>`;
    }
  }
};

function escapeHtml(unsafe) {
  return unsafe
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#039;");
}
