import { sendCommentData, sendPostData } from "../api/send_threads.js";

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
        sendCommentData(commentInput);
      } else {
        commentInput.classList.add("error");
        setTimeout(() => {
          commentInput.classList.remove("error");
        }, 2000);
      }
    });
};


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
        let data = { categories: categories, content: postContent.value }
        await sendPostData(data)
      } else {
        postContent.classList.add("error");
        setTimeout(() => {
          postContent.classList.remove("error");
        }, 2000);
        return;
      }
    });
  };