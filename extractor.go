package bot

import "github.com/amaterasutears/bot/models"

func ExtractUserIDFromUpdate(upd *models.Update) (int64, bool) {
	if upd != nil {
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
		}
	}

	return 0, false
}

func ExtractChatIDFromUpdate(upd *models.Update) (int64, bool) {
	if upd != nil {
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
		}
	}

	return 0, false
}
