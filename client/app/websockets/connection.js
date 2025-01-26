import {
  handleIncommingMessage,
  sendPrivateMessage,
  typingStatus,
} from "./handle_connection.js";
import { nav } from "../utils/navigation.js";

export const connectToServer = () => {
  const conn = new WebSocket("ws://localhost:1999/ws");
  const sendMessageBtn = document.getElementById("btn-message");
  const logoutBtn = document.getElementById("logout-button");
  conn.onopen = () => {
    console.log("Connected");
  };

  conn.onmessage = (event) => {
    let res = JSON.parse(event.data);
    handleIncommingMessage(res);
  };

  conn.onerror = (error) => {
    console.error("WebSocket error:", error);
  };

  conn.onclose = (event) => {
    console.log("WebSocket connection closed:", event.code, event.reason);
  };

  sendMessageBtn.addEventListener("click", () => {
    sendPrivateMessage(conn);
  });

  typingStatus(conn);

  logoutBtn.addEventListener("click", async () => {
    conn.close(1000, "logout");
    await fetch("/logout");
    nav("/login");
  });
};
