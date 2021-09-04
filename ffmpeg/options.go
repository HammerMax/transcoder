package ffmpeg

import (
	"fmt"
	"reflect"
)

// Options defines allowed FFmpeg arguments
type Options struct {
	Aspect                *string           `flag:"-aspect"`
	Resolution            *string           `flag:"-s"`
	VideoBitRate          *string           `flag:"-b:v"`
	VideoBitRateTolerance *int              `flag:"-bt"`
	VideoMaxBitRate       *int              `flag:"-maxrate"`
	VideoMinBitrate       *int              `flag:"-minrate"`
	VideoCodec            *string           `flag:"-c:v"`
	Vframes               *int              `flag:"-vframes"`
	FrameRate             *int              `flag:"-r"`
	AudioRate             *int              `flag:"-ar"`
	KeyframeInterval      *int              `flag:"-g"`
	AudioCodec            *string           `flag:"-c:a"`
	AudioBitrate          *string           `flag:"-ab"`
	AudioChannels         *int              `flag:"-ac"`
	AudioVariableBitrate  *bool             `flag:"-q:a"`
	BufferSize            *int              `flag:"-bufsize"`
	Threadset             *bool             `flag:"-threads"`
	Threads               *int              `flag:"-threads"`
	Preset                *string           `flag:"-preset"`
	Tune                  *string           `flag:"-tune"`
	AudioProfile          *string           `flag:"-profile:a"`
	VideoProfile          *string           `flag:"-profile:v"`
	Target                *string           `flag:"-target"`
	Duration              *string           `flag:"-t"`
	Qscale                *int              `flag:"-qscale"`
	Crf                   *int              `flag:"-crf"`
	Strict                *int              `flag:"-strict"`
	MuxDelay              *string           `flag:"-muxdelay"`
	SeekTime              *string           `flag:"-ss"`
	SeekUsingTimestamp    *bool             `flag:"-seek_timestamp"`
	MovFlags              *string           `flag:"-movflags"`
	HideBanner            *bool             `flag:"-hide_banner"`
	OutputFormat          *string           `flag:"-f"`
	CopyTs                *bool             `flag:"-copyts"`
	NativeFramerateInput  *bool             `flag:"-re"`
	InputInitialOffset    *string           `flag:"-itsoffset"`
	RtmpLive              *string           `flag:"-rtmp_live"`
	HlsPlaylistType       *string           `flag:"-hls_playlist_type"`
	HlsListSize           *int              `flag:"-hls_list_size"`
	HlsSegmentDuration    *int              `flag:"-hls_time"`
	HlsMasterPlaylistName *string           `flag:"-master_pl_name"`
	HlsSegmentFilename    *string           `flag:"-hls_segment_filename"`
	HTTPMethod            *string           `flag:"-method"`
	HTTPKeepAlive         *bool             `flag:"-multiple_requests"`
	Hwaccel               *string           `flag:"-hwaccel"`
	StreamIds             map[string]string `flag:"-streamid"`
	VideoFilter           *string           `flag:"-vf"`
	AudioFilter           *string           `flag:"-af"`
	SkipVideo             *bool             `flag:"-vn"`
	SkipAudio             *bool             `flag:"-an"`
	CompressionLevel      *int              `flag:"-compression_level"`
	MapMetadata           *string           `flag:"-map_metadata"`
	Metadata              map[string]string `flag:"-metadata"`
	EncryptionKey         *string           `flag:"-hls_key_info_file"`
	Bframe                *int              `flag:"-bf"`
	PixFmt                *string           `flag:"-pix_fmt"`
	WhiteListProtocols    []string          `flag:"-protocol_whitelist"`
	Overwrite             *bool             `flag:"-y"`
	FilterComplex         *string           `flag:"-filter_complex"`
	Map         *string           `flag:"-map"`
	ExtraArgs             map[string]interface{}
}

func NewOptions() *Options {
	return &Options{}
}

// GetStrArguments ...
func (opts *Options) GetStrArguments() []string {
	f := reflect.TypeOf(*opts)
	v := reflect.ValueOf(*opts)

	values := []string{}

	for i := 0; i < f.NumField(); i++ {
		flag := f.Field(i).Tag.Get("flag")
		value := v.Field(i).Interface()

		if !v.Field(i).IsNil() {

			if _, ok := value.(*bool); ok {
				values = append(values, flag)
			}

			if vs, ok := value.(*string); ok {
				values = append(values, flag, *vs)
			}

			if va, ok := value.([]string); ok {

				for i := 0; i < len(va); i++ {
					item := va[i]
					values = append(values, flag, item)
				}
			}

			if vm, ok := value.(map[string]interface{}); ok {
				for k, v := range vm {
					values = append(values, k, fmt.Sprintf("%v", v))
				}
			}

			if vi, ok := value.(*int); ok {
				values = append(values, flag, fmt.Sprintf("%d", *vi))
			}

		}
	}

	return values
}

func (opts *Options) SetAspect(aspect string) *Options {
	opts.Aspect = &aspect
	return opts
}

func (opts *Options) SetResolution(resolution string) *Options {
	opts.Resolution = &resolution
	return opts
}

func (opts *Options) SetVideoBitRate(videoBitRate string) *Options {
	opts.VideoBitRate = &videoBitRate
	return opts
}

func (opts *Options) SetVideoBitRateTolerance(videoBitRateTolerance int) *Options {
	opts.VideoBitRateTolerance = &videoBitRateTolerance
	return opts
}
func (opts *Options) SetVideoMaxBitRate(videoMaxBitRate int) *Options {
	opts.VideoMaxBitRate = &videoMaxBitRate
	return opts
}
func (opts *Options) SetVideoMinBitrate(videoMinBitrate int) *Options {
	opts.VideoMinBitrate = &videoMinBitrate
	return opts

}
func (opts *Options) SetVideoCodec(videoCodec string) *Options {
	opts.VideoCodec = &videoCodec
	return opts

}
func (opts *Options) SetVframes(vframes int) *Options {
	opts.Vframes = &vframes
	return opts

}
func (opts *Options) SetFrameRate(frameRate int) *Options {
	opts.FrameRate = &frameRate
	return opts

}
func (opts *Options) SetAudioRate(audioRate int) *Options {
	opts.AudioRate = &audioRate
	return opts

}
func (opts *Options) SetKeyframeInterval(keyframeInterval int) *Options {
	opts.KeyframeInterval = &keyframeInterval
	return opts

}
func (opts *Options) SetAudioCodec(audioCodec string) *Options {
	opts.AudioCodec = &audioCodec
	return opts

}
func (opts *Options) SetAudioBitrate(audioBitrate string) *Options {
	opts.AudioBitrate = &audioBitrate
	return opts

}
func (opts *Options) SetAudioChannels(audioChannels int) *Options {
	opts.AudioChannels = &audioChannels
	return opts
}
func (opts *Options) SetAudioVariableBitrate(audioVariableBitrate bool) *Options {
	opts.AudioVariableBitrate = &audioVariableBitrate
	return opts

}
func (opts *Options) SetBufferSize(bufferSize int) *Options {
	opts.BufferSize = &bufferSize
	return opts

}
func (opts *Options) SetThreadset(threadset bool) *Options {
	opts.Threadset = &threadset
	return opts

}
func (opts *Options) SetThreads(Threads int) *Options {
	opts.Threads = &Threads
	return opts

}
func (opts *Options) SetPreset(Preset string) *Options {
	opts.Preset = &Preset
	return opts

}
func (opts *Options) SetTune(Tune string) *Options {
	opts.Tune = &Tune
	return opts

}
func (opts *Options) SetAudioProfile(audioProfile string) *Options {
	opts.AudioProfile = &audioProfile
	return opts

}
func (opts *Options) SetVideoProfile(videoProfile string) *Options {
	opts.VideoProfile = &videoProfile
	return opts

}
func (opts *Options) SetTarget(target string) *Options {
	opts.Target = &target
	return opts

}
func (opts *Options) SetDuration(duration string) *Options {
	opts.Duration = &duration
	return opts

}
func (opts *Options) SetQscale(qscale int) *Options {
	opts.Qscale = &qscale
	return opts

}
func (opts *Options) SetCrf(crf int) *Options {
	opts.Crf = &crf
	return opts

}
func (opts *Options) SetStrict(strict int) *Options {
	opts.Strict = &strict
	return opts

}
func (opts *Options) SetMuxDelay(muxDelay string) *Options {
	opts.MuxDelay = &muxDelay
	return opts

}
func (opts *Options) SetSeekTime(ss string) *Options {
	opts.SeekTime = &ss
	return opts
}
func (opts *Options) SetSeekUsingTimestamp(seekUsingTimestamp bool) *Options {
	opts.SeekUsingTimestamp = &seekUsingTimestamp
	return opts
}
func (opts *Options) SetMovFlags(movFlags string) *Options {
	opts.MovFlags = &movFlags
	return opts

}
func (opts *Options) SetHideBanner(hideBanner bool) *Options {
	opts.HideBanner = &hideBanner
	return opts

}
func (opts *Options) SetOutputFormat(outputFormat string) *Options {
	opts.OutputFormat = &outputFormat
	return opts

}
func (opts *Options) SetCopyTs(copyTs bool) *Options {
	opts.CopyTs = &copyTs
	return opts

}
func (opts *Options) SetNativeFramerateInput(nativeFramerateInput bool) *Options {
	opts.NativeFramerateInput = &nativeFramerateInput
	return opts

}
func (opts *Options) SetInputInitialOffset(itsOffset string) *Options {
	opts.InputInitialOffset = &itsOffset
	return opts
}
func (opts *Options) SetRtmpLive(rtmpLive string) *Options {
	opts.RtmpLive = &rtmpLive
	return opts
}
func (opts *Options) SetHlsPlaylistType(hlsPlaylistType string) *Options {
	opts.HlsPlaylistType = &hlsPlaylistType
	return opts
}
func (opts *Options) SetHlsListSize(hlsListSize int) *Options {
	opts.HlsListSize = &hlsListSize
	return opts
}
func (opts *Options) SetHlsSegmentDuration(hlsSegmentDuration int) *Options {
	opts.HlsSegmentDuration = &hlsSegmentDuration
	return opts
}
func (opts *Options) SetHlsMasterPlaylistName(hlsMasterPlaylistName string) *Options {
	opts.HlsMasterPlaylistName = &hlsMasterPlaylistName
	return opts
}
func (opts *Options) SetHlsSegmentFilename(hlsSegmentFilename string) *Options {
	opts.HlsSegmentFilename = &hlsSegmentFilename
	return opts
}
func (opts *Options) SetHTTPMethod(hTTPMethod string) *Options {
	opts.HTTPMethod = &hTTPMethod
	return opts
}
func (opts *Options) SetHTTPKeepAlive(hTTPKeepAlive bool) *Options {
	opts.HTTPKeepAlive = &hTTPKeepAlive
	return opts
}
func (opts *Options) SetHwaccel(hwaccel string) *Options {
	opts.Hwaccel = &hwaccel
	return opts
}
func (opts *Options) SetStreamIds(streamIds map[string]string) *Options {
	opts.StreamIds = streamIds
	return opts
}
func (opts *Options) SetVideoFilter(videoFilter string) *Options {
	opts.VideoFilter = &videoFilter
	return opts
}
func (opts *Options) SetAudioFilter(audioFilter string) *Options {
	opts.AudioFilter = &audioFilter
	return opts
}
func (opts *Options) SetSkipVideo(skipVideo bool) *Options {
	opts.SkipVideo = &skipVideo
	return opts
}
func (opts *Options) SetSkipAudio(skipAudio bool) *Options {
	opts.SkipAudio = &skipAudio
	return opts
}
func (opts *Options) SetCompressionLevel(compressionLevel int) *Options {
	opts.CompressionLevel = &compressionLevel
	return opts
}
func (opts *Options) SetMapMetadata(mapMetadata string) *Options {
	opts.MapMetadata = &mapMetadata
	return opts
}
func (opts *Options) SetMetadata(metadata map[string]string) *Options {
	opts.Metadata = metadata
	return opts
}
func (opts *Options) SetEncryptionKey(encryptionKey string) *Options {
	opts.EncryptionKey = &encryptionKey
	return opts
}
func (opts *Options) SetBframe(bframe int) *Options {
	opts.Bframe = &bframe
	return opts
}
func (opts *Options) SetPixFmt(pixFmt string) *Options {
	opts.PixFmt = &pixFmt
	return opts

}
func (opts *Options) SetWhiteListProtocols(whiteListProtocols []string) *Options {
	opts.WhiteListProtocols = whiteListProtocols
	return opts

}
func (opts *Options) SetOverwrite(overwrite bool) *Options {
	opts.Overwrite = &overwrite
	return opts

}
func (opts *Options) SetExtraArgs(extra map[string]interface{}) *Options {
	opts.ExtraArgs = extra
	return opts

}

func (opts *Options) SetFilterComplex(filterComplex string) *Options {
	opts.FilterComplex = &filterComplex
	return opts
}

func (opts *Options) SetMap(strMap string) *Options {
	opts.Map = &strMap
	return opts
}