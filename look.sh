if [[ $(find . -name 'a*') ]]; then
	find . -name 'a*'
elif [[ $(find . -type f -name '*z') ]]; then
	find . -type f -name '*z'
else
	find . -type f -name 'z*a!'
fi
