import { getRequest, fillComment } from "../utils/helpers.js";

export const getComment = async (postId) => {
  const commentsContainer = document.querySelector(".comments-container");
  const comments = await getRequest(`/api/post/${postId}/comment`);
  commentsContainer.innerHTML = ""
  if (comments) {
    comments.forEach((cmt) => {
      commentsContainer.prepend(fillComment(cmt));
    });
  }
};
