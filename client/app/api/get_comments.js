import { getRequest, fillComment } from "../utils/helpers.js";
export const getComment = async (postId, commentsOffset) => {
  const comments = await getRequest(`/api/post/${postId}/comment?` + new URLSearchParams({ offset: commentsOffset }));
  if (comments) {
    const commentsContainer = document.querySelector(".comments-container");
    comments.forEach((cmt) => {
      commentsContainer.append(fillComment(cmt));
    });
    commentsOffset += 10
  }
  return commentsOffset
};
