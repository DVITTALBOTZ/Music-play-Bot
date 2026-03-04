/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025-2026 TEAMDEV
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/justfortestingnothibghere/TgMusicBot
 */

package types

import "ashokshau/tgmusic/src/vc/ntgcalls"

type PendingConnection struct {
	MediaDescription ntgcalls.MediaDescription
	Payload          string
	Presentation     bool
}
