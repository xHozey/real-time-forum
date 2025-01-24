import {
  postRequest,
  fillComment,
  targetPost,
  fillPost,
} from "../utils/helpers.js";

export const sendCommentData = async (element) => {
  const commentsContainer = document.querySelector(".comments-container");
  let insertedComment = await postRequest(
    { content: element.value, post_id: targetPost },
    "/api/addcomment"
  );

  let comment = fillComment(insertedComment);
  commentsContainer.prepend(comment);
  element.value = "";
  commentsContainer.scrollTop = 0;
};

export const sendPostData = async (data) => {
  const popupOverlay = document.getElementById("popup-overlay");
  const postContent = document.getElementById("postContent");
  const characterCounter = document.querySelector(".character-counter");
  const postContainer = document.querySelector(".post-container")
  let post = await postRequest(data, "/api/addpost");
  postContainer.prepend(fillPost(post));
  postContent.value = "";
  characterCounter.textContent = "0/500";
  popupOverlay.classList.add("hidden");
  postContainer.scrollTop = 0
};
