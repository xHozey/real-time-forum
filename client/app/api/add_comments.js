import { postRequest, fillComment, targetPost } from "../utils/helpers.js";
export const addComments = () => {
  const commentInput = document.getElementById("comment-input");
  document.getElementById("close-comments").addEventListener("click", () => {
    const commentsOverlay = document.getElementById("comments-overlay");
    commentsOverlay.classList.add("hidden");
    commentInput.classList.remove("error");
  });

  document
    .getElementById("submit-comment")
    .addEventListener("click", async () => {
      const commentText = commentInput.value.trim();
      if (commentText) {
        const commentsContainer = document.querySelector(".comments-container");
        let insertedComment = await postRequest(
          { content: commentText, post_id: targetPost },
          "/api/addcomment"
        );
        let comment = fillComment(insertedComment);
        commentsContainer.prepend(comment);
        commentInput.value = "";
        commentsContainer.scrollTop = commentsContainer.scrollHeight;
      } else {
        commentInput.classList.add("error");
        setTimeout(() => {
          commentInput.classList.remove("error");
        },2000);
      }
    });
};
