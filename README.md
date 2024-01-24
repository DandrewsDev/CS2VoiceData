# CS2 Voice extractor

Example code for exporting players voices from CS2 demos into WAV files.

**Valve Matchmaking demos do not contain voice audio data, as such there is nothing to extract from MM demo files.**


## Purpose and goals
The intention of this project is not to be an end user tool for bulk or batch processing demos and extracting voice data. 

However, this should serve as a guideline for how to process the audio data as pulled by 
[demoinfocs-golang](https://github.com/markus-wa/demoinfocs-golang). People using that tool to process their demos 
who wish to also pull voice data can leverage this sample to build that audio processing into their demo processing
tools.


## Setup and processing
1. Pulling all the required dependencies.
   2. `go get ./...`
3. Update the cs2-voide-data.go file with the path to your unzipped demo file.
4. Running the sample
   5. `go run cs2-voice-data.go`


## Dependencies
This project does have a dependency on lib opus, which is easy to install on mac/linux.

Linux:
```sh
sudo apt-get install pkg-config libopus-dev libopusfile-dev
```

Mac:
```sh
brew install pkg-config opus opusfile
```

As for direct application dependencies those are all handled by the go.mod and are all pulled doing the `go get ./...` from step 2 above.

# Acknowledgements

Thanks to [@rumblefrog](https://github.com/rumblefrog) for all their help in getting this working. Check out this excellent blog post about [Reversing Steam Voice Codec](https://zhenyangli.me/posts/reversing-steam-voice-codec/) and their work on [Source Chat Relay](https://github.com/rumblefrog/source-chat-relay)

This sample relies on [demoinfocs-golang](https://github.com/markus-wa/demoinfocs-golang). Thank you to [@markus-wa](https://github.com/markus-wa), [@akiver](https://github.com/akiver) and all the contributors there.
