# we need to run this:
# sed -e "s/\(#cgo darwin LDFLAGS: -L\).*\(src\/github.com\/c4milo\/govix\)/\1 GOPATH\2/" -i '' *.go
# except "GOPATH" needs to be the $GOPATH with forward-slashes escaped

# TODO: perhaps allow an argument to this script, and use that instead of $GOPATH
# TODO: somehow use /Applications/VMware Fusion.app/Contents/Public and ./libvx ?? peace doesn't understand...

PATH_SEP="/"
PATH_SEP_ESCAPED="\\/"
GOPATH_ESCAPED="${GOPATH//$PATH_SEP/$PATH_SEP_ESCAPED}"

sed -e "s/\(#cgo darwin LDFLAGS: -L\).*\(src\/github.com\/c4milo\/govix\)/\1 $GOPATH_ESCAPED\/\2/" -i '' *.go
