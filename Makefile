export ASSET_DIRECTORY=${HOME}/data/scratchprogrammingassets

start:
	find . -name \*.go -o -name \*.html | entr -r go run .

download:
	cd ${ASSET_DIRECTORY}/videos && yt-dlp https://www.youtube.com/@ScratchTeam --verbose -f "bv[ext=webm]+ba[ext=webm]" --merge-output-format webm

download2:
	cd ${ASSET_DIRECTORY}/videos && yt-dlp https://www.youtube.com/@Zinnea --verbose -f "bv[ext=webm]+ba[ext=webm]" --merge-output-format webm
