import { getRequest } from "../utils/helpers.js";
import { fillPost } from "../utils/helpers.js";
export const getPosts = async () => {
  let container = document.querySelector(".post-container");
  const posts = await getRequest("/api/post");

  if (posts) {
    posts.forEach((element) => {
      container.append(fillPost(element));
    });
  }
};

