export ASSET_DIRECTORY?=/mnt/scratchprogrammingassets

start:
	find . -name \*.go -o -name \*.html | entr -r go run .

dl: dl-@scratchcatthings dl-@ScratchTeam dl-@JaredOwen dl-@griffpatch dl-@Zinnea dl-@KritaOrgPainting

YTDLP := $(shell pwd)/vbin/yt-dlp

dl-%: ${ASSET_DIRECTORY}/videos/%
	cd $< && ${YTDLP} "https://www.youtube.com/$*" -f "bv[ext=webm]+ba[ext=webm]" --merge-output-format webm

ruby: ${ASSET_DIRECTORY}/videos/ruby-playlist
	cd $< && ${YTDLP} "https://www.youtube.com/playlist?list=PLDp1sDeXBYvuViLOv4RcnLPBSAqkaM6dS" -f "bv[ext=webm]+ba[ext=webm]" --merge-output-format webm

.SECONDARY: ${ASSET_DIRECTORY}/videos/%
${ASSET_DIRECTORY}/videos/%:
	mkdir $@

tunnel:
	docker run cloudflare/cloudflared:latest tunnel --no-autoupdate run --token eyJhIjoiMjI0ZDczOGUwMGI0OWE0N2Q5YTIwNTQ2MzUxZTdiNTciLCJ0IjoiMTUxMTNiMTAtZGRiOS00YWNhLWI2ZGMtZTY1OGZkY2RkNmRlIiwicyI6IlkyTTFaVFl3TTJZdE5UazJOeTAwTW1FNExUbGxNalV0T0dJNE5tSTFaRFl6T1RrMCJ9
