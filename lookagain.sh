# find . -type f -name '*.sh' | rev | cut -c 4- | rev | cut -c 3- | sort -r
find . -type f -name '*.sh' | rev | cut -c 4- | rev    |    rev | cut -d '/' -f 1 | rev    | sort -r
