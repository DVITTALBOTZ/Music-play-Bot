/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025-2026 TEAMDEV
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/justfortestingnothibghere/TgMusicBot
 */

package ntgcalls

type SegmentPartRequest struct {
	SegmentID     int64
	PartID        int32
	Limit         int32
	Timestamp     int64
	QualityUpdate bool
	ChannelID     int32
	Quality       MediaSegmentQuality
}
