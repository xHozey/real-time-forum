import { getRequest, fillComment, fillPost } from "../utils/helpers.js";

export const getComment = async (postId, commentsOffset) => {
  const comments = await getRequest(
    `/api/post/${postId}/comment?` +
      new URLSearchParams({ offset: commentsOffset })
  );
  if (comments) {
    const commentsContainer = document.querySelector(".comments-container");
    comments.forEach((cmt) => {
      commentsContainer.append(fillComment(cmt));
    });
    commentsOffset += 10;
  }
  return commentsOffset;
};

let offset = 0;
export const getPosts = async () => {
  let container = document.querySelector(".post-container");
  const posts = await getRequest(
    "/api/post?" + new URLSearchParams({ offset: offset })
  );
  if (posts) {
    posts.forEach((element) => {
      container.append(fillPost(element));
    });
    offset += 10;
  }
};
