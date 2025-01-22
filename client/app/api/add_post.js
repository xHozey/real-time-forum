import { fillPost, postRequest } from "../utils/helpers.js";
export const showPostPanel = () => {
  const openAddPostInput = document.getElementById("open-add-post");
  const popupOverlay = document.getElementById("popup-overlay");

  openAddPostInput.addEventListener("click", () => {
    popupOverlay.classList.remove("hidden");
  });

  const closeAddPostButton = document.getElementById("close-add-post");

  closeAddPostButton.addEventListener("click", () => {
    popupOverlay.classList.add("hidden");
  });

  popupOverlay.addEventListener("click", (event) => {
    if (event.target === popupOverlay) {
      popupOverlay.classList.add("hidden");
    }
  });

  document.querySelectorAll(".checkbox").forEach((checkbox) => {
    checkbox.addEventListener("click", () => {
      checkbox.classList.toggle("selected");
    });
  });

  const postContent = document.getElementById("postContent");
  const characterCounter = document.querySelector(".character-counter");

  postContent.addEventListener("input", () => {
    const length = postContent.value.length;
    characterCounter.textContent = `${length}/500`;
  });

  const postButton = document.getElementById("Posting");

  postButton.addEventListener("click",async () => {
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
      document.querySelector(".post-container").prepend(fillPost(post))
      postContent.value = "";
      characterCounter.textContent = "0/500";
      popupOverlay.classList.add("hidden");
    } else {
      alert("Please write something before posting.");
      return;
    }
  });
};
