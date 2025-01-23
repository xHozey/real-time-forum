import { getRequest } from "../utils/helpers.js";
import { fillPost } from "../utils/helpers.js";
let offset = 0
export const getPosts = async () => {
  let container = document.querySelector(".post-container");
  const posts = await getRequest("/api/post?" + new URLSearchParams({ offset: offset }));  
  if (posts) {
    posts.forEach((element) => {
      container.append(fillPost(element));
    });
    offset += 10
  }
};