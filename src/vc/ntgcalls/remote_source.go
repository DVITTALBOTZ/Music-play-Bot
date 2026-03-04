/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025-2026 TEAMDEV
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/justfortestingnothibghere/TgMusicBot
 */

package ntgcalls

type RemoteSource struct {
	Ssrc   uint32
	State  StreamStatus
	Device StreamDevice
}
