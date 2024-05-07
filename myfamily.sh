if [[ $HERO_ID ]]; then
#	i=0
	curl -s https://01.gritlab.ax/assets/superhero/all.json | jq --arg n "$HERO_ID" '.[$n | tonumber - 1].connections.relatives' | tr -d \"
fi
