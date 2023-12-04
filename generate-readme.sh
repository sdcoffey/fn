set -eof pipefail

go test -v .

cp ./examples_test.go tmp/examples.go

go build -o ./bin/dochelp ./cmd/dochelp/main.go

example_functions=$(go doc tmp | ./bin/dochelp fnname)

echo "# fn\n" > tmp/readme.md
go doc fn | grep -ve "^func.*" | grep -ve "^type.*" | tail -n +3 >> tmp/readme.md

cat >> tmp/readme.md <<- EOF
## Installation

\`\`\`
go get github.com/sdcoffey/fn
\`\`\`
EOF

echo "## In this package" >> tmp/readme.md
for example_name in ${example_functions}; do
  src_name=$(echo $example_name | sed -e "s/^Example//")
  echo "* [$src_name](#$src_name)" >> tmp/readme.md
done

echo "\n" >> tmp/readme.md

echo "Generating documentation for:"
for example_name in ${example_functions}; do
  src_name=$(echo $example_name | sed -e "s/^Example//")

  echo "### $src_name" >> tmp/readme.md
  go doc $src_name | ./bin/dochelp extract-doc >> tmp/readme.md
  echo "\n\`\`\`go" >> tmp/readme.md
  go doc -src tmp $example_name | tail -n +3 >> tmp/readme.md
  echo "\`\`\`\n" >> tmp/readme.md

  echo "- $src_name"
done

mv tmp/readme.md ./README.md

echo "\ndone"
