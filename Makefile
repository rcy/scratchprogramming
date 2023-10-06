export ASSET_DIRECTORY?=/media/rcy/data/scratchprogrammingassets

start:
	find . -name \*.go -o -name \*.html | entr -r go run .

dl-%: ${ASSET_DIRECTORY}/videos/%
	cd $< && pipx run yt-dlp https://www.youtube.com/$* -f "bv[ext=webm]+ba[ext=webm]" --merge-output-format webm

.SECONDARY: ${ASSET_DIRECTORY}/videos/%
${ASSET_DIRECTORY}/videos/%:
	mkdir $@
