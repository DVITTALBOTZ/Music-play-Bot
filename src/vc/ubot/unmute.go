/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025-2026 TEAMDEV
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/justfortestingnothibghere/TgMusicBot
 */

package ubot

func (ctx *Context) Unmute(chatId int64) (bool, error) {
	return ctx.binding.UnMute(chatId)
}
