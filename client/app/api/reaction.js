import { postRequest } from "../utils/helpers.js";
import { likeReaction, dislikeReaction } from "../services/reaction.js";
export const sendReaction = async (reactionObj, docs) => {
  const res = await postRequest(reactionObj, "/api/reaction");
  if (res) {
    const { reaction } = reactionObj;
    if (reaction === 1) {
      likeReaction(docs);
    } else {
      dislikeReaction(docs);
    }
  }
};
