set -eof pipefail

last_tag=$(git tag --list | tail -n1)
echo "Last tag: $last_tag"

echo "New version?"
read new_tag

if [[ $new_tag =~ ^v[0-9]+\.[0-9]+ ]]; then
  ./generate-readme.sh
  git add README.md
  git commit -m "Releasing $new_tag: update README.md" || true
  git tag $new_tag
  git push origin $new_tag
  git push origin main

  GOPROXY=proxy.golang.org go list -m "github.com/sdcoffey/fn@$new_tag"
else
  echo "Invalid tag: $new_tag -- must be in the form v0.0.0"
  exit 1
fi

