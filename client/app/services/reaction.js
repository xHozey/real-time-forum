const updateCounter = (counter, increment = true) => {
  let count = parseInt(counter.innerText, 10) || 0;
  counter.innerText = increment ? count + 1 : Math.max(count - 1, 0);
};

const toggleClasses = (element, addClass, removeClass) => {
  element.classList.add(addClass);
  element.classList.remove(removeClass);
};

export const likeReaction = (docs) => {
  const { like, dislike, likeCn, dislikeCn } = docs;

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
};

export const dislikeReaction = (docs) => {
  const { like, dislike, likeCn, dislikeCn } = docs;

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
};
