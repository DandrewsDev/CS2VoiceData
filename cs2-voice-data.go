package main

import (
	"CS2VoiceData/decoder"
	"fmt"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/msgs2"
	"log"
	"os"
	"strconv"
)

func main() {
	// Create a map of a users to voice data.
	// Each chunk of voice data is a slice of bytes, store all those slices in a grouped slice.
	var voiceDataPerPlayer = map[string][][]byte{}

	// The file path to an unzipped demo file.
	file, err := os.Open("1-34428882-6181-4c75-a24b-4982764122e2.dem")
	if err != nil {
		log.Fatal("Failed to open demo file")
	}
	defer file.Close()

	parser := dem.NewParser(file)

	// Add a parser register for the VoiceData net message.
	parser.RegisterNetMessageHandler(func(m *msgs2.CSVCMsg_VoiceData) {
		// Get the users Steam ID 64.
		steamId := strconv.Itoa(int(m.GetXuid()))
		// Append voice data to map
		voiceDataPerPlayer[steamId] = append(voiceDataPerPlayer[steamId], m.Audio.VoiceData)
	})

	// Parse the full demo file.
	err = parser.ParseToEnd()

	// For each users data, create a wav file containing their voice comms.
	for playerId, voiceData := range voiceDataPerPlayer {
		wavFilePath := fmt.Sprintf("%s.wav", playerId)
		convertAudioDataToWavFiles(voiceData, wavFilePath)
	}

	defer parser.Close()
}

func convertAudioDataToWavFiles(payloads [][]byte, fileName string) {
	// This sample rate can be set using data from the VoiceData net message.
	// But every demo processed has used 24000 and is single channel.
	voiceDecoder, err := decoder.NewOpusDecoder(24000, 1)

	if err != nil {
		fmt.Println(err)
	}

	o := make([]int, 0, 1024)

	for _, payload := range payloads {
		c, err := decoder.DecodeChunk(payload)

		if err != nil {
			fmt.Println(err)
		}

		// Not silent frame
		if len(c.Data) > 0 {
			pcm, err := voiceDecoder.Decode(c.Data)

			if err != nil {
				fmt.Println(err)
			}

			converted := make([]int, len(pcm))
			for i, v := range pcm {
				// Float32 buffer implementation is wrong in go-audio, so we have to convert to int before encoding
				converted[i] = int(v * 2147483647)
			}

			o = append(o, converted...)
		}
	}

	outFile, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Encode new wav file, from decoded opus data.
	enc := wav.NewEncoder(outFile, 24000, 32, 1, 1)

	buf := &audio.IntBuffer{
		Data: o,
		Format: &audio.Format{
			SampleRate:  24000,
			NumChannels: 1,
		},
	}

	// Write voice data to the file.
	if err := enc.Write(buf); err != nil {
		fmt.Println(err)
	}

	enc.Close()
}
