package bot

import "github.com/amaterasutears/bot/models"

func ExtractUserIDFromUpdate(upd *models.Update) (int64, bool) {
	if upd == nil {
		return 0, false
	}

	switch {
	case upd.Message != nil && upd.Message.From != nil:
		return upd.Message.From.ID, true
	case upd.EditedMessage != nil && upd.EditedMessage.From != nil:
		return upd.EditedMessage.From.ID, true
	case upd.ChannelPost != nil && upd.ChannelPost.From != nil:
		return upd.ChannelPost.From.ID, true
	case upd.EditedChannelPost != nil && upd.EditedChannelPost.From != nil:
		return upd.EditedChannelPost.From.ID, true
	case upd.BusinessConnection != nil:
		return upd.BusinessConnection.User.ID, true
	case upd.BusinessMessage != nil && upd.BusinessMessage.From != nil:
		return upd.BusinessMessage.From.ID, true
	case upd.EditedBusinessMessage != nil && upd.EditedBusinessMessage.From != nil:
		return upd.EditedBusinessMessage.From.ID, true
	case upd.MessageReaction != nil && upd.MessageReaction.User != nil:
		return upd.MessageReaction.User.ID, true
	case upd.InlineQuery != nil && upd.InlineQuery.From != nil:
		return upd.InlineQuery.From.ID, true
	case upd.ChosenInlineResult != nil:
		return upd.ChosenInlineResult.From.ID, true
	case upd.CallbackQuery != nil:
		return upd.CallbackQuery.From.ID, true
	case upd.ShippingQuery != nil && upd.ShippingQuery.From != nil:
		return upd.ShippingQuery.From.ID, true
	case upd.PreCheckoutQuery != nil && upd.PreCheckoutQuery.From != nil:
		return upd.PreCheckoutQuery.From.ID, true
	case upd.PurchasedPaidMedia != nil:
		return upd.PurchasedPaidMedia.From.ID, true
	case upd.PollAnswer != nil && upd.PollAnswer.User != nil:
		return upd.PollAnswer.User.ID, true
	case upd.MyChatMember != nil:
		return upd.MyChatMember.From.ID, true
	case upd.ChatMember != nil:
		return upd.ChatMember.From.ID, true
	case upd.ChatJoinRequest != nil:
		return upd.ChatJoinRequest.From.ID, true
	default:
		return 0, false
	}
}

func ExtractChatIDFromUpdate(upd *models.Update) (int64, bool) {
	if upd == nil {
		return 0, false
	}

	switch {
	case upd.Message != nil:
		return upd.Message.Chat.ID, true
	case upd.EditedMessage != nil:
		return upd.EditedMessage.Chat.ID, true
	case upd.ChannelPost != nil:
		return upd.ChannelPost.Chat.ID, true
	case upd.EditedChannelPost != nil:
		return upd.EditedChannelPost.Chat.ID, true
	case upd.MessageReaction != nil:
		return upd.MessageReaction.Chat.ID, true
	case upd.CallbackQuery != nil && upd.CallbackQuery.Message.Message != nil:
		return upd.CallbackQuery.Message.Message.Chat.ID, true
	case upd.BusinessConnection != nil:
		return upd.BusinessConnection.UserChatID, true
	case upd.BusinessMessage != nil:
		return upd.BusinessMessage.Chat.ID, true
	case upd.EditedBusinessMessage != nil:
		return upd.EditedBusinessMessage.Chat.ID, true
	case upd.DeletedBusinessMessages != nil:
		return upd.DeletedBusinessMessages.Chat.ID, true
	case upd.MessageReaction != nil:
		return upd.MessageReaction.Chat.ID, true
	case upd.MessageReactionCount != nil:
		return upd.MessageReactionCount.Chat.ID, true
	case upd.PollAnswer != nil && upd.PollAnswer.VoterChat != nil:
		return upd.PollAnswer.VoterChat.ID, true
	case upd.MyChatMember != nil:
		return upd.MyChatMember.Chat.ID, true
	case upd.ChatMember != nil:
		return upd.ChatMember.Chat.ID, true
	case upd.ChatJoinRequest != nil:
		return upd.ChatJoinRequest.Chat.ID, true
	case upd.ChatBoost != nil:
		return upd.ChatBoost.Chat.ID, true
	case upd.RemovedChatBoost != nil:
		return upd.RemovedChatBoost.Chat.ID, true
	default:
		return 0, false
	}
}

func ExtractUsernameFromUpdate(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	switch {
	case upd.Message != nil && upd.Message.From != nil:
		return upd.Message.From.Username, true
	case upd.EditedMessage != nil && upd.EditedMessage.From != nil:
		return upd.EditedMessage.From.Username, true
	case upd.ChannelPost != nil && upd.ChannelPost.From != nil:
		return upd.ChannelPost.From.Username, true
	case upd.EditedChannelPost != nil && upd.EditedChannelPost.From != nil:
		return upd.EditedChannelPost.From.Username, true
	case upd.BusinessConnection != nil:
		return upd.BusinessConnection.User.Username, true
	case upd.BusinessMessage != nil && upd.BusinessMessage.From != nil:
		return upd.BusinessMessage.From.Username, true
	case upd.EditedBusinessMessage != nil && upd.EditedBusinessMessage.From != nil:
		return upd.EditedBusinessMessage.From.Username, true
	case upd.MessageReaction != nil && upd.MessageReaction.User != nil:
		return upd.MessageReaction.User.Username, true
	case upd.InlineQuery != nil && upd.InlineQuery.From != nil:
		return upd.InlineQuery.From.Username, true
	case upd.ChosenInlineResult != nil:
		return upd.ChosenInlineResult.From.Username, true
	case upd.CallbackQuery != nil:
		return upd.CallbackQuery.From.Username, true
	case upd.ShippingQuery != nil && upd.ShippingQuery.From != nil:
		return upd.ShippingQuery.From.Username, true
	case upd.PreCheckoutQuery != nil && upd.PreCheckoutQuery.From != nil:
		return upd.PreCheckoutQuery.From.Username, true
	case upd.PurchasedPaidMedia != nil:
		return upd.PurchasedPaidMedia.From.Username, true
	case upd.PollAnswer != nil && upd.PollAnswer.User != nil:
		return upd.PollAnswer.User.Username, true
	case upd.MyChatMember != nil:
		return upd.MyChatMember.From.Username, true
	case upd.ChatMember != nil:
		return upd.ChatMember.From.Username, true
	case upd.ChatJoinRequest != nil:
		return upd.ChatJoinRequest.From.Username, true
	default:
		return "", false
	}
}

func ExtractFirstNameFromUpdate(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	switch {
	case upd.Message != nil && upd.Message.From != nil:
		return upd.Message.From.FirstName, true
	case upd.EditedMessage != nil && upd.EditedMessage.From != nil:
		return upd.EditedMessage.From.FirstName, true
	case upd.ChannelPost != nil && upd.ChannelPost.From != nil:
		return upd.ChannelPost.From.FirstName, true
	case upd.EditedChannelPost != nil && upd.EditedChannelPost.From != nil:
		return upd.EditedChannelPost.From.FirstName, true
	case upd.BusinessConnection != nil:
		return upd.BusinessConnection.User.FirstName, true
	case upd.BusinessMessage != nil && upd.BusinessMessage.From != nil:
		return upd.BusinessMessage.From.FirstName, true
	case upd.EditedBusinessMessage != nil && upd.EditedBusinessMessage.From != nil:
		return upd.EditedBusinessMessage.From.FirstName, true
	case upd.MessageReaction != nil && upd.MessageReaction.User != nil:
		return upd.MessageReaction.User.FirstName, true
	case upd.InlineQuery != nil && upd.InlineQuery.From != nil:
		return upd.InlineQuery.From.FirstName, true
	case upd.ChosenInlineResult != nil:
		return upd.ChosenInlineResult.From.FirstName, true
	case upd.CallbackQuery != nil:
		return upd.CallbackQuery.From.FirstName, true
	case upd.ShippingQuery != nil && upd.ShippingQuery.From != nil:
		return upd.ShippingQuery.From.FirstName, true
	case upd.PreCheckoutQuery != nil && upd.PreCheckoutQuery.From != nil:
		return upd.PreCheckoutQuery.From.FirstName, true
	case upd.PurchasedPaidMedia != nil:
		return upd.PurchasedPaidMedia.From.FirstName, true
	case upd.PollAnswer != nil && upd.PollAnswer.User != nil:
		return upd.PollAnswer.User.FirstName, true
	case upd.MyChatMember != nil:
		return upd.MyChatMember.From.FirstName, true
	case upd.ChatMember != nil:
		return upd.ChatMember.From.FirstName, true
	case upd.ChatJoinRequest != nil:
		return upd.ChatJoinRequest.From.FirstName, true
	default:
		return "", false
	}
}

func ExtractTextFromMessage(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	if upd.Message == nil {
		return "", false
	}

	return upd.Message.Text, true
}

func ExtractTextFromEditedMessage(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	if upd.EditedMessage == nil {
		return "", false
	}

	return upd.EditedMessage.Text, true
}

func ExtractTextFromChannelPost(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	if upd.ChannelPost == nil {
		return "", false
	}

	return upd.ChannelPost.Text, true
}

func ExtractTextFromEditedChannelPost(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	if upd.EditedChannelPost == nil {
		return "", false
	}

	return upd.EditedChannelPost.Text, true
}

func ExtractTextFromBusinessMessage(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	if upd.BusinessMessage == nil {
		return "", false
	}

	return upd.BusinessMessage.Text, true
}

func ExtractTextFromEditedBusinessMessage(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	if upd.EditedBusinessMessage == nil {
		return "", false
	}

	return upd.EditedBusinessMessage.Text, true
}

func ExtractTextFromReplyToMessage(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	if upd.Message == nil {
		return "", false
	}

	if upd.Message.ReplyToMessage == nil {
		return "", false
	}

	return upd.Message.ReplyToMessage.Text, true
}

func ExtractIDFromMessage(upd *models.Update) (int, bool) {
	if upd == nil {
		return 0, false
	}

	if upd.Message == nil {
		return 0, false
	}

	return upd.Message.ID, true
}

func ExtractIDFromEditedMessage(upd *models.Update) (int, bool) {
	if upd == nil {
		return 0, false
	}

	if upd.EditedMessage == nil {
		return 0, false
	}

	return upd.EditedMessage.ID, true
}

func ExtractIDFromChannelPost(upd *models.Update) (int, bool) {
	if upd == nil {
		return 0, false
	}

	if upd.ChannelPost == nil {
		return 0, false
	}

	return upd.ChannelPost.ID, true
}

func ExtractIDFromEditedChannelPost(upd *models.Update) (int, bool) {
	if upd == nil {
		return 0, false
	}

	if upd.EditedChannelPost == nil {
		return 0, false
	}

	return upd.EditedChannelPost.ID, true
}

func ExtractIDFromBusinessMessage(upd *models.Update) (int, bool) {
	if upd == nil {
		return 0, false
	}

	if upd.BusinessMessage == nil {
		return 0, false
	}

	return upd.BusinessMessage.ID, true
}

func ExtractIDFromEditedBusinessMessage(upd *models.Update) (int, bool) {
	if upd == nil {
		return 0, false
	}

	if upd.EditedBusinessMessage == nil {
		return 0, false
	}

	return upd.EditedBusinessMessage.ID, true
}

func ExtractMessageIDFromCallbackQuery(upd *models.Update) (int, bool) {
	if upd == nil {
		return 0, false
	}

	if upd.CallbackQuery == nil {
		return 0, false
	}

	if upd.CallbackQuery.Message.Message == nil {
		return 0, false
	}

	return upd.CallbackQuery.Message.Message.ID, true
}

func ExtractDataFromCallbackQuery(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	if upd.CallbackQuery == nil {
		return "", false
	}

	return upd.CallbackQuery.Data, true
}

func ExtractCallbackQueryIDFromUpdate(upd *models.Update) (string, bool) {
	if upd == nil {
		return "", false
	}

	if upd.CallbackQuery == nil {
		return "", false
	}

	return upd.CallbackQuery.ID, true
}
