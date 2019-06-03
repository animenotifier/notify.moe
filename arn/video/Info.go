package video

import (
	"os"
	"strings"
	"time"

	"github.com/akyoto/go-matroska/matroska"
)

// Info includes some general information about the video file.
type Info struct {
	Duration time.Duration `json:"duration"`
	FileSize int64         `json:"fileSize"`
	Video    videoInfo     `json:"video"`
	Audio    audioInfo     `json:"audio"`
}

type videoInfo struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	FPS    float64 `json:"fps"`
	Codec  string  `json:"codec"`
}

type audioInfo struct {
	BitDepth          int     `json:"bitDepth"`
	SamplingFrequency float64 `json:"samplingFrequency"`
	Channels          int     `json:"channels"`
	Codec             string  `json:"codec"`
}

// GetInfo returns the information about the given video file.
func GetInfo(file string) (*Info, error) {
	stat, err := os.Stat(file)

	if err != nil {
		return nil, err
	}

	doc, err := matroska.Decode(file)

	if err != nil {
		return nil, err
	}

	video := doc.Segment.Tracks[0].Entries[0]
	audio := doc.Segment.Tracks[0].Entries[1]
	frameDuration := video.DefaultDuration
	fps := 1.0 / frameDuration.Seconds()
	duration := time.Duration(doc.Segment.Info[0].Duration) * doc.Segment.Info[0].TimecodeScale
	videoCodecID := normalizeCodecID(video.CodecID)
	audioCodecID := normalizeCodecID(audio.CodecID)

	info := &Info{
		Duration: duration,
		FileSize: stat.Size(),

		Video: videoInfo{
			Width:  video.Video.Width,
			Height: video.Video.Height,
			FPS:    fps,
			Codec:  videoCodecID,
		},

		Audio: audioInfo{
			BitDepth:          audio.Audio.BitDepth,
			SamplingFrequency: audio.Audio.SamplingFreq,
			Channels:          audio.Audio.Channels,
			Codec:             audioCodecID,
		},
	}

	return info, nil
}

func normalizeCodecID(codecID string) string {
	if strings.HasPrefix(codecID, "V_") {
		codecID = strings.TrimPrefix(codecID, "V_")
	}

	if strings.HasPrefix(codecID, "A_") {
		codecID = strings.TrimPrefix(codecID, "A_")
	}

	return strings.ToLower(codecID)
}
