if [[ $HERO_ID ]]; then
	HERO_ID=$((HERO_ID)); curl -s https://01.gritlab.ax/assets/superhero/all.json | jq '.[] | select(.id=='"$HERO_ID"') | .connections.relatives' | tr -d "\""
fi
