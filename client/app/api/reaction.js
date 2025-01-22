import { postRequest } from "../utils/helpers.js";

export const sendReaction = async (reactionObj, docs) => {
  const { like, dislike, likeCn, dislikeCn } = docs;
  const toggleClasses = (element, addClass, removeClass) => {
    element.classList.add(addClass);
    element.classList.remove(removeClass);
  };
  const updateCounter = (counter, increment = true) => {
    let count = parseInt(counter.innerText, 10) || 0;
    counter.innerText = increment ? count + 1 : Math.max(count - 1, 0);
  };
  const res = await postRequest(reactionObj, "/api/reaction");
  if (res) {
    const { reaction } = reactionObj;
    if (reaction === 1) {
      if (like.classList.contains("not-liked")) {
        if (dislike.classList.contains("disliked")) {
          updateCounter(dislikeCn, false);
        }
        toggleClasses(like, "liked", "not-liked");
        toggleClasses(dislike, "not-disliked", "disliked");
        updateCounter(likeCn, true);
      } else {
        toggleClasses(like, "not-liked", "liked");
        updateCounter(likeCn, false);
      }
    } else {
      if (dislike.classList.contains("not-disliked")) {
        if (like.classList.contains("liked")) {
          updateCounter(likeCn, false);
        }
        toggleClasses(dislike, "disliked", "not-disliked");
        toggleClasses(like, "not-liked", "liked");
        updateCounter(dislikeCn, true);
      } else {
        toggleClasses(dislike, "not-disliked", "disliked");
        updateCounter(dislikeCn, false);
      }
    }
  }
};
