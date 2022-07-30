foo=`cat foo.json | openssl dgst -sha256 -binary | base64`
bar=`cat bar.json | openssl dgst -sha256 -binary | base64`

http PUT localhost:11000/objects/test "Digest: SHA-256=${foo}" < foo.json
http PUT localhost:11000/objects/test "Digest: SHA-256=${bar}" < bar.json

http localhost:11000/locate/${foo}
http localhost:11000/locate/${bar}

http localhost:11000/versions/test
http localhost:11000/objects/test
http "localhost:11000/objects/test?version=1"

http DELETE localhost:11000/objects/test
http localhost:11000/versions/test
http localhost:11000/objects/test
http "localhost:11000/objects/test?version=1"