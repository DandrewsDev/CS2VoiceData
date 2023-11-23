# CS2 Voice extractor

Example code for exporting players voices from CS2 demos into WAV files.

**Valve Matchmaking demos do not contain voice audio data, as such there is nothing to extract from MM demo files.**


## Setup and processing
1. Pulling all the required dependencies.
   2. `go get ./...`
3. Update the cs2-voide-data.go file with the path to your unzipped demo file.
4. Running the sample
   5. `go run cs2-voide-data.go`


# Acknowledgements

Thanks to [@rumblefrog](https://github.com/rumblefrog) for all their help in getting this working. Check out this excellent blog post about [Reversing Steam Voice Codec](https://zhenyangli.me/posts/reversing-steam-voice-codec/) and their work on [Source Chat Relay](https://github.com/rumblefrog/source-chat-relay)