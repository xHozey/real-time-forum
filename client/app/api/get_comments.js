import { getRequest } from "../utils/helpers.js"

export const getComment = async (postId) => {
   const comments = await getRequest(`/api/post/${postId}/comment`)
   comments.forEach(cmt => {
        
   });
}