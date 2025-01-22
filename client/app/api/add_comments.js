import { postRequest, fillComment, targetPost } from "../utils/helpers.js";

export const addComments = () => {
  function hideCommentsPopup() {
    const commentsOverlay = document.getElementById("comments-overlay");
    commentsOverlay.classList.add("hidden");
  }
  document
    .getElementById("close-comments")
    .addEventListener("click", hideCommentsPopup);

  document.getElementById("submit-comment").addEventListener("click", async () => {
    const commentInput = document.getElementById("comment-input");
    const commentText = commentInput.value.trim();
    if (commentText) {
      const commentsContainer = document.querySelector(".comments-container");
      let insertedComment = await postRequest({content: commentText, post_id: targetPost},"/api/addcomment")
      let comment = fillComment(insertedComment)
      commentsContainer.prepend(comment);
      commentInput.value = "";
      commentsContainer.scrollTop = commentsContainer.scrollHeight;
    }
  });
};

