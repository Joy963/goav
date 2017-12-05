package main

import (
	"github.com/Joy963/goav/avcodec"
	"github.com/Joy963/goav/avdevice"
	"github.com/Joy963/goav/avfilter"
	"github.com/Joy963/goav/avformat"
	"github.com/Joy963/goav/avutil"
	"github.com/Joy963/goav/swresample"
	"github.com/Joy963/goav/swscale"
	"log"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
