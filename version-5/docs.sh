BASEDIR=$(dirname "$0")
GO111MODULE=off swagger generate spec -o $BASEDIR/swagger/swagger.yml --scan-models 
GO111MODULE=off swagger generate spec -o $BASEDIR/swagger/swagger.json --scan-models 