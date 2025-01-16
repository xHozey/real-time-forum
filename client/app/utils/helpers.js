export const escapeHtml = (unsafe) => {
    return unsafe
      .replace(/&/g, "&amp;")
      .replace(/</g, "&lt;")
      .replace(/>/g, "&gt;")
      .replace(/"/g, "&quot;")
      .replace(/'/g, "&#039;");
  }

  export const sendMessage = (message, name) => {
    let messagesContainer = document.querySelector(".messages-container");
        let div = document.createElement("div");
        div.classList.add("message");
        let author = document.createElement("div");
        author.classList.add("message-author");
        author.innerText = name;
        div.append(author);
        let content = document.createElement("div");
        content.classList.add("message-content");
        content.innerText = message;
        div.append(content);
        messagesContainer.append(div);
  }