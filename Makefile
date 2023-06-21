start:
	find . -name \*.go -o -name \*.html | entr -r go run .

download:
	cd assets/videos && yt-dlp https://www.youtube.com/@ScratchTeam --verbose -f "bv[ext=webm]+ba[ext=webm]" --merge-output-format webm
