import { loadMessages, targetId } from "../api/users.js";
import { sendReaction } from "../api/reaction.js";
import { getComment } from "../api/get_comments.js";
export let targetPost
export const escapeHtml = (unsafe) => {
  return unsafe
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#039;");
};

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
  messagesContainer.scroll(0, messagesContainer.scrollHeight);
};

export const prepandMessage = (message, name) => {
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
  messagesContainer.prepend(div);
};

export const createUser = (user) => {
  const usersList = document.querySelector(".users-list");
  const div = document.createElement("div");
  div.classList.add("user-item");
  div.id = user.id;
  div.setAttribute("onclick", `target(${user.id})`);
  div.setAttribute("data-nickname", `${user.nickname}`);
  div.innerHTML = `${user.nickname} <span id="status-${user.id}" class="${
    user.status ? "online" : "offline"
  }"></span>`;
  usersList.append(div);
};

export const postRequest = async (data, url) => {
  try {
    const response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    if (!response.ok) {
      console.error(response.status);
      return;
    }
    try {
      let insertedData = await response.json();
      return insertedData;
    } catch (err) {
      return true;
    }
  } catch (error) {
    console.error("Error:", error);
  }
};

export const getRequest = async (url) => {
  try {
    const res = await fetch(url);
    const data = await res.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
  }
};

export const throttle = (element) => {
  let timer;
  document
    .querySelector(".messages-container")
    .addEventListener("scroll", (event) => {
      const container = document.querySelector(".messages-container");
      clearTimeout(timer);
      if (container.scrollTop == 0) {
        timer = setTimeout(() => {
          loadMessages(container, targetId);
        }, 1000);
      }
    });
};

export const fillPost = (postInfo) => {
  const post = document.createElement("div");
  const categories = document.createElement("div");
  post.classList.add("post");

  if (postInfo.categories) {
    postInfo.categories.forEach((categorie) => {
      let span = document.createElement("span");
      span.innerText = categorie;
      span.classList.add("post-categorie");
      categories.appendChild(span);
    });
  }
  post.innerHTML = `<div class="post-header">
        <div class="profile-img icons"></div>
        <span class="username">${postInfo.author}</span>
        <span class="post-time">${new Date(
          postInfo.creation_date
        ).toLocaleDateString("en-US")}</span>
      </div>
      <div class="categories">
      ${categories.innerHTML ? categories.innerHTML : ""}
      </div>
      <p class="post-text">${postInfo.content}</p>
      <div class="post-actions">
          <button id="like-${postInfo.id}" class="icons actionIcon"></button>
          <span id="like-counter" class="nbr">${postInfo.likes}</span>
          <button id="dislike-${postInfo.id}" class="icons actionIcon"></button>
          <span id="dislike-counter" class="nbr">${postInfo.dislikes}</span>
          <button class="comment-icon icons actionIcon"></button>
      </div>`;

  const likeButton = post.querySelector(`#like-${postInfo.id}`);
  const dislikeButton = post.querySelector(`#dislike-${postInfo.id}`);
  const likeCounter = post.querySelector("#like-counter");
  const dislikeCounter = post.querySelector("#dislike-counter");
  post.querySelector(`.comment-icon`).onclick = async () => {
    targetPost = postInfo.id
    await getComment(postInfo.id)
    document.getElementById("comments-overlay").classList.remove("hidden");
  };
  switch (postInfo.isliked) {
    case 1:
      likeButton.classList.add("liked");
      dislikeButton.classList.add("not-disliked");
      break;
    case -1:
      likeButton.classList.add("not-liked");
      dislikeButton.classList.add("disliked");
    default:
      likeButton.classList.add("not-liked");
      dislikeButton.classList.add("not-disliked");
      break;
  }

  likeButton.onclick = () =>
    sendReaction(
      {
        thread_id: postInfo.id,
        thread_type: "post",
        reaction: 1,
      },
      {
        like: likeButton,
        dislike: dislikeButton,
        likeCn: likeCounter,
        dislikeCn: dislikeCounter,
      }
    );
  dislikeButton.onclick = () =>
    sendReaction(
      {
        thread_id: postInfo.id,
        thread_type: "post",
        reaction: -1,
      },
      {
        like: likeButton,
        dislike: dislikeButton,
        likeCn: likeCounter,
        dislikeCn: dislikeCounter,
      }
    );

  return post;
};

export const fillComment = (commentInfo) => {
  const comment = document.createElement('div')
   comment.innerHTML = `<div class="comment-post">
  <div class="comment-author">${commentInfo.author}</div>
  <div class="comment-text">${commentInfo.content}</div>
  <div class="comment-footer">
      <span class="comment-datetime">${new Date(commentInfo.creation_date).toLocaleDateString(
        "en-US"
      )}</span>
      <div class="comment-actions">
          <button id="like-${
            commentInfo.id
          }" class="icons actionIcon"></button>
        <span id="like-counter" class="nbr">${commentInfo.likes}</span>
        <button id="dislike-${
          commentInfo.id
        }" class="icons actionIcon"></button>
        <span id="dislike-counter" class="nbr">${commentInfo.dislikes}</span>
      </div>
  </div>
</div>
      </div>`;
  const likeButton = comment.querySelector(`#like-${commentInfo.id}`);
  const dislikeButton = comment.querySelector(`#dislike-${commentInfo.id}`);
  const likeCounter = comment.querySelector("#like-counter");
  const dislikeCounter = comment.querySelector("#dislike-counter");
  switch (commentInfo.isliked) {
    case 1:
      likeButton.classList.add("liked");
      dislikeButton.classList.add("not-disliked");
      break;
    case -1:
      likeButton.classList.add("not-liked");
      dislikeButton.classList.add("disliked");
    default:
      likeButton.classList.add("not-liked");
      dislikeButton.classList.add("not-disliked");
      break;
  }

  likeButton.onclick = () =>
    sendReaction(
      {
        thread_id: commentInfo.id,
        thread_type: "comment",
        reaction: 1,
      },
      {
        like: likeButton,
        dislike: dislikeButton,
        likeCn: likeCounter,
        dislikeCn: dislikeCounter,
      }
    );
  dislikeButton.onclick = () =>
    sendReaction(
      {
        thread_id: commentInfo.id,
        thread_type: "comment",
        reaction: -1,
      },
      {
        like: likeButton,
        dislike: dislikeButton,
        likeCn: likeCounter,
        dislikeCn: dislikeCounter,
      }
    );
  return comment;
};
