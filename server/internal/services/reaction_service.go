package services

import (
	"forum/server/internal/types"
)

func (db *ServiceLayer) CheckReactInput(react types.Reaction) error {
	if (react.Reaction != -1 && react.Reaction != 1) || (react.Thread_type != "post" && react.Thread_type != "comment") || react.Thread_id < 0 || !db.ServiceDB.MiddlewareData.CheckIfThreadExists(react) {
		return types.ErrInvalidReaction
	}
	return nil
}

func (db *ServiceLayer) ReactionService(react types.Reaction, user_id int) error {
	if react.Thread_type == "post" {
		err := db.postReaction(react.Thread_id, user_id, react.Reaction)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := db.commentReaction(react.Thread_id, user_id, react.Reaction)
		if err != nil {
			return err
		}
		return nil
	}
}

func (db *ServiceLayer) postReaction(post_id, user_id, react int) error {
	var exists bool
	exists, err := db.ServiceDB.MiddlewareData.CheckPostReaction(user_id, post_id)
	if err != nil {
		return err
	}
	if !exists {
		err := db.ServiceDB.MiddlewareData.InsertReactPost(user_id, post_id, react)
		if err != nil {
			return err
		}
	} else {
		var like_type int
		like_type, err := db.ServiceDB.MiddlewareData.GetReactionTypePost(user_id, post_id)
		if err != nil {
			return err
		}
		if like_type == react {
			err := db.ServiceDB.MiddlewareData.DeleteReactionPost(user_id, post_id)
			if err != nil {
				return err
			}
		} else {
			err := db.ServiceDB.MiddlewareData.DeleteReactionPost(user_id, post_id)
			if err != nil {
				return err
			}
			err = db.ServiceDB.MiddlewareData.InsertReactPost(user_id, post_id, react)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (db *ServiceLayer) commentReaction(comment_id, user_id, react int) error {
	var exists bool
	exists, err := db.ServiceDB.MiddlewareData.CheckCommentReaction(user_id, comment_id)
	if err != nil {
		return err
	}
	if !exists {
		err := db.ServiceDB.MiddlewareData.InsertReactComment(user_id, comment_id, react)
		if err != nil {
			return err
		}
	} else {
		var isLiked int
		isLiked, err := db.ServiceDB.MiddlewareData.GetReactionTypeComment(user_id, comment_id)
		if err != nil {
			return err
		}
		if isLiked == react {
			err := db.ServiceDB.MiddlewareData.DeleteReactionComment(user_id, comment_id)
			if err != nil {
				return err
			}
		} else {
			err := db.ServiceDB.MiddlewareData.DeleteReactionComment(user_id, comment_id)
			if err != nil {
				return err
			}
			err = db.ServiceDB.MiddlewareData.InsertReactComment(user_id, comment_id, react)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
