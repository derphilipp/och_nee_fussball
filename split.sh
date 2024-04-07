# Split the file into chunks of 50 lines each.
split -l 80 paarungen.txt

# Initialize a counter for your file naming
counter=0

# Loop through the files created by split (assuming they are named 'xaa', 'xab', etc.)
for file in x*; do
  # Format the counter to a three-digit number with leading zeros
  printf -v num "%03d" $counter
  # Rename the file to your specified pattern
  mv "$file" "paarungen_$num.txt"
  # Increment the counter
  #mastodon-filter delete "Och_nee_Fussball_$num"
  mastodon-filter create "Och_nee_Fussball_$num" --context home,public,thread --action warn "paarungen_$num.txt"
  ((counter++))
done
