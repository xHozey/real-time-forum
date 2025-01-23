import { fillPost, postRequest } from "../utils/helpers.js";
export const showPostPanel = () => {
  const openAddPostInput = document.getElementById("open-add-post");
  const popupOverlay = document.getElementById("popup-overlay");
  const postContent = document.getElementById("postContent");
  const characterCounter = document.querySelector(".character-counter");
  const closeAddPostButton = document.getElementById("close-add-post");
  const postButton = document.getElementById("Posting");

  openAddPostInput.addEventListener("click", () => {
    popupOverlay.classList.remove("hidden");
  });

  closeAddPostButton.addEventListener("click", () => {
    popupOverlay.classList.add("hidden");
    postContent.classList.remove("error");
  });

  document.querySelectorAll(".checkbox").forEach((checkbox) => {
    checkbox.addEventListener("click", () => {
      checkbox.classList.toggle("selected");
    });
  });

  postContent.addEventListener("input", () => {
    const length = postContent.value.length;
    characterCounter.textContent = `${length}/500`;
  });

  postButton.addEventListener("click", async () => {
    const postText = postContent.value.trim();
    if (postText) {
      let categories = [];
      document.querySelectorAll(".selected").forEach((categorie) => {
        categories.push(categorie.innerText);
      });
      let post = await postRequest(
        { categories: categories, content: postContent.value },
        "/api/addpost"
      );
      document.querySelector(".post-container").prepend(fillPost(post));
      postContent.value = "";
      characterCounter.textContent = "0/500";
      popupOverlay.classList.add("hidden");
    } else {
      postContent.classList.add("error");
      setTimeout(() => {
        postContent.classList.remove("error");
      }, 2000);
      return;
    }
  });
};
