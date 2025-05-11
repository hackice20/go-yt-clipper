# Video and Audio Download/Trim Tool

This Go-based tool allows users to download and manipulate videos and audio files. It leverages `yt-dlp` for downloading and `ffmpeg` for video/audio trimming and conversion. The tool provides options to:

![Screeenshot](https://github.com/hackice20/go-yt-clipper/blob/main/Screenshot%20From%202025-05-11%2017-07-43.png)


* Download full video
* Download audio only
* Trim videos to a specific start and end time
* Trim audio to a specific start and end time

The files are saved to a `downloads/` folder in the current directory, and all downloads are organized by title or timestamp.

## Features

* **Download Full Video**: Downloads the full video from a provided URL.
* **Download Audio**: Downloads only the audio from a provided URL.
* **Trimmed Video**: Trims the downloaded video between a start and end time, then saves it as a new video file.
* **Audio Trim**: Trims the audio file between a start and end time, then saves it as a new audio file.
* **File Organization**: Downloads and outputs are stored in a `downloads/` folder to prevent clutter in the working directory.
* **Timestamped Output**: Trimmed files are saved with timestamps to avoid overwriting existing files.

## Requirements

To use this tool, you need to install the following dependencies:

1. **Go Programming Language**:

   * Go (Golang) is required to build and run this tool.
   * Download Go from the official site: [go](https://golang.org/dl/)

2. **yt-dlp** (for video/audio downloading):

   * `yt-dlp` is a command-line tool that allows downloading videos from various websites.
   * You can install `yt-dlp` by following the instructions at: [yt-dlp](https://github.com/yt-dlp/yt-dlp#installation)

3. **FFmpeg** (for video/audio trimming and conversion):

   * FFmpeg is a powerful multimedia processing tool that enables video and audio manipulation.
   * Install FFmpeg by following the instructions here: [ffmpeg](https://ffmpeg.org/download.html)



## Usage

Once the program runs, you will be prompted to choose an action:

1. **Download Full Video**: Downloads a full video from the provided URL.
2. **Download Audio**: Downloads only the audio from the provided URL.
3. **Trimmed Video**: Allows you to trim a video between a start and end time.
4. **Audio Trim**: Allows you to trim audio between a start and end time.

After selecting the option, follow the on-screen prompts to provide the necessary URL and other details like start and end times.

## Output Files

All downloads and processed files will be saved in the `downloads/` folder in the current directory:

* Videos will be saved as `downloads/[video-title].webm`.
* Audio files will be saved as `downloads/[audio-title].opus`.
* Trimmed videos will be saved with a timestamp (e.g., `downloads/trimmed_YYYYMMDD_HHMMSS.webm`).
* Trimmed audio will be saved as `downloads/output.opus`.

## Cleanup

After processing, temporary files (such as the initial downloaded files) will be deleted automatically.

## MIT LICENSE

---
