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
	Qscale                *uint32           `flag:"-qscale"`
	Crf                   *uint32           `flag:"-crf"`
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
	FilterComplex                *string           `flag:"-filter_complex"`
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

func (opts *Options) SetAspect() {

}
func (opts *Options) SetResolution() {

}
func (opts *Options) SetVideoBitRate() {

}
func (opts *Options) SetVideoBitRateTolerance() {

}
func (opts *Options) SetVideoMaxBitRate() {

}
func (opts *Options) SetVideoMinBitrate() {

}
func (opts *Options) SetVideoCodec() {

}
func (opts *Options) SetVframes() {

}
func (opts *Options) SetFrameRate() {

}
func (opts *Options) SetAudioRate() {

}
func (opts *Options) SetKeyframeInterval() {

}
func (opts *Options) SetAudioCodec() {

}
func (opts *Options) SetAudioBitrate() {

}
func (opts *Options) SetAudioChannels() {

}
func (opts *Options) SetAudioVariableBitrate() {

}
func (opts *Options) SetBufferSize() {

}
func (opts *Options) SetThreadset() {

}
func (opts *Options) SetThreads() {

}
func (opts *Options) SetPreset() {

}
func (opts *Options) SetTune() {

}
func (opts *Options) SetAudioProfile() {

}
func (opts *Options) SetVideoProfile() {

}
func (opts *Options) SetTarget() {

}
func (opts *Options) SetDuration() {

}
func (opts *Options) SetQscale() {

}
func (opts *Options) SetCrf() {

}
func (opts *Options) SetStrict() {

}
func (opts *Options) SetMuxDelay() {

}
func (opts *Options) SetSeekTime(ss string) *Options {
	opts.SeekTime = &ss
	return opts
}
func (opts *Options) SetSeekUsingTimestamp() {

}
func (opts *Options) SetMovFlags() {

}
func (opts *Options) SetHideBanner() {

}
func (opts *Options) SetOutputFormat() {

}
func (opts *Options) SetCopyTs() {

}
func (opts *Options) SetNativeFramerateInput() {

}
func (opts *Options) SetInputInitialOffset(itsOffset string) *Options {
	opts.InputInitialOffset = &itsOffset
	return opts
}
func (opts *Options) SetRtmpLive() {

}
func (opts *Options) SetHlsPlaylistType() {

}
func (opts *Options) SetHlsListSize() {

}
func (opts *Options) SetHlsSegmentDuration() {

}
func (opts *Options) SetHlsMasterPlaylistName() {

}
func (opts *Options) SetHlsSegmentFilename() {

}
func (opts *Options) SetHTTPMethod() {

}
func (opts *Options) SetHTTPKeepAlive() {

}
func (opts *Options) SetHwaccel() {

}
func (opts *Options) SetStreamIds() {

}
func (opts *Options) SetVideoFilter() {

}
func (opts *Options) SetAudioFilter() {

}
func (opts *Options) SetSkipVideo() {

}
func (opts *Options) SetSkipAudio() {

}
func (opts *Options) SetCompressionLevel() {

}
func (opts *Options) SetMapMetadata() {

}
func (opts *Options) SetMetadata() {

}
func (opts *Options) SetEncryptionKey() {

}
func (opts *Options) SetBframe() {

}
func (opts *Options) SetPixFmt() {

}
func (opts *Options) SetWhiteListProtocols() {

}
func (opts *Options) SetOverwrite() {

}
func (opts *Options) SetExtraArgs() {

}

func (opts *Options) SetFilterComplex(filterComplex string) *Options {
	opts.FilterComplex = &filterComplex
	return opts
}
