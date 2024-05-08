ls -l | awk '  BEGIN { FS="\n" }
{
	i=1;
	while ($(i)) {
		if (NR % 2 == 0) {
			print $i;
		}
		i++;
	}
}

'
